package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func mainMenu(name string, s tcell.Screen) {
	s.Clear()
	s.PutStr(0, 0, "Hello "+name+" ! >> Char  Color  Settings  Exit")
	s.Put(12+len(name), 0, "C", tcell.Style.Underline(tcell.StyleDefault, true))
	s.Put(19+len(name), 0, "o", tcell.Style.Underline(tcell.StyleDefault, true))
	s.Put(25+len(name), 0, "S", tcell.Style.Underline(tcell.StyleDefault, true))
	s.Put(35+len(name), 0, "E", tcell.Style.Underline(tcell.StyleDefault, true))
}

func mainMenuActions(ev *tcell.EventKey, state *string) {
	if ev.Str() == "c" {
		*state = "char"
	} else if ev.Str() == "o" {
		*state = "color"
	} else if ev.Str() == "s" {
		*state = "settings"
	} else if ev.Str() == "e" {
		fmt.Printf("\033[0;0H\033[0J")
		os.Exit(0)
	}
}

func charMenu(s tcell.Screen) {
	s.Clear()
	s.PutStr(0, 0, "Characters >> Exit")
	s.Put(14, 0, "E", tcell.Style.Underline(tcell.StyleDefault, true))
}

func charMenuActions(ev *tcell.EventKey, state *string) {
	if ev.Str() == "e" {
		*state = "mainMenu"
	}
}

func colorMenu(s tcell.Screen) {
	s.Clear()
	s.PutStr(0, 0, "Colors >> Exit")
	s.Put(10, 0, "E", tcell.Style.Underline(tcell.StyleDefault, true))
}

func colorMenuActions(ev *tcell.EventKey, state *string) {
	if ev.Str() == "e" {
		*state = "mainMenu"
	}
}

func settingsMenu(name string, s tcell.Screen) {
	s.Clear()
	s.PutStr(0, 0, "Settings >> Exit")
	s.Put(12, 0, "E", tcell.Style.Underline(tcell.StyleDefault, true))
	s.PutStr(0, 2, "Name: "+name)
	s.Put(0, 2, "N", tcell.Style.Underline(tcell.StyleDefault, true))
}

func settingsMenuActions(s tcell.Screen, ev *tcell.EventKey, state *string, name *string) {
	if ev.Str() == "e" {
		*state = "mainMenu"
	} else if ev.Str() == "n" {
		changeName(s, name)
	}
}

func changeName(s tcell.Screen, name *string) {
	tempName := ""
	exit := false
	w, _ := s.Size()
	s.PutStrStyled(6, 2, "_"+strings.Repeat(" ", w-6), tcell.StyleDefault.Background(color.NewRGBColor(200, 200, 200)).Foreground(color.NewRGBColor(25, 25, 25)))
	s.PutStr(6, 3, "> Current name: "+*name)
	s.PutStr(6, 4, "> between 1 and 30 characters")
	s.PutStr(6, 5, "> Leave blank to dismiss")
	s.Show()
	for {
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter && len(tempName) <= 30 {
				exit = true
			}
			if ev.Key() == tcell.KeyBackspace && len(tempName) != 0 {
				tempName = tempName[:len(tempName)-1]
			}
			if len(tempName) < 30 {
				tempName = tempName + ev.Str()
			}
			s.PutStrStyled(6, 2, tempName+"_"+strings.Repeat(" ", w-5), tcell.StyleDefault.Background(color.NewRGBColor(200, 200, 200)).Foreground(color.NewRGBColor(25, 25, 25)))
			s.Show()
		}
		if exit {
			if tempName != "" {
				*name = tempName
			}
			break
		}
	}
}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// set default text and style
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
	s.SetStyle(defStyle)

	state := "mainMenu"
	name := "user"
	mainMenu(name, s)

	for {
		switch state {
		case "mainMenu":
			mainMenu(name, s)
		case "char":
			charMenu(s)
		case "color":
			colorMenu(s)
		case "settings":
			settingsMenu(name, s)
		}

		// update screen
		s.Show()

		// poll event
		ev := <-s.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch state {
			case "mainMenu":
				mainMenuActions(ev, &state)
			case "char":
				charMenuActions(ev, &state)
			case "color":
				colorMenuActions(ev, &state)
			case "settings":
				settingsMenuActions(s, ev, &state, &name)
			}
		}
	}
}
