package set

type ISet interface {
	Collections() []int
	Add(elem int) bool
	Contains(elem int) bool
	Remove(elem int) bool
	Size() int
	Union(set *Set) *Set
	Intersection(set *Set) *Set
	Difference(set *Set) *Set
	Subset(set *Set) bool
}

var _ ISet = (*Set)(nil)

type Set struct {
	collections []int
}

func NewSet() *Set {
	return new(Set)
}

func (s *Set) Collections() []int {
	return s.collections
}

func (s *Set) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s.collections = append(s.collections, elem)
	return true
}

func (s *Set) Contains(elem int) bool {
	for _, value := range s.collections {
		if elem == value {
			return true
		}
	}
	return false
}

func (s *Set) Remove(elem int) bool {
	if !s.Contains(elem) {
		return false
	}

	for key, value := range s.collections {
		if elem == value {
			s.collections[key] = s.collections[s.Size()-1]
			s.collections = s.collections[:s.Size()-1]
			return true
		}
	}
	return false
}

func (s *Set) Size() int {
	return len(s.collections)
}

func (s *Set) Union(set *Set) *Set {
	newSet := Set{}

	newSet.collections = append(newSet.collections, s.collections...)
	newSet.collections = append(newSet.collections, set.collections...)
	return &newSet
}

func (s *Set) Intersection(set *Set) *Set {
	//TODO implement me
	panic("implement me")
}

func (s *Set) Difference(set *Set) *Set {
	//TODO implement me
	panic("implement me")
}

func (s *Set) Subset(set *Set) bool {
	//TODO implement me
	panic("implement me")
}
