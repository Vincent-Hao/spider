package engine

import (
    "fmt"
    "github.com/Vincent-Hao/spider/fetcher"
    "github.com/Vincent-Hao/spider/model"
)


func Run(seeds ...model.Request){
    var requests []model.Request
    for _,seed:= range seeds{
        requests = append(requests,seed)
    }
    for len(requests) > 0 {
        request := requests[0]
        requests = requests[1:]
        body,err := fetcher.Fetch(request.Url)
        if err != nil{
            fmt.Printf("fetcher url:%s error:%s",request.Url,err)
            continue
        }
        fmt.Printf("Fetching: %s\n",request.Url)
        result := request.ParseFunc(body)
        if result.Requests != nil{
            requests = append(requests,result.Requests...)
        }
        for _,item := range result.Items{
            fmt.Printf("request.url:%s, get item:%v",request.Url,item)
            fmt.Println()
        }
    }
    
}