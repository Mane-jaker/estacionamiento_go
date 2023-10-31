package controllers

import (
	"estacionamiento/models"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

const (
	maxCapacity = 20
)

type ParkingController struct {
	parkingLot *models.ParkingLot
	app        fyne.App
	win        fyne.Window
	container  *fyne.Container
}

func NewParkingController(app fyne.App, win fyne.Window, container *fyne.Container) *ParkingController {
	return &ParkingController{
		parkingLot: models.NewParkingLot(maxCapacity),
		app:        app,
		win:        win,
		container:  container,
	}
}

func (p *ParkingController) AddVehicle() {
	car := models.NewVehicle()
	car.SetPox(0)
	car.SetPoy(280)
	car.Img().Resize(fyne.NewSize(50, 50))
	car.Img().Move(fyne.NewPos(0, 280))
	p.container.Add(car.Img())
	p.container.Refresh()

	p.parkingLot.Enter(car, car.Img(), p.container)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	p.parkingLot.Exit(car, car.Img(), p.container)
}
