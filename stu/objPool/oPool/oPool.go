package oPool

import (
	"errors"
	"time"
)

type Object struct {
}

type Opool struct {
	BufferChan chan *Object
	timeOut    time.Duration
}

func GetPool(num int) (op *Opool) {

	op.BufferChan = make(chan *Object, num)
	for i := 0; i < num; i++ {
		op.BufferChan <- &Object{}
	}
	return
}

func (op *Opool) GetObject() (o *Object, err error) {
	select {
	case o = <-op.BufferChan:
		return
	case <-time.After(op.timeOut):
		return nil, errors.New("请求超时")
	}
}

func (op *Opool) ReleaseObject(o *Object) error {
	select {
	case op.BufferChan <- o:
		return nil
	default:
		return errors.New("over flow")
	}
}
