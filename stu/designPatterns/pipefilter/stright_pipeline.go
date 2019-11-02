package pipefilter

import "errors"

var NewStrightPipelineParmsError = errors.New("parms should be []Filter")

type StrightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStrightPipeline(name string, filters ...Filter) *StrightPipeline {
	return &StrightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (sp *StrightPipeline) Process(data Request) (Response, error) {

	var (
		ret interface{}
		err error
	)
	for _, filter := range *sp.Filters {

		ret, err = filter.Process(data)
		if err != nil {
			return nil, err
		}
		data = ret

	}
	return ret, err
}
