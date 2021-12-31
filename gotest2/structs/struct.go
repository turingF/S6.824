package main

import "fmt"

type user struct {
	name string
	age  int
	id   int
}


func (u user) printUsr() {
	fmt.Printf("%s : %d", u.name,u.age)
}

func main() {
	bob := user{
		name: "xy",
		age:  18,
		id:   1001,
	}

	bob.printUsr()

}
