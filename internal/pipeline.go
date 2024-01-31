package internal

import (
	"context"
	"sync"
)

type RunableFunc func() (uint32, error)
type CancelFunc func() error

type PipelineIFace interface {
	Add(ctx context.Context, action RunableFunc, cancel RunableFunc) error
	Start(ctx context.Context)
}

type Pipeline struct {
	List []struct {
		Action RunableFunc
		Cancel CancelFunc
	}
	CancelList []CancelFunc
	log        LoggerIFace
	err        []error
}

func NewPipeline(log LoggerIFace) *Pipeline {
	return &Pipeline{
		log: log,
	}
}

func (p *Pipeline) Add(ctx context.Context, action RunableFunc, cancel CancelFunc) error {
	p.List = append(p.List,
		struct {
			Action RunableFunc
			Cancel CancelFunc
		}{
			Action: action,
			Cancel: cancel,
		},
	)
	return nil
}

func (p *Pipeline) Start(ctx context.Context) (code uint32, err error) {
	defer func() {
		if code != 0 {
			_ = p.Cancel(ctx)
		}
	}()

	for _, v := range p.List {
		if v.Action != nil {
			code, err = v.Action()

			if code != 0 {
				return code, nil
			}

			if v.Cancel != nil {
				p.CancelList = append(p.CancelList, v.Cancel)
			}

			if err != nil {
				return code, err
			}
		}
	}

	return
}

func (p *Pipeline) Cancel(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(p.CancelList))

	for _, v := range p.CancelList {
		wg.Add(1)
		go func(v func() error) {
			defer wg.Done()
			errCh <- v()
		}(v)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			p.err = append(p.err, err)
		}
	}

	return nil
}
