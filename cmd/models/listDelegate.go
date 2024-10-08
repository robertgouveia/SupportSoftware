package models

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Delegate struct{}

var (
	itemStyle         = lipgloss.NewStyle().MarginLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func (d Delegate) Height() int                             { return 1 }
func (d Delegate) Spacing() int                            { return 0 }
func (d Delegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d Delegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Option)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
