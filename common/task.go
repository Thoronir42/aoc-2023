package common

import (
	"io"
)

type AocTask interface {
	Pt1(in io.Reader) string
	Pt2(in io.Reader) string
}
