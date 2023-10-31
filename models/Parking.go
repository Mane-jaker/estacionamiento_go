package models

import (
	"fmt"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type ParkingLot struct {
	capacity int
	occupied int
	gate     sync.Mutex
	spaces   []bool
	spacesX  []int
	spacesY  []int
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		spaces:   make([]bool, capacity),
		spacesX: []int{
			175,
			272,
			370,
			465,
			560,
			655,
			755,
			850,
			945,
			1040,
			150,
			245,
			345,
			440,
			537,
			633,
			730,
			830,
			930,
			1025,
		},
		spacesY: []int{
			100,
			580,
		},
	}
}

func (p *ParkingLot) Enter(car *Vehicle, carImage *canvas.Image, sceneContainer *fyne.Container) {
	p.gate.Lock()
	defer p.gate.Unlock()

	time.Sleep(time.Second)
	for {
		if p.occupied < p.capacity {
			break
		}
		sceneContainer.Remove(carImage)
		sceneContainer.Refresh()
		return
	}

	for i := 0; i < p.capacity; i++ {
		if !p.spaces[i] {
			p.spaces[i] = true
			p.occupied++

			if i < 10 {
				for j := car.poy; j < car.poy+100; j += 2 {
					carImage.Move(fyne.NewPos(float32(0), float32(j)))
					time.Sleep(time.Millisecond)
				}
				carImage.Move(fyne.NewPos(float32(p.spacesX[i]), float32(p.spacesY[0])))
				car.SetInIndex(i)
				fmt.Printf("Vehicle parked in space %d\n", i+1)
				break
			}
			for j := car.poy; j < car.poy+100; j += 2 {
				carImage.Move(fyne.NewPos(float32(0), float32(j)))
				time.Sleep(time.Millisecond)
			}
			carImage.Move(fyne.NewPos(float32(p.spacesX[i]), float32(p.spacesY[1])))
			car.SetInIndex(i)
			fmt.Printf("Vehicle parked in space %d\n", i+1)
			break
		}
	}
}

func (p *ParkingLot) Exit(car *Vehicle, carImage *canvas.Image, sceneContainer *fyne.Container) {
	p.gate.Lock()
	defer p.gate.Unlock()

	for i := p.capacity - 1; i >= 0; i-- {
		if i == car.inIndex {
			p.spaces[i] = false
			p.occupied--

			if i < 10 {
				for j := p.spacesY[0]; j < p.spacesY[0]+100; j += 2 {
					carImage.Move(fyne.NewPos(float32(p.spacesX[i]), float32(j)))
					time.Sleep(time.Millisecond)
				}
				sceneContainer.Remove(carImage)
				sceneContainer.Refresh()
				break
			}

			for j := p.spacesY[1]; j > p.spacesY[1]-100; j -= 2 {
				carImage.Move(fyne.NewPos(float32(p.spacesX[i]), float32(j)))
				time.Sleep(time.Millisecond)
			}
			sceneContainer.Remove(carImage)
			sceneContainer.Refresh()
			fmt.Printf("Vehicle left space %d\n", i+1)
			break
		}
	}
}
