package indexmap

import (
	"fmt"
	"testing"
)

type Student struct{
	no int32
	age int32
	name string
}

func NewStudent(name string, age int32, no int32)(*Student){
	return &Student{
		no : no,
		age : age,
		name : name,
	}
}

func TestIndexMap(t *testing.T) {
	indexmanage := &IndexManage{
		Mainindexname : "IndexNameTwo",
		Indexs : map[string]Index{
			"IndexNameOne" : NewUniqueIndex(
				UniqueIndexType,
				"IndexNameOne",
				func(value interface{})(rev []interface{}){
					stu := value.(Student)
					rev = append(rev, (stu).name)
					rev = append(rev, (stu).no)
					return
				},
				func(value interface{},key []interface{}) bool{
					if 2 != len(key) {
						return false
					}
					stu := value.(Student)
					if (stu).name == key[0] && (stu).no == key[1] {
						return true
					}

					return false
				}),
			"IndexNameTwo" : NewUniqueIndex(
				MainIndexType,
				"IndexNameTwo",
				func(value interface{})(rev []interface{}){
					stu := value.(Student)
					rev = append(rev, (stu).no)
					return
				},
				func(value interface{},key []interface{}) bool{
					if 1 != len(key) {
						return false
					}
					stu := value.(Student)

					if (stu).no == key[0] {
						return true
					}

					return false
				}),
			"IndexNameThree" :NewNormalIndex(
				NormalIndexType,
				"IndexNameThree",
				func(value interface{})(rev []interface{}){
					stu := value.(Student)
					rev = append(rev, (stu).age)
					return
				},
				func(value interface{},key []interface{}) bool{
					if 1 != len(key) {
						return false
					}
					stu := value.(Student)

					if (stu).age == key[0] {
						return true
					}

					return false
				}),
		},
	}

	stuOne := NewStudent("Zouqiang1", 30, 1)
	stuTwo := NewStudent("Zouqiang2", 30, 2)
	stuThree := NewStudent("Zouqiang3", 30, 3)

	indexmanage.AddData("IndexNameOne", *stuOne)
	indexmanage.AddData("IndexNameTwo", *stuTwo)
	indexmanage.AddData("IndexNameThree", *stuThree)

	fmt.Println("----------------------------------------------------")
	// 索引1打印数据
	zouqiang1 := indexmanage.GetData("IndexNameOne","Zouqiang1", int32(1))
	zouqiang2 := indexmanage.GetData("IndexNameOne","Zouqiang2", int32(2))
	zouqiang3 := indexmanage.GetData("IndexNameOne","Zouqiang3", int32(3))
	fmt.Println("IndexOne Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexOne Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexOne Zouqiang3 : ", zouqiang3)

	// 索引2打印数据
	zouqiang1 = indexmanage.GetData("IndexNameTwo", int32(1))
	zouqiang2 = indexmanage.GetData("IndexNameTwo", int32(2))
	zouqiang3 = indexmanage.GetData("IndexNameTwo", int32(3))
	fmt.Println("IndexTwo Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexTwo Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexTwo Zouqiang3 : ", zouqiang3)

	// 索引3打印数据
	zouqiang := indexmanage.GetData("IndexNameThree", int32(30))
	fmt.Println("IndexThree Zouqiang : ", zouqiang)

	fmt.Println("----------------------------------------------------")

	fmt.Println("----------------------------------------------------")
	// 删除数据
	indexmanage.DeleteData("IndexNameOne", "Zouqiang1", int32(1))

	// 删除Zouqiang1之后
	// 索引1打印数据
	zouqiang1 = indexmanage.GetData("IndexNameOne","Zouqiang1", int32(1))
	zouqiang2 = indexmanage.GetData("IndexNameOne","Zouqiang2", int32(2))
	zouqiang3 = indexmanage.GetData("IndexNameOne","Zouqiang3", int32(3))

	fmt.Println("IndexOne Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexOne Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexOne Zouqiang3 : ", zouqiang3)

	// 索引2打印数据
	zouqiang1 = indexmanage.GetData("IndexNameTwo", int32(1))
	zouqiang2 = indexmanage.GetData("IndexNameTwo", int32(2))
	zouqiang3 = indexmanage.GetData("IndexNameTwo", int32(3))

	fmt.Println("IndexTwo Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexTwo Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexTwo Zouqiang3 : ", zouqiang3)

	// 索引3打印数据
	zouqiang = indexmanage.GetData("IndexNameThree", int32(30))
	fmt.Println("IndexThree Zouqiang : ", zouqiang)

	fmt.Println("----------------------------------------------------")


	fmt.Println("----------------------------------------------------")

	// 修改数据
	stu := NewStudent("Zouqiang2", 50, 2)

	indexmanage.UpdateData("IndexNameOne", *stu)
	zouqiang1 = indexmanage.GetData("IndexNameOne","Zouqiang1", int32(1))
	zouqiang2 = indexmanage.GetData("IndexNameOne","Zouqiang2", int32(2))
	zouqiang3 = indexmanage.GetData("IndexNameOne","Zouqiang3", int32(3))

	fmt.Println("IndexOne Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexOne Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexOne Zouqiang3 : ", zouqiang3)

	// 索引2打印数据
	zouqiang1 = indexmanage.GetData("IndexNameTwo", int32(1))
	zouqiang2 = indexmanage.GetData("IndexNameTwo", int32(2))
	zouqiang3 = indexmanage.GetData("IndexNameTwo", int32(3))

	fmt.Println("IndexTwo Zouqiang1 : ", zouqiang1)
	fmt.Println("IndexTwo Zouqiang2 : ", zouqiang2)
	fmt.Println("IndexTwo Zouqiang3 : ", zouqiang3)

	// 索引3打印数据
	zouqiang = indexmanage.GetData("IndexNameThree", int32(30))
	fmt.Println("IndexThree Zouqiang : ", zouqiang)
	fmt.Println("----------------------------------------------------")
}
