package record

import (
	"errors"
)

type Range interface {
	Start() Time
	End() Time
	Duration() Duration
}

type timeRange struct {
	start Time
	end   Time
}

func NewRange(start Time, end Time) (Range, error) {
	if !end.IsAfterOrEqual(start) {
		return nil, errors.New("ILLEGAL_RANGE")
	}
	return timeRange{
		start: start,
		end:   end,
	}, nil
}

func (tr timeRange) Start() Time {
	return tr.start
}

func (tr timeRange) End() Time {
	return tr.end
}

func (tr timeRange) Duration() Duration {
	start := tr.Start().MidnightOffset().InMinutes()
	end := tr.End().MidnightOffset().InMinutes()
	return NewDuration(0, end-start)
}