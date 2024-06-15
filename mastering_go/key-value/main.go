package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var Data = make(map[string]myElement)

func ADD(key string, n myElement) bool {
	if key == "" {
		return false
	}

	if LOOKUP(key) == nil {
		Data[key] = n
		return true
	}
	return false
}

func DELETE(key string) bool {
	if LOOKUP(key) != nil {
		delete(Data, key)
		return true
	}
	return false
}

func LOOKUP(key string) *myElement {
	n, ok := Data[key]
	if ok {
		return &n
	}
	return nil
}

func CHANGE(key string, n myElement) bool {
	if LOOKUP(key) != nil {
		Data[key] = n
		return true
	}
	return false
}

func PRINT() {
	for k, v := range Data {
		fmt.Println("key:", k, "value:", v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed.")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed.")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n == nil {
				fmt.Println("The element is not found.")
			} else {
				fmt.Println("Name:", n.Name, "Surname:", n.Surname, "Id:", n.Id)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Change operation failed.")
			}
		default:
			fmt.Println("Unknown command.")
		}
	}
}
