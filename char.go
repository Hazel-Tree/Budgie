package main

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
)

func char(s tcell.Screen, state *string) {
	clear()
	fmt.Printf("Characters\n\r%sxit\n\n\r", ulb("E"))

	for *state == "char" {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Str() == "e" {
				*state = "mainMenu"
			}
		}
	}
}
