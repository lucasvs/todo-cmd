package main

import (
	"fmt"
	"log"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

type Task struct {	
	Description string	
}

var tasks = []Task{
	{"foo"},
	{"bar"},
}

func main() {
	idx, err := fuzzyfinder.FindMulti(
		tasks,
		func(i int) string {
			return tasks[i].Description
		},
		)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}