package main

import (
	"fmt"
)

var graph = map[string][]string{
	"you":    {"tobi", "uzor", "joseph"},
	"tobi":   {"elijah", "enoch"},
	"uzor":   {"tireni"},
	"joseph": {"dalevz", "kelly"},
}

func main() {
	if Search("you") {
		fmt.Println("Done!")
	}
}

func Search(name string) bool {

	queue := graph[name]
	searched := map[string]struct{}{}

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]
		if _, ok := searched[person]; ok { //stores two values in two variables a boolean and the value mapped to the key "person", if person exists- ok is true (a boolean), we don't care about thee value hence the underscore (_)
			continue //If we’ve already seen this person before, skip them immediately and move on
		}
		if person == "enoch" {
			fmt.Printf("The king has been located! His name is %s\n", person)
			return true
		} else {
			queue = append(queue, graph[person]...)
			searched[person] = struct{}{}
		}

	}

	return false
}
