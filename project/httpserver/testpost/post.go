package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type User struct {
	Name     string
	Age      int
	Password string
}

type JsonReq struct {
	Version string
	Token   string
	Users   []User
}

func GenJsonReq() string {
	version := "0.0.0"
	token := "aGVsbG8gZ28sIEknbSB0ZXN0aW5nIGh0dHAu"
	users := make([]User, 1)
	for i := range users {
		users[i] = *GenUser()
	}
	jsonReq := JsonReq{Version: version, Token: token, Users: users}
	reqBody, err := json.Marshal(jsonReq)
	if err != nil {
		log.Printf("GenJsonReq json Marshal")
	}
	return *(*string)(unsafe.Pointer(&reqBody)) //减少了一次拷贝  但是 unsafe
}

func GenUser() *User {
	return &User{Name: randStr(10), Age: 18, Password: randStr(15)}
}

func PostRequest(url, reqBody string) (string, error) {
	// log.Printf("url:%s,body:%s", url, reqBody)
	resp, err := http.Post(url, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return "post", err
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return "body", err1
	}
	return string(body), nil
}

type Request struct {
	url  string
	body string
}

var RequestQueue chan Request

type Worker struct {
	WorkerPool     chan chan Request
	RequestChannel chan Request
	quit           chan bool
}

func NewWorker(workerPool chan chan Request) Worker {
	return Worker{
		WorkerPool:     workerPool,
		RequestChannel: make(chan Request),
		quit:           make(chan bool)}
}
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.RequestChannel
			select {
			case Request := <-w.RequestChannel:
				body, err := PostRequest(Request.url, Request.body)
				if err != nil {
					log.Println(body, err)
					continue
				}
				// _ = Request.url[:]
				log.Println("success post")
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

type WorkerPool struct {
	workers    []Worker
	pool       chan chan Request
	maxWorkers int
	runEnd     chan bool
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	pool := make(chan chan Request, maxWorkers)
	runEnd := make(chan bool)
	return &WorkerPool{pool: pool, maxWorkers: maxWorkers, runEnd: runEnd}
}

func (d *WorkerPool) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.pool)
		worker.Start()
		d.workers = append(d.workers, worker)
	}
	go d.dispatch()
}

func (d *WorkerPool) dispatch() {
	for request := range RequestQueue {
		// log.Println("test")
		requestChannel := <-d.pool
		requestChannel <- request
	}
	for _, worker := range d.workers {
		worker.Stop()
	}
	d.runEnd <- true
}
func InitLog() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("post.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func Init() {
	InitLog()
	RequestQueue = make(chan Request, 10000)
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("usage:please using %s url(string) postNum(int)\n", args[0])
		return
	}
	// RequestQueue = make(chan Request, 1000)
	// go func() {
	// 	for {
	// 		<-RequestQueue
	// 	}
	// }()
	// url := "http://localhost:8000/post/json"
	url := args[1]
	RequestTimes, err := strconv.Atoi(args[2])
	if err != nil {
		RequestTimes = 10000
	}

	Init()
	pool := NewWorkerPool(100)
	pool.Run()

	start := time.Now()
	for i := 0; i < RequestTimes; i++ {
		reqBody := GenJsonReq()
		RequestQueue <- Request{url: url, body: reqBody}
		if i%10000 == 0 {
			duration := time.Since(start)
			fmt.Printf("qps:%f\n", float64(i)/float64(duration.Milliseconds())*1000)
		}
	}
	close(RequestQueue)
	<-pool.runEnd
	duration := time.Since(start)
	fmt.Printf("using time: %d ms\n", duration.Milliseconds())

}
