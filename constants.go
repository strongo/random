package random

import (
	"math/rand"
	"sync"
	"time"
)

var (
	r   = rand.New(rand.NewSource(time.Now().UnixNano()))
	rMu sync.Mutex
)

const digits = "0123456789"
const lowerCase = "abcdefghijklmnopqrstuvwxyz"
const upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const chars = lowerCase + upperCase + digits
