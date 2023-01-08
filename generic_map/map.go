package generic_map

import (
	"github.com/dgryski/go-wyhash"
	"math/rand"
)

const (
	loadFactorNum = 13
	loadFactorDen = 2

	ptrSize = 4 << (^uintptr(0) >> 63)
)

type (
	HashMap[T any] interface {
		Get(key string) T
		Get2(key string) (T, bool)
		Put(key string, value T)
		Delete(key string)
		Range(f func(k string, v T) bool)
		Len() int
	}

	hmap[T any] struct {
		len  int
		b    uint8
		seed uint64

		buckets []bucket[T]
	}
)

func (h hmap[T]) Get(key string) T {
	topHash, targetBucket := h.locateBucket(key)

	v, _ := h.buckets[targetBucket].Get(key, topHash)
	return v
}

func (h hmap[T]) Get2(key string) (T, bool) {
	topHash, targetBucket := h.locateBucket(key)

	return h.buckets[targetBucket].Get(key, topHash)
}

func (h hmap[T]) Put(key string, value T) {
	tophash, targetBucket := h.locateBucket(key)

	if h.buckets[targetBucket].Put(key, tophash, value) {
		h.len++
	}
}

func (h hmap[T]) Delete(key string) {
	//TODO implement me
	panic("implement me")
}

func (h hmap[T]) Range(f func(k string, v T) bool) {
	//TODO implement me
	panic("implement me")
}

func (h hmap[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (h hmap[T]) locateBucket(key string) (tophash uint8, targetBucket uint64) {
	hash := hash(key, h.seed)
	tophash = topHash(hash)
	mask := bucketMask(h.b)

	targetBucket = hash & mask

	return tophash, targetBucket
}

func New[T any](size int) HashMap[T] {
	h := new(hmap[T])

	h.seed = generateSeed()

	b := uint8(0)

	for overLoadFactor(size, b) {
		b++
	}

	h.b = b
	h.buckets = make([]bucket[T], bucketsNum(h.b))

	return h
}

func hash(key string, seed uint64) uint64 {
	return wyhash.Hash([]byte(key), seed)
}

func topHash(val uint64) uint8 {
	tophash := uint8(val >> (ptrSize*8 - 8))
	if tophash < minTopHash {
		tophash += minTopHash
	}
	return tophash
}

func bucketMask(b uint8) uint64 {
	return bucketsNum(b) - 1
}

func overLoadFactor(size int, b uint8) bool {
	return size > bucketSize && uint64(size) > loadFactorNum*(bucketsNum(b)/loadFactorDen)
}

func bucketsNum(b uint8) uint64 {
	return 1 << b
}

func generateSeed() uint64 {
	return rand.Uint64()
}
