// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the timeformat checker.

package a

import (
	"time"

	"b"
)

func hasError() {
	a, _ := time.Parse("2006-02-01 15:04:05", "2021-01-01 00:00:00") // want `2006-02-01 should be 2006-01-02`
	a.Format(`2006-02-01`)                                           // want `2006-02-01 should be 2006-01-02`
	a.Format("2006-02-01 15:04:05")                                  // want `2006-02-01 should be 2006-01-02`

	const c = "2006-02-01"
	a.Format(c) // want `2006-02-01 should be 2006-01-02`
}

func notHasError() {
	a, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")
	a.Format("2006-01-02")

	const c = "2006-01-02"
	a.Format(c)

	v := "2006-02-01"
	a.Format(v) // Allowed though variables.

	m := map[string]string{
		"y": "2006-02-01",
	}
	a.Format(m["y"])

	s := []string{"2006-02-01"}
	a.Format(s[0])

	a.Format(badFormat())

	o := b.Parse("2006-02-01 15:04:05", "2021-01-01 00:00:00")
	o.Format("2006-02-01")
}

func badFormat() string {
	return "2006-02-01"
}
