package indexmap

import "github.com/mitchellh/hashstructure"

type UniqueIndex struct {
	indexType int32
	indexName string
	pFuncGetField PFuncGetField
	pFuncCmpField PFuncCmpField
	indexData map[uint64]([]*Record)
}

func NewUniqueIndex(indextype int32, indexname string, pfuncgetfield PFuncGetField, pfunccmpfield  PFuncCmpField) *UniqueIndex{
	return &UniqueIndex{
		indexType : indextype,
		indexName : indexname,
		pFuncGetField : pfuncgetfield,
		pFuncCmpField : pfunccmpfield,
		indexData : make(map[uint64]([]*Record), 0),
	}
}

func (index *UniqueIndex) Insert(data *Record){
	// 获取索引字段
	hashfield := index.pFuncGetField(data.Data)

	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(hashfield, nil)

	// 添加数据
	index.indexData[hash] = append(index.indexData[hash], data)
}

func (index *UniqueIndex) Delete(field ...interface{}){
	// 计算索引字段的HASH值
	hash, _ := hashstructure.Hash(field, nil)

	// 查找对应的数据
	for i, v := range index.indexData[hash] {
		if index.pFuncCmpField(v.Data, field){
			index.indexData[hash][i].IsValid = false
		}
	}
}

func (index *UniqueIndex) Query(field ...interface{})(data []interface{}){
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
			return []interface{}{v.Data}
		}
	}

	return nil
}

func (index *UniqueIndex) Modify(data *Record, field ...interface{}){
	index.Delete(field...)
	index.Insert(data)
}

func (index *UniqueIndex) IndexType() int32{
	return index.indexType
}

func (index *UniqueIndex) IndexName() string{
	return index.indexName
}

func (index *UniqueIndex) PFuncGetField()PFuncGetField{
	return index.pFuncGetField
}
