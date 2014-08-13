package cuid

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Counter interface {
	Next() int64
}

const (
	BLOCK_SIZE = 4
	BASE       = 36
)

var (
	defaultCounter Counter = nil
	defaultRandom          = rand.New(rand.NewSource(time.Now().Unix()))
	discreteValues         = int64(math.Pow(BASE, BLOCK_SIZE))
	padding                = strings.Repeat("0", BLOCK_SIZE)
	fingerprint            = ""
)

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "dummy-host"
	}
	acc := int64(len(hostname) + BASE)
	for i := range hostname {
		acc = acc + int64(hostname[i])
	}

	hostID := pad(strconv.FormatInt(int64(os.Getpid()), BASE), 2)
	host := pad(strconv.FormatInt(acc, 10), 2)
	fingerprint = hostID + host
}

func New() string {
	if defaultCounter == nil {
		defaultCounter = NewDefaultCounter()
	}

	timestampBlock := strconv.FormatInt(time.Now().Unix()*1000, BASE)
	counterBlock := pad(strconv.FormatInt(defaultCounter.Next(), BASE), BLOCK_SIZE)
	randomBlock1 := pad(strconv.FormatInt(defaultRandom.Int63n(discreteValues), BASE), BLOCK_SIZE)
	randomBlock2 := pad(strconv.FormatInt(defaultRandom.Int63n(discreteValues), BASE), BLOCK_SIZE)

	return "c" + timestampBlock + counterBlock + fingerprint + randomBlock1 + randomBlock2
}

func pad(str string, size int) string {
	if len(str) == size {
		return str
	}

	if len(str) < size {
		str = padding + str
	}

	i := len(str) - size

	return str[i:]
}
