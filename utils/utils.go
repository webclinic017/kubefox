// Copyright 2023 XigXog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

var (
	RegexpNameSpecialChar  = regexp.MustCompile(`[^a-z0-9]`)
	RegexpLabelSpecialChar = regexp.MustCompile(`[^a-z0-9A-Z-_\.]`)
	RegexpLabelPrefix      = regexp.MustCompile(`^[^a-z0-9A-Z]*`)
	RegexpLabelSuffix      = regexp.MustCompile(`[^a-z0-9A-Z-_\.]*[^a-z0-9A-Z]*$`)
)

func ResolveFlag(curr, envVar, def string) string {
	if curr != "" {
		return curr
	}
	if e := os.Getenv(envVar); e != "" {
		return e
	} else {
		return def
	}
}

func ResolveFlagBool(curr bool, envVar string, def bool) bool {
	if curr != def {
		return curr
	}
	if e, err := strconv.ParseBool(os.Getenv(envVar)); err == nil {
		return e
	} else {
		return def
	}
}

func ResolveFlagInt(curr int, envVar string, def int) int {
	if curr != def {
		return curr
	}
	if e, err := strconv.ParseInt(os.Getenv(envVar), 10, 0); err == nil {
		return int(e)
	} else {
		return def
	}
}

func CheckRequiredFlag(n, p string) {
	if p == "" {
		fmt.Fprintf(os.Stderr, "The flag \"%s\" is required.\n\n", n)
		flag.Usage()
		os.Exit(1)
	}
}

func EnvDef(name, def string) string {
	e, _ := os.LookupEnv(name)
	if e == "" {
		return def
	}
	return e
}

func UIntToByteArray(i uint64) []byte {
	data := *(*[unsafe.Sizeof(i)]byte)(unsafe.Pointer(&i))
	return data[:]
}

func ByteArrayToUInt(b []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(&b[0]))
}

func ShortHash(hash string) string {
	if len(hash) < 7 {
		return ""
	}

	return hash[0:7]
}

// First returns the first non-empty string. If all strings are empty then empty
// string is returned.
func First(strs ...string) string {
	for _, s := range strs {
		if s != "" {
			return s
		}
	}

	return ""
}

// CleanName returns name with all special characters replaced with dashes and
// set to lowercase. The string is truncated if longer than 63 characters. If
// name is a path only the basename is used.
//
// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
func CleanName(name string) string {
	cleaned := filepath.Base(name)
	cleaned = strings.ToLower(cleaned)
	cleaned = RegexpNameSpecialChar.ReplaceAllLiteralString(cleaned, "-")
	cleaned = strings.TrimPrefix(strings.TrimSuffix(cleaned, "-"), "-")
	if len(cleaned) > 63 {
		cleaned = cleaned[:63]
	}

	return cleaned
}

func IsValidName(name string) bool {
	if name == "" {
		return false
	}
	return name == CleanName(name)
}

// CleanLabel returns the label value with all special characters replaced with
// dashes and any character that is not [a-z0-9A-Z] trimmed from start and end.
// If name is a path only the basename is used.
//
// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
func CleanLabel(value any) string {
	cleaned := filepath.Base(fmt.Sprint(value))
	// Remove special chars.
	cleaned = RegexpLabelSpecialChar.ReplaceAllLiteralString(cleaned, "-")
	// Ensure value begins and ends with [a-z0-9A-Z].
	cleaned = RegexpLabelPrefix.ReplaceAllLiteralString(cleaned, "")
	cleaned = RegexpLabelSuffix.ReplaceAllLiteralString(cleaned, "")
	return cleaned
}

// Join concatenates the elements of its second argument to create a single
// string. The separator string sep is placed between elements in the resulting
// string. Empty elements are ignored.
func Join(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	var b strings.Builder
	for _, s := range elems {
		if s != "" {
			b.WriteString(sep)
			b.WriteString(s)
		}
	}
	if b.Len() == 0 {
		return ""
	}

	return b.String()[1:]
}

func SetBit(n uint32, pos uint) uint32 {
	return n | (1 << pos)
}

func ClearBit(n uint32, pos uint) uint32 {
	return n & ^(1 << pos)
}

func HasBit(n uint32, pos uint) bool {
	return (n & (1 << pos)) > 0
}
