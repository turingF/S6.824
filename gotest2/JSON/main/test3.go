package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
  @Description:
*/

func main() {
	c := make(map[string]interface{})
	c["foo"] = "bar"
	c["hello"] = "world"
	c["xy"] = "xy2"

	// data 返回的是byte[],即序列化后的json字符串
	data,err := json.MarshalIndent(c,"","		")
	if err!= nil {
		log.Println(err)
		return
	}

	fmt.Println(string(data))
}
