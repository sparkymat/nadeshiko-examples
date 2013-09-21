package main

import "./nadeshiko"

type HelloWorldActivity struct {
	Greeting string
}

func (a HelloWorldActivity) Start(connection *nadeshiko.Connection) {
	connection.JQuery("body").Append(a.Greeting)
}

func main() {
	nadeshiko.Start(HelloWorldActivity{Greeting: "Hello World"})
}
