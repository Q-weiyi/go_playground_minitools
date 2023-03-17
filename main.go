package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/Q-weiyi/go_playground_minitools/drawing"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My GUI Application")

	var csc drawing.CallbackSignatureCheck
	callbackContent := csc.GetContent(myWindow)

	// Create the second tab content
	var bs drawing.BodySigner
	bodySigner := bs.GetContent(myWindow)

	// Create AppTabs container and add tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Callback Signature Check", callbackContent),
		container.NewTabItem("Sign Request", bodySigner),
	)

	// Set the window content and size, then show the window
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
