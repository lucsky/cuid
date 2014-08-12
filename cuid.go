package cuid

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	BLOCK_SIZE = 4
	BASE       = 36
)

var (
	counter        int64 = 0
	discreteValues       = int64(math.Pow(BASE, BLOCK_SIZE))
	fingerprint          = ""
	randomSource         = rand.NewSource(time.Now().Unix())
	random               = rand.New(randomSource)
	counterChan          = make(chan int64)
)

func init() {
	pidPart := pad(strconv.FormatInt(int64(os.Getpid()), BASE), 2)

	hostname, _ := os.Hostname()
	acc := int64(len(hostname) + BASE)
	for i := range hostname {
		acc = acc + int64(hostname[i])
	}
	hostnamePart := pad(strconv.FormatInt(acc, 10), 2)

	fingerprint = pidPart + hostnamePart

	go counterFunc()
}

func New() string {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, BASE)
	counter := pad(strconv.FormatInt(<-counterChan, BASE), BLOCK_SIZE)
	randomBlock1 := newRandomBlock()
	randomBlock2 := newRandomBlock()

	return "c" + timestamp + counter + fingerprint + randomBlock1 + randomBlock2
}

func NewSlug() string {
	return "slug"
}

func counterFunc() {
	var count int64 = 0
	for {
		counterChan <- count
		counter = counter + 1
		if counter >= discreteValues {
			counter = 0
		}
	}
}

func newRandomBlock() string {
	return pad(strconv.FormatInt(random.Int63n(discreteValues), BASE), BLOCK_SIZE)
}

func pad(str string, size int) string {
	s := "0000000000" + str
	i := len(s) - size
	return s[i:]
}
