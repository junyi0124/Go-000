package main

import (
	"fmt"
)

var msg = "Hello World!"

func NewMessage() Message {
	return Message(msg)
}

// 要说的内容的抽象
type Message string

func NewPeople(m Message) People {
	return People{name: "小明", message: m}
}

// 小明这个人的抽象
type People struct {
	name    string
	message Message
}

// 小明这个人会说话
func (p People) SayHello() string {
	msg := fmt.Sprintf("%s 对世界说:%s\n", p.name, p.message)
	return msg
}

func NewEvent(p People) Event {
	return Event{people: p}
}

// 小明去说话这个行为抽象为一个事件
type Event struct {
	people People
}

func (e Event) start() {
	msg := e.people.SayHello()
	fmt.Println(msg)
}

func main() {
	e := InitializeEvent()
	e.start()
}
