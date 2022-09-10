package main

import (
	tl "github.com/JoelOtter/termloop"
)

const SNAKE_BODY_PART_SIZE = 2
const SNAKE_COLOR = tl.ColorBlack

const (
	right int = iota
	left
	up
	down
)

type Snake struct {
	*tl.Entity
	coordinates  []Coordinate
	lastMoviment int
}

func NewSnake() *Snake {
	return &Snake{
		Entity:       tl.NewEntity(5, 5, 1, 1),
		lastMoviment: right,
		coordinates: []Coordinate{
			{3, 5},
			{4, 5},
			{5, 5},
		},
	}
}

func (snake *Snake) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()

	snake.move()
	snake.checkCollision(screenWidth, screenHeight)
	snake.render(screen)

}

func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.lastMoviment != left {
				snake.lastMoviment = right
			}

		case tl.KeyArrowLeft:
			if snake.lastMoviment != right {
				snake.lastMoviment = left
			}

		case tl.KeyArrowUp:
			if snake.lastMoviment != down {
				snake.lastMoviment = up
			}

		case tl.KeyArrowDown:
			if snake.lastMoviment != up {
				snake.lastMoviment = down
			}
		}
	}
}

func (snake *Snake) move() {
	head := snake.head()

	headX, headY := head.X, head.Y

	switch snake.lastMoviment {
	case right:
		headX += 1
	case left:
		headX -= 1
	case up:
		headY -= 1
	case down:
		headY += 1
	}

	snake.coordinates = append(snake.coordinates[1:], Coordinate{headX, headY})
}

func (snake *Snake) render(screen *tl.Screen) {
	for _, bodyPart := range snake.coordinates {
		screen.RenderCell(bodyPart.X, bodyPart.Y, &tl.Cell{
			Fg: tl.ColorBlack,
			Ch: 'o',
		})
	}
}

func (snake *Snake) checkCollision(screenWidth int, screenHeight int) {
	if snake.checkCollisionWithFood() {
		score.IncreaseScoreValue()
		food.GenerateRandomPosition(screenWidth, screenHeight)
		snake.coordinates = append(snake.coordinates, snake.head())
		return
	}

	if snake.checkBorderCollision() {
		EndGame(screenWidth, screenHeight)
		return
	}

	if snake.checkCollisionWithItself() {
		EndGame(screenWidth, screenHeight)
		return
	}

}

func (snake *Snake) checkCollisionWithFood() bool {
	head := snake.head()

	foodX, foodY := food.Position()

	if head.X == foodX && head.Y == foodY {
		return true
	}

	return false
}

func (snake *Snake) checkBorderCollision() bool {
	head := snake.head()

	return border.Contains(head)
}

func (snake *Snake) checkCollisionWithItself() bool {
	head := snake.head()

	for i := 0; i < len(snake.coordinates)-1; i++ {
		coord := snake.coordinates[i]
		if head.X == coord.X && head.Y == coord.Y {
			return true
		}
	}

	return false
}

func (snake *Snake) head() Coordinate {
	return snake.coordinates[len(snake.coordinates)-1]
}
