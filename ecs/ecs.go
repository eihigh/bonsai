package ecs

import "sync/atomic"

type Entity interface {
	Delete()
	Deleted() bool
}

type DeletedFlag struct {
	deleted bool
}

func (d *DeletedFlag) Delete() { d.deleted = true }

func (d *DeletedFlag) Deleted() bool { return d.deleted }

type DeletedFlagAtomic struct {
	deleted atomic.Bool
}

func (d *DeletedFlagAtomic) Delete() { d.deleted.Store(true) }

func (d *DeletedFlagAtomic) Deleted() bool { return d.deleted.Load() }

type slice[E Entity] []E

func Slice[E Entity](s []E) *slice[E] { return &s }

func (s *slice[E]) SweepAll(yield func(int, E) bool) bool {
	cont := true
	j := 0
	for i, x := range *s {
		if !x.Deleted() {
			(*s)[j] = x
			j++
			if yield != nil && cont {
				cont = yield(i, x)
			}
		}
	}
	*s = (*s)[:j]
	return true
}

type map_[K comparable, V Entity] map[K]V

func Map[K comparable, V Entity](m map[K]V) map_[K, V] { return m }

func (m map_[K, V]) SweepAll(yield func(K, V) bool) bool {
	cont := true
	for k, v := range m {
		if v.Deleted() {
			delete(m, k)
		} else if yield != nil && cont {
			cont = yield(k, v)
		}
	}
	return true
}
