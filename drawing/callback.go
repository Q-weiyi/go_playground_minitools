package drawing

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Q-weiyi/go_playground_minitools/helper"
)

type CallbackSignatureCheck struct {
}

func (m CallbackSignatureCheck) GetContent(myWindow fyne.Window) (content *fyne.Container) {
	// Create the first tab content
	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("Enter body text here...")
	body.SetMinRowsVisible(10)

	signature := widget.NewMultiLineEntry()
	signature.SetPlaceHolder("Enter CPG-Signature text here...")

	signKey := widget.NewEntry()
	signKey.SetPlaceHolder("Enter sign key here...")

	callbackTypes := []string{"directdeposit", "invoicedeposit", "settlement", "refund", "withdrawal"}
	dropdown := widget.NewSelect(callbackTypes, func(selected string) {
		var exists = false
		for _, ks := range callbackTypes {
			if ks == selected {
				exists = true
				break
			}
		}
		if !exists {
			dialog.ShowError(errors.New("please select a callbacktype"), myWindow)
		}
	})
	compareButton := widget.NewButton("Compare", func() {
		m.compareSignature(myWindow, dropdown.Selected, body.Text, signKey.Text, signature.Text)
	})

	content = container.NewVBox(
		widget.NewLabel("Callback Type"),
		dropdown,
		widget.NewLabel("Body"),
		body,
		widget.NewLabel("Signature"),
		signature,
		widget.NewLabel("Sign Key"),
		signKey,
		compareButton,
	)
	return
}

func (m CallbackSignatureCheck) compareSignature(w fyne.Window, callbackType, bodyText, signKeyText, signature string) {
	data, err := helper.ParseJson(bodyText, callbackType)
	if err != nil {
		dialog.ShowError(errors.New("invalid body. not a proper json check the comma symbol"), w)
		return
	}

	hash, err := helper.CreateSignatureHeader([]byte(signKeyText), data)
	if err != nil {
		dialog.ShowError(err, w)
		return
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

	m.displayResult(w, mb)
}

type CustomResultMessageBox struct {
	Title   string
	Dismiss string
	Message string
}

func (m CallbackSignatureCheck) displayResult(w fyne.Window, param CustomResultMessageBox) {
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
