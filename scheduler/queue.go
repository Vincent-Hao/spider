package scheduler

import (
    "github.com/Vincent-Hao/spider/model"
)

type QueueScheduler struct {
    requestChan chan model.Request
    workChan chan chan model.Request
}

func (q *QueueScheduler) Submit(request model.Request){
    go func() {
        q.requestChan <- request
    }()
}
func (q *QueueScheduler) GetWorkChan() chan model.Request{
    return make(chan model.Request)
}

func (q *QueueScheduler) WorkReady(w chan model.Request){
    q.workChan <- w
}

func (q *QueueScheduler) Run(){
    q.requestChan = make(chan model.Request)
    q.workChan = make(chan chan model.Request)
    go func() {
        var requestQ  []model.Request
        var workerQ   []chan model.Request
        for{
            var activeRequest model.Request
            var activeWork  chan model.Request
            if len(requestQ) > 0 && len(workerQ) > 0 {
                activeRequest = requestQ[0]
                activeWork = workerQ[0]
            }
            select{
            case r := <- q.requestChan:
                requestQ = append(requestQ,r)
            case w := <- q.workChan:
                workerQ = append(workerQ,w)
            case activeWork <- activeRequest:
                requestQ = requestQ[1:]
                workerQ = workerQ[1:]
            }
        }
    }()
}