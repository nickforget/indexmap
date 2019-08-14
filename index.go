package indexmap

// 数据结构
type Record struct {
	IsValid bool
	Data interface{}
}

// 获取字段函数
type PFuncGetField = func(interface{})([]interface{})

// 比较字段函数
type PFuncCmpField = func(interface{}, []interface{}) bool

type Index interface {
 	Insert(data *Record)
	Delete(field ...interface{})
	Query(field ...interface{})(data []interface{})
	Modify(data *Record, field ...interface{})
	IndexType() int32
 	IndexName() string
	PFuncGetField()PFuncGetField
}
