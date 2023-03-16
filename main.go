package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Q-weiyi/go_playground_minitools/helper"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My GUI Application")

	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("Enter body text here...")
	body.SetMinRowsVisible(10)

	signature := widget.NewMultiLineEntry()
	signature.SetPlaceHolder("Enter CPG-Signature text here...")

	signKey := widget.NewEntry()
	signKey.SetPlaceHolder("Enter sign key here...")

	compareButton := widget.NewButton("Compare", func() {
		compareSignature(myWindow, body.Text, signKey.Text, signature.Text)
	})

	content := container.NewVBox(
		widget.NewLabel("Body"),
		body,
		widget.NewLabel("Signature"),
		signature,
		widget.NewLabel("Sign Key"),
		signKey,
		compareButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func compareSignature(w fyne.Window, bodyText, signKeyText, signature string) {
	data, err := helper.MinifyJSON(bodyText)
	if err != nil {
		dialog.ShowError(err, w)
	}

	hash, err := helper.CreateSignatureHeader([]byte(signKeyText), data)
	if err != nil {
		dialog.ShowError(err, w)
	}

	result := hash == signature
	mb := CustomResultMessageBox{
		Dismiss: "Close",
	}

	if !result {
		mb.Title = "Signature NOT MATCH"
		mb.Message = fmt.Sprintf("The signature is NOT equal.\nComputed Signature:%s\n\nProvided Signature:%s", hash, signature)
	} else {
		mb.Title = "Passed!!"
		mb.Message = fmt.Sprintf("The signature is equal.\n\nComputed Signature:%s\n\nProvided Signature:%s", hash, signature)
	}

	displayResult(w, mb)
}

type CustomResultMessageBox struct {
	Title   string
	Dismiss string
	Message string
}

func displayResult(w fyne.Window, param CustomResultMessageBox) {
	entry := widget.NewMultiLineEntry()
	entry.SetText(param.Message)
	entry.SetMinRowsVisible(5)
	content := container.NewVBox(
		widget.NewLabel("Output"),
		entry,
	)

	// Set the window size to 50% of the screen size.
	// Set the window size to 50% of the approximate screen size.
	d := dialog.NewCustom(param.Title, param.Dismiss, content, w)
	d.Resize(fyne.NewSize(750, 300))
	d.Show()
}
