package indexmap

import (
	"github.com/mitchellh/hashstructure"
)

type NormalIndex struct {
	indexType int32
	indexName string
	pFuncGetField PFuncGetField
	pFuncCmpField PFuncCmpField
	indexData map[uint64]([]*Record)
}

func NewNormalIndex(indextype int32, indexname string, pfuncgetfield PFuncGetField, pfunccmpfield  PFuncCmpField) *NormalIndex{
	return &NormalIndex{
		indexType : indextype,
		indexName : indexname,
		pFuncGetField : pfuncgetfield,
		pFuncCmpField : pfunccmpfield,
		indexData : make(map[uint64]([]*Record), 0),
	}
}

func (index *NormalIndex) Insert(data *Record){
	// 获取索引字段
	hashfield := index.pFuncGetField(data.Data)

	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 添加数据
	index.indexData[hash] = append(index.indexData[hash], data)
}

func (index *NormalIndex) Delete(field ...interface{}){
	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(field, nil)

	// 查找对应的数据
	for i, v := range index.indexData[hash] {
		if index.pFuncCmpField(v.Data, field){
			index.indexData[hash][i].IsValid = false
		}
	}
}

func (index *NormalIndex) Query(field ...interface{})(data []interface{}){
	// 计算HASH值
	hash, _ := hashstructure.Hash(field, nil)

	// 查找数据
	_, ok:= index.indexData[hash]

	// 没找到返回nil
	if !ok {
		return nil
	}

	// 比较内容
	for _, v := range index.indexData[hash] {
		if v.IsValid && index.pFuncCmpField(v.Data, field){
			data = append(data, v.Data)
		}
	}

	return data
}

func (index *NormalIndex) Modify(data *Record, field ...interface{}){
	index.Delete(field...)
	index.Insert(data)
}

func (index *NormalIndex) IndexType() int32{
	return index.indexType
}

func (index *NormalIndex) IndexName() string{
	return index.indexName
}

func (index *NormalIndex) PFuncGetField()PFuncGetField{
	return index.pFuncGetField
}

