package themes

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func ThemeLagoon() *huh.Theme {
	t := huh.ThemeBase()

	var (
		normalFg        = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		lagoonBlue      = lipgloss.AdaptiveColor{Light: "#4578e5", Dark: "#4578e5"}
		cream           = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		lagoonLightBlue = lipgloss.Color("#2094f3")
		errorText       = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
		accents         = lipgloss.AdaptiveColor{Light: "#364c86", Dark: "#fafafc"}
	)

	t.Focused.Base = t.Focused.Base.BorderForeground(lipgloss.Color("238"))
	t.Focused.Card = t.Focused.Base
	t.Focused.Title = t.Focused.Title.Foreground(lagoonBlue).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(lagoonBlue).Bold(true).MarginBottom(1)
	t.Focused.Directory = t.Focused.Directory.Foreground(lagoonBlue)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(errorText)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(errorText)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(accents)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(lagoonLightBlue)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(lagoonLightBlue)
	t.Focused.Option = t.Focused.Option.Foreground(normalFg)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(accents)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(normalFg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(cream).Background(lagoonLightBlue)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(accents)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(accents)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.Card = t.Blurred.Base
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	t.Group.Title = t.Focused.Title
	t.Group.Description = t.Focused.Description
	return t
}
