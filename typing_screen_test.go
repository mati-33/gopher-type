package main

import (
	"testing"
	"time"
)

func Test_calculateWpm(t *testing.T) {
	tests := []struct {
		name    string
		runesNo int
		elapsed time.Duration
		want    int
	}{
		{
			name:    "case1",
			runesNo: 50,
			elapsed: time.Second * 60,
			want:    10,
		},
		{
			name:    "case2",
			runesNo: 250,
			elapsed: time.Second * 33,
			want:    91,
		},
		{
			name:    "case3",
			runesNo: 133,
			elapsed: time.Second * 66,
			want:    24,
		},
		{
			name:    "case4",
			runesNo: 200,
			elapsed: time.Second * 5,
			want:    480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateWpm(tt.runesNo, tt.elapsed)
			if got != tt.want {
				t.Errorf("calculateWpm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAccuracy(t *testing.T) {
	tests := []struct {
		name     string
		runesNo  int
		errorsNo int
		want     float64
	}{
		{
			name:     "case1",
			runesNo:  100,
			errorsNo: 1,
			want:     0.99,
		},
		{
			name:     "case2",
			runesNo:  100,
			errorsNo: 0,
			want:     1.0,
		},
		{
			name:     "case3",
			runesNo:  100,
			errorsNo: 10,
			want:     0.9,
		},
		{
			name:     "case4",
			runesNo:  133,
			errorsNo: 5,
			want:     0.9624,
		},
		{
			name:     "case5",
			runesNo:  456,
			errorsNo: 59,
			want:     0.8706,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateAccuracy(tt.runesNo, tt.errorsNo)
			if got != tt.want {
				t.Errorf("calculateAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}
