package main

import (
	"log"

	"github.com/gdamore/tcell/v3"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	state := "mainMenu"
	name := "user"

	for {
		switch state {
		case "mainMenu":
			mainMenu(s, &state, name)
		case "settings":
			settings(s, &state, &name)
		case "color":
			color(s, &state)
		case "char":
			char(s, &state)
		}
	}
}
