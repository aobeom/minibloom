package minibloom

import (
	"log"

	"github.com/damnever/bitarray"
	"github.com/spaolacci/murmur3"
)

// MiniBloom
type MiniBloom struct {
	bitmap     *bitarray.BitArray
	size       int
	hashCounts int
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
