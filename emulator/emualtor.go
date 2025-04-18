package emulator

import (
	"fmt"
	"io"
	"strings"

	"github.com/aitsuki/avds/cmd"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// var (
// 	appStyle   = lipgloss.NewStyle().Padding(1, 2)
// 	titleStyle = lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#FFFDF5")).
// 			Background(lipgloss.Color("#25A065")).
// 			Padding(0, 1)
// 	statusMessageStyle = lipgloss.NewStyle().
// 				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
// 				Render
// )

const listHeight = 14
const defaultWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
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

type model struct {
	list     list.Model
	selected string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selected = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func initModel(avds []string) model {
	items := make([]list.Item, len(avds))
	for i, avd := range avds {
		items[i] = item(avd)
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.SetShowTitle(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{list: l}
}

func Run() error {
	emulatorPath, err := cmd.GetEmulatorPath()
	if err != nil {
		return fmt.Errorf("get emulator path: %w", err)
	}

	avds, err := cmd.ListAvds(emulatorPath)
	if err != nil {
		return fmt.Errorf("list avds: %w", err)
	}

	p := tea.NewProgram(initModel(avds))
	m, err := p.Run()
	if err != nil {
		return fmt.Errorf("run emualtor program: %w", err)
	}

	finalModel := m.(model)
	if finalModel.selected != "" {
		avdName := finalModel.selected
		err := cmd.StartAvd(emulatorPath, avdName)
		if err != nil {
			return fmt.Errorf("start avd: %w", err)
		}
		fmt.Printf("Starting %s, please wait...\n", avdName)
	}
	return nil
}
