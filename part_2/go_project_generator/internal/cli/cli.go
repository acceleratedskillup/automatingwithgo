package cli

// An example demonstrating an application with multiple views.
//
// Note that this example was produced before the Bubbles progress component
// was available (github.com/charmbracelet/bubbles/progress) and thus, we're
// implementing a progress bar from scratch here.

import (
	"fmt"
	"goprojgen/internal/projgenerator"
	"goprojgen/internal/termstyle"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const padding = 2

const (
	ProjectNamePage int = iota
	ConfigTypePage
	RouterTypePage
)

type View interface {
	Update(msg tea.Msg, m *Model) tea.Cmd
	View() string
}
type AppConfigurable interface {
	ToAppConfig() projgenerator.AppConfig
}
type Model struct {
	Views           []View
	CurrentViewIdx  int
	CreateProject   bool
	ProgressPercent float64
	Progress        progress.Model
	Quitting        bool
	ProgressChannel chan tea.Msg
}
type createProjectMsg struct{}

func Start() {

	initialModel := Model{
		Views: []View{
			NewProjectNameView(),
			NewMultiChoiceView("Read configuration settings from:", []string{"Command-line flags", "Environment variables"}, configTypeConfigurator),
			NewMultiChoiceView("Pick your preferred router:", []string{"Gorilla Mux", "HttpRouter"}, routerTypeConfigurator),
		},
		CurrentViewIdx:  0,
		Progress:        progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
		ProgressChannel: make(chan tea.Msg),
	}

	if _, err := tea.NewProgram(initialModel).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start program: %v\n", err)

		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Main update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("msg:", msg)

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEsc || msg.Type == tea.KeyCtrlC {
			close(m.ProgressChannel)
			return m, tea.Quit
		}
	case createProjectMsg:
		log.Println("creating project")
		appConfig := populateAppConfig(m.Views)
		go projgenerator.GenerateProject(appConfig, m.ProgressChannel)
		return m, listenToProgress(m.ProgressChannel)
	case projgenerator.ProjectCreationInProgressMsg:
		m.ProgressPercent = float64(msg)
		return m, listenToProgress(m.ProgressChannel)
	case projgenerator.ProjectCreatedMsg:
		m.ProgressPercent = 100
		return m, tea.Quit
	}
	log.Println("m.CurrentViewIdx < len(m.Views):", m.CurrentViewIdx < len(m.Views))
	if m.CurrentViewIdx < len(m.Views) {
		cmd = m.Views[m.CurrentViewIdx].Update(msg, &m)
	}

	return m, cmd
}

// The main view, which just calls the appropriate sub-view
func (m Model) View() string {
	log.Println("cli View() > m.CreateProject:", m.CreateProject)
	if m.Quitting {
		return "\n  See you later!\n\n"
	} else if m.CreateProject {
		pad := strings.Repeat(" ", padding)
		return "\n" +
			pad + m.Progress.ViewAs(m.ProgressPercent) + "\n\n" +
			pad + termstyle.HelpStyle("Press any key to quit")
	}
	log.Println("cli View() m.CurrentViewIdx:", m.CurrentViewIdx)
	log.Println("m.Views[m.CurrentViewIdx]:", m.Views[m.CurrentViewIdx])
	return m.Views[m.CurrentViewIdx].View()
}

func populateAppConfig(views []View) projgenerator.AppConfig {
	appConfig := projgenerator.AppConfig{}

	for _, view := range views {
		if appConfigurable, ok := view.(AppConfigurable); ok {
			// Assuming a function that merges two AppConfigs
			appConfig = mergeAppConfigs(appConfig, appConfigurable.ToAppConfig())
		}
	}

	return appConfig
}

func mergeAppConfigs(dest, source projgenerator.AppConfig) projgenerator.AppConfig {
	// Simplified example: overwrite fields in 'a' with non-zero value fields from 'b'
	if source.ProjectName != "" {
		dest.ProjectName = source.ProjectName
	}

	if source.UseRouter {
		dest.UseRouter = source.UseRouter
		dest.RouterType = source.RouterType
		dest.RouterImportPath = source.RouterImportPath
		dest.RouterConstructor = source.RouterConstructor
	}
	// Add more fields as necessary
	return dest
}

// This function starts listening to the progressCh channel and returns messages back to the Bubble Tea program.
func listenToProgress(progressChannel <-chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-progressChannel
		if !ok {
			return nil // Channel closed, no more messages
		}
		return msg
	}
}
