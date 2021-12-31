package structs

import "fmt"

func main(){
	//m1 := make(map[string]int)
	m2 := map[string]int {"hello":1,"world":2}

	for k,v := range m2 {
		fmt.Printf("k:%s , v:%d\n",k,v)
	}
}