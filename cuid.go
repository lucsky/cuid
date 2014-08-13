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
	counter        Counter = nil
	discreteValues         = int64(math.Pow(BASE, BLOCK_SIZE))
	fingerprint            = ""
	randomSource           = rand.NewSource(time.Now().Unix())
	random                 = rand.New(randomSource)
	padding                = strings.Repeat("0", BLOCK_SIZE)
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
	if counter == nil {
		counter = NewDefaultCounter()
	}

	timestampBlock := strconv.FormatInt(time.Now().Unix()*1000, BASE)
	counterBlock := pad(strconv.FormatInt(counter.Next(), BASE), BLOCK_SIZE)
	randomBlock1 := pad(strconv.FormatInt(random.Int63n(discreteValues), BASE), BLOCK_SIZE)
	randomBlock2 := pad(strconv.FormatInt(random.Int63n(discreteValues), BASE), BLOCK_SIZE)

	return "c" + timestampBlock + counterBlock + fingerprint + randomBlock1 + randomBlock2
}

func pad(str string, size int) string {
	s := padding + str
	i := len(s) - size
	return s[i:]
}
