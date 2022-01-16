package main_page

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/extensions/bindings"
	"github.com/wooseopkim/ck/v2/usecases"
)

const (
	http  = "http://"
	https = "https://"
)

func NewWidget(
	inferRemoteTime *usecases.InferRemoteTime,
	testRemoteTime *usecases.TestRemoteTime,
) fyne.CanvasObject {
	urlEntry := widget.NewEntry()
	protocolSelect := widget.NewSelect([]string{https, http}, func(s string) {})
	protocolSelect.SetSelected(protocolSelect.Options[0])
	inputContainer := fyne.NewContainerWithLayout(
		layout.NewFormLayout(),
		protocolSelect,
		urlEntry,
	)
	panelLabel := widget.NewLabel("")
	panelLabel.Alignment = fyne.TextAlignCenter
	targetLabel := widget.NewLabel("")
	targetLabel.Alignment = fyne.TextAlignCenter
	targetLabel.TextStyle = fyne.TextStyle{Bold: true}
	submitButton := widget.NewButton("Go", func() {})
	testResultLabel := widget.NewLabel("")
	testResultLabel.Alignment = fyne.TextAlignCenter
	testButton := widget.NewButton("", func() {})
	testerLayout := container.NewVBox(
		layout.NewSpacer(),
		testResultLabel,
		layout.NewSpacer(),
		testButton,
	)
	testerLayout.Hide()
	layout := container.NewVBox(
		inputContainer,
		layout.NewSpacer(),
		targetLabel,
		panelLabel,
		layout.NewSpacer(),
		submitButton,
		testerLayout,
	)

	vm := NewViewModel(
		inferRemoteTime,
		testRemoteTime,
	)
	submitButton.OnTapped = func() {
		vm.OnSubmit(protocolSelect.Selected + urlEntry.Text)
	}
	testButton.OnTapped = func() {
		vm.OnTest()
	}
	inputEnabled := vm.InputEnabled()
	bindings.OnBoolChange(inputEnabled, func(value bool) {
		if value {
			submitButton.Enable()
			testButton.Enable()
		} else {
			submitButton.Disable()
			testButton.Disable()
		}
	})
	panel := vm.Panel()
	bindings.OnStringChange(panel, func(value string) {
		panelLabel.Text = value
		panelLabel.Refresh()
		// Maybe because the panel is updated too quickly,
		// another widget must be `Refresh`ed for the panel to be updated.
		submitButton.Refresh()
	})
	target := vm.Target()
	bindings.OnStringChange(target, func(value string) {
		targetLabel.Text = value
		targetLabel.Refresh()
		testButton.SetText(fmt.Sprintf("Test %s", value))
	})
	testResult := vm.TestResult()
	bindings.OnUntypedChange(testResult, func(value interface{}) {
		if value == nil {
			return
		}
		testResult := value.(adapters.TestResult)
		testResultLabel.Text = fmt.Sprintf(
			"Client: %s\nClock: %s\nRemote: %s",
			testResult.ClientTime.Format(timeTemplate),
			testResult.ClockTime.Format(timeTemplate),
			testResult.RemoteTime.Format(timeTemplate),
		)
		testResultLabel.Refresh()
		layout.Refresh()
	})
	testerReady := vm.TesterReady()
	bindings.OnBoolChange(testerReady, func(value bool) {
		if value {
			testerLayout.Show()
		} else {
			testerLayout.Hide()
		}
		testerLayout.Refresh()
		layout.Refresh()
	})

	return layout
}
