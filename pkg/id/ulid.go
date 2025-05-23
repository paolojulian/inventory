package id

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

var entropy *ulid.MonotonicEntropy

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy = ulid.Monotonic(rand.New(source), 0)
}

func NewULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
