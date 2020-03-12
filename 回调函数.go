package main

//  测试开启协程加锁问题

import (
	// "fmt"
	"os"
	"strings"
	// "io/ioutil"
	"log"
	// "crypto/md5"
	// "encoding/json"
	// "reflect"
	// "time"
	"sync"
	"path/filepath"
)

var waitRcsv sync.WaitGroup //定义一个同步等待读csv文件

func main(){
	dataall := make(map[string]map[string]string)
	ch := make(chan map[string]string, 0) 

	go findCsv("2-合并按列拆分的表", ch) 

	for data := range ch{ // 一直阻塞读，直到写入端close
		code := data["JTCYBM"]  + data["YCDM"]
		dataall[code] = data
	}
}

// 获取目录下csv文件
func findCsv(root string, ch chan map[string]string){
	
	filepath.Walk(root,	func(path string, info os.FileInfo, err error) error {
		log.Println("findCsv_1")
		if strings.Index(path, ".csv") > 0 {
			go readCsv(path, ch)
		}
		return nil
	})

	waitRcsv.Wait() // 阻塞到数据读完
	close(ch)
}

func readCsv(_ string, ch chan map[string]string){
	log.Println("readCsv")
	waitRcsv.Add(1)
	defer waitRcsv.Done() 

	// for _, row := range rows{
	// 	ch <- data
	// }
}