package fp_test

import (
	"testing"

	"github.com/bcarnazzi/eulergo/pkg/fp"
)

func Test_Unfold(t *testing.T) {
	ch := fp.Unfold(1, func(i int) int { return 2 * i })
	t1 := <-ch
	t2 := <-ch
	t3 := <-ch
	t4 := <-ch
	t5 := <-ch
	if t1 != 1 {
		t.Error("incorrect result: expected 1, got", t1)
	}
	if t2 != 2 {
		t.Error("incorrect result: expected 2, got", t2)
	}
	if t3 != 4 {
		t.Error("incorrect result: expected 4, got", t3)
	}
	if t4 != 8 {
		t.Error("incorrect result: expected 8, got", t4)
	}
	if t5 != 16 {
		t.Error("incorrect result: expected 16, got", t5)
	}
}

func Test_Unfold2(t *testing.T) {
	ch := fp.Unfold2(0, 1, func(i, j int) int { return 2*i + j })
	t1 := <-ch
	t2 := <-ch
	t3 := <-ch
	t4 := <-ch
	t5 := <-ch
	if t1 != 0 {
		t.Error("incorrect result: expected 0, got", t1)
	}
	if t2 != 1 {
		t.Error("incorrect result: expected 1, got", t2)
	}
	if t3 != 1 {
		t.Error("incorrect result: expected 1, got", t3)
	}
	if t4 != 3 {
		t.Error("incorrect result: expected 3, got", t4)
	}
	if t5 != 5 {
		t.Error("incorrect result: expected 5, got", t5)
	}
}

func Test_Fold(t *testing.T) {
	ch := make(chan int, 4)

	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
		}
		close(ch)
	}()

	result := fp.Fold(ch, 0, func(i, j int) int { return i + j })
	if result != 10 {
		t.Error("incorrect result: expected 10, got", result)
	}

}

func Test_TakeWhile(t *testing.T) {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- 2 * i
		}
		close(ch)
	}()

	result := fp.TakeWhile(ch, func(i int) bool {
		return i <= 7
	})
	i := 1
	for v := range result {
		if v != 2*i {
			t.Errorf("incorrect result: expected %d, got %d", 2*i, v)
		}
		i++
	}

}

func Test_Map(t *testing.T) {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	result := fp.Map(ch, func(i int) int {
		return 3 * i
	})

	i := 1
	for v := range result {
		if v != 3*i {
			t.Errorf("incorrect result: expected %d, got %d", 3*i, v)
		}
		i++
	}
}

func Test_Filter(t *testing.T) {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- 3 * i
		}
		close(ch)
	}()

	result := fp.Filter(ch, func(i int) bool {
		return i%2 == 0
	})

	for v := range result {
		if v%2 != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 0, v%2)
		}
	}
}
