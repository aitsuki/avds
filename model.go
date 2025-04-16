package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	avds     []string
	cursor   int
	selected string
	offset   int
	maxLines int
}

func initialModel(emulatorPath string) model {
	avds, err := listAvds(emulatorPath)
	if err != nil {
		log.Fatalln(err)
	}
	return model{avds: avds, maxLines: 10}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.avds) - 1
				m.offset = max(0, m.cursor-m.maxLines+1)
			}
			if m.cursor < m.offset {
				m.offset = m.cursor
			}
		case "down", "j":
			if m.cursor < len(m.avds)-1 {
				m.cursor++
			} else {
				m.cursor = 0
				m.offset = 0
			}
			if m.cursor >= m.offset+m.maxLines {
				m.offset = m.cursor - m.maxLines + 1
			}
		case "enter":
			m.selected = m.avds[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if len(m.avds) == 0 {
		return "No AVDs available\n"
	}

	var lines []string
	start := m.offset
	end := min(m.offset+m.maxLines, len(m.avds))

	if m.offset > 0 {
		lines = append(lines, "  ...")
	}

	for i := start; i < end; i++ {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		lines = append(lines, fmt.Sprintf("%s %s", cursor, m.avds[i]))
	}

	if end < len(m.avds) {
		lines = append(lines, "  ...")
	}

	return strings.Join(lines, "\n") + "\n"
}
