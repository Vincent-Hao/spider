package scheduler

import "github.com/Vincent-Hao/spider/model"

type SimpleScheduler struct {
    WorkChan chan model.Request
}

func (s *SimpleScheduler) Submit(request model.Request) {
    go func() {
        s.WorkChan <- request
    }()
}
func (s *SimpleScheduler) GetWorkChan() chan model.Request {
    return s.WorkChan
}

func (s *SimpleScheduler) WorkReady(in chan model.Request){

}
func (s *SimpleScheduler) Run(){
    s.WorkChan = make(chan model.Request)
}
