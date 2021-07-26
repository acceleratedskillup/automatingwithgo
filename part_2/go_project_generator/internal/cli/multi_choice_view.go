package cli

import (
	"goprojgen/internal/projgenerator"
	"goprojgen/internal/termstyle"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type MultiChoiceView struct {
	Prompt          string
	Options         []string
	Selected        int
	AppConfigurator func(int) projgenerator.AppConfig
}

func NewMultiChoiceView(prompt string, options []string, appConfigurator func(int) projgenerator.AppConfig) *MultiChoiceView {
	return &MultiChoiceView{
		Prompt:          prompt,
		Options:         options,
		Selected:        0,
		AppConfigurator: appConfigurator,
	}
}

func (v *MultiChoiceView) View() string {
	log.Println("in MultiChoiceView.View()")
	var builder strings.Builder

	// Building the prompt
	builder.WriteString(v.Prompt + "\n\n")

	// Iterating over options to create the selection list
	for index, option := range v.Options {
		// Use the Checkbox function for each option
		checkbox := Checkbox(option, index == v.Selected)
		builder.WriteString(checkbox + "\n")
	}

	// Adding instructions
	instructions := termstyle.Subtle("enter: choose") + termstyle.Dot + termstyle.Subtle("esc or ctrl+c: quit")
	builder.WriteString("\n" + instructions)

	return builder.String()
}

func (v *MultiChoiceView) Update(msg tea.Msg, m *Model) tea.Cmd {

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			v.Selected++
			if v.Selected >= len(v.Options) {
				v.Selected = len(v.Options) - 1
			}
		case tea.KeyUp:
			v.Selected--
			if v.Selected < 0 {
				v.Selected = 0
			}
		case tea.KeyEnter:
			//next step is to create project
			if m.CurrentViewIdx == len(m.Views)-1 {
				m.CreateProject = true
				log.Println("returning createProjectMsg")
				cmd = createProject()
			} else { //move to next page
				m.CurrentViewIdx++
			}
		}
	}
	return cmd
}

func (v *MultiChoiceView) ToAppConfig() projgenerator.AppConfig {
	if v.AppConfigurator != nil {
		return v.AppConfigurator(v.Selected)
	}
	return projgenerator.AppConfig{}
}

func createProject() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return createProjectMsg{}
	})
}
