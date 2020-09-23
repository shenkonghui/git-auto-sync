package common

import (
	"time"
)

type GitOptions struct {
	Interval time.Duration
	Directory string
	CommitName string
	CommitEmail string
}
