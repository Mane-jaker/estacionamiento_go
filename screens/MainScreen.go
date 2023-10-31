package screens

import (
	"estacionamiento/controllers"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainScreen struct {
	app            fyne.App
	win            fyne.Window
	sceneContainer *fyne.Container
}

func NewMainScreen() *MainScreen {
	myApp := app.New()
	appWindow := myApp.NewWindow("ESTACIONAMIENTO")
	appWindow.Resize(fyne.NewSize(1366, 768))
	return &MainScreen{
		app:            myApp,
		win:            appWindow,
		sceneContainer: container.NewWithoutLayout(),
	}
}

func (m *MainScreen) Start() {
	rect := canvas.NewRectangle(color.White)
	rect.Resize(fyne.NewSize(1920, 1080))
	background := canvas.NewImageFromFile("assets/estacionamiento.png")
	background.Resize(fyne.NewSize(1200, 668))
	background.Move(fyne.NewPos(100, 50))
	m.sceneContainer.Add(rect)
	m.sceneContainer.Add(background)
	m.win.SetContent(m.sceneContainer)

	vc := controllers.NewVehicleController(m.app, m.win, m.sceneContainer)

	go vc.Create()

	m.win.CenterOnScreen()
	m.win.ShowAndRun()
}
