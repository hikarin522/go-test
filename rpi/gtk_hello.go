package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	// gtk Start
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Hello World!")
	window.Connect("destroy", gtk.MainQuit)

	hbox := gtk.NewHBox(false, 0)

	button := gtk.NewButtonWithLabel("Click Here")

	button.Clicked(func() {
		msgDlg := gtk.NewMessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"こんにちは世界")

		msgDlg.Response(msgDlg.Destroy)

		msgDlg.Run()
	})

	hbox.Add(button)

	button2 := gtk.NewButtonWithLabel("Button2")

	button2.Clicked(func() {
		msgDlg := gtk.NewMessageDialog(
			button2.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"Button2")

		msgDlg.Response(msgDlg.Destroy)

		msgDlg.Run()
	})

	hbox.Add(button2)

	window.Add(hbox)

	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
