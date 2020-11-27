package main

import (
	//"TestClient/dbop"
	"fmt"
	"time"
)

type userlist struct {
	msglist []string
}

/*
var pool *redis.Pool

const (
	REDISIP = "192.168.20.160:6379"
)

type Pair struct {
	Value string
	Key   int
}
type PairList []Pair

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 20,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", REDISIP)

			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	fmt.Println("init successful")

}
func GetUserId(aid int) map[int]int {

	var mp map[int]int = make(map[int]int, 0)
	var plist PairList

	redkey := fmt.Sprintf("Set:Active:ActiveUser:AID:183963")
	redkey1 := fmt.Sprintf("Set:Active:ActiveUser:AID:%d by Hash:Active:UserData:AID:%d:UID:*", aid, aid)

	fmt.Println(redkey1)

	fmt.Println(redkey)

	conn := pool.Get()

	defer conn.Close()

	strval, _ := redis.Ints(conn.Do("SMEMBERS", redkey))

	fmt.Println(strval)

	fmt.Println("GetUserid...")

	//sort1, _ := conn.Do("SORT", redkey1)

	for _, value := range strval {

		var pair Pair

		redkey2 := fmt.Sprintf("Hash:Active:UserData:AID:%d:UID:%d", aid, value)

		valuestr, _ := redis.Strings(conn.Do("HMGET", redkey2, "factor"))

		pair.Value = valuestr[0]

		pair.Key = value

		plist = append(plist, pair)
	}

	fmt.Println(plist)
	//排序
	sort.Sort(plist)
	fmt.Println(plist)

	//排序完成后进行排名
	plen := len(plist) - 1
	rank := 1
	for index := 0; index < plen; index++ {

		mp[plist[index].Key] = rank
		if 0 != strings.Compare(plist[index].Value, plist[index+1].Value) {
			rank++
		}

		if plen == index+1 {
			mp[plist[index+1].Key] = rank
		}
	}

	fmt.Println(mp)
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	result := strings.Compare(p[i].Value, p[j].Value)
	if result == -1 || result == 0 {
		return false
	}
	return true
}
*/
func main() {
	workerChan := make(chan string)
	exitChan := make(chan int)

	go func() {
		<-time.After(time.Second * 1)
		select {
		case workerChan <- "Task 1":
		}
		fmt.Println("Task 1 exit")
	}()

	go func() {
		select {
		case workerChan <- "Task 2":
		case <-exitChan:
		}
		fmt.Println("Task 2 exit")
	}()

	go func() {
		select {
		case <-time.After(time.Second * 2):
		}
		fmt.Println("Close exitChan")
		close(exitChan)
		<-time.After(time.Second * 4)
		close(workerChan)
	}()

	<-time.After(time.Second * 3)
	fmt.Println("Start receive from workerChan")
loop:
	for {
		select {
		case i, ok := <-workerChan:
			if ok {
				fmt.Println("Receive:", i)
			} else {
				break loop
			}

		}
	}

	<-time.After(time.Second * 1)
}
/*
Task1和Task2是两个生产者，它们都向workerChan发送消息，其中Task2立即发送，Task1有一定延时，workerChan是一个阻塞的go channel。 
同时，有一个go channel发送结束信号（关闭exitChan）。随后开启消费者，接收workerChan的消息， 
Task1和Task2的区别是Task2在select中多了一个对exitChan的监听。

从结果可以看出，当exitChan被关闭时，Task2结束对workerChan的阻塞，取消了像worker发送信号，同时结束了自身。 
而没有监听exitChan的Task1依然在阻塞，直到被读取后才退出。

示例说明了可以通过对exitChan的使用来结束对阻塞go channel的等待。需要说明的是，在真实场景中， 
消费者在发出结束的意图后可能并不会去处理尚未被处理的消息，所以像示例中的Task1是无法正常结束的。

*/
