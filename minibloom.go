package minibloom

import (
	"log"
	"math"

	"github.com/damnever/bitarray"
	"github.com/spaolacci/murmur3"
)

// MiniBloom
type MiniBloom struct {
	bitmap     *bitarray.BitArray
	size       int
	hashCounts int
}

// Calculate
//	n: the expected number of items
//	p: accepted false positive rate (0.01 = 1%)
func Calculate(n int, p float64) (size, hashCounts int) {
	//	formulas:
	//	m: (size) bloom filter size (bit)
	//	k: (hashCounts) number of hash functions
	//
	//	m = -n*ln(p) / (ln(2)^2)
	//	k = m/n * ln(2)
	nf := float64(n)
	m := -nf * math.Log(p) / math.Pow(math.Log(2), 2)
	k := m / nf * math.Log(2)

	size = int(math.Ceil(m))
	hashCounts = int(math.Ceil(k))
	return
}

// New create a bitmap
func New(size int, counts int) *MiniBloom {
	return &MiniBloom{
		bitmap:     bitarray.New(size),
		size:       size,
		hashCounts: counts,
	}
}

// Add add data to bitmap
func (b *MiniBloom) Add(data []byte) {
	for i := 0; i < b.hashCounts; i++ {
		point := b.hashPoint(data, i)
		_, err := b.bitmap.Put(point, 1)
		if err != nil {
			log.Panic(err)
		}
	}
}

// In check data in bitmap
func (b *MiniBloom) In(data []byte) bool {
	for i := 0; i < b.hashCounts; i++ {
		point := b.hashPoint(data, i)
		p, err := b.bitmap.Get(point)
		if err != nil {
			log.Panic(err)
		}
		if p == 0 {
			return false
		}
	}
	return true
}

// hashPoint calc data points in bitmap
func (b *MiniBloom) hashPoint(data []byte, count int) int {
	seed := uint32(100 * count)
	hashCode := murmur3.Sum32WithSeed(data, seed)
	return int(hashCode) % b.size
}
