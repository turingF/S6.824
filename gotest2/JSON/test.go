package JSON

import (
	"encoding/json"
	"fmt"
)

/*
  @Description:
*/
type (
	res struct {
		AData string `json:"a"`
	}

	resp struct {
		Datas []res `json:"data"`
	}
)

func main() {
	test_string := `{
	"data": [
		{
			"a": "hello"
		},
		{
			"a": "world"
		}
	]
}`
	var test_resp resp

	err := json.Unmarshal([]byte(test_string), &test_resp)
	if err != nil {
		println("ERROR", err)
		return
	}
	fmt.Println(test_resp)
}
