package screens

import (
	"estacionamiento/controllers"
	"math/rand"
	"time"

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
	background := canvas.NewImageFromFile("assets/estacionamiento.png")
	background.Resize(fyne.NewSize(1200, 668))
	background.Move(fyne.NewPos(100, 50))
	m.sceneContainer.Add(background)
	m.win.SetContent(m.sceneContainer)

	vc := controllers.NewVehicleController(m.app, m.win, m.sceneContainer)
	pc := controllers.NewParkingController(m.app, m.win, m.sceneContainer)

	m.createVehicles(vc, pc)

	m.win.CenterOnScreen()
	m.win.ShowAndRun()
}

func (m *MainScreen) createVehicles(vc *controllers.VehicleController, pc *controllers.ParkingController) {
	for i := 0; i < 100; i++ {
		go m.addVehicle(vc, pc)
		time.Sleep(time.Duration(rand.ExpFloat64()/0.5) * time.Second)
	}
}

func (m *MainScreen) addVehicle(vc *controllers.VehicleController, pc *controllers.ParkingController) {
	car := vc.CreateVehicle()
	pc.AddVehicle(car)
}
