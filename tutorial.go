package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type product struct {
	name  string
	price float64
}

var products = []product{
	{"Onions", 1.99},
	{"Carrots", 1.99},
	{"Bread", 3.99},
	{"Chicken", 8.99},
	{"Ribeye steaks", 16.99},
	{"Salmon", 12.99},
}

type model struct {
	cursor        int
	selected      map[int]bool
	orderComplete bool
}

func initialModel() model {
	return model{
		cursor:   0,
		selected: make(map[int]bool),
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
			if m.cursor < len(products)-1 {
				m.cursor++
			}
		case " ":
			m.selected[m.cursor] = !m.selected[m.cursor]
		case "enter":
			m.orderComplete = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select the products you'd like to buy (space to select)\n\n"

	s += fmt.Sprintf("%-5s %-22s | %8s\n", "", "Product", "Price")
	s += fmt.Sprintf("%-5s %-22s | %8s\n", "", strings.Repeat("-", 22), strings.Repeat("-", 8))

	for i, product := range products {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected[i] {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%-1s] %-22s | $%7.2f\n", cursor, checked, product.name, product.price)
	}

	s += "\nEnter to complete order, q to quit\n"

	if m.orderComplete && len(m.selected) > 0 {
		var totalPrice float64
		var selectedItems []string

		for i, selected := range m.selected {
			if selected {
				selectedItems = append(selectedItems, products[i].name)
				totalPrice += products[i].price
			}
		}

		s += fmt.Sprintf("\nBought: %s\nTotal price: $%.2f\n", strings.Join(selectedItems, ", "), totalPrice)
	}

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
