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
	if b.isNewGame() {
		b.Games = append(b.Games, b.createNewGame())
	}
	b.Games[currentFrameIndex] = b.Games[currentFrameIndex].Hit(hit)

	b.calculateBonus(hit)
	b.calculateScore()
	b.increaseExtraFrame()
	b.Status = Thrown

	if b.Games[currentFrameIndex].NoMoreHit() || b.FrameNumber > 10 {
		b.ApplyReloadEvent()
	}
}

func (b *Bowling) ApplyReloadEvent() {
	if b.FrameNumber-b.ExtraFrame >= 10 {
		b.Status = FrameFinished
		return
	}
	b.FrameNumber++
}

func (b *Bowling) isNewGame() bool {
	return len(b.Games) != int(b.FrameNumber)
}

func (b *Bowling) createNewGame() Game {
	if b.FrameNumber > FrameWithExtraBonus {
		return NewGameWithoutExtraBonus()
	}
	return NewGame()
}

func (b *Bowling) calculateScore() {
	var score uint32
	for _, g := range b.Games {
		score = score + g.Score
	}
	b.Score = score
}

func (b *Bowling) calculateBonus(hit uint32) {
	currentFrameIndex := b.FrameNumber - 1
	for i := 1; i <= 2; i++ {
		bonusFrameIndex := int(currentFrameIndex) - i
		if bonusFrameIndex >= 0 {
			b.Games[bonusFrameIndex] = b.Games[bonusFrameIndex].Bonus(hit)
		}
	}
}

func (b *Bowling) increaseExtraFrame() {
	if b.FrameNumber < standardFrameNumber {
		return
	}
	if b.hasExtraFrame() {
		b.ExtraFrame += 1
	}
}

func (b *Bowling) hasExtraFrame() bool {
	currentFrameIndex := b.FrameNumber - 1
	noOpen := b.FrameNumber == standardFrameNumber && b.Games[currentFrameIndex].Status != Open
	strikeTwice := b.FrameNumber == standardFrameNumber+1 && b.Games[currentFrameIndex].Status == Strike
	return noOpen || strikeTwice
}
