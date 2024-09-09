package workpool

import (
	"context"
	"fmt"
)

const (
	defaultWorkPool = 10
)

type processmap map[int]joberMap
type joberMap map[int]Jober

type WorkerPool struct {
	maxPool       int
	cancelProcess map[int]context.CancelFunc
	jobs          chan Jober
}

type Job struct {
	id int
	JobFunc
}

type JobFunc func(context.Context) error

type Jober interface {
	Run(context.Context)
}

func (jf JobFunc) Run(ctx context.Context) {
	select {
	// case <-ctx.Done():
	default:
		err:=jf(ctx)
		if err != nil {
			fmt.Println("err",err)
		}
	}

}

func (j *Job) Run(ctx context.Context) {
	j.JobFunc(ctx)
}

func NewWorkerPool(max int) *WorkerPool {
	return &WorkerPool{
		maxPool: max,
		jobs:    make(chan Jober, max*3),
	}
}

func (p *WorkerPool) Start() error {
	for v := range p.maxPool {
		go work(p.jobs, v)
	}

	return nil
}

func (p *WorkerPool) Add(job Jober) error {
	p.jobs <- job
	//add id ?
	return nil
}

func (p *WorkerPool) CancelJob(id int) error {
	p.cancelProcess[id]()
	return nil
}

func (p *WorkerPool) End() error {
	return nil
}

func CreatJob(jf func(context.Context) error) JobFunc {
	return JobFunc(jf)
}

func work(jobs chan Jober, processId int) {
	fmt.Println("procces id", processId)
	for job := range jobs {
		fmt.Println("job running")
		ctx := context.TODO()
		job.Run(ctx)
	}
}
