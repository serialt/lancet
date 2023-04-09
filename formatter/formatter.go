// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package formatter implements some functions to format string, struct.
package formatter

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/serialt/lancet/convertor"
	"github.com/serialt/lancet/strutil"
	"github.com/serialt/lancet/validator"
	"golang.org/x/exp/constraints"
)

// Comma add comma to a number value by every 3 numbers from right. ahead by symbol char.
// if value is invalid number string eg "aa", return empty string
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
// Play: https://go.dev/play/p/eRD5k2vzUVX
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string {
	if validator.IsInt(value) {
		v, err := convertor.ToInt(value)
		if err != nil {
			return ""
		}
		return symbol + commaInt(v)
	}

	if validator.IsFloat(value) {
		v, err := convertor.ToFloat(value)
		if err != nil {
			return ""
		}
		return symbol + commaFloat(v)
	}

	if strutil.IsString(value) {
		v := fmt.Sprintf("%v", value)
		if validator.IsNumberStr(v) {
			return symbol + commaStr(v)
		}
		return ""
	}

	return ""
}

// Pretty data to JSON string.
// Play: todo
func Pretty(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// PrettyToWriter pretty encode data to writer.
// Play: todo
func PrettyToWriter(v any, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}
