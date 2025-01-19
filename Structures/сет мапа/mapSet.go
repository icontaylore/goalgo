package main

import "fmt"

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{Elements: make(map[string]struct{})}
}

func (s *Set) Add(e string) {
	s.Elements[e] = struct{}{}
}

func (s *Set) Contains(e string) bool {
	_, exists := s.Elements[e]
	return exists
}

func (s *Set) Delete(e string) {
	delete(s.Elements, e)
}

func (s *Set) List() []string {
	keys := make([]string, 0, len(s.Elements))

	for key := range s.Elements {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	m1 := NewSet()

	m1.Add("earth")
	m1.Add("earth2")
	m1.Add("earth3")

	arr := m1.List()
	fmt.Println(arr)
}
