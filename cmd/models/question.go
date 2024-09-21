package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func Question(items []string, title string, message string, count, filter bool) string {

	data := []list.Item{}
	for _, db := range items {
		option := Option(db)
		data = append(data, option)
	}

	outcome := func(choice string, quitTextStyle lipgloss.Style) string {
		return quitTextStyle.Render(fmt.Sprintf("%s: %s", message, choice))
	}

	return Model{}.New(data, title, count, filter, outcome)
}
