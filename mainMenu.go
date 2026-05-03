package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v3"
)

func mainMenu(s tcell.Screen, state *string, name string) {
	clear()
	fmt.Printf("Hello %s !!!!\n\r%shar  C%slor  %settings  %sxit\n\n\rFavorites", name, ulb("C"), ulb("o"), ulb("S"), ulb("E"))

	for *state == "mainMenu" {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Str() == "c" {
				*state = "char"
			} else if ev.Str() == "o" {
				*state = "color"
			} else if ev.Str() == "s" {
				*state = "settings"
			} else if ev.Str() == "e" {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
