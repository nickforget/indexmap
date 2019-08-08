package indexmap

import (
	"github.com/mitchellh/hashstructure"
)

type NormalIndex struct {
	IndexType int32
	IndexName string
	GetField GetIndexField
	CompareField CompareIndexField
	IndexData map[interface{}]([]*Record)
}

func NewNormalIndex(indextype int32, indexname string, getfield GetIndexField, comparefield CompareIndexField) *NormalIndex{
	return &NormalIndex{
		IndexType : indextype,
		IndexName : indexname,
		GetField : getfield,
		CompareField : comparefield,
		IndexData : make(map[interface{}]([]*Record), 0),
	}
}

func (index *NormalIndex) GetIndexType() int32{
	return index.IndexType
}

func (index *NormalIndex) GetIndexName() string{
	return index.IndexName
}

func (index *NormalIndex) GetIndexField()GetIndexField{
	return index.GetField
}

func (index *NormalIndex) Add(data *Record){
	// 获取索引字段
	hashfield := index.GetField(data.Data)

	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 添加数据
	index.IndexData[hash] = append(index.IndexData[hash], data)
}

func (index *NormalIndex) Delete(hashfield ...interface{}){
	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 查找对应的数据
	for i, v := range index.IndexData[hash] {
		if index.CompareField(v.Data, hashfield){
			index.IndexData[hash][i].IsValid = false
		}
	}
}

func (index *NormalIndex) Get(hashfield ...interface{})(value []interface{}){
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
			value = append(value, v.Data)
		}
	}

	return value
}

func (index *NormalIndex) Update(data *Record, hashfield ...interface{}){
	index.Delete(hashfield...)
	index.Add(data)
}