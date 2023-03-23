package graphicinterface

import (
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
	MasterLabel     *widget.Label
	ResponseChannel chan ReturnTuple
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

func (gui GUI) AddRequest(request string) {
	gui.MasterLabel.SetText(request)
}

// poll until a response was given
func (gui *GUI) AwaitResponse() ReturnTuple {

	ReturnTuple := <-gui.ResponseChannel

	return ReturnTuple
}

func (gui *GUI) SetupGUI() {

	gui.MainWindow.Resize(fyne.NewSize(1000, 600))
	myCanvas := gui.MainWindow.Canvas()

	gui.MasterLabel = widget.NewLabel("")
	gui.MasterLabel.Alignment = fyne.TextAlignCenter
	gui.MasterLabel.TextStyle = fyne.TextStyle{Bold: true}

	//Build first row of keys
	keyboardRow1 := container.New(layout.NewGridLayout(12))
	keyboardRow1.Add(layout.NewSpacer())
	for _, button := range keyboardMapping1 {
		keyboardRow1.Add(widget.NewButton(button, func() {}))
	}
	keyboardRow1.Add(layout.NewSpacer())

	//Build second row including erase button
	keyboardRow2 := container.New(layout.NewGridLayout(12))
	keyboardRow2.Add(layout.NewSpacer())
	keyboardRow2.Add(widget.NewButton("Erase", func() {}))
	for _, button := range keyboardMapping2 {
		keyboardRow2.Add(widget.NewButton(button, func() {}))
	}
	keyboardRow2.Add(layout.NewSpacer())

	keyboardRow3 := container.New(layout.NewGridLayout(12))
	keyboardRow3.Add(layout.NewSpacer())
	keyboardRow3.Add(widget.NewButton("Enter", func() {}))
	keyboardRow3.Add(widget.NewButton("Mic", func() {}))
	for _, button := range keyboardMapping3 {
		keyboardRow3.Add(widget.NewButton(button, func() {}))
	}
	keyboardRow3.Add(layout.NewSpacer())

	parentContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), gui.MasterLabel, layout.NewSpacer(), keyboardRow1, keyboardRow2, keyboardRow3, layout.NewSpacer())
	myCanvas.SetContent(parentContainer)

}
