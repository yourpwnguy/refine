package types

import "time"

type Params struct {
	OutputFile       string
	TotalLinesCount  int
	UniqueLinesCount int
	TimeTaken        time.Duration
}
