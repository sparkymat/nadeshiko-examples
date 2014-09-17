package main

import "github.com/kirillrdy/nadeshiko"

type HelloWorldActivity struct {
	Greeting string
}

func (a HelloWorldActivity) Start(connection *nadeshiko.Connection) {
	connection.JQuery("body").Append(a.Greeting)
}

func main() {
	nadeshiko.StartActivity(HelloWorldActivity{Greeting: "Hello World"})
}
