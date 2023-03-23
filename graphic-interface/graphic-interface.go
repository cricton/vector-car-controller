package graphicinterface

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	ACCEPTED = 0
	DECLINED = 1
	INFO     = 2
)

type ReturnType struct {
	Content string
	Code    uint8
}

type GUI struct {
	MainWindow      fyne.Window
	MasterLabel     *widget.Label
	ResponseChannel chan ReturnType
}

func (gui GUI) AddRequest(request string) {
	gui.MasterLabel.SetText(request)
}

// poll until a response was given
func (gui *GUI) AwaitResponse() ReturnType {

	returnType := <-gui.ResponseChannel

	return returnType
}

func (gui *GUI) SetupGUI() {

	gui.MainWindow.Resize(fyne.NewSize(1000, 600))
	myCanvas := gui.MainWindow.Canvas()

	//var buttons []widget.Button
	/*for i := 0; i < 8; i++ {
		append(buttons, widget.NewButton("click me"))
	}*/

	button1 := widget.NewButton("Confirm", func() {
		gui.ResponseChannel <- ReturnType{Content: "", Code: ACCEPTED}
	})

	button2 := widget.NewButton("Decline", func() {
		gui.ResponseChannel <- ReturnType{Content: "", Code: DECLINED}
	})

	gui.MasterLabel = widget.NewLabel("text")
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
