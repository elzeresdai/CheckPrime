package pkg

import (
	"reflect"
	"testing"
)

func TestCheckPrimes(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    []string
	}{
		{
			name:    "test passed all numbers are prime",
			numbers: []int{2, 3, 5, 7},
			want:    []string{"true", "true", "true", "true"},
		},
		{
			name:    "test passed all numbers are not prime",
			numbers: []int{22, 33, 63, 72},
			want:    []string{"false", "false", "false", "false"},
		},
		{
			name:    "test passed mixed",
			numbers: []int{22, 2, 1, 17},
			want:    []string{"false", "true", "false", "true"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPrimes(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {

	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "test passed with prime number",
			args: 2,
			want: "true",
		},
		{
			name: "test passed with not prime number",
			args: 22,
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
