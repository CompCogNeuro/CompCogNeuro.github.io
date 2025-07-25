// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package yaegisims

//go:generate ./make

import (
	"reflect"
)

var Symbols = map[string]map[string]reflect.Value{}
