package pipefilter

import "errors"

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	var sum int
	for _, v := range parts {
		sum += v
	}
	return sum, nil
}
