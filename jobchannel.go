package main

import (
	//	"fmt"
	//	"reflect"
	//"runtime"
	"flag"
	"net/http"
	//"time"
)

var (
	MaxWorker = 10
)

var JobQueue chan Job

type Payload struct {
	w http.ResponseWriter
}

type Job struct {
	Payloaded Payload
}

func NewWorker(workerPool chan chan Job, no int) Worker {
	fmt.Println("创建一个工作单元")
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	no         int
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			fmt.Println("w.workerpool <- w.Jobchannel", w)
			select {
			case job := <-w.JobChannel:
				fmt.Println("Job:=<-w.Jobchannel")
				fmt.Println(job)
				time.Sleep(100 * time.Second)
			case <-w.quit:
				return
			}
		}

	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}
func (d *Dispatcher) Run() {
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			//fmt.Println("job := <-JobQueue:")
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				//	fmt.Println("jobChannel := <-d.WorkerPool", reflect.TypeOf(jobChannel))
				jobChannel <- job
				//fmt.Println("jobChannel <- job")
			}(job)
		}
	}
}

func main() {
	JobQueue = make(chan Job, 10)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	time.Sleep(1 * time.Second)

	go GetRequest()

	time.Sleep(1000 * time.Second)

}
func GetRequest() {

	host := flag.String("host", "127.0.0.1", "listen host")
	port := flag.String("port", "80", "listen port")

	http.HandleFunc("/GetJson", GetJson)

	err := http.ListenAndServe(*host+":"+*port, nil)

	if err != nil {
		panic(err)
	}
}

func GetJson(w http.ResponseWriter, req *http.Request) {
	//解析json数据
	w.Write([]byte("Hello World"))

	for {
		//job任务
		payLoad := Payload{Num: i}
		work := Job{Payloaded: payLoad}

		JobQueue <- work
		//fmt.Println("JobQueue <- work", i)
		//fmt.Println("当前协程数目", runtime.NumGoroutine())

		time.Sleep(100 * time.Millisecond)
	}

}

//客户端程序
