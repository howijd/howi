// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.package flags

package flags

import (
	"errors"
	"testing"
)

func TestOptionFlag(t *testing.T) {
	flag, _ := NewOptionFlag("some-flag", []string{"a", "b", "c"}, "s")
	if ok, err := flag.Parse(&[]string{"--some-flag=a"}); !ok || err != nil {
		t.Error("expected option flag parser to return ok, ", ok, err)
	}

	if flag.Value().String() != "a" {
		t.Error("expected option value to be \"a\" got ", flag.Value().String())
	}
}

func TestOptionFlagFalse(t *testing.T) {
	flag, _ := NewOptionFlag("some-flag", []string{"a", "b", "c"}, "s")
	if present, err := flag.Parse(&[]string{"--some-flag=d"}); !present || err == nil {
		t.Error("expected option flag parser to return !present and err, ", present, err)
	}

	if flag.Value().String() != "d" {
		t.Error("expected option value to be \"d\" got ", flag.Value().String())
	}
}

func TestOptionFlagEmpty(t *testing.T) {
	flag, _ := NewOptionFlag("some-flag", []string{"a", "b", "c"}, "s")
	if present, err := flag.Parse(&[]string{"--some-flag"}); !present || err == nil {
		t.Error("expected option flag parser to return present and err, ", present, err)
	}

	if flag.Value().String() != "" {
		t.Error("expected option value to be \"\" got ", flag.Value().String())
	}
}

func TestOptions(t *testing.T) {
	var tests = []struct {
		name   string
		opts   []string
		defval interface{}
		val    string
		err    error
	}{
		{"basic1", nil, nil, "", nil},
		{"basic2", []string{"opt1", "opt2"}, nil, "opt3", ErrInvalidValue},
		{"basic3", []string{"opt1", "opt2"}, nil, "opt2", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewOptionFlag(tt.name, tt.opts)
			if len(tt.opts) == 0 {
				if !errors.Is(err, ErrMissingOption) {
					t.Error("expected error while creating opt flag got: ", err)
				}
				return
			}

			if len(tt.opts) > 0 && err != nil {
				t.Error("did not expect error while creating opt flag got: ", err)
				return
			}

			args := []string{"--" + tt.name, tt.val}
			_, err = flag.Parse(&args)
			if !errors.Is(err, tt.err) {
				t.Errorf("expected error %q got %q", tt.err, err)
			}
		})
	}
}
