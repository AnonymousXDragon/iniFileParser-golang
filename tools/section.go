package tools

import "fmt"

type Section struct {
	properties map[string]string
}

func NewSection() *Section {
	return &Section{
		properties: make(map[string]string),
	}
}

func (s *Section) SetProperty(key string, value string) {
	s.properties[key] = value
}

func (s *Section) ToString() string {
	return fmt.Sprintf("<Section %v >", s.properties)
}
