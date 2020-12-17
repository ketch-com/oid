package oid

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"strings"
	"time"
)

var zeros = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var zero, base, maxInt *big.Int

// NewOID creates a new OID
func NewOID() (string, error) {
	n, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return "", err
	}

	data := make([]byte, 8, 16)

	binary.BigEndian.PutUint64(data, uint64(time.Now().UnixNano()))
	data = append(data, n.Bytes()...)
	if len(data) < 16 {
		data = append(data, zeros[0:16-len(data)]...)
	}

	n = big.NewInt(0)
	n.SetBytes(data)

	buf := new(strings.Builder)
	m := big.NewInt(0)

	for n.Cmp(zero) != 0 {
		n.DivMod(n, base, m)
		z := m.Int64()
		buf.WriteString(alphabet[z:z+1])
	}

	return reverse(buf.String()), nil
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func init() {
	zero = big.NewInt(0)
	base = big.NewInt(58)
	maxInt = big.NewInt(0)
	maxInt.SetString("18446744073709551616", 10)
}
