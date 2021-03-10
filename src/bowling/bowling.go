package bowling

const (
	standardFrameNumber = 10
	maxExtraFrameNumber = 2
	FrameWithExtraBonus = 9
)

type Bowling struct {
	ID          string
	FrameNumber uint32
	Status      BowlingStatus
	Score       uint32
	ExtraFrame  uint32
	Games       []Game
}

func NewBowling() *Bowling {
	b := &Bowling{
		ID:          "0",
		FrameNumber: 1,
		Status:      GamePrepared,
		Games:       make([]Game, 0, standardFrameNumber+maxExtraFrameNumber),
	}
	return b
}

func (b *Bowling) ApplyThrownEvent(hit uint32) {
	currentFrameIndex := b.FrameNumber - 1
	if len(b.Games) == int(b.FrameNumber) {
		b.Games[currentFrameIndex] = b.Games[currentFrameIndex].Hit(hit)
	} else {
		newGame := NewGame()
		if b.FrameNumber > FrameWithExtraBonus {
			newGame = NewGameWithoutExtraBonus()
		}
		b.Games = append(b.Games, newGame.Hit(hit))
	}

	for i := 1; i <= 2; i++ {
		bonusFrameIndex := int(currentFrameIndex) - i
		if bonusFrameIndex >= 0 {
			b.Games[bonusFrameIndex] = b.Games[bonusFrameIndex].Bonus(hit)
		}
	}
	b.calculateScore()
	b.Status = Thrown

	if b.Games[currentFrameIndex].NoMoreHit() || b.FrameNumber > 10 {
		b.ApplyReloadEvent()
	}
}

func (b *Bowling) calculateScore() {
	switch b.Games[b.FrameNumber-1].Status {
	case Spare:
		if b.FrameNumber == standardFrameNumber {
			b.ExtraFrame += 1
		}
	case Strike:
		if b.FrameNumber == standardFrameNumber {
			b.ExtraFrame += 2
		}
	}

	var score uint32
	for _, g := range b.Games {
		score = score + g.Score
	}
	b.Score = score
}

func (b *Bowling) ApplyReloadEvent() {
	if b.FrameNumber-b.ExtraFrame >= 10 {
		b.Status = FrameFinished
		return
	}
	b.FrameNumber++
}
