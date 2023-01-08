package generic_map

const (
	bucketSize = 8

	emptyRest  = 0
	emptyCell  = 1
	minTopHash = 3
)

type bucket[T any] struct {
	topHash [bucketSize]uint8

	keys   [bucketSize]string
	values [bucketSize]T

	overflow *bucket[T]
}

func (b *bucket[T]) Get(key string, topHash uint8) (T, bool) {
	for i := range b.topHash {
		if b.topHash[i] != topHash {
			// константа означает что все следующие ячейки пустые -- выходим из цикла.
			if b.topHash[i] == emptyRest {
				break
			}
			continue
		}

		if !isCellEmpty(b.topHash[i]) && b.keys[i] == key {
			return b.values[i], true
		}
	}

	// проверим бакет overflow
	if b.overflow != nil {
		return b.overflow.Get(key, topHash)
	}

	return *new(T), false
}

func (b *bucket[T]) Put(key string, topHash uint8, value T) (isAdded bool) {
	var insertIdx *int

	for i := range b.topHash {
		// сравниваем topHash а не ключ, т.к. это позволяет нам использовать флаги и это работает быстрее
		if b.topHash[i] != topHash {
			if b.topHash[i] == emptyRest {
				insertIdx = new(int)
				*insertIdx = i
				break
			}

			if insertIdx == nil && isCellEmpty(b.topHash[i]) {
				insertIdx = new(int)
				*insertIdx = i
			}
			continue
		}

		// topHash значения не уникальны, по этому при совпадении проверяем дополнительно и сам ключ
		if b.keys[i] != key {
			continue
		}

		b.values[i] = value
		return false
	}

	// проверяем overflow бакет.
	if insertIdx == nil || b.overflow != nil {
		if b.overflow == nil {
			b.overflow = &bucket[T]{}
		}

		return b.overflow.Put(key, topHash, value)
	}

	// сохраняем ключ, значение и topHash
	b.keys[*insertIdx] = key
	b.values[*insertIdx] = value
	b.topHash[*insertIdx] = topHash

	// возвращаем признак успеха. по нему калькулируем количество элементов в мапе.
	return true
}

func isCellEmpty(val uint8) bool {
	return val <= emptyCell
}
