package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

//资源池，goroutine之间安全地共享资源

var errPoolClosed = errors.New("pool is closed")

type pool struct {
	factory func()(io.Closer, error)
	resource chan io.Closer//所有close函数都实现这个接口

	mtx sync.Mutex
	closed bool
}
//pool constructor
func New(factory func()(io.Closer, error), size uint) (*pool, error) {
	if size <0 {
		return nil,errors.New("size is error")
	}

	return &pool{
		factory:  factory,
		resource: make(chan io.Closer, size),
		closed: false
	}, nil
}

func  (p *pool)AquireResource() (io.Closer, error) {
	select {
	case resource, ok := <-p.resource:
		if !ok{//就代表该通道已经关闭
			return nil,errPoolClosed
		}
		return resource, nil
	default:
		return p.factory()
	}

}

func (p *pool) ReleaseResource(resource io.Closer){
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if p.closed{
		resource.Close()
		return
	}

	select {
	case p.resource <- resource://如果能够丢回去
		fmt.Println("release the resource")
	default:
		fmt.Println("resource closed")
		resource.Close()
		
	}
}

func (p *pool)Close(){
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if p.closed{
		return
	}
	p.closed = true
	close(p.resource)

	for resource := range p.resource{
		resource.Close()
	}
}