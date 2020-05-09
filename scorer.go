package main

type Scorer struct {
	ball            int
	itsThrows       [21]int
	itsCurrentThrow int
}

func NewScorer() *Scorer {
	return &Scorer{}
}

func (s *Scorer) AddThrow(pins int) {
	s.itsThrows[s.itsCurrentThrow] = pins
	s.itsCurrentThrow += 1
}

func (s *Scorer) ScoreForFrame(theFrame int) int {
	s.ball = 0
	var score int
	for currentFrame := 0; currentFrame < theFrame; currentFrame++ {
		if s.strike() {
			score += (10 + s.nextTwoBallsForStrike())
			s.ball++
		} else if s.spare() {
			score += (10 + s.nextBallsForSpare())
			s.ball += 2
		} else {
			score += s.twoBallsInFrame()
			s.ball += 2
		}
	}
	return score
}

func (s *Scorer) strike() bool {
	return s.itsThrows[s.ball] == 10
}

func (s *Scorer) spare() bool {
	return (s.itsThrows[s.ball] + s.itsThrows[s.ball+1]) == 10
}

func (s *Scorer) nextTwoBallsForStrike() int {
	return s.itsThrows[s.ball+1] + s.itsThrows[s.ball+2]
}

func (s *Scorer) nextBallsForSpare() int {
	return s.itsThrows[s.ball+2]
}

func (s *Scorer) twoBallsInFrame() int {
	return s.itsThrows[s.ball] + s.itsThrows[s.ball+1]
}
