package make

import (
	"context"
	"fmt"
	car "diction-car/pb/car"
	"log"
	"strconv"
	"strings"
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
	car.UnimplementedMakerServer
}

func (s *Serviceserver) MakeCar(ctx context.Context, in *car.CarRequest) (*car.CarReply, error) {
	log.Printf("Received: %s", in.GetKind())

	tirech := make(chan *Car)
	paintch := make(chan *Car)
	starch := make(chan *[]string)

	wg.Add(3)
	go MakeBody(tirech)
	go InstallTire(tirech, paintch)
	go PaintCar(paintch, starch)

	return &car.CarReply{Message: "hellow"}, nil
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}

	}
}

func InstallTire(tirech, paintch chan *Car) {
	for car := range tirech {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintch <- car
	}
	wg.Done()
	close(paintch)
}

func PaintCar(paintch chan *Car, starch chan *[]string) {
	for car := range paintch {
		var str []string
		var strs []string
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		str = append(str, strconv.Itoa(int(duration)))
		str = append(str, car.Tire)
		str = append(str, car.Body)
		str = append(str, car.Color)
		strs = append(strs, strings.Join(str, " "))
		fmt.Printf("%.4f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Body)
		starch <- &strs
	}
	wg.Done()
}
