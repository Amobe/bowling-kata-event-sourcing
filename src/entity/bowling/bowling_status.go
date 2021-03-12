package bowling

type BowlingStatus string

const (
	FrameInitialed BowlingStatus = "FrameInitialed"
	GamePrepared   BowlingStatus = "GamePrepared"
	Thrown         BowlingStatus = "Thrown"
	Reloaded       BowlingStatus = "Reloaded"
	FrameFinished  BowlingStatus = "FrameFinished"
)
