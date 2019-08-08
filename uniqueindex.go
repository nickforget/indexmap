package indexmap

import "github.com/mitchellh/hashstructure"

type UniqueIndex struct {
	IndexType int32
	IndexName string
	GetField GetIndexField
	CompareField CompareIndexField
	IndexData map[interface{}]([]*Record)
}

func NewUniqueIndex(indextype int32, indexname string, getfield GetIndexField, comparefield CompareIndexField) *UniqueIndex{
	return &UniqueIndex{
		IndexType : indextype,
		IndexName : indexname,
		GetField : getfield,
		CompareField : comparefield,
		IndexData : make(map[interface{}]([]*Record), 0),
	}
}

func (index *UniqueIndex) GetIndexType() int32{
	return index.IndexType
}

func (index *UniqueIndex) GetIndexName() string{
	return index.IndexName
}

func (index *UniqueIndex) GetIndexField()GetIndexField{
	return index.GetField
}

func (index *UniqueIndex) Add(data *Record){
	// 获取索引字段
	hashfield := index.GetField(data.Data)

	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 添加数据
	index.IndexData[hash] = append(index.IndexData[hash], data)
}

func (index *UniqueIndex) Delete(hashfield ...interface{}){
	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 查找对应的数据
	for i, v := range index.IndexData[hash] {
		if index.CompareField(v.Data, hashfield){
			index.IndexData[hash][i].IsValid = false
		}
	}
}

func (index *UniqueIndex) Get(hashfield ...interface{})(value []interface{}){
	// 计算HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 查找数据
	_, ok:= index.IndexData[hash]

	// 没找到返回nil
	if !ok {
		return nil
	}

	// 比较内容
	for _, v := range index.IndexData[hash] {
		if v.IsValid && index.CompareField(v.Data, hashfield){
			return []interface{}{v.Data}
		}
	}

	return nil
}

func (index *UniqueIndex) Update(data *Record, hashfield ...interface{}){
	index.Delete(hashfield...)
	index.Add(data)
}
