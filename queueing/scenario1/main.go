package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID int

	SystemIn  time.Time
	SystemOut time.Time

	service        time.Time
	processingTime time.Duration
}

func incTaskID() (id int) {
	id = taskID
	taskID++
	return
}
func NewTask() *Task {
	return &Task{
		ID: incTaskID(),

		SystemIn: time.Now(),

		processingTime: 10 * time.Millisecond}
}

func (t *Task) WaitingTime() time.Duration {
	return t.service.Sub(t.SystemIn)
}
func (t *Task) Service() {
	t.service = time.Now()
}

type queue chan (*Task)

var numWorkers, numQueues, numTasks, queueSize int

var queues []queue
var outQueue queue
var taskID int

func dequeue(w int, q queue) {
	for {
		if t, ok := <-q; ok {
			t.Service()
			// fmt.Printf("dequeue task %d with worker %d\n", t.ID, w)
			time.Sleep(t.processingTime)
			outQueue <- t
		} else {
			return
		}
	}
}
func enqueueTo(q queue, n int) {
}

func enqueue() {
	q := 0
	for i := 0; i < numTasks; i++ {
		queues[q] <- NewTask()
		q++
		if q == numQueues {
			q = 0
		}
	}
}

func do() {
	queues = make([]queue, numQueues)
	outQueue = make(queue, numTasks)

	for q := 0; q < numQueues; q++ {
		queues[q] = make(queue, queueSize)
	}
	enqueue()

	for w := 0; w < numWorkers; w++ {
		go dequeue(w, queues[w])
	}

	var waitingTime time.Duration
	var processingTime time.Duration
	for i := 0; i < numTasks; i++ {
		t := <-outQueue
		t.SystemOut = time.Now()
		wt := t.WaitingTime()
		d := t.processingTime
		waitingTime += wt
		processingTime += d
	}
	for q := 0; q < numQueues; q++ {
		close(queues[q])
	}

	fmt.Printf("numTasks %v\n", numTasks)
	fmt.Printf("numQueues %v\n", numQueues)
	fmt.Printf("numWorkers %v\n", numWorkers)
	fmt.Printf("sum waitingTime %v\n", waitingTime)
	fmt.Printf("avg waitingTime %v\n", waitingTime/time.Duration(numTasks))
	fmt.Printf("sum processingTime %v\n", processingTime)
	fmt.Printf("avg processingTime %v\n\n\n", processingTime/time.Duration(numTasks))

}

func main() {
	numTasks = 10000
	queueSize = 10000

	var workers = []int{1, 2, 3, 4, 5, 10, 20, 50, 100, 1000}

	for _, w := range workers {
		numWorkers = w
		numQueues = w

		do()
	}
}

/*
numTasks 10000
numQueues 1
numWorkers 1
sum waitingTime 152h9m19.585499282s
avg waitingTime 54.775958549s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 2
numWorkers 2
sum waitingTime 76h21m54.154876624s
avg waitingTime 27.491415487s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 3
numWorkers 3
sum waitingTime 50h40m41.148476042s
avg waitingTime 18.244114847s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 4
numWorkers 4
sum waitingTime 38h7m0.633965812s
avg waitingTime 13.722063396s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 5
numWorkers 5
sum waitingTime 30h22m52.086352917s
avg waitingTime 10.937208635s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 10
numWorkers 10
sum waitingTime 15h12m10.519516589s
avg waitingTime 5.473051951s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 20
numWorkers 20
sum waitingTime 7h36m28.041815865s
avg waitingTime 2.738804181s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 50
numWorkers 50
sum waitingTime 3h3m28.090872345s
avg waitingTime 1.100809087s
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 100
numWorkers 100
sum waitingTime 1h30m38.949729578s
avg waitingTime 543.894972ms
sum processingTime 1m40s
avg processingTime 10ms


numTasks 10000
numQueues 1000
numWorkers 1000
sum waitingTime 11m36.265156071s
avg waitingTime 69.626515ms
sum processingTime 1m40s
avg processingTime 10ms
*/
