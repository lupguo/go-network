package mutext

import (
	"sync"
	"testing"
)

func Test_decrease(t *testing.T) {
	type args struct {
		num  int
		m    *sync.Mutex
		lock bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{num: 100, m: new(sync.Mutex),}},
		{"t2", args{num: 1000, m: new(sync.Mutex),}},
		{"t3", args{num: 10000, m: new(sync.Mutex), lock: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decrease(&tt.args.num, tt.args.m, tt.args.lock)
			t.Log(tt.name, tt.args.num)
		})
	}
}

func BenchmarkSimulatedDec100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if 0 != SimulatedDec(100, false) {
			b.Error("error, simulated fail")
		}
	}
}

func BenchmarkSimulatedDecUnlock(b *testing.B) {
	benchmarks := []struct {
		name  string
		count int
	}{
		{name: "bm0", count: 10},
		{name: "bm1", count: 100},
		{name: "bm2", count: 1000},
		{name: "bm3", count: 10000},
		{name: "bm4", count: 100000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if c := SimulatedDec(bm.count, false); c != 0 {
					b.Error("error, simulated fail, c=", c)
				}
			}
		})
	}
}

func BenchmarkSimulatedDecLock(b *testing.B) {
	benchmarks := []struct {
		name  string
		count int
	}{
		{name: "bm0", count: 10},
		{name: "bm1", count: 100},
		{name: "bm2", count: 1000},
		{name: "bm3", count: 10000},
		{name: "bm4", count: 100000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if c := SimulatedDec(bm.count, true); c != 0 {
					b.Error("error, simulated fail, c=", c)
				}
			}
		})
	}
}
