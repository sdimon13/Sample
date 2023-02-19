package helpy

import (
	"fmt"
	"strings"
)

type Size uint64

const (
	Byte Size = 1 << (10 * iota)
	KibiByte
	MebiByte
	GibiByte
	TebiByte
	PebiByte
	ExbiByte
	// overflow
)

func (s Size) split() (low, high Size) {
	return s % 1024, s / 1024
}

func (s Size) String() string {
	var b, kb, mb, gb, tb, pb, eb Size
	var res []string

	b = s

	b, kb = b.split()
	kb, mb = kb.split()
	mb, gb = mb.split()
	gb, tb = gb.split()
	tb, pb = tb.split()
	pb, eb = pb.split()

	if eb > 0 {
		res = append(res, fmt.Sprintf("%dEiB", eb))
	}
	if pb > 0 {
		res = append(res, fmt.Sprintf("%dPiB", pb))
	}
	if tb > 0 {
		res = append(res, fmt.Sprintf("%dTiB", tb))
	}
	if gb > 0 {
		res = append(res, fmt.Sprintf("%dGiB", gb))
	}
	if mb > 0 {
		res = append(res, fmt.Sprintf("%dMiB", mb))
	}
	if kb > 0 {
		res = append(res, fmt.Sprintf("%dKiB", kb))
	}
	if b > 0 || len(res) == 0 {
		res = append(res, fmt.Sprintf("%dB", b))
	}

	return strings.Join(res, " ")
}
