package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Vehicle struct {
	in      bool
	inIndex int
	img     *canvas.Image
	pox     int
	poy     int
}

func NewVehicle() *Vehicle {
	return &Vehicle{
		in:  false,
		img: canvas.NewImageFromFile("assets/car.png"),
		pox: 0,
		poy: 0,
	}
}

func (c *Vehicle) InIndex() int {
	return c.inIndex
}

func (c *Vehicle) SetInIndex(inIndex int) {
	c.inIndex = inIndex
}

func (c *Vehicle) Pox() int {
	return c.pox
}

func (c *Vehicle) SetPox(pox int) {
	c.pox = pox
}

func (c *Vehicle) Poy() int {
	return c.poy
}

func (c *Vehicle) SetPoy(poy int) {
	c.poy = poy
}

func (c *Vehicle) In() bool {
	return c.in
}

func (c *Vehicle) SetIn(in bool) {
	c.in = in
}

func (c *Vehicle) Img() *canvas.Image {
	return c.img
}

func (c *Vehicle) SetImg(img *canvas.Image) {
	c.img = img
}
