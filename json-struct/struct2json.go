package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       int64
	Name     string
	Birthday string
}

type Cache struct {
	Data interface{}
}

func (i *Cache) UnmarshalBinary(data []byte) error {
	fmt.Println("MarshalBinary")
	err := json.Unmarshal(data, i)
	return err
}

func (i *Cache) MarshalBinary() (data []byte, err error) {
	fmt.Println("UnmarshalBinary")
	data, err = json.Marshal(i)
	return data, err
}

func main() {
	u := User{Id: 1, Name: "hello", Birthday: "666"}
	cache := Cache{Data: u}
	data, err := json.Marshal(cache)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	cacheout := &Cache{Data: u}
	err1 := json.Unmarshal(data, cacheout)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(cacheout)
}
