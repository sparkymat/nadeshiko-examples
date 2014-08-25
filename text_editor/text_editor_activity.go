package main

import (
	"fmt"
	"log"

	"github.com/kirillrdy/nadeshiko"
)

type TextEditorActivity struct {
	cursor Cursor
}

func (activity TextEditorActivity) Start(connection *nadeshiko.Connection) {

	//go func() {
	//	blinkCursor(connection)
	//}()

	for i := 0; i < 100; i++ {
		connection.JQuery("body").Append(fmt.Sprintf("<div id='%d'></div>", i))
	}

	event := nadeshiko.EventSubscription{EventName: "editor_update", Connection: connection}
	nadeshiko.SubscribeToEvent <- event

	connection.JQuery("body").Keydown(func(key int) {
		onKeyDown(connection, key)
	})

	connection.JQuery("body").Keypress(func(key int) {
		activity.onKeyPress(connection, key)
	})
}

const LEFT_KEY = 37
const UP_KEY = 38
const RIGHT_KEY = 39
const DOWN_KEY = 40
const ENTER_KEY = 13

func onKeyDown(connection *nadeshiko.Connection, key int) {
	log.Printf("key down: %d \n", key)

	//if key == ENTER_KEY {
	//	y = y + 1
	//	addNewLine(connection, y)
	//	moveCursorToLine(connection, y)
	//} else if key == 8 {
	//	connection.JQuery("#cursor").PrevRemove()
	//} else if key == UP_KEY {
	//	y = y - 1
	//	moveCursorToLine(connection, y)
	//} else if key == DOWN_KEY {
	//	y = y + 1
	//	moveCursorToLine(connection, y)
	//}
}

func (activity *TextEditorActivity) onKeyPress(old_connection *nadeshiko.Connection, key int) {
	textBuffer.InsertChar(activity.cursor.x, activity.cursor.y, key)
	activity.cursor.x += 1

	nadeshiko.TriggerEvent("editor_update", func(connection *nadeshiko.Connection) {
		connection.JQuery(fmt.Sprintf("#%d", activity.cursor.y)).SetText(textBuffer.data[activity.cursor.y])
	})

	log.Printf("key press: %s \n", string(key))
}
