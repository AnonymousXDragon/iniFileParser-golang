package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"iniParser/tools"
)

// var source = `

// [general]
// appname = configparser
// version = 0.1

// [author]
// name = rohitbr
// email = rohitbr@gmail.com

// `

func ReadSource() (string, error) {
	var output string
	file, err := os.Open(os.Args[1])
	if err != nil {
		return "", errors.New("expecting a file to parse ")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}

	return output, nil
}

func main() {
	fmt.Println("welcome to ini parser ")
	source, err := ReadSource()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	ini := tools.Parser(source)
	fmt.Println("testing..1", ini.GetProperty("general", "appname"))
	fmt.Println("testing..2", ini.GetProperty("author", "name"))
	fmt.Println("testing..3", ini.GetProperty("author", "email"))
	fmt.Println(ini)
}
