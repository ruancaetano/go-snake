package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

type Food struct {
	*tl.Entity
}

func NewFood() *Food {
	food := new(Food)

	food.Entity = tl.NewEntity(0, 0, 1, 1)
	food.GenerateRandomPosition(20, 20)

	return food
}

func (food *Food) Draw(screen *tl.Screen) {
	x, y := food.Position()

	screen.RenderCell(x, y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '*',
	})

}

func (food *Food) GenerateRandomPosition(screenWidth int, screenHeight int) {
	maxX := screenWidth - 2  // - border width
	maxY := screenHeight - 2 // - border width
	minValue := 5

	x := (rand.Intn(maxX-minValue) + minValue)
	y := (rand.Intn(maxY-minValue) + minValue)

	food.SetPosition(x, y)
}
