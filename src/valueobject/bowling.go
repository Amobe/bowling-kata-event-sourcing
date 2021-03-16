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
