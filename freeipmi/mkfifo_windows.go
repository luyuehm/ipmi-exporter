// Copyright 2024 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows

package freeipmi

import "errors"

// mkfifo is unsupported on Windows (no syscall.Mkfifo). The freeipmi exec path
// is unreachable on Windows; native mode (--native-ipmi, go-ipmi) does not call
// this. Returning an error keeps the package cross-compilable.
func mkfifo(path string, mode uint32) error {
	return errors.New("mkfifo is not supported on windows (use --native-ipmi)")
}
