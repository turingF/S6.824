package JSON

import (
	"encoding/json"
	"fmt"
)

/*
  @Description: 使用map[string] interface{} 来进行通用的json解析
*/

func main() {
	response := `{
		"code":200,
		"data":{
			"foo":"bar",
			"hello":"world"
		}
}`
	// 将json转换为一个map
	var mapRes map[string] interface{}

	err := json.Unmarshal([]byte(response), &mapRes)
	if err != nil {
		println(err)
		return
	}

	// 在解析map元素时需要使用强转，不适合嵌套较多的元素
	vars  := mapRes["data"].(map[string]interface{})["foo"]

	fmt.Println(mapRes)
	fmt.Println(vars)
}