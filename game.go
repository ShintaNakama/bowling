package main

import "math"

type Game struct {
	itsCurrentFrame   int
	firstThrowInFrame bool
	itsScorer         *Scorer
}

func NewGame() *Game {
	return &Game{firstThrowInFrame: true, itsScorer: NewScorer()}
}

func (g *Game) Score() int {
	return g.ScoreForFrame(g.itsCurrentFrame)
}

func (g *Game) Add(pins int) {
	g.itsScorer.AddThrow(pins)
	g.adjustCurrentFrame(pins)
}

func (g *Game) adjustCurrentFrame(pins int) {
	if g.lastBallInFrame(pins) {
		g.advanceFrame()
	} else {
		g.firstThrowInFrame = false
	}
}

func (g *Game) lastBallInFrame(pins int) bool {
	return g.strike(pins) || g.firstThrowInFrame
}

func (g *Game) strike(pins int) bool {
	return (g.firstThrowInFrame && pins == 10)
}

func (g *Game) advanceFrame() {
	g.itsCurrentFrame = int(math.Min(10, float64(g.itsCurrentFrame+1)))
}

func (g *Game) ScoreForFrame(theFrame int) int {
	return g.itsScorer.ScoreForFrame(theFrame)
}
