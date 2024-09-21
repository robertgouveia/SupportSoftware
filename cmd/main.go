package main

import (
	"bufio"
	"database/cmd/models"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

var input string

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Example"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, cmd
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	input = m.textInput.Value()
	return fmt.Sprintf("Question: %s", m.textInput.View())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	os.Setenv("SERVER", models.Question(models.GetConnections(), "Select a Server", "You chose", true, true))

	conn, db := models.Connection{}.Open()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	os.Setenv("DATABASE", models.Question(conn.Databases(db), "Select a Database", "You chose", true, true))
	conn, db = models.Connection{}.Open()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	imports := models.ImportList{}.Get()
	queryName := models.Question(models.ImportNames(imports), "Select an Export", "You chose", false, false)

	var sqlFile string
	for _, i := range imports {
		if queryName == models.ImportToName(i.Name) {
			sqlFile = i.Path
		}
	}

	if sqlFile == "" {
		log.Fatal("No SQL File Selected")
	}

	file, err := os.Open("./imports/" + sqlFile)
	if err != nil {
		log.Fatalf("Unable to read SQL File: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatalf("Unable to read SQL File: %v", err)
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "DECLARE") {
			parts := strings.Split(line, "=")
			if len(parts) > 1 {
				value := strings.ReplaceAll(strings.TrimSpace(parts[1]), ";", "")
				fmt.Println(strings.ReplaceAll(line, value, input))
			}
		}
	}

	models.Execute(sqlFile, db, queryName)
}
