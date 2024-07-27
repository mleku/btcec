package bech32

import (
	"bytes"
	"os"

	"ec.mleku.dev/v2/lol"
)

type (
	B = []byte
	S = string
)

var (
	log, chk, errorf = lol.New(os.Stderr)
	equals           = bytes.Equal
)
