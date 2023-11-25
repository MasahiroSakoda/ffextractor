// Package util provides function for utility
package util

import "golang.org/x/mod/semver"

// Version is value for semver
type Version string

// Newer returns whether up to date or not
func (v Version) Newer(r Version) bool {
	return semver.Compare(string(v), string(r)) == 1
}
