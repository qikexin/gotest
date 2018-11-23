package parser

import (
	"test/crawller/engine"
	"regexp"
	"strconv"
	"test/crawller/model"
)
//以下的正则匹配profile里的每个字段（除name字段外）
var ageRe    = regexp.MustCompile(`<td><span class="label">年龄: </span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高: </span>(\d+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重: </span><span fileld="">(\d+)KG</span></td>`)

var incomeRe = regexp.MustCompile(`<td><span class="label">月收入: </span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别: </span><span fileld="">([^<]+)</span></td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座: </span><span fileld="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况: </span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历: </span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业: </span><span fileld="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯: </span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件: </span><span fileld="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车: </span><span fileld="">([^<]+)</span></td>`)


func ParseProfile(contents []byte,name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents,ageRe))  //extractString返回的是字符串，所以年龄需要转成数字
	if err !=nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents,heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents,weightRe))
	if err != nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents,incomeRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Xinzuo = extractString(contents,xingzuoRe)
	profile.Marriage = extractString(contents,marriageRe)
	profile.Education = extractString(contents,educationRe)
	profile.Occupation = extractString(contents,occupationRe)
	profile.Hokou = extractString(contents,hokouRe)
	profile.House = extractString(contents,houseRe)
	profile.Car = extractString(contents,carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte,re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}else {
		return ""
	}
}