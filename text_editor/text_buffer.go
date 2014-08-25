package main

type TextBuffer struct {
	data []string
}

var textBuffer = TextBuffer{}

func (buffer *TextBuffer) InsertChar(x, y, char int) {

	for len(buffer.data) <= y {
		buffer.data = append(buffer.data, "")
	}

	for len(buffer.data[y]) <= x {
		buffer.data[y] += " "
	}

	buffer.data[y] = buffer.data[y][0:x] + string(char) + buffer.data[y][x:len(buffer.data[y])]
}
