package valueobject

type BowlingStatus string

const (
	GamePrepared  BowlingStatus = "GamePrepared"
	Thrown        BowlingStatus = "Thrown"
	FrameFinished BowlingStatus = "FrameFinished"
)

type BowlingGameStatus string

const (
	Strike BowlingGameStatus = "Strike"
	Spare  BowlingGameStatus = "Spare"
	Open   BowlingGameStatus = "Open"
)

type BowlingGame struct {
	FrameNumber       uint32
	ThrowNumber       uint32
	Score             uint32
	Left              uint32
	Status            BowlingGameStatus
	WithoutExtraBonus bool
	ExtraBonus        uint32
}
