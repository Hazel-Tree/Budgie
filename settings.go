package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v3"
)

func changeName(s tcell.Screen, name *string) {
	w, _ := s.Size()
	fmt.Printf("\033[4;7H\033[5m\033[7m_%s\033[m", strings.Repeat(" ", w))
	fmt.Printf("\033[5;7HCurrent name: %s | Must be between 1 and 30 characters | Leave blank to dismiss%s\033[m", *name, strings.Repeat(" ", w))
	tempName := ""
changedName:
	for {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				if tempName != "" {
					*name = tempName
				}
				fmt.Printf("\033[4;7H%s%s", *name, strings.Repeat(" ", w))
				fmt.Printf("\033[5;7H%s", strings.Repeat(" ", w))
				break changedName
			}
			if ev.Key() == tcell.KeyBackspace && len(tempName) > 0 {
				tempName = tempName[:len(tempName)-1]
			}
			if len(tempName) < 30 {
				tempName += ev.Str()
			}
			fmt.Printf("\033[4;7H\033[7m%s\033[5m_ \033[m", tempName)
		}
	}
}

func settings(s tcell.Screen, state *string, name *string) {
	clear()
	fmt.Printf("Settings\n\r%sxit\n\n\r", ulb("E"))
	fmt.Printf("%same: %s", ulb("N"), *name)

	for *state == "settings" {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Str() == "e" {
				*state = "mainMenu"
			} else if ev.Str() == "n" {
				changeName(s, name)
			}
		}
	}
}
