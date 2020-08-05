package gox

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1288834974657

	// NodeBits holds the number of bits to use for Snowflake
	// Remember, you have a total 22 bits to share between Snowflake/Step
	NodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Snowflake/Step
	StepBits uint8 = 12

	mu        sync.Mutex
	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

const encodeBase32Map = "ybndrfg8ejkmcpqxot1uwisza345h769"

var decodeBase32Map [256]byte

const encodeBase58Map = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var decodeBase58Map [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid Id is provided.
type JSONSyntaxError struct{ original []byte }

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake Id %q", string(j.original))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32. This speeds up the process tremendously.
func init() {
	for i := 0; i < len(encodeBase58Map); i++ {
		decodeBase58Map[i] = 0xFF
	}

	for i := 0; i < len(encodeBase58Map); i++ {
		decodeBase58Map[encodeBase58Map[i]] = byte(i)
	}

	for i := 0; i < len(encodeBase32Map); i++ {
		decodeBase32Map[i] = 0xFF
	}

	for i := 0; i < len(encodeBase32Map); i++ {
		decodeBase32Map[encodeBase32Map[i]] = byte(i)
	}
}

// A Snowflake struct holds the basic information needed for a snowflake generator node
type Snowflake struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

// An Id is a custom type used for a snowflake Id.  This is used so we can attach methods onto the Id.
type Id int64

// NewSnowflake returns a new snowflake node that can be used to generate snowflake IDs
func NewSnowflake(node int64) (*Snowflake, error) {
	mu.Lock()
	nodeMax = -1 ^ (-1 << NodeBits)
	nodeMask = nodeMax << StepBits
	stepMask = -1 ^ (-1 << StepBits)
	timeShift = NodeBits + StepBits
	nodeShift = StepBits
	mu.Unlock()

	s := Snowflake{}
	s.node = node
	s.nodeMax = -1 ^ (-1 << NodeBits)
	s.nodeMask = s.nodeMax << StepBits
	s.stepMask = -1 ^ (-1 << StepBits)
	s.timeShift = NodeBits + StepBits
	s.nodeShift = StepBits

	if s.node < 0 || s.node > s.nodeMax {
		return nil, errors.New("Snowflake number must be between 0 and " + strconv.FormatInt(s.nodeMax, 10))
	}

	var curTime = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	s.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &s, nil
}

// Next creates and returns a unique snowflake Id
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node Id
func (s *Snowflake) Next() Id {
	s.mu.Lock()

	now := time.Since(s.epoch).Nanoseconds() / 1000000

	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if s.step == 0 {
			for now <= s.time {
				now = time.Since(s.epoch).Nanoseconds() / 1000000
			}
		}
	} else {
		s.step = 0
	}

	s.time = now

	r := Id((now)<<s.timeShift |
		(s.node << s.nodeShift) |
		(s.step),
	)

	s.mu.Unlock()

	return r
}

// NextId 下一个长整形Id
func (s *Snowflake) NextId() int64 {
	return s.Next().Int64()
}

// NextString 下一个字符串形式的Id
func (s *Snowflake) NextString() string {
	return s.Next().String()
}

// Int64 returns an int64 of the snowflake Id
func (f Id) Int64() int64 {
	return int64(f)
}

// ParseInt64 converts an int64 into a snowflake Id
func ParseInt64(id int64) Id {
	return Id(id)
}

// String returns a string of the snowflake Id
func (f Id) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// ParseString converts a string into a snowflake Id
func ParseString(id string) (Id, error) {
	i, err := strconv.ParseInt(id, 10, 64)

	return Id(i), err

}

// Base2 returns a string base2 of the snowflake Id
func (f Id) Base2() string {
	return strconv.FormatInt(int64(f), 2)
}

// ParseBase2 converts a Base2 string into a snowflake Id
func ParseBase2(id string) (Id, error) {
	i, err := strconv.ParseInt(id, 2, 64)

	return Id(i), err
}

// Base32 uses the z-base-32 character set but encodes and decodes similar to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when doing any interoperation.
func (f Id) Base32() string {
	if f < 32 {
		return string(encodeBase32Map[f])
	}

	b := make([]byte, 0, 12)
	for f >= 32 {
		b = append(b, encodeBase32Map[f%32])
		f /= 32
	}
	b = append(b, encodeBase32Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

// ParseBase32 parses a base32 []byte into a snowflake Id
// NOTE: There are many different base32 implementations so becareful when doing any interoperation.
func ParseBase32(b []byte) (Id, error) {
	var id int64

	for i := range b {
		if decodeBase32Map[b[i]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		id = id*32 + int64(decodeBase32Map[b[i]])
	}

	return Id(id), nil
}

// Base36 returns a base36 string of the snowflake Id
func (f Id) Base36() string {
	return strconv.FormatInt(int64(f), 36)
}

// ParseBase36 converts a Base36 string into a snowflake Id
func ParseBase36(id string) (Id, error) {
	i, err := strconv.ParseInt(id, 36, 64)

	return Id(i), err
}

// Base58 returns a base58 string of the snowflake Id
func (f Id) Base58() string {
	if f < 58 {
		return string(encodeBase58Map[f])
	}

	b := make([]byte, 0, 11)
	for f >= 58 {
		b = append(b, encodeBase58Map[f%58])
		f /= 58
	}
	b = append(b, encodeBase58Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

// ParseBase58 parses a base58 []byte into a snowflake Id
func ParseBase58(b []byte) (Id, error) {
	var id int64

	for i := range b {
		if decodeBase58Map[b[i]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		id = id*58 + int64(decodeBase58Map[b[i]])
	}

	return Id(id), nil
}

// Base64 returns a base64 string of the snowflake Id
func (f Id) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake Id
func ParseBase64(id string) (Id, error) {
	b, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return -1, err
	}

	return ParseBytes(b)

}

// Bytes returns a byte slice of the snowflake Id
func (f Id) Bytes() []byte {
	return []byte(f.String())
}

// ParseBytes converts a byte slice into a snowflake Id
func ParseBytes(id []byte) (Id, error) {
	i, err := strconv.ParseInt(string(id), 10, 64)

	return Id(i), err
}

// IntBytes returns an array of bytes of the snowflake Id, encoded as a big endian integer.
func (f Id) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))

	return b
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake Id
func ParseIntBytes(id [8]byte) Id {
	return Id(int64(binary.BigEndian.Uint64(id[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake Id time
func (f Id) Time() int64 {
	return (int64(f) >> timeShift) + Epoch
}

// Snowflake returns an int64 of the snowflake Id node number
func (f Id) Node() int64 {
	return int64(f) & nodeMask >> nodeShift
}

// Step returns an int64 of the snowflake step (or sequence) number
func (f Id) Step() int64 {
	return int64(f) & stepMask
}

// MarshalJSON returns a json byte array string of the snowflake Id.
func (f Id) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')

	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake Id into an Id type.
func (f *Id) UnmarshalJSON(b []byte) error {
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return JSONSyntaxError{b}
	}

	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = Id(i)

	return nil
}
