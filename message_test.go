package main

import "testing"

func TestMessage(t *testing.T) {

	m := Message{data: []string{"Hello", "world"}}
	b := m.String()

	if b != "data: Hello\ndata: world\n\n\n" {
		t.Error("Didn't get what I expected")
	}
}

func TestSetEvent(t *testing.T) {
	m := Message{data: []string{"one"}}
	m.SetEvent("hello")

	if m.String() != "event: hello\ndata: one\n\n\n" {
		t.Error("Should have an event set.")
	}
}

func TestMessageSetDAta(t *testing.T) {
	m := new(Message)
	m.SetData("one", "two")

	if m.String() != "data: one\ndata: two\n\n\n" {
		t.Error("Should then have one and two.")
	}
}

func TestMessageAppend(t *testing.T) {

	m := Message{data: []string{"one"}}
	if m.String() != "data: one\n\n\n" {
		t.Error("Should start with one.")
	}

	m.AppendData("two")
	if m.String() != "data: one\ndata: two\n\n\n" {
		t.Error("Should then have one and two.")
	}
}
