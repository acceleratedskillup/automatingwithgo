package cli

import (
	"fmt"
	"goprojgen/internal/termstyle"
	"math"
	"strings"

	"github.com/muesli/termenv"
)

var (
	progressEmpty = termstyle.Subtle(progressEmptyChar)
	// Gradient colors we'll use for the progress bar
	ramp = termstyle.MakeRamp("#B14FFF", "#00FFA3", progressBarWidth)
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

func Checkbox(label string, checked bool) string {
	if checked {
		return termstyle.ColorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

func Progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += termenv.String(progressFullChar).Foreground(termstyle.Term.Color(ramp[i])).String()
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
