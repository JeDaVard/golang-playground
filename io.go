package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type io string

var ioOperation io = ""

func (i io) save(filename string, data deck) error {
	return ioutil.WriteFile("card_files/"+filename, []byte(strings.Join([]string(data), ",\n")), 0666)
}

func (i io) load(filename string) []string {
	bs, err := ioutil.ReadFile("card_files/" + filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",\n")
}
func (i io) print(deck deck) {
	for i, card := range deck {
		fmt.Println(i+1, card)
	}
}
