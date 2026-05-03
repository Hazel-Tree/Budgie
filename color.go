package main

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
)

func color(s tcell.Screen, state *string) {
	clear()
	fmt.Printf("Colors\n\r%sxit\n\n\r", ulb("E"))

	for *state == "color" {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Str() == "e" {
				*state = "mainMenu"
			}
		}
	}
}
