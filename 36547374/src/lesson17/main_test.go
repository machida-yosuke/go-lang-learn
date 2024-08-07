package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2=3",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "2+2=4",
			args: args{x: 2, y: 2},
			want: 4,
		},

		{
			name: "0+0=0",
			args: args{x: 0, y: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "10/5=2",
			args: args{x: 10, y: 5},
			want: 2,
		},
		{
			name: "0/0=0",
			args: args{x: 0, y: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
