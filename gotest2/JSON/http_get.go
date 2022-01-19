package JSON

import (
	"encoding/json"
	"fmt"
)

/*
 @Description:
*/

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"UnescapedURL"`
		URL                string `json:"URL"`
		VisibleURL         string `json:"VisibleURL"`
		CacheURL           string `json:"CacheURL"`
		Title              string `json:"Title"`
		TitleNoFormatting  string `json:"TitleNoFormatting"`
		Content            string `json:"Content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)


func main() {
	//uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	str := `{
   "responseData": {
       "results": [
           {
               "GsearchResultClass": "GwebSearch",
               "unescapedUrl": "https://www.reddit.com/r/golang",
               "url": "https://www.reddit.com/r/golang",
               "visibleUrl": "www.reddit.com",
               "cacheUrl": "http://www.google.com/search?q=cache:W...",
               "title": "r/<b>Golang</b> - Reddit",
               "titleNoFormatting": "r/Golang - Reddit",
               "content": "First Open Source <b>Golang\u0000.."
           },
           {
               "GsearchResultClass": "GwebSearch",
               "unescapedUrl": "http://tour.golang.org/",
               "url": "http://tour.golang.org/",
               "visibleUrl": "tour.golang.org",
               "cacheUrl": "http://www.google.com/search?q=cache:O...",
               "title": "A Tour of Go",
               "titleNoFormatting": "A Tour of Go",
               "content": "Welcome to a tour of the Go programming ..."
           }
       ]
   }
}`

	var gr gResponse
	err := json.Unmarshal([]byte(str),&gr)
	if err != nil {
		println("ERROR",err)
		return
	}

	fmt.Println(gr)
}
