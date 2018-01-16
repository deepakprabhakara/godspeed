// Copyright 2014-2015 PagerDuty, Inc, et al. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package godspeed

import (
	"bytes"
	"strings"
)

// stats names can't include :, |, or @
var reservedReplacer = strings.NewReplacer(":", "_", "|", "_", "@", "_")

func trimReserved(s string) string {
	return reservedReplacer.Replace(s)
}

var pipesReplacer = strings.NewReplacer("|", "")

// function to make sure tags are unique
func uniqueTags(gt, t []string) []string {
	allTags := make([]string, 0, len(gt)+len(t))
	allTags = append(allTags, gt...)
	allTags = append(allTags, t...)
	t = allTags

	// if the tag slice is empty avoid allocation
	if len(t) < 1 {
		return nil
	}

	// build a map to track which values we've seen
	s := make(map[string]bool)

	// loop over each string provided
	// if the value is not in the map then replace
	// the value at t[len(s)] so that we always have
	// only unique tags at the beginning of the slice
	for i, v := range t {
		if _, x := s[v]; !x {
			// only change the value if needed
			if i != len(s) {
				t[len(s)] = v
			}

			s[v] = true
		}
	}

	// based on the size of the map we know
	// how many unique tags there were
	// so return that slice
	return []string(t[:len(s)])
}

func writeUniqueTags(buf *bytes.Buffer, replacer *strings.Replacer, gt, t []string) {
	if len(gt)+len(t) > 0 {
		buf.WriteString("|#")

		// build a map to track which values we've seen
		s := make(map[string]bool)

		needsComma := false
		for _, tags := range [][]string{gt, t} {
			for _, v := range tags {
				if s[v] {
					continue
				}
				s[v] = true

				if needsComma {
					buf.WriteString(",")
				} else {
					needsComma = true
				}
				if replacer != nil {
					replacer.WriteString(buf, v)
				} else {
					buf.WriteString(v)
				}
			}
		}
	}
}
