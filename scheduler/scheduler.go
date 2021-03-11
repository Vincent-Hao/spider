package scheduler

import (
    "github.com/Vincent-Hao/spider/model"
)

type Scheduler interface {
    Submit(request model.Request)
    GetWorkChan() chan model.Request
    WorkReady(chan model.Request)
    Run()
}