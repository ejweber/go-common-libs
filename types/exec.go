package types

import (
	"time"
)

const (
	BinaryCryptsetup = "cryptsetup"
	BinaryFstrim     = "fstrim"
	BinaryFsfreeze   = "fsfreeze"
)

const (
	ExecuteNoTimeout      = time.Duration(-1)
	ExecuteDefaultTimeout = time.Minute
)
