package parser

import (
    "fmt"
    "github.com/Vincent-Hao/spider/model"
    "regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]*)</a>`

func ParserCityList(content []byte) model.RequestResult{
    re := regexp.MustCompile(CityListRe)
    matches := re.FindAllStringSubmatch(string(content),-1)
    var requests []model.Request
    var items []interface{}
    for _,m := range matches{
        items = append(items,m[2])
        requests = append(requests,model.Request{m[1],ParseProfile})
        fmt.Println(m[2])
    }
    return model.RequestResult{requests,items}
}