package main

import (
	tl "github.com/JoelOtter/termloop"
)

var game = tl.NewGame()
var border = NewBorder()
var snake = NewSnake()
var food = NewFood()
var score = NewScore()

func StartGame() {
	game.SetDebugOn(true)
	game.Screen().SetFps(10)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.RgbTo256Color(112, 255, 148),
		Fg: tl.ColorBlack,
	})

	level.AddEntity(border)
	level.AddEntity(score)
	level.AddEntity(snake)
	level.AddEntity(food)

	game.Screen().SetLevel(level)

	game.Start()
}

func EndGame(screenWidth int, screenHeight int) {
	game.Screen().Level().RemoveEntity(snake)
	game.Screen().Level().RemoveEntity(food)
	game.Screen().Level().AddEntity(NewGameOver(screenWidth, screenHeight))
}
