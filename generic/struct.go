package generic

import (
    "cmp"
    "sync"
)

// ref https://segmentfault.com/a/1190000041634906

func Div[T int | float64](a, b T) T {
    return a / b
}

func Add[T cmp.Ordered](a, b T) T {
    return a + b
}

type SyncMap[K comparable, V comparable] struct {
    mux sync.RWMutex
    m   map[K]V
}

func (s *SyncMap[K, V]) Get(key K) (val V) {
    s.mux.RLock()
    defer s.mux.RUnlock()
    return s.m[key]
}

func (s *SyncMap[K, V]) Set(key K, val V) {
    s.mux.Lock()
    defer s.mux.Unlock()
    s.m[key] = val
}
