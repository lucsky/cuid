package cuid

import (
	cryptoRand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	blockSize = 4
	base      = 36
)

var (
	mutex          sync.Mutex
	counter        Counter
	random         *rand.Rand
	discreteValues = int32(math.Pow(base, blockSize))
	padding        = strings.Repeat("0", blockSize)
	fingerprint    = ""
	format         = regexp.MustCompile(fmt.Sprintf("c[0-9a-z]{%d}", 6*blockSize))
)

func init() {
	SetRandomSource(rand.NewSource(time.Now().Unix()))
	SetCounter(&DefaultCounter{})

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "dummy-host"
	}

	acc := len(hostname) + base
	for i := range hostname {
		acc = acc + int(hostname[i])
	}

	hostID := pad(strconv.FormatInt(int64(os.Getpid()), base), 2)
	host := pad(strconv.FormatInt(int64(acc), base), 2)
	fingerprint = hostID + host
}

func SetRandomSource(src rand.Source) {
	SetRandom(rand.New(src))
}

func SetRandom(rnd *rand.Rand) {
	mutex.Lock()
	random = rnd
	mutex.Unlock()
}

func SetCounter(cnt Counter) {
	mutex.Lock()
	counter = cnt
	mutex.Unlock()
}

func New() string {
	// Global random generation functions from the math/rand package use a global
	// locked source, custom Rand objects need to be manually synchronized to avoid
	// race conditions.

	mutex.Lock()
	randomInt1 := int64(random.Int31n(discreteValues))
	randomInt2 := int64(random.Int31n(discreteValues))
	mutex.Unlock()

	return assembleCUID(randomInt1, randomInt2)
}

func NewCrypto(reader io.Reader) (string, error) {
	r1, err := cryptoRand.Int(reader, big.NewInt(int64(discreteValues)))
	if err != nil {
		return "", err
	}

	r2, err := cryptoRand.Int(reader, big.NewInt(int64(discreteValues)))
	if err != nil {
		return "", err
	}

	cuid := assembleCUID(r1.Int64(), r2.Int64())

	return cuid, nil
}

func Slug() string {
	timestamp := strconv.FormatInt(makeTimestamp(), base)
	counter := strconv.FormatInt(int64(counter.Next()), base)

	mutex.Lock()
	randomStr := strconv.FormatInt(int64(random.Int31n(discreteValues)), base)
	mutex.Unlock()

	timestampBlock := timestamp[len(timestamp)-2:]
	printBlock := fingerprint[0:1] + fingerprint[len(fingerprint)-1:]
	var counterBlock string
	var randomBlock string

	if len(counter) < 4 {
		counterBlock = counter
	} else {
		counterBlock = counter[len(counter)-4:]
	}

	if len(randomStr) < 4 {
		randomBlock = randomStr
	} else {
		randomBlock = randomStr[len(randomStr)-4:]
	}

	return timestampBlock + counterBlock + printBlock + randomBlock
}

func IsCuid(c string) error {
	if !format.MatchString(c) {
		return errors.New("Incorrect format")
	}
	return nil
}

func IsSlug(s string) error {
	if len(s) < 6 || len(s) > 12 {
		return errors.New("Incorrect format")
	}
	return nil
}

// Utility functions

func assembleCUID(randomInt1, randomInt2 int64) string {
	timestampBlock := strconv.FormatInt(makeTimestamp(), base)
	counterBlock := pad(strconv.FormatInt(int64(counter.Next()), base), blockSize)
	randomBlock1 := pad(strconv.FormatInt(randomInt1, base), blockSize)
	randomBlock2 := pad(strconv.FormatInt(randomInt2, base), blockSize)

	return "c" + timestampBlock + counterBlock + fingerprint + randomBlock1 + randomBlock2
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
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

// Default counter implementation

type Counter interface {
	Next() int32
}

type DefaultCounter struct {
	count int32
	mutex sync.Mutex
}

func (c *DefaultCounter) Next() int32 {
	c.mutex.Lock()

	counterValue := c.count

	c.count = c.count + 1
	if c.count >= discreteValues {
		c.count = 0
	}

	c.mutex.Unlock()

	return counterValue
}
