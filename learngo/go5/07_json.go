package main

import (
	"encoding/json"
	"fmt"
)

type It_temp struct {
	Company  string   `json:"company"` // 二次编码
	Subjects []string `json:"-"`       //忽略该类型
	IsOk     bool     `json:",string"` //改变类型
	Price    float64
}

type It struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func (this It) Encoder() (res string) {
	str, err := json.MarshalIndent(this, "", " ")
	// str, err := json.Marshal(this)
	if err != nil {
		// fmt.Println(err)
		res = err.Error()
		return
	}
	res = string(str)
	return
	// fmt.Printf("%s\n", str)
}

func (this *It) Decoder(str string) (err error) {
	err = json.Unmarshal([]byte(str), this)
	return
}
func main() {

	// struct to json
	s := It{"itcast", []string{"c++", "go"}, true, 666.6}

	str := s.Encoder()
	fmt.Println(str)

	//创建一个map to json
	m := make(map[string]interface{}, 4)
	m["company"] = "hello"
	m["subject"] = []string{"go", "c++", "python"}
	m["isok"] = true
	m["price"] = 3.333
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	str = string(res)
	fmt.Println(str)
	// #################### JSON to struct #####################
	str = `{
		"Company": "itcast",
		"Subjects": [
		 "c++",
		 "go"
		],
		"IsOk": true,
		"Price": 666.6
	   }`
	var itt It
	itt.Decoder(str)
	fmt.Printf("%+v", itt)
}
