package make

import (
	"context"
	parallel "diction-car/pb/parallel"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var startTime = time.Now()

type Car struct {
	Body  string
	Tire  string
	Color string
}

type Serviceserver struct {
	parallel.UnimplementedParallelServer
}

func (s *Serviceserver) Paralleler(ctx context.Context, in *parallel.ParallelRequest) (*parallel.ParallelReply, error) {

	var values []int

	for i := 0; i < 10000; i++ {
		values = append(values, rand.Intn(1000))
	}
	NormalResult, NormalTime := Normal(values)
	GoResult, GoTime := Goroutine(values, 4)

	results := map[string]string{
		"NormalResult": strconv.Itoa(NormalResult),
		"NormalTime":   NormalTime,
		"Goresult":     strconv.Itoa(GoResult),
		"GoTime":       GoTime,
	}
	return &parallel.ParallelReply{Minimums: results}, nil
}

func Normal(a []int) (int, string) {
	startTime := time.Now()
	if len(a) == 0 {
		return 0, "0"
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	endtime := time.Now().Sub(startTime).String()
	return min, endtime
}

func Goroutine(a []int, n int) (int, string) {

	startTime := time.Now()

	if len(a) < n {
		return Normal(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i], _ = Normal(a[begin:end])
		}(i)
	}
	wg.Wait()

	end, _ := Normal(mins)

	endtime := time.Now().Sub(startTime).String()
	return end, endtime
}
