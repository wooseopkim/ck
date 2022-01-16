package bindings

import (
	"fyne.io/fyne/v2/data/binding"
)

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
