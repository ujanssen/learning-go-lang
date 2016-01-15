package main

func main() {
	p1 := newPipeline()
	p2 := newPipeline()
	p3 := newPipeline()

	p1.In = make(chan int)
	p2.In = p1.Out
	p3.In = p2.Out

	p1.Run()
	p2.Run()
	p3.Run()

	p1.In <- 42
	println(<-p3.Out)
}

type pipeline struct {
	In, Out chan int
}

func newPipeline() *pipeline {
	return &pipeline{
		Out: make(chan int)}
}

func (p *pipeline) Run() {
	go func() {
		i := <-p.In
		i++
		p.Out <- i
	}()
}
