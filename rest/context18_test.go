// +build go1.8

package rest

import (
	"context"
	"net/url"
)

func errIsContextCanceled(err error) bool {
	ue, ok := err.(*url.Error)
	if !ok {
		return false
	}
	return ue.Err == context.Canceled
}
