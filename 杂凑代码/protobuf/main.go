package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"protobuf/src/test"

	"github.com/golang/protobuf/proto"
)

func write() {
	p1 := &test.Person{
		Id:   1,
		Name: "小张",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "111111111"},
			{Type: test.PhoneType_WORK, Number: "222222222"},
		},
	}
	p2 := &test.Person{
		Id:   2,
		Name: "小王",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "333333333"},
			{Type: test.PhoneType_WORK, Number: "444444444"},
		},
	}

	//创建地址簿
	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./test.txt")
	book := &test.ContactBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
