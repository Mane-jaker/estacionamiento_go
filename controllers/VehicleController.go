package controllers

import (
	"estacionamiento/models"

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

func (v *VehicleController) CreateVehicle() *models.Vehicle {
	car := models.NewVehicle()
	car.SetPox(0)
	car.SetPoy(280)
	car.Img().Resize(fyne.NewSize(96, 96))
	car.Img().Move(fyne.NewPos(0, 280))
	v.container.Add(car.Img())
	v.container.Refresh()
	return car
}
