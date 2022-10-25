package main

import (
	"encoding/json"
	"log"
	"os"
)

const MaxLength int64 = 2048
const MaxWorker int = 64
const MaxLengthQueue int = 1024

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

type Job struct {
	User User
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// 线程池 从JobQueue中取出任务，然后看是否有空闲的 worker 有的话，取出一个执行。
// Worker 空闲的话加入到 Dispatcher 的 WorkerPool, 其实是 Worker 中的 chan Job 通道加入到 WorkerPool。
// 分配任务同 写入到 JobChannel 中实现， worker 读取 JobChannel ,然后执行。
type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Job
}

func InitWorker() {
	JobQueue = make(chan Job, MaxLengthQueue)
	log.Println("start server...")
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				buf, err1 := json.Marshal(job.User)
				if err1 != nil {
					log.Printf("json.Marshal error:%s\n", err1)
				}
				filename := "post/" + job.User.Name
				fp, err2 := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
				if err2 != nil {
					log.Printf("%v Open file %s error : %s\n", job, job.User.Name, err2)
				}
				fp.WriteString(string(buf))
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)

	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool
				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
