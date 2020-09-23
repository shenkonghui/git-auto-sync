package common

import (
	"time"
)

type GitOptions struct {
	CommitInterval time.Duration
	PushInterval time.Duration
	Directory string
	CommitName string
	CommitEmail string
}
