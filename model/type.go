package model
type Request struct {
    Url string
    ParseFunc func([]byte) RequestResult
}

type RequestResult struct{
    Requests []Request
    Items    []interface{}
}

