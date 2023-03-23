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

	button1 := widget.NewButton("Confirm", func() {
		if len(gui.MasterLabel.Text) > 0 {
			gui.ResponseChannel <- ReturnTuple{Content: "", Code: ACCEPTED}
			gui.MasterLabel.SetText("")
		}

	})

	button2 := widget.NewButton("Decline", func() {
		if len(gui.MasterLabel.Text) > 0 {
			gui.ResponseChannel <- ReturnTuple{Content: "", Code: DECLINED}
			gui.MasterLabel.SetText("")
		}
	})

	gui.MasterLabel = widget.NewLabel("")
	gui.MasterLabel.Alignment = fyne.TextAlignCenter
	gui.MasterLabel.TextStyle = fyne.TextStyle{Bold: true}

	content := container.New(layout.NewGridLayout(5), layout.NewSpacer(), button1, layout.NewSpacer(), button2, layout.NewSpacer())
	parentContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), gui.MasterLabel, layout.NewSpacer(), content, layout.NewSpacer())
	myCanvas.SetContent(parentContainer)

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	gui.MasterLabel.SetText("You good?")
	// }()

}
