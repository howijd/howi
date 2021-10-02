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

	if flag.Value().String() != "" {
		t.Error("expected option value to be \"\" got ", flag.Value().String())
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
		{"basic3", []string{"opt1", "opt2"}, nil, "", ErrMissingOption},
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

func TestOptionName(t *testing.T) {
	for _, tt := range testflags() {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewOptionFlag(tt.name, []string{"a"})
			if !tt.valid {
				if err == nil {
					t.Errorf("invalid flag %q expected error got <nil>", tt.name)
				}
				if flag != nil {
					t.Errorf("invalid flag %q should be <nil> got %#v", tt.name, flag)
				}
			}
		})
	}
}

func TestMultiOpt(t *testing.T) {
	var tests = []struct {
		name string
		opts []string
	}{
		{"basic1", []string{"opt1", "opt2"}},
		{"basic1", []string{"opt1", "opt2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, _ := NewOptionFlag(tt.name, tt.opts)

			var args []string
			for _, o := range tt.opts {
				args = append(args, []string{"--" + tt.name, o}...)
			}

			ok, err := flag.Parse(&args)
			if !ok {
				t.Errorf("expected to parse opt multi flag %q got %q", tt.name, err)
			}
			opts := flag.Options()
			if len(opts) != len(tt.opts) {
				t.Errorf("epexted flag %q to have %d opts got %d", tt.name, len(tt.opts), len(opts))
			}
		})
	}
}

func TestOptionFlagDefaults(t *testing.T) {
	var tests = []struct {
		name     string
		opts     []string
		defaults []string
	}{
		{"basic1", []string{"opt1", "opt2"}, []string{"opt1"}},
		{"basic2", []string{"opt1", "opt2"}, []string{"opt1", "opt2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, _ := NewOptionFlag(tt.name, tt.opts)
			flag.Default(tt.defaults)

			var args []string
			_, err := flag.Parse(&args)
			if err != nil {
				t.Errorf("expected to parse opt flag %q got %q", tt.name, err)
			}
			opts := flag.Options()
			if len(opts) != len(tt.defaults) {
				t.Errorf("epexted flag %q to have %d opts got %d", tt.name, len(tt.defaults), len(opts))
			}
		})
	}
}
