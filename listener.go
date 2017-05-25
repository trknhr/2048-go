package main

type Message struct {
	data string
}

var listeners = make(map[string]func(*Message))

func add(event string, listener func(*Message)) {
	listeners[event] = listener
}

func dispatch(event string, m *Message) {
	listeners[event](m)
}
