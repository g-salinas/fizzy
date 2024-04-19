package main

import (
	"errors"
	"testing"
)

func TestBuildMessage(t *testing.T) {

	tests := []struct {
		name        string
		input       Input
		want        string
		wantedError error
	}{{
		name: "example",
		input: Input{
			Int1:  3,
			Int2:  5,
			Limit: 100,
			Str1:  "fizz",
			Str2:  "buzz",
		},
		want: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz",
	},
		{
			name: "using 1",
			input: Input{
				Int1:  1,
				Int2:  5,
				Limit: 20,
				Str1:  "a",
				Str2:  "b",
			},
			want: "a,a,a,a,ab,a,a,a,a,ab,a,a,a,a,ab,a,a,a,a,ab",
		},
		{
			name: "nonsense limit",
			input: Input{
				Int1:  1,
				Int2:  5,
				Limit: -10,
				Str1:  "a",
				Str2:  "b",
			},
			want: "",
		},
		{
			name: "nonsense divider",
			input: Input{
				Int1:  1,
				Int2:  0,
				Limit: 20,
				Str1:  "a",
				Str2:  "b",
			},
			want:        "",
			wantedError: ErrorOnDivider,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildMessage(tt.input)
			if got != tt.want {
				t.Errorf("buildMessage() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.wantedError) {
				t.Errorf("error = %v, want %v", err, tt.wantedError)
			}
		})
	}
}
