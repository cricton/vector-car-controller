package graphicinterface

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	MainWindow  fyne.Window
	MasterLabel widget.Label
}

func (gui GUI) AddRequest(request string) {

}

func (gui GUI) SetupGUI() {

	gui.MainWindow.Resize(fyne.NewSize(1000, 600))
	myCanvas := gui.MainWindow.Canvas()

	//var buttons []widget.Button
	/*for i := 0; i < 8; i++ {
		append(buttons, widget.NewButton("click me"))
	}*/

	button1 := widget.NewButton("Confirm", func() {
		dialog.ShowConfirm("Test", "You good bro?", func(b bool) {}, gui.MainWindow)
	})
	button2 := widget.NewButton("Decline", func() {
		log.Println("2")
	})

	textPrompt := widget.NewLabel("text")
	textPrompt.Alignment = fyne.TextAlignCenter
	textPrompt.TextStyle = fyne.TextStyle{Bold: true}

	content := container.New(layout.NewGridLayout(5), layout.NewSpacer(), button1, layout.NewSpacer(), button2, layout.NewSpacer())
	parentContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), textPrompt, layout.NewSpacer(), content, layout.NewSpacer())
	myCanvas.SetContent(parentContainer)
	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})
	go func() {
		time.Sleep(time.Second * 2)
		textPrompt.SetText("You good?")
	}()

}
