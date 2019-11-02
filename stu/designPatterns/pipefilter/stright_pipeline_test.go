package pipefilter

import "testing"

func TestStrightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStrightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	t.Log(ret)
	t.Log(err)
}
