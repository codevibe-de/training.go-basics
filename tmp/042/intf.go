package main

import "fmt"

type Session interface {
	Id() string
	GetAttribute(key string) (string, bool)
}

type session struct {
	id    string
	attrs map[string]string
}

func (s *session) Id() string {
	return s.id
}

func (s *session) GetAttribute(key string) (string, bool) {
	v, exists := s.attrs[key]
	return v, exists
}

func main() {
	var s Session = NewSession()
	fmt.Println(s.Id())
}

func NewSession() Session {
	s := session{
		id:    "abcd-1234",
		attrs: make(map[string]string),
	}
	return &s
}
