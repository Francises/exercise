package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	_ "strconv"
	"strings"
	"sync"
	"time"
)

type UUID [16]byte

// Used in string method conversion
const dash byte = '-'

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}
func initClockSequence() {
	buf := make([]byte, 2)
	safeRandom(buf)
	clockSequence = binary.BigEndian.Uint16(buf)
}
func initHardwareAddr() {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if len(iface.HardwareAddr) >= 6 {
				copy(hardwareAddr[:], iface.HardwareAddr)
				return
			}
		}
	}

	// Initialize hardwareAddr randomly in case
	// of real network interfaces absence
	safeRandom(hardwareAddr[:])

	// Set multicast bit as recommended in RFC 4122
	hardwareAddr[0] |= 0x01
}
func initStorage() {
	initClockSequence()
	initHardwareAddr()
}

// Difference in 100-nanosecond intervals between
// UUID epoch (October 15, 1582) and Unix epoch (January 1, 1970).
const epochStart = 122192928000000000

// Returns difference in 100-nanosecond intervals between
// UUID epoch (October 15, 1582) and current time.
// This is default epoch calculation function.
func unixTimeFunc() uint64 {
	return epochStart + uint64(time.Now().UnixNano()/100)
}

// UUID v1/v2 storage.
var (
	storageMutex  sync.Mutex
	storageOnce   sync.Once
	epochFunc     = unixTimeFunc
	clockSequence uint16
	lastTime      uint64
	hardwareAddr  [6]byte
	posixUID      = uint32(os.Getuid())
	posixGID      = uint32(os.Getgid())
)

// Returns UUID v1/v2 storage state.
// Returns epoch timestamp, clock sequence, and hardware address.
func getStorage() (uint64, uint16, []byte) {
	storageOnce.Do(initStorage)

	storageMutex.Lock()
	defer storageMutex.Unlock()

	timeNow := epochFunc()
	// Clock changed backwards since last UUID generation.
	// Should increase clock sequence.
	if timeNow <= lastTime {
		clockSequence++
	}
	lastTime = timeNow

	return timeNow, clockSequence, hardwareAddr[:]
}

// SetVersion sets version bits.
func (u *UUID) SetVersion(v byte) {
	u[6] = (u[6] & 0x0f) | (v << 4)
}

// SetVariant sets variant bits as described in RFC 4122.
func (u *UUID) SetVariant() {
	u[8] = (u[8] & 0xbf) | 0x80
}
func NewV1() UUID {
	u := UUID{}//u:00000000-0000-0000-0000-000000000000
	timeNow, clockSeq, hardwareAddr := getStorage()

	binary.BigEndian.PutUint32(u[0:], uint32(timeNow))
	fmt.Println(u)
	binary.BigEndian.PutUint16(u[4:], uint16(timeNow>>32))
	fmt.Println(u)
	binary.BigEndian.PutUint16(u[6:], uint16(timeNow>>48))
	fmt.Println(u)
	binary.BigEndian.PutUint16(u[8:], clockSeq)
	fmt.Println(u)

	copy(u[10:], hardwareAddr)
	fmt.Println(u)

	u.SetVersion(1)
	fmt.Println(u)
	u.SetVariant()
	fmt.Println(u)

	return u
}

// Returns canonical string representation of UUID:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (u UUID) String() string {
	buf := make([]byte, 36)
	fmt.Println("_______________________")
	fmt.Println(buf)

	hex.Encode(buf[0:8], u[0:4])
	fmt.Println(buf)
	buf[8] = dash
	fmt.Println(buf)
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], u[10:])
	fmt.Println(buf)

	return string(buf)
}

// GetStrGUID 生成GUID
func GetStrGUID() string {
	u := NewV1().String()
	fmt.Println("))))))))))))))))))))))))))))))))))))))0")
	fmt.Println(u)
	u = strings.Replace(u, "-", "", -1)
	fmt.Println(u)
	return StrMd5([]byte(u))
}

// StrMd5 16进制MD5
func StrMd5(bt []byte) string {
	h := md5.New()
	fmt.Println("***********************************************")
	fmt.Println(h)
	h.Write(bt[:])
	fmt.Println(h)
	fmt.Println(h.Sum(nil))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	str := GetStrGUID()
	fmt.Println(str)
}
