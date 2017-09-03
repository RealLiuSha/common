// +build !go1.8

package rest

import "strings"

func errIsContextCanceled(err error) bool {
	return strings.Contains(err.Error(), "request canceled")
}
