// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stderr, "github.com/dupoxy/go-tour-fr/gotour has moved to github.com/dupoxy/go-tour-fr\n")
	os.Exit(1)
}
