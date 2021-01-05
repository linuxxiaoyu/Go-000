package main

import (
	"container/ring"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numBuckets         = 3
	timeInMilliseconds = 300
)

type Bucket struct {
	StartTime time.Time
	Success   int
}

func (b *Bucket) String() string {
	return fmt.Sprintf("bucket: %d\"%02d: %d",
		b.StartTime.Second(),
		b.StartTime.Nanosecond()/1e7,
		b.Success,
	)
}

func NewWindowWithContext(ctx context.Context, timeInMilliseconds int, numBuckets int) *window {
	w := &window{
		r:   ring.New(numBuckets),
		cap: numBuckets,
	}
	w.r.Value = &Bucket{
		StartTime: time.Now().Local(),
	}
	ctx, cancle := context.WithCancel(ctx)
	w.cancleFunc = cancle
	t := time.NewTicker(time.Millisecond * time.Duration(timeInMilliseconds/numBuckets))
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				curTime := <-t.C
				w.Sum()
				w.mux.Lock()
				w.r = w.r.Next()
				w.r.Value = &Bucket{StartTime: curTime}
				w.mux.Unlock()
			}
		}
	}()
	return w
}

type window struct {
	cancleFunc context.CancelFunc
	r          *ring.Ring
	cap        int
	mux        sync.RWMutex
}

func (w *window) IncSuccess() {
	w.mux.Lock()
	defer w.mux.Unlock()
	b := w.r.Value.(*Bucket)
	if b != nil {
		b.Success++
	}
}

func (w *window) Success() int {
	w.mux.RLock()
	defer w.mux.RUnlock()
	return w.r.Value.(*Bucket).Success
}

func (w *window) Shutdown() {
	w.cancleFunc()
}

func (w *window) Print() {
	w.mux.RLock()
	defer w.mux.RUnlock()

	p := w.r
	for i := 0; i < w.cap; i++ {
		if p.Value == nil {
			break
		}
		fmt.Println(p.Value.(*Bucket))
		p = p.Prev()
	}
}

func (w *window) Sum() {
	w.mux.RLock()
	defer w.mux.RUnlock()

	sum := 0
	p := w.r
	for i := 0; i < w.cap; i++ {
		if p.Value == nil {
			break
		}
		b := p.Value.(*Bucket)
		sum += b.Success
		p = p.Prev()
	}

	var startTime time.Time
	startTime = p.Next().Value.(*Bucket).StartTime
	fmt.Printf(
		"window: %d\"%02d: %d\n\n",
		startTime.Second(),
		startTime.Nanosecond()/1e7,
		sum,
	)
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	m := numBuckets + 1
	w := NewWindowWithContext(context.Background(), timeInMilliseconds, numBuckets)
	for i := 0; i < m; i++ {
		n := rand.Intn(10)
		var wg sync.WaitGroup
		wg.Add(n)
		for j := 0; j < n; j++ {
			go func() {
				defer wg.Done()
				w.IncSuccess()
			}()
		}
		wg.Wait()
		fmt.Printf(
			"%d\"%02d: +%d\n",
			time.Now().Second(),
			time.Now().Nanosecond()/1e7,
			n,
		)
		if i < m-1 {
			time.Sleep(time.Duration(timeInMilliseconds/numBuckets) * time.Millisecond)
		}
	}
	w.Shutdown()
	w.Print()
	w.Sum()
}
