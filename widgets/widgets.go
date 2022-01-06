package widgets

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func AppendText(label *widget.Label, text string) {
	label.SetText(label.Text + text)
}

func OnStringChange(item binding.String, callback func(string)) {
	item.AddListener(binding.NewDataListener(func() {
		value, _ := item.Get()
		callback(value)
	}))
}

func OnBoolChange(item binding.Bool, callback func(bool)) {
	item.AddListener(binding.NewDataListener(func() {
		value, _ := item.Get()
		callback(value)
	}))
}

func OnUntypedChange(item binding.Untyped, callback func(interface{})) {
	item.AddListener(binding.NewDataListener(func() {
		value, _ := item.Get()
		callback(value)
	}))
}
