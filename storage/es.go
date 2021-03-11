package storage

import (
    "context"
    "fmt"
    "github.com/Vincent-Hao/spider/model"
    "github.com/olivere/elastic"
    "log"
    "os"
)

type EsStorage struct {
    ItemChan chan interface{}
    client *elastic.Client
}

func NewEsStorage() *EsStorage{
    return &EsStorage{make(chan interface{}),NewElasticsearchClient()}
}
func NewElasticsearchClient() *elastic.Client{
    host := "http://192.168.43.195:9200/"
    errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
    client, err := elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host),elastic.SetSniff(false))
    if err != nil {
        panic(err)
    }
    info, code, err := client.Ping(host).Do(context.Background())
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
    
    esversion, err := client.ElasticsearchVersion(host)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch version %s\n", esversion)
    fmt.Println(client)
    return client
}
func (es *EsStorage)Save(in model.Profile) (id string,err error){
    //item :=  <- es.ItemChan
    //client,err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL("http://192.168.0.195:9200"))
    resp,err := es.client.Index().Index("data").Type("zhenai").Id(in.Id).BodyJson(in).Do(context.Background())
    if err != nil {
        fmt.Printf("save %v error: %s",in,err.Error())
        return "",err
    }
    fmt.Printf("save %s successful.",resp.Id)
    return resp.Id,nil
}
