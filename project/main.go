package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name string
	Delay time.Duration
	Number int
}

type Worker struct {
	Id int
	JobQueue chan Job
	QuitChan chan bool
	WorkerPool chan chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id: id,
		JobQueue: make(chan Job),
		WorkerPool: workerPool,
		QuitChan: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker id %d started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker id %d was finished, result %d", w.Id, fib)

			case <- w.QuitChan:
				fmt.Printf("Worker id %d was finished\n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	fibonacci := 0
	memo := make(map[int]int)

	for i := 0; i < number - 1; i++ {
		if i <= 1 {
			memo[i] = i
		} else {
			memo[i] = memo[i - 1] + memo[i - 2]
		}

		fibonacci = fibonacci + memo[i]
	}

	return fibonacci + 1
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorker int
	JobQueue chan Job
}

func NewDispatcher(queue chan Job, workers int) *Dispatcher {
	return &Dispatcher{
		JobQueue: queue,
		MaxWorker: workers,
		WorkerPool: make(chan chan Job, workers),
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <- d.JobQueue:
			go func() {
				wq := <-d.WorkerPool
				wq <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(res http.ResponseWriter, req *http.Request, queue chan Job) {
	// We only accept post requests.
	if req.Method != "POST" {
		res.Header().Set("Allow", "POST")
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(req.FormValue("delay"))

	if err != nil {
		log.Println(err)
		http.Error(res, "Invalid delay", http.StatusUnprocessableEntity)
		return
	}

	value, err := strconv.Atoi(req.FormValue("value"))
	if err != nil {
		http.Error(res, "Invalid value", http.StatusUnprocessableEntity)
		return
	}

	name := req.FormValue("name")
	if name == "" {
		http.Error(res, "Invalid name", http.StatusUnprocessableEntity)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	queue <- job
	res.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers = 4
		maxQueueSize = 20
		appPort = ":8081"
	)

	queue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(queue, maxWorkers)
	dispatcher.Run()

	http.HandleFunc("/fib", func (res http.ResponseWriter, req *http.Request) {
		RequestHandler(res, req, queue)
	})

	println("Running on port 8081")
	log.Fatal(http.ListenAndServe(appPort, nil))
}
