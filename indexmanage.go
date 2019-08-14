package indexmap

import "sync"

const (
	MainIndexType int32 = iota
	UniqueIndexType
	NormalIndexType
)

type IndexManage struct {
	mutex sync.RWMutex
	MainIndexName string
	Indexs map[string]Index
}

func (this *IndexManage) InsertData(name string, data interface{}) bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()

	_, ok := this.Indexs[name]

	if !ok {
		return false
	}

	v, ok := this.Indexs[this.MainIndexName]

	if !ok {
		return false
	}

	rev := v.Query(v.PFuncGetField()(data)...)

	if nil != rev {
		return false
	}

	datatmp := &Record{
		IsValid : true,
		Data : data,
	}

	for _, v := range this.Indexs {
		v.Insert(datatmp)
	}

	return true
}

func (this *IndexManage) DeleteData(name string, key ...interface{}){
	this.mutex.Lock()
	defer this.mutex.Unlock()

	_, ok := this.Indexs[name]

	if !ok {
		return
	}

	data := this.Indexs[name].Query(key...)

	if nil == data {
		return
	}

	this.Indexs[name].Delete(key...)
}

func (this *IndexManage) QueryData(name string, field ...interface{})([]interface{}){
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	_, ok := this.Indexs[name]

	if !ok {
		return nil
	}

	data := this.Indexs[name].Query(field...)

	return data
}

func (this *IndexManage) ModifyData(name string, data interface{})bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 查找是否存在索引
	index, ok := this.Indexs[name]

	if !ok {
		return false
	}

	// 判断索引是不是唯一键或者主键
	if UniqueIndexType != index.IndexType() && MainIndexType != index.IndexType() {
		return false
	}

	// 查找是否存在主键
	v, ok := this.Indexs[this.MainIndexName]

	if !ok {
		return false
	}

	// 查找是否有数据
	revdata := v.Query(v.PFuncGetField()(data)...)
	if nil == revdata {
		return false
	}

	// 通过主索引删除数据
	for _, tmp := range revdata {
		v.Delete(v.PFuncGetField()(tmp)...)
	}

	for _, v := range this.Indexs {
		datatmp := &Record{
			IsValid : true,
			Data : data,
		}

		v.Insert(datatmp)
	}

	return true
}