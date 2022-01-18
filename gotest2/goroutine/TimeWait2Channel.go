package test

import "time"

/*
  @Description: test time
*/

func main() {

	// send a time object to channel
	var timeout <- chan time.Time
	timeout = time.After(10 * time.Second)

	select{
	case <- timeout:
		println("get from channel")
	}

}
