package termstyle

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

var (
	Term      = termenv.EnvColorProfile()
	Subtle    = makeFgStyle("241")
	Dot       = ColorFg(" • ", "236")
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
)

// Color a string's foreground with the given value.
func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(Term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(Term.Color(color)).Styled
}

// Generate a blend of colors.
func MakeRamp(colorA, colorB string, steps float64) (s []string) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, ColorToHex(c))
	}
	return
}

// Convert a colorful.Color to a hexadecimal format compatible with termenv.
func ColorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", ColorFloatToHex(c.R), ColorFloatToHex(c.G), ColorFloatToHex(c.B))
}

// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func ColorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
