package while

import "context"

type While struct {
	Ctx        context.Context
	Cancel     context.CancelFunc
	MaxWorkNum int
}

func NewWhile(maxWorkNum int) While {
	ctx, cancel := context.WithCancel(context.Background())
	return While{
		Ctx:        ctx,
		Cancel:     cancel,
		MaxWorkNum: maxWorkNum,
	}
}

func (w *While) For(f func()) {
	var count int
loop:
	for {
		f()
		count++
		if count >= w.MaxWorkNum {
			break
		}

		select {
		case <-w.Ctx.Done():
			break loop
		default:
			continue
		}
	}
}

func (w *While) Break() {
	w.Cancel()
}
