package cake_test

import (
	"cake"
	"testing"
	"time"
)

var person = cake.Person{
	MixingTime:   time.Millisecond * 10,
	BakingTime:   time.Millisecond * 10,
	FrostingTime: time.Millisecond * 10,
}

func BenchmarkMakeSequentially(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cake.MakeSequentially(person)
	}
}

func BenchmarkMakeConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cake.MakeConcurrently(person, person, person)
	}
}


func BenchmarkMakeConcurrentlyWithVariousSpeed(b *testing.B) {
	p := cake.Person{
		MixingTime:   time.Millisecond * 10,
		BakingTime:   time.Millisecond * 10,
		FrostingTime: time.Millisecond * 10,
		StandardDeviation: time.Millisecond * 3,
	}
	for i := 0; i < b.N; i++ {
		cake.MakeConcurrently(p, p, p)
	}
}

func BenchmarkMakeConcurrentlyWithVariousSpeedAndBuffer(b *testing.B) {
	p := cake.Person{
		MixingTime:   time.Millisecond * 10,
		BakingTime:   time.Millisecond * 10,
		FrostingTime: time.Millisecond * 10,
		StandardDeviation: time.Millisecond * 3,
	}
	for i := 0; i < b.N; i++ {
		cake.MakeConcurrentlyWithBuffer(p, p, p)
	}
}
