package main

import "github.com/kirillrdy/nadeshiko"

func handler(document nadeshiko.Document) {
	document.JQuery("body").Append("Hello World !!!")
}

func main() {
	nadeshiko.Nadeshiko("/", handler)
	nadeshiko.Start()
}
