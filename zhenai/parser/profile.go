package parser

import (
	"github.com/Vincent-Hao/spider/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<tr><td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身[^<]+高：</span>([\d]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="grayL">月   薪：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(``)
var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="grayL">学   历：</span>([^<]+)/td>`)
var marrigeRe = regexp.MustCompile(`<tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
var nameRe = regexp.MustCompile(`<tr><th><a href="(http://[0-9a-z]+.zhenai.com/u/)([0-9]+)" target="_blank">([^<]+)</a></th></tr>`)

var cityPageRe = regexp.MustCompile(`<li class="paging-item"><a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+/[0-9]+)">下一页</a> </li>`)
var manRe = regexp.MustCompile(`<div class="content"><table><tbody><tr><th><a href="http://[a-z0-9]+.zhenai.com/u/([0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td> <td><span class="grayL">居住地：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([0-9]+)</td>  <td><span class="grayL">月[^<]+薪：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td> <td width="180"><span class="grayL">身[^<]+高：</span>([0-9]+)</td></tr></tbody></table> <div class="introduce">([^<]+)</div></div>`)
var womanRe = regexp.MustCompile(`<div class="content"><table><tbody><tr><th><a href="http://[a-z0-9]+.zhenai.com/u/([0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td> <td><span class="grayL">居住地：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([0-9]+)</td> <td><span class="grayL">学[^<]+历：</span>([^<]+)</td> </tr> <tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td> <td width="180"><span class="grayL">身[^<]+高：</span>([0-9]+)</td></tr></tbody></table> <div class="introduce">([^<]+)</div></div>`)
func ParseProfile(content []byte) model.RequestResult {
	profile := model.Profile{}
	var items []interface{}
	vipman := manRe.FindAllStringSubmatch(string(content),-1)
	for  _,vip := range vipman{
		profile.Id = vip[1]
		profile.Name = vip[2]
		profile.Gender = vip[3]
		profile.Address = vip[4]
		age,_ := strconv.Atoi(vip[5])
		profile.Age = age
		profile.Incoming = vip[6]
		profile.Marriage = vip[7]
		height,_ := strconv.Atoi(vip[8])
		profile.Height = height
		profile.Label = vip[9]
		items = append(items,profile)
		//fmt.Println(profile)
	}
	vipWoman := womanRe.FindAllStringSubmatch(string(content),-1)
	for  _,vip := range vipWoman{
		profile.Id = vip[1]
		profile.Name = vip[2]
		profile.Gender = vip[3]
		profile.Address = vip[4]
		age,_ := strconv.Atoi(vip[5])
		profile.Age = age
		profile.Education = vip[6]
		profile.Marriage = vip[7]
		height,_ := strconv.Atoi(vip[8])
		profile.Height = height
		profile.Label = vip[9]
		//fmt.Println(profile)
		items = append(items,profile)
	}
	var requests []model.Request
	pageUrl := cityPageRe.FindStringSubmatch(string(content))
	if pageUrl != nil{
		requests = append(requests,model.Request{pageUrl[1],ParseProfile})
		return model.RequestResult{Requests:requests,Items:items}
	}
	return model.RequestResult{nil,items}
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(string(content))
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
