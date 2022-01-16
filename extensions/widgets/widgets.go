package widgets

import (
	"fyne.io/fyne/v2/widget"
)

func AppendText(label *widget.Label, text string) {
	label.SetText(label.Text + text)
}
