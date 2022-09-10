package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Border struct {
	*tl.Entity
	coordinates map[Coordinate]bool
}

func NewBorder() *Border {
	border := new(Border)

	border.Entity = tl.NewEntity(0, 0, 0, 0)
	border.coordinates = make(map[Coordinate]bool)

	return border
}

func (border *Border) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()

	for i := 0; i < screenWidth; i++ {
		border.coordinates[Coordinate{i, 0}] = true
		border.coordinates[Coordinate{i, screenHeight - 1}] = true
		// top
		screen.RenderCell(i, 0, &tl.Cell{
			Bg: tl.RgbTo256Color(83, 87, 84),
		})
		// bottom
		screen.RenderCell(i, screenHeight-1, &tl.Cell{
			Bg: tl.RgbTo256Color(83, 87, 84),
		})
	}

	for i := 1; i < screenHeight; i++ {
		border.coordinates[Coordinate{0, i}] = true
		border.coordinates[Coordinate{screenWidth - 1, i}] = true
		// left
		screen.RenderCell(0, i, &tl.Cell{
			Bg: tl.RgbTo256Color(83, 87, 84),
		})
		// right
		screen.RenderCell(screenWidth-1, i, &tl.Cell{
			Bg: tl.RgbTo256Color(83, 87, 84),
		})
	}
}

func (border *Border) Contains(coordinate Coordinate) bool {
	return border.coordinates[coordinate] || false
}
