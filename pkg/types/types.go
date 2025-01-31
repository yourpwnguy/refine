package types

import "time"

type Params struct {
	Filename         string
	OutputFile       string
	TotalLinesCount  int
	UniqueLinesCount int
	TimeTaken        time.Duration
}
