package main_page

import "fyne.io/fyne/v2/data/binding"

type ViewModel interface {
	OnSubmit(url string)
	Panel() binding.String
	InputEnabled() binding.Bool
}
