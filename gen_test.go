package pipelines

import (
	"sync"
	"testing"
)

var ts_gen = []struct {
	src []int
	dst []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 3}},
}

var ts_sq = []struct {
	src []int
	dst []int
}{
	{[]int{1, 2, 3}, []int{1, 4, 9}},
}

var ts_sq_merge = []struct {
	src []int
	sum int
}{
	{[]int{1, 2, 3}, 14},
}

func TestGen(t *testing.T) {
	for _, tests := range ts_gen {
		ch_gen := gen(tests.src...)

		for _, d := range tests.dst {
			r, ok := <-ch_gen
			if !ok {
				t.Errorf("failed to read on ch element %d", d)
			}
			if r != d {
				t.Errorf("expected %d, got %d", d, r)
			}
		}
	}
}

func TestSq(t *testing.T) {
	for _, tests := range ts_sq {
		ch_gen := gen(tests.src...)
		ch_sq := sq(ch_gen)

		for _, d := range tests.dst {
			r, ok := <-ch_sq
			if !ok {
				t.Errorf("failed to read on ch element %d", d)
			}
			if r != d {
				t.Errorf("expected %d, got %d", d, r)
			}
		}
	}
}

func TestSum(t *testing.T) {
	for _, tests := range ts_sq_merge {
		ch_gen := gen(tests.src...)
		ch_sq1 := sq(ch_gen)
		ch_sq2 := sq(ch_gen)
		ch_sq := merge(ch_sq1, ch_sq2)

		var sum int
		for _, s := range tests.src {
			r, ok := <-ch_sq
			if !ok {
				t.Errorf("failed to read on ch element %d", s)
			}
			sum += r
		}

		if sum != tests.sum {
			t.Errorf("expected %d, got %d", tests.sum, sum)
		}
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
