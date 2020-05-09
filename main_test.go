package main_test

import (
	"testing"

	bowling "github.com/ShintaNakama/bowling"
)

func TestTwoThrowsNoMark(t *testing.T) {
	g := bowling.NewGame()
	g.Add(5)
	g.Add(4)
	res := g.Score()
	if res != 9 {
		t.Errorf("expected: 9, result: %v", res)
	}
}

func TestFourThrowsNoMark(t *testing.T) {
	g := bowling.NewGame()
	g.Add(5)
	g.Add(4)
	g.Add(7)
	g.Add(2)
	res := g.Score()
	if res != 18 {
		t.Errorf("score expected: 18, result: %v", res)
	}
	frame1Score := g.ScoreForFrame(1)
	if frame1Score != 9 {
		t.Errorf("frame1 expected: 9, result: %v", frame1Score)
	}
	frame2Score := g.ScoreForFrame(2)
	if frame2Score != 18 {
		t.Errorf("frame2 expected: 18, result: %v", frame2Score)
	}
}

func TestSimpleSpare(t *testing.T) {
	g := bowling.NewGame()
	g.Add(3)
	g.Add(7)
	g.Add(3)
	frame1Score := g.ScoreForFrame(1)
	if frame1Score != 13 {
		t.Errorf("expected: 13, result: %v", frame1Score)
	}
}

func TestSimpleFrameAfterSpare(t *testing.T) {
	g := bowling.NewGame()
	g.Add(3)
	g.Add(7)
	g.Add(3)
	g.Add(2)
	frame1Score := g.ScoreForFrame(1)
	if frame1Score != 13 {
		t.Errorf("expected: 13, result: %v", frame1Score)
	}
	frame2Score := g.ScoreForFrame(2)
	if frame2Score != 18 {
		t.Errorf("expected: 18, result: %v", frame2Score)
	}
	res := g.Score()
	if res != 18 {
		t.Errorf("score expected: 18, result: %v", res)
	}
}

func TestSimpleStrike(t *testing.T) {
	g := bowling.NewGame()
	g.Add(10)
	g.Add(3)
	g.Add(6)
	frame1Score := g.ScoreForFrame(1)
	if frame1Score != 19 {
		t.Errorf("expected: 19, result: %v", frame1Score)
	}
	res := g.Score()
	if res != 28 {
		t.Errorf("score expected: 28, result: %v", res)
	}
}

func TestPerfectGame(t *testing.T) {
	g := bowling.NewGame()
	for i := 0; i < 12; i++ {
		g.Add(10)
	}
	res := g.Score()
	if res != 300 {
		t.Errorf("score expected: 300, result: %v", res)
	}
}

func TestEndOfArray(t *testing.T) {
	g := bowling.NewGame()
	for i := 0; i < 9; i++ {
		g.Add(0)
		g.Add(0)
	}
	g.Add(2)
	g.Add(8)  // 10番目のフレームがスペア
	g.Add(10) // 最後にストライク
	res := g.Score()
	if res != 20 {
		t.Errorf("score expected: 20, result: %v", res)
	}
}

func TestSampleGame(t *testing.T) {
	g := bowling.NewGame()
	g.Add(1)
	g.Add(4)
	g.Add(4)
	g.Add(5)
	g.Add(6)
	g.Add(4)
	g.Add(5)
	g.Add(5)
	g.Add(10)
	g.Add(0)
	g.Add(1)
	g.Add(7)
	g.Add(3)
	g.Add(6)
	g.Add(4)
	g.Add(10)
	g.Add(2)
	g.Add(8)
	g.Add(6)
	res := g.Score()
	if res != 133 {
		t.Errorf("score expected: 133, result: %v", res)
	}
}

func TestHeartBreak(t *testing.T) {
	g := bowling.NewGame()
	for i := 0; i < 11; i++ {
		g.Add(10)
	}
	g.Add(9)
	res := g.Score()
	if res != 299 {
		t.Errorf("score expected: 299, result: %v", res)
	}
}

func TestTenthFrameSpare(t *testing.T) {
	g := bowling.NewGame()
	for i := 0; i < 9; i++ {
		g.Add(10)
	}
	g.Add(9)
	g.Add(1)
	g.Add(1)
	res := g.Score()
	if res != 270 {
		t.Errorf("score expected: 270, result: %v", res)
	}
}
