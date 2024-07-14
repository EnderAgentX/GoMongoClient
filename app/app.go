package app

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func App() {
	gtk.Init(nil)
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	err = b.AddFromFile("Chat_gui.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	//Объекты
	objReg := GetObject(b, "regWin")

	//Окна
	winReg := objReg.(*gtk.Window)
	winReg.Connect("delete-event", func() {
		gtk.MainQuit()

	})

	//Элементы GUI
	objReg, _ = b.GetObject("login_reg_btn")
	loginRegBtn := objReg.(*gtk.Button)

	//Обработка событий
	loginRegBtn.Connect("clicked", func() {
		fmt.Println("test")
	})

	winReg.ShowAll()
	gtk.Main()
}

func GetObject(builder *gtk.Builder, objectId string) glib.IObject {
	obj, err := builder.GetObject(objectId)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	return obj
}
