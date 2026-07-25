package modules

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"pc-monitoring/helpers"

	"github.com/gogpu/systray"
)

type Tray struct {
	server *Server
	tray   *systray.SystemTray
	action string
	icon   string
}

func NewTray(server *Server) *Tray {
	tray := systray.New()
	icon := "icon.png"
	
	instance := &Tray{
		server: server,
		tray:   tray,
		icon:   icon,
		action: "Stop",
	}

	menu := instance.setMenu()
	image := instance.setIcon()

	tray.SetTooltip("PC Monitoring").SetIcon(image.Bytes()).SetMenu(menu)

	return instance
}

func (t *Tray) Show() {
	fmt.Println("Starting Application...")
	t.tray.Show()

	go func() {
		for msg := range helpers.ServerStatus {
			if msg {
				t.action = "Stop"
			} else {
				t.action = "Start"
			}
		}
	}()

	// Run the platform message loop (blocks until Quit)
	if err := t.tray.Run(); err != nil {
		fmt.Println("error:", err)
	}
}

func(t *Tray) setMenu() *systray.Menu {
	menu := systray.NewMenu()

	// TODO: t.action not changing Tray Menu Label
    menu.Add("Start/Stop", func() {
		switch t.action {
			case "Stop":
				t.action = "Start"
				t.server.Stop()
			case "Start":
				t.action = "Stop"
				t.server.Start()
		}
	})

	menu.Add("Open", func () {
		err := helpers.OpenBrowser("http://localhost:9003")

		if err != nil {
			fmt.Println("Error to open Web Browser: ", err)
		}
	})

    menu.AddSeparator()
    menu.Add("Quit", func() {
		t.tray.Remove()
        os.Exit(0)
    })

	return menu
}

func(t *Tray) setIcon() bytes.Buffer {
	var buffer bytes.Buffer
	file, err := os.Open("static/"+t.icon)

	if err != nil {
		fmt.Println("Cannot Open Tray Icon File:", err)
		return buffer
	}

	defer func() {
		err = file.Close()

		if err != nil {
			fmt.Println("Cannot Close Tray Icon File:", err)
			return
		}
	}()

	image, _, err := image.Decode(file)

	if err != nil {
		fmt.Println("Cannot Decode Tray Icon File:", err)
		return buffer
	}

	err = png.Encode(&buffer, image)

	if err != nil {
		fmt.Println("Cannot Convert Tray Icon File to Bytes:", err)
		return buffer
	}

	return buffer 
}
