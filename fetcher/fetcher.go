package fetcher

import (
    "bufio"
    "golang.org/x/net/html/charset"
    "golang.org/x/text/encoding"
    "golang.org/x/text/transform"
    "io/ioutil"
    "net/http"
)

func Fetch(url string) ([]byte,error){
    resp,err := http.Get(url)
    if err != nil{
        return nil,err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK{
        return nil,err
    }
    bufferReader := bufio.NewReader(resp.Body)
    e := determineEncoding(bufferReader)
    utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
    
    return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
    bytes,err := r.Peek(1024)
    if err != nil{
       panic(err)
    }
    e,_,_:= charset.DetermineEncoding(bytes,"")
    return e
}