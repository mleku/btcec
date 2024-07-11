package chaincfg

import (
	"bytes"
	"os"

	"github.com/mleku/btcec/lol"
)

type (
	B = []byte
	S = string
)

var (
	log, chk, errorf = lol.New(os.Stderr)
	equals           = bytes.Equal
)
