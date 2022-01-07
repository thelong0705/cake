package cake

import (
	"math/rand"
	"time"
)

const CakeNums = 20

type Person struct {
	MixingTime   time.Duration
	BakingTime        time.Duration
	FrostingTime      time.Duration
	StandardDeviation time.Duration
}
type cake struct{}

func (p *Person) mix() {
	sleepWithStandardDeviation(p.MixingTime, p.StandardDeviation)
}

func (p *Person) bake() {
	sleepWithStandardDeviation(p.BakingTime, p.StandardDeviation)
}

func (p *Person) frost() {
	sleepWithStandardDeviation(p.FrostingTime, p.StandardDeviation)
}

func MakeSequentially(p Person) {
	for i := 0; i < CakeNums; i++ {
		p.mix()
		p.bake()
		p.frost()
	}
}

func sleepWithStandardDeviation(d, deviation time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(deviation))
	time.Sleep(delay)
}

func MakeConcurrently(mixer, baker, froster Person) {
	mixChannel := make(chan cake)

	go func() {
		for i := 0; i < CakeNums; i++ {
			mixer.mix()
			mixChannel <- cake{}
		}
		close(mixChannel)
	}()

	bakeChannel := make(chan cake)
	go func() {
		for i := 0; i < CakeNums; i++ {
			<-mixChannel
			baker.bake()
			bakeChannel <- cake{}
		}
	}()

	for i := 0; i < CakeNums; i++ {
		<-bakeChannel
		froster.frost()
	}
}

var (
	mixBuffer = 10
	bakeBuffer = 10
)
func MakeConcurrentlyWithBuffer(mixer, baker, froster Person) {
	mixChannel := make(chan cake, mixBuffer)

	go func() {
		for i := 0; i < CakeNums; i++ {
			mixer.mix()
			mixChannel <- cake{}
		}
		close(mixChannel)
	}()

	bakeChannel := make(chan cake, bakeBuffer)
	go func() {
		for i := 0; i < CakeNums; i++ {
			<-mixChannel
			baker.bake()
			bakeChannel <- cake{}
		}
	}()

	for i := 0; i < CakeNums; i++ {
		<-bakeChannel
		froster.frost()
	}
}


