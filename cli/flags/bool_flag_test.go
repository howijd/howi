// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import "testing"

func TestBoolFlagPresent(t *testing.T) {
	flag, _ := NewBoolFlag("some-flag")
	if present, err := flag.Parse(&[]string{"--some-flag"}); !present || err != nil {
		t.Error("expected bool flag parser to return present, nil got ", present, err)
	}

	if !flag.Present() {
		t.Error("expected bool flag to be present")
	}

	if flag.Value().Bool() != true {
		t.Error("expected bool value to be true got ", flag.Value().Bool())
	}

	if flag.Value().String() != "true" {
		t.Error("expected bool value to be \"true\" got ", flag.Value().String())
	}
	flag.Unset()
	if flag.Present() {
		t.Error("expected bool flag not to be present")
	}
}

func TestBoolFlagNotPresent(t *testing.T) {
	flag, _ := NewBoolFlag("some-flag")
	if ok, err := flag.Parse(&[]string{"--some-flag-2"}); ok || err != nil {
		t.Error("expected bool flag parser to return not ok, ", err)
	}

	if flag.Present() {
		t.Error("expected bool flag not to be present")
	}

	if flag.Value().Bool() != false {
		t.Error("expected bool value to be false got ", flag.Value().Bool())
	}

	if flag.Value().String() != "false" {
		t.Error("expected bool value to be \"false\" got ", flag.Value().String())
	}
}

func TestBooltName(t *testing.T) {
	for _, tt := range testflags() {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewBoolFlag(tt.name)
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
