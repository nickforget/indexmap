package indexmap

type Record struct {
	IsValid bool
	Data interface{}
}

type GetIndexField = func(interface{})([]interface{})
type CompareIndexField = func(interface{}, []interface{}) bool

type Index interface {
 	Add(data *Record)
	Delete(hashfield ...interface{})
	Get(hashfield ...interface{})(value []interface{})
	Update(data *Record, hashfield ...interface{})
	GetIndexType() int32
 	GetIndexName() string
	GetIndexField()GetIndexField
}
