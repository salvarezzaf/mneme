package main

import (
	"os"



	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
    
	os.Setenv("FYNE_THEME","light")

	mneme := app.New()
    mnemeWindow := mneme.NewWindow("Mneme - Kindle Highlights and Notes Manager")
   

	hello := widget.NewLabel("Test")
	mnemeWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			dialog.ShowFileOpen(func(uc fyne.URIReadCloser, e error) {
				dialog.ShowInformation("Your file selection",uc.URI().Path(),mnemeWindow)
			},mnemeWindow)
		}),
	))
    mnemeWindow.Resize(fyne.NewSize(500,700))
	mnemeWindow.ShowAndRun()


	//fetchBookMatadata(os.Args[1],os.Args[2])
}