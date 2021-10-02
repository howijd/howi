// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.package flags

package flags

import "testing"

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
