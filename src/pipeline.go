package threshold

import (
	"context"
	"log"

	"github.com/SuperJourney/threshold/util"
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
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
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
			// 通过了，添加回滚
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
	for _, v := range p.CancelList {
		v := v
		go util.PanicHandle(func(ctx context.Context) {
			err := v()
			if err != nil {
				log.Fatalf("cancel error: %v", err)
			}
		})(ctx)
	}
	return nil
}
