package cli

import (
	"goprojgen/internal/projgenerator"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectNameView struct {
	textinput.Model
}

func NewProjectNameView() *ProjectNameView {
	ti := textinput.New()
	ti.Placeholder = "Project Name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &ProjectNameView{Model: ti}
}
func (v *ProjectNameView) Update(msg tea.Msg, m *Model) tea.Cmd {
	log.Println("ProjectNameView.Update() > msg:", msg)
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.CurrentViewIdx++
		}
	}
	v.Model, cmd = v.Model.Update(msg)
	return cmd
}

func (v *ProjectNameView) View() string {
	return "Enter Project Name:\n" + v.Model.View()
}

func (v *ProjectNameView) ToAppConfig() projgenerator.AppConfig {
	return projgenerator.AppConfig{
		ProjectName: v.Model.Value(),
	}
}
