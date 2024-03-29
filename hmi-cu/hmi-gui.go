package hmicu

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/cricton/types"
)

type GUI struct {
	MainWindow      fyne.Window
	RequestLabel    *widget.Label
	ResponseChannel chan types.ReturnTuple
	UserEntry       strings.Builder
	inputLabel      *widget.Label

	//--------for testing--------//
	enterButton *widget.Button
}

func (gui GUI) GetConfirmation(requestor string, request string) {

	d := dialog.NewConfirm(requestor, request,
		func(b bool) {
			if b {
				gui.ResponseChannel <- types.ReturnTuple{Content: "", Code: types.ACCEPTED}
			} else {
				gui.ResponseChannel <- types.ReturnTuple{Content: "", Code: types.DECLINED}
			}
		}, gui.MainWindow)
	d.Show()

}

func (gui GUI) ShowInfo(requestor string, request string) {
	d := dialog.NewInformation(requestor, request, gui.MainWindow)

	d.SetOnClosed(func() {
		gui.ResponseChannel <- types.ReturnTuple{Content: "", Code: types.INFO}
	})
	d.Show()

}

// poll until a response was given
func (gui *GUI) AwaitResponse() types.ReturnTuple {

	ReturnTuple := <-gui.ResponseChannel

	return ReturnTuple
}

// Displays a text prompt for the user to respond to
func (gui GUI) GetString(requestor string, request string) {
	var strBuilder strings.Builder
	strBuilder.Write([]byte(requestor))
	strBuilder.Write([]byte(": "))
	strBuilder.Write([]byte(request))
	gui.RequestLabel.SetText(strBuilder.String())
}

func (gui *GUI) makeKeyboardRows() (*fyne.Container, *fyne.Container, *fyne.Container) {
	// Build first row of keys
	keyboardRow1 := container.New(layout.NewGridLayout(12))
	keyboardRow1.Add(layout.NewSpacer())
	for _, button := range types.KeyboardMapping1 {
		letter := button

		keyboardRow1.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				gui.inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow1.Add(layout.NewSpacer())

	// Build second row including erase button
	keyboardRow2 := container.New(layout.NewGridLayout(12))
	keyboardRow2.Add(layout.NewSpacer())
	keyboardRow2.Add(widget.NewButton("Erase", func() {
		str := gui.UserEntry.String()
		if len(str) > 0 {
			str = str[:len(str)-1]
		}
		gui.UserEntry.Reset()
		gui.UserEntry.WriteString(str)
		gui.inputLabel.SetText(gui.UserEntry.String())
	}))

	for _, button := range types.KeyboardMapping2 {
		letter := button
		keyboardRow2.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				gui.inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow2.Add(layout.NewSpacer())

	// Build third row including enter and microphone buttons
	keyboardRow3 := container.New(layout.NewGridLayout(12))
	keyboardRow3.Add(layout.NewSpacer())
	gui.enterButton = widget.NewButton("Enter", func() {
		userEntry := gui.UserEntry.String()
		gui.UserEntry.Reset()
		gui.RequestLabel.SetText("")
		gui.ResponseChannel <- types.ReturnTuple{Content: userEntry, Code: types.STRING}
		fmt.Println(userEntry)
		gui.inputLabel.SetText(gui.UserEntry.String())

	})
	keyboardRow3.Add(gui.enterButton)

	keyboardRow3.Add(widget.NewButton("Mic", func() {}))
	for _, button := range types.KeyboardMapping3 {
		letter := button
		keyboardRow3.Add(widget.NewButton(button, func() {
			if len(gui.RequestLabel.Text) > 0 {
				gui.UserEntry.WriteString(letter)
				gui.inputLabel.SetText(gui.UserEntry.String())
			}
		}))
	}
	keyboardRow3.Add(layout.NewSpacer())

	return keyboardRow1, keyboardRow2, keyboardRow3
}
func (gui *GUI) SetupGUI() {

	gui.MainWindow.Resize(fyne.NewSize(800, 400))
	myCanvas := gui.MainWindow.Canvas()

	//Label for requests
	gui.RequestLabel = widget.NewLabel("")
	gui.RequestLabel.Alignment = fyne.TextAlignCenter
	gui.RequestLabel.TextStyle = fyne.TextStyle{Bold: true}

	//Label for keyboard input
	gui.inputLabel = widget.NewLabel("")
	gui.inputLabel.Resize(fyne.NewSize(300, 100))
	inputLabelLayout := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), gui.inputLabel, layout.NewSpacer())

	//Add horizontal dividers
	rect1 := canvas.NewRectangle(color.White)
	rect1.SetMinSize(fyne.NewSize(2, 2))
	rect2 := canvas.NewRectangle(color.White)
	rect2.SetMinSize(fyne.NewSize(2, 2))

	parentContainer := container.NewVBox(
		layout.NewSpacer(),
		rect1,
		gui.RequestLabel,
		inputLabelLayout,
		rect2,
		layout.NewSpacer(),
		container.NewVBox(gui.makeKeyboardRows()),
		layout.NewSpacer())

	myCanvas.SetContent(parentContainer)

}

// Used for testing, returns reference to the keyboard enter button
func (gui GUI) GetEnterButton() *widget.Button {
	return gui.enterButton
}
