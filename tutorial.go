package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

var products = []string{"Onions", "Carrots", "Bread", "Chicken thighs", "Ribeye steak", "Salmon"}

func initialModel() model {
	return model{
		choices:  products,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
} 

func (m model) View() string {
	s := "Select the products you'd like to buy\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		_, ok := m.selected[i]
		if ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	_, err := p.Run()
	if err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
}
