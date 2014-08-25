package main

import "github.com/kirillrdy/nadeshiko"

type User struct {
	connection *nadeshiko.Connection
	cursor     Cursor
}
