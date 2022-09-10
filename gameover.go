package main

import (
	tl "github.com/JoelOtter/termloop"
)

type GameOver struct {
	*tl.Text
}

func NewGameOver(screenWidth int, screenHeight int) Score {
	text := "GAME OVER"
	return Score{
		Text: tl.NewText(
			(screenWidth/2)-len(text)/2,
			screenHeight/2,
			text,
			tl.ColorRed,
			tl.ColorWhite,
		),
	}
}
