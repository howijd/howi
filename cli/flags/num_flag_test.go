// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"testing"
)

func TestNumFlag(t *testing.T) {
	flag, _ := NewNumFlag("some-flag", "a")
	if present, err := flag.Parse(&[]string{"--some-flag"}); !present || err == nil {
		t.Error("expected num flag parser to return not ok, ", present, err)
	}
	if flag.Value().String() != "0" {
		t.Error("expected num value to be \"0\" got ", flag.Value().String())
	}
}

func TestNumFlagInt(t *testing.T) {
	var tests = []struct {
		name   string
		defval interface{}
		val    int
	}{
		{"basic", 1, 2},
		{"basic", 5, 2},
		{"basic", 4, 0},
		{"basic", 100, 2},
		{"basic", 123456789, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewNumFlag(tt.name)
			if err != nil {
				t.Errorf("valid flag %q did not expect error got %q", tt.name, err)
			}
			if tt.defval != nil {
				flag.Default(tt.defval)
			}
			args := []string{"--" + flag.Name(), fmt.Sprint(tt.val)}
			flag.Parse(&args)
			if flag.Value().Int() != tt.val {
				t.Errorf("expected flag %q Int to eq %d got %d", tt.name, tt.val, flag.Value().Int())
			}

			if flag.String() != fmt.Sprint(tt.val) {
				t.Errorf("expected flag %q to eq %q got %q", tt.name, fmt.Sprint(tt.val), flag.String())
			}
			if tt.defval != nil {
				flag.Unset()
				if flag.Value().Int() != tt.defval.(int) {
					t.Errorf("expected flag %q Int to eq %d got %d", tt.name, tt.defval, flag.Value().Int())
				}
				if flag.String() != fmt.Sprint(tt.defval) {
					t.Errorf("expected flag %q to eq %q got %q", tt.name, fmt.Sprint(tt.defval), flag.String())
				}
			}
		})
	}
}

func TestNumFlagFloat(t *testing.T) {
	var tests = []struct {
		name   string
		defval interface{}
		val    float64
	}{
		{"basic", 1.0, 2.0},
		{"basic", 1.0, 0.0},
		{"basic", 0.5, 2.0},
		{"basic", 100.05, 2.0},
		{"basic", 123456789.01, 0.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewNumFlag(tt.name)
			if err != nil {
				t.Errorf("valid flag %q did not expect error got %q", tt.name, err)
			}
			if tt.defval != nil {
				flag.Default(tt.defval)
			}
			args := []string{"--" + flag.Name(), fmt.Sprint(tt.val)}
			flag.Parse(&args)
			if flag.Value().Float64() != tt.val {
				t.Errorf("expected flag %q Float64 to eq %f got %f", tt.name, tt.val, flag.Value().Float64())
			}

			if flag.String() != fmt.Sprint(tt.val) {
				t.Errorf("expected flag %q to eq %q got %q", tt.name, fmt.Sprint(tt.val), flag.String())
			}
			if tt.defval != nil {
				flag.Unset()
				if flag.Value().Float64() != tt.defval {
					t.Errorf("expected flag %q Float64 to eq %f got %f", tt.name, float64(tt.defval.(int)), flag.Value().Float64())
				}
				if flag.String() != fmt.Sprint(tt.defval) {
					t.Errorf("expected flag %q to eq %q got %q", tt.name, fmt.Sprint(tt.defval), flag.String())
				}
			}
		})
	}
}
