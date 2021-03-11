package engine

import (
	"fmt"
	"github.com/Vincent-Hao/spider/fetcher"
	"github.com/Vincent-Hao/spider/model"
	"github.com/Vincent-Hao/spider/scheduler"
	"github.com/Vincent-Hao/spider/storage"
)

type ConcurrentEngine struct {
	Scheduler scheduler.Scheduler
	WorkerNum int
	ItemSaver storage.Storage
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	e.Scheduler.Run()
	resultChan := make(chan model.RequestResult)
	//e.Scheduler.RecChanSet(requestChan)
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	for i := 0; i < e.WorkerNum; i++ {
		e.CreateWorker(resultChan, e.Scheduler)
	}
	e.Consumer(resultChan)
}

func (e *ConcurrentEngine) Consumer(resultChan chan model.RequestResult) {
	var count int
	go func() {
		for {
			result := <-resultChan
			for _, re := range result.Requests {
				e.Scheduler.Submit(re)
			}
			for _, item := range result.Items {
				fmt.Println("item: ",item)
				in, ok := item.(model.Profile)
				if ok {
					fmt.Println("save item count:", count)
					id, err := e.ItemSaver.Save(in)
					if err != nil {
						fmt.Println("save error:", id)
					}
					count++
				}

			}
		}
	}()
}
func (e *ConcurrentEngine) CreateWorker(resultChan chan model.RequestResult, s scheduler.Scheduler) {
	workChan := s.GetWorkChan()

	go func() {
		for {
			s.WorkReady(workChan)
			request := <-workChan
			body, err := fetcher.Fetch(request.Url)
			if err != nil {
				fmt.Println("fetch URL error:", err)
				continue
			}
			resultChan <- request.ParseFunc(body)
		}
	}()
}
