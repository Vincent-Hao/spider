package engine

import (
    "fmt"
    "github.com/Vincent-Hao/spider/fetcher"
    "github.com/Vincent-Hao/spider/model"
)

type SimpleEngine struct {

}

func (e *SimpleEngine)Run(seeds ...model.Request){
    var requests []model.Request
    for _,seed:= range seeds{
        requests = append(requests,seed)
    }
    for len(requests) > 0 {
        request := requests[0]
        requests = requests[1:]
        result,err := e.Worker(request)
        if err != nil {
            fmt.Printf("worker error:%s ",err.Error())
            continue
        }
        if result.Requests != nil {
            requests = append(requests,result.Requests...)
        }
        for _,item := range result.Items{
            fmt.Printf("request.url:%s, get item:%v",request.Url,item)
            fmt.Println()
        }
    }
}

func (e *SimpleEngine)Worker(request model.Request) (model.RequestResult,error){
    body,err := fetcher.Fetch(request.Url)
    if err != nil{
        fmt.Printf("fetcher url:%s error:%s",request.Url,err)
        return model.RequestResult{},err
    }
    fmt.Printf("Fetching: %s\n",request.Url)
    return request.ParseFunc(body),nil
}