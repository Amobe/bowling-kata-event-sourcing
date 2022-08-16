package bowling

type Status string

const (
	GamePrepared  Status = "GamePrepared"
	Thrown        Status = "Thrown"
	FrameFinished Status = "FrameFinished"
)
