package main

import "./nadeshiko"
import "fmt"
import "time"
import "strconv"

type TextEditorActivity struct {
}

func (t TextEditorActivity) Start(connection *nadeshiko.Connection) {

	x := 0
	y := 0

	line_id := fmt.Sprintf("line%d", y)
	line_div := fmt.Sprintf("<div id='%s'></div>", line_id)
	jquery_line_id := "#" + line_id
	connection.JQuery("body").Append(line_div)

	cursor_div := "<span id='cursor'>|</span>"
	connection.JQuery(jquery_line_id).Append(cursor_div)

	go func() {
		c := time.Tick(500 * time.Millisecond)
		show_cursor := false
		for _ = range c {
			show_cursor = !show_cursor
			if show_cursor {
				connection.JQuery("#cursor").SetText("|")
			} else {
				connection.JQuery("#cursor").SetText("")
			}
		}
	}()

	connection.JQuery("body").Keydown(func(key int) {
		if strconv.IsPrint(rune(key)) {
			x = x + 1
			span_to_add := fmt.Sprintf("<span>%s</span>", string(key))
			connection.JQuery("#cursor").Before(span_to_add)
		}
		if key == 13 {
			y = y + 1
			line_id := fmt.Sprintf("line%d", y)
			line_div := fmt.Sprintf("<div id='%s'></div>", line_id)
			jquery_line_id := "#" + line_id
			connection.JQuery("body").Append(line_div)
			connection.JQuery("#cursor").Remove()
			connection.JQuery(jquery_line_id).Append(cursor_div)
			x = 0
		} else if key == 8 {
			connection.JQuery("#cursor").PrevRemove()
		} else {

		}
	})
}

func main() {
	nadeshiko.Start(TextEditorActivity{})
}
