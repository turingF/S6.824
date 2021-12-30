package main

/*
  @Description:
*/

type notifier interface {
	notify()
}

type user struct {
	name string
	age int
}

func (u *user) notify(){
	println(u.name)
}

func sendNotification(n notifier){
	n.notify()
}

func main() {
	foo := user{"bar",24}
	sendNotification(&foo)
}