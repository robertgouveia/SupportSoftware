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
	defaultWidth    = 20
	listHeight      = 14
	outcome         func(string, lipgloss.Style) string
	choice          string
)

type Model struct {
	List     list.Model
	choice   string
	quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.List.SelectedItem().(Option)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.choice != "" {
		choice = m.choice
		return outcome(m.choice, quitTextStyle)
	}
	if m.quitting {
		os.Exit(0)
		return quitTextStyle.Render("Session Exit")
	}

	return "\n" + m.List.View()
}

func (m Model) New(items []list.Item, title string, showCount bool, filterAllowed bool, outcomeInput func(string, lipgloss.Style) string) string {
	l := list.New(items, Delegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(showCount)
	l.SetFilteringEnabled(filterAllowed)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	m.List = l

	outcome = outcomeInput

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running the program: ", err)
		os.Exit(1)
	}

	return choice
}
