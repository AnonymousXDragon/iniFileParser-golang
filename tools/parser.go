package tools

import (
	"fmt"
	"strings"
)

const (
	CreateSection = iota
	CreateKeyValue
)

func ReadSection(ini *Ini, sectionName string) {
	ini.SetSection(sectionName, Section{
		properties: make(map[string]string),
	})
}

func ReadKV(ini *Ini, sectionName string, key string, value string) {
	ini.SetProperty(sectionName, key, value)
}

func RemoveWhiteSpace(arr []string, condition func(char string) bool) {
	for index := range arr {
		if condition(arr[index]) {
			fmt.Println("...", arr[index])
			arr = append(arr, arr[index:]...)
			index = 0
		}
	}
}

func Parser(source string) Ini {
	ini := NewIni()
	currentSectionName := " "
	lines := strings.Split(source, "\n")
	state := CreateSection

	for _, line := range lines {

		if strings.ReplaceAll(line, " ", "") == "" {
			continue
		}

		if strings.Contains(line, ";") || strings.Contains(line, "#") {
			panic("invalid characters found")
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			state = CreateSection
		}

		if state == CreateSection {
			currentSectionName = line[1 : len(line)-1]
			ReadSection(ini, currentSectionName)
			state = CreateKeyValue
			continue
		}

		if state == CreateKeyValue {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				ReadKV(ini, currentSectionName, strings.ReplaceAll(parts[0], " ", ""), strings.ReplaceAll(parts[1], " ", ""))
			}
		}
	}

	return *ini
}
