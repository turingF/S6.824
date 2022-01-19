package IO

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
  @Description: 单元测试文件名必须以_test为结尾
*/

const checkMark = "\u2713"
const ballotX = "\u2717"

// 函数名必须以Test开头，接受一个(t *testing.T)的参数，返回nil
func TestHello(t *testing.T) {
	t.Logf("Hello world %v %v",checkMark,ballotX)
}

func TestCurl(t *testing.T) {
	uri := "http://www.baidu.com"

	resp,err := http.Get(uri)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("get uri success\n")
	println(resp)
}


func mockServer() *httptest.Server {

	feed := `{"foo":"bar"}`
	f := func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w,feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestTables(t *testing.T){
	tables := []struct{
		uri string
		code int
	}{
		{"http://www.baidu.com",http.StatusOK},
		{"http://www.baidu.com/a.pdf",http.StatusNotFound},
	}

	for _,item := range tables {
		t.Logf("Checking uri : [%s] ---> code : [%d]",item.uri,item.code)
		resp,err := http.Get(item.uri)
		if err != nil {
			t.Fatal("Get uri false")
		}
		defer resp.Body.Close()

		if resp.StatusCode == item.code {
			t.Logf("Get Correct Code from [%s]  %v",item.uri,checkMark)
		}else{
			t.Errorf("Get Error Code from [%s]  %v",item.uri,ballotX)
		}

	}
}