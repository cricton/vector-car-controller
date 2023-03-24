package graphicinterface

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ReturnType uint8

const (
	ACCEPTED ReturnType = 0
	DECLINED ReturnType = 1
	INFO     ReturnType = 2
	ERROR    ReturnType = 3
)

type ReturnTuple struct {
	Content string
	Code    ReturnType
}

type GUI struct {
	MainWindow      fyne.Window
	RequestLabel    *widget.Label
	ResponseChannel chan ReturnTuple
	UserEntry       strings.Builder
}

func (gui GUI) GetConfirmation(request string) {
	d := dialog.NewConfirm("Please confirm!", request,
		func(b bool) {
			if b {
				gui.ResponseChannel <- ReturnTuple{Content: "", Code: ACCEPTED}
			} else {
				gui.ResponseChannel <- ReturnTuple{Content: "", Code: DECLINED}
			}
		}, gui.MainWindow)
	d.Show()

}

func (gui GUI) ShowInfo(request string) ReturnTuple {
	d := dialog.NewInformation("Info!", request, gui.MainWindow)
	d.Show()

	return ReturnTuple{Content: "", Code: INFO}
}

// poll until a response was given
func (gui *GUI) AwaitResponse() ReturnTuple {

	ReturnTuple := <-gui.ResponseChannel

	return ReturnTuple
}

func (gui GUI) GetString(request string) {
	gui.RequestLabel.SetText(request)
}

func (gui *GUI) SetupGUI() {

	gui.MainWindow.Resize(fyne.NewSize(1000, 600))
	myCanvas := gui.MainWindow.Canvas()

	//Label for requests
	gui.RequestLabel = widget.NewLabel("")
	gui.RequestLabel.Alignment = fyne.TextAlignCenter
	gui.RequestLabel.TextStyle = fyne.TextStyle{Bold: true}

	//Label for keyboard input
	inputLabel := widget.NewLabel("")
	inputLabel.Resize(fyne.NewSize(300, 100))
	inputLabelLayout := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), inputLabel, layout.NewSpacer())

	//Build first row of keys
	keyboardRow1 := container.New(layout.NewGridLayout(12))
	keyboardRow1.Add(layout.NewSpacer())
	for _, button := range keyboardMapping1 {
		letter := button

		keyboardRow1.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow1.Add(layout.NewSpacer())

	//Build second row including erase button
	keyboardRow2 := container.New(layout.NewGridLayout(12))
	keyboardRow2.Add(layout.NewSpacer())
	keyboardRow2.Add(widget.NewButton("Erase", func() {
		str := gui.UserEntry.String()
		if len(str) > 0 {
			str = str[:len(str)-1]
		}
		gui.UserEntry.Reset()
		gui.UserEntry.WriteString(str)
		inputLabel.SetText(gui.UserEntry.String())
	}))
	for _, button := range keyboardMapping2 {
		letter := button
		keyboardRow2.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow2.Add(layout.NewSpacer())

	//Build third row including enter and microphone buttons
	keyboardRow3 := container.New(layout.NewGridLayout(12))
	keyboardRow3.Add(layout.NewSpacer())
	keyboardRow3.Add(widget.NewButton("Enter", func() {
		userEntry := gui.UserEntry.String()
		gui.UserEntry.Reset()
		gui.RequestLabel.SetText("")
		gui.ResponseChannel <- ReturnTuple{Content: userEntry, Code: INFO}
		fmt.Println(userEntry)
		inputLabel.SetText(gui.UserEntry.String())

	}))

	keyboardRow3.Add(widget.NewButton("Mic", func() {}))
	for _, button := range keyboardMapping3 {
		letter := button
		keyboardRow3.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow3.Add(layout.NewSpacer())

	sep := widget.NewSeparator()
	parentContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), gui.RequestLabel, inputLabelLayout, layout.NewSpacer(), keyboardRow1, keyboardRow2, keyboardRow3, layout.NewSpacer())
	myCanvas.SetContent(parentContainer)
	sep.Show()

}
