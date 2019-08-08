package indexmap

import "sync"

const (
	MainIndexType int32 = iota
	UniqueIndexType
	NormalIndexType
)

type IndexManage struct {
	mutex sync.RWMutex
	Mainindexname string
	Indexs map[string]Index
}

func (this *IndexManage) AddData(name string, data interface{}) bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()

	_, ok := this.Indexs[name]

	if !ok {
		return false
	}

	v, ok := this.Indexs[this.Mainindexname]

	if !ok {
		return false
	}

	rev := v.Get(v.GetIndexField()(data)...)

	if nil != rev {
		return false
	}

	datatmp := &Record{
		IsValid : true,
		Data : data,
	}

	for _, v := range this.Indexs {
		v.Add(datatmp)
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

	revdata := this.Indexs[name].Get(key...)

	if nil == revdata {
		return
	}

	this.Indexs[name].Delete(key...)
}

func (this *IndexManage) GetData(name string, hashfield ...interface{})([]interface{}){
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	_, ok := this.Indexs[name]

	if !ok {
		return nil
	}

	revdata := this.Indexs[name].Get(hashfield...)

	if nil == revdata {
		return nil
	}

	return revdata
}

func (this *IndexManage) UpdateData(name string, data interface{})bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 查找是否存在索引
	index, ok := this.Indexs[name]

	if !ok {
		return false
	}

	// 判断索引是不是唯一键或者主键
	if UniqueIndexType != index.GetIndexType() && MainIndexType != index.GetIndexType() {
		return false
	}

	// 查找是否存在主键
	v, ok := this.Indexs[this.Mainindexname]

	if !ok {
		return false
	}

	// 查找是否有数据
	rev := v.Get(v.GetIndexField()(data)...)
	if nil == rev {
		return false
	}

	// 通过主索引删除数据
	for _, tmp := range rev {
		v.Delete(v.GetIndexField()(tmp)...)
	}

	for _, v := range this.Indexs {
		datatmp := &Record{
			IsValid : true,
			Data : data,
		}

		v.Add(datatmp)
	}

	return true
}