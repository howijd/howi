// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import "testing"

func TestNumFlag(t *testing.T) {
	flag, _ := NewNumFlag("some-flag", "a")
	if present, err := flag.Parse(&[]string{"--some-flag"}); !present || err == nil {
		t.Error("expected num flag parser to return not ok, ", present, err)
	}
	if flag.Value().String() != "0" {
		t.Error("expected num value to be \"0\" got ", flag.Value().String())
	}
}
