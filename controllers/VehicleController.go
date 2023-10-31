package controllers

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

type VehicleController struct {
	app       fyne.App
	win       fyne.Window
	container *fyne.Container
}

func NewVehicleController(app fyne.App, win fyne.Window, container *fyne.Container) *VehicleController {
	return &VehicleController{
		app:       app,
		win:       win,
		container: container,
	}
}

func (v *VehicleController) Create() {
	pc := NewParkingController(v.app, v.win, v.container)

	for i := 0; i < 100; i++ {
		go pc.AddVehicle()
		time.Sleep(time.Duration(rand.ExpFloat64()/0.5) * time.Second)
	}

	select {}
}
