package main

import (
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type Cursor struct {
	x, y int
}

func blinkCursor(connection *nadeshiko.Connection) {
	tick := time.Tick(500 * time.Millisecond)
	show_cursor := false
	for _ = range tick {
		show_cursor = !show_cursor
		if show_cursor {
			connection.JQuery("#cursor").SetText("|")
		} else {
			connection.JQuery("#cursor").SetText("")
		}
	}
}
