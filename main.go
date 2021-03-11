package main

import (
	"github.com/Vincent-Hao/spider/engine"
	"github.com/Vincent-Hao/spider/model"
	"github.com/Vincent-Hao/spider/scheduler"
	"github.com/Vincent-Hao/spider/storage"
	"github.com/Vincent-Hao/spider/zhenai/parser"
	"net/http"
)

func main() {

	//f, err := os.Create("cpuprofile")
	//if err != nil {
	//	fmt.Println("could not create CPU profile: ", err)
	//}
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	fmt.Println("could not start CPU profile: ", err)
	//}
	//defer pprof.StopCPUProfile()
	//
	//// ... rest of the program ...
	//
	//f1, err := os.Create("memprofile")
	//if err != nil {
	//	fmt.Println("could not create memory profile: ", err)
	//}
	//runtime.GC() // get up-to-date statistics
	//if err := pprof.WriteHeapProfile(f1); err != nil {
	//	fmt.Println("could not write memory profile: ", err)
	//}
	//f1.Close()

	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	panic(resp.StatusCode)
	//}
	////var res []byte
	////n,err := resp.Body.Read(res)
	////fmt.Println(res,n)
	////result,err := httputil.DumpResponse(resp,true)
	////fmt.Println(string(result))
	////fmt.Println(resp.Header.Get("Content-Type"))
	//result2, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("-------------------------")
	////fmt.Println(string(result2))
	//
	////text := "my email is haohaozhang@163.com.cn "
	////re := regexp.MustCompile(`.+@.+\..+`)
	////括号内分割
	////re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\\.[a-zA-Z0-9.]+)")
	////email := re.FindAllStringSubmatch(text,-1)
	////fmt.Println(email)
	////s := GetCityList(result2)
	////fmt.Println(len(s))
	////for i,v := range s{
	////    fmt.Println(i,v)
	////}
	////fmt.Println(s)
	////n := runtime.Stack(buf[:], false)
	////idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	////id, err := strconv.Atoi(idField)
	////if err != nil {
	////    panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	////}
	//requestResult := parser.ParserCityList(result2)
	//fmt.Println(requestResult)
	//e1 := &engine.SimpleEngine{}
	//e1.Run(engine.Request{"http://www.zhenai.com/zhenghun",parser.ParserCityList})
	//engine.Run(engine.Request{"http://www.zhenai.com/zhenghun/aba",parser.ParseProfile})
	e := &engine.ConcurrentEngine{&scheduler.QueueScheduler{},10,storage.NewEsStorage()}
	e.Run(model.Request{"http://www.zhenai.com/zhenghun",parser.ParserCityList})
	http.ListenAndServe("",nil)
	
	
	
}


