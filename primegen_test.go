package primegen

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

type primegenTest struct {
	high     uint64
	checksum string
}

// Test checksums are listed at http://cr.yp.to/primegen.html
var golden = []primegenTest{
	{1000000, "c13929ee9d2aea8f83aa076236079e94"},
	{10000000, "60e34d268bad671a5f299e1ecc988ff6"},
	{100000000, "4e2b0027288a27e9c99699364877c9db"},
}

func TestSieve(t *testing.T) {
	h := md5.New()
	for _, test := range golden {
		h.Reset()
		Write(h, 1, test.high)
		sum := hex.EncodeToString(h.Sum(nil))
		if sum != test.checksum {
			t.Errorf("Invalid checksum for high value %d: Expected %s, got %s\n", test.high, test.checksum, sum)
		}
	}
}

func TestCount(t *testing.T) {
	sieve := New()
	count := sieve.Count(1000000000)
	if count != 50847534 {
		t.Errorf("Incorrect count %d of primes up to 1000000000, expected 50847534\n", count)
	}
}

func BenchmarkSieve(b *testing.B) {
	b.StopTimer()
	sieve := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sieve.Reset()
		for j := 0; j < 1000000; j++ {
			sieve.Next()
		}
	}
}
