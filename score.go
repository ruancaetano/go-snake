package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

type Score struct {
	*tl.Text
	value int
}

func NewScore() Score {
	return Score{
		Text:  tl.NewText(2, 0, "Score: 0", tl.ColorWhite, tl.RgbTo256Color(83, 87, 84)),
		value: 0,
	}
}

func (score *Score) IncreaseScoreValue() {
	score.value += 1
	score.SetText(fmt.Sprint("Score: ", score.value, " "))
}

func (score *Score) ResetScoreValue() {
	score.value = 0
	score.SetText(fmt.Sprint("Score: ", score.value, " "))
}
