package models

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type DBList struct {
	List     list.Model
	choice   string
	quitting bool
}

func (d DBList) Init() tea.Cmd {
	return nil
}

func (d DBList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		d.List.SetWidth(msg.Width)
		return d, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			d.quitting = true
			return d, tea.Quit
		case "enter":
			i, ok := d.List.SelectedItem().(Option)
			if ok {
				d.choice = string(i)
			}
			return d, tea.Quit
		}
	}
	var cmd tea.Cmd
	d.List, cmd = d.List.Update(msg)
	return d, cmd
}

func (d DBList) View() string {
	if d.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("You chose: %s", d.choice))
	}
	if d.quitting {
		return quitTextStyle.Render("Session Exit")
	}

	return "\n" + d.List.View()
}

func (d DBList) New(items []list.Item, title string) {
	const defaultWidth = 20
	const listHeight = 14

	l := list.New(items, DatabaseDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	d.List = l

	if _, err := tea.NewProgram(d).Run(); err != nil {
		fmt.Println("Error running the program: ", err)
		os.Exit(1)
	}
}
