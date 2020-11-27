package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func gso(w http.ResponseWriter, r *http.Request) {
	r1 := r.URL.String()
	r2 := r.RequestURI
	fmt.Println(r1)                         //go?a=111&b=456
	fmt.Println(r2)                         //go?a=111&b=456
	u, _ := url.Parse(r1)                   //将string解析成*URL格式
	fmt.Println(u)                          // go?a=111&b=456
	fmt.Println(u.RawQuery)                 //编码后的查询字符串，没有'?' a=111&b=456
	values, _ := url.ParseQuery(u.RawQuery) //返回Values类型的字典
	fmt.Println(values)                     // map[a:[111] b:[456]]
	fmt.Println(values["a"])                //[111]
	fmt.Println(values.Get("b"))            //456
}

func main() {
	http.HandleFunc("/go", gso)
	http.ListenAndServe(":8080", nil)
}
