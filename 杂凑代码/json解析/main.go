package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	studentInfo struct {
		Name  string `json:"name"`
		Hobby string `json:"hobby"`
	}

	class struct {
		Group string `json:"group"`
		Info  []studentInfo `json:"info"`
	}
)

func main(){
	var data class

	// 读取json文件内容 返回字节切片
	bytes,_ := ioutil.ReadFile("data.json")

	// 将字节切片映射到指定结构体上
	json.Unmarshal(bytes,&data)

	// 打印对象结构
	fmt.Println(data)

	// 循环所有学生的信息
	for idx,val := range data.Info{
        fmt.Printf("index: %v 学生名：%v,爱好: %v\n",idx,val.Name,val.Hobby)
	}

	// 更改值
	data.Group = "1013"
}