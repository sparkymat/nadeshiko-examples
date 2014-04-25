package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type TextBuffer struct {
	storage []byte
}

type TextEditorActivity struct {
	buffer TextBuffer
	x, y   int
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

func moveCursorToLine(connection *nadeshiko.Connection, line_number int) {
	line_id := fmt.Sprintf("line%d", line_number)
	cursor_div := "<span id='cursor'>|</span>"
	connection.JQuery("#cursor").Remove()
	connection.JQuery("#" + line_id).Append(cursor_div)
}

func addNewLine(connection *nadeshiko.Connection, line_number int) {
	line_id := fmt.Sprintf("line%d", line_number)
	line_div := fmt.Sprintf("<div id='%s'></div>", line_id)
	connection.JQuery("body").Append(line_div)
}

const UP_KEY = 38
const DOWN_KEY = 40

func (t TextEditorActivity) Start(connection *nadeshiko.Connection) {

	y := 0

	addNewLine(connection, y)
	moveCursorToLine(connection, y)

	go func() {
		blinkCursor(connection)
	}()

	connection.JQuery("body").Keydown(func(key int) {
		onKeyDown(connection, key)
	})

	connection.JQuery("body").Keypress(func(key int) {

		span_to_add := fmt.Sprintf("<span>%s</span>", string(key))
		connection.JQuery("#cursor").Before(span_to_add)
	})
}

func onKeyDown(connection *nadeshiko.Connection, key int) {
	log.Printf("key: %d \n", key)

	if key == 13 {
		y = y + 1
		addNewLine(connection, y)
		moveCursorToLine(connection, y)
	} else if key == 8 {
		connection.JQuery("#cursor").PrevRemove()
	} else if key == UP_KEY {
		y = y - 1
		moveCursorToLine(connection, y)
	} else if key == DOWN_KEY {
		y = y + 1
		moveCursorToLine(connection, y)
	}
}

func main() {
	nadeshiko.Start(TextEditorActivity{})
}
