package sync_map

import "sync"

type SyncMap struct {
    mux sync.RWMutex
    m   map[string]string
}

func NewSyncMap() *SyncMap {
    return &SyncMap{
        m: make(map[string]string),
    }
}

func (s *SyncMap) Get(key string) string {
    s.mux.RLock()
    defer s.mux.RUnlock()
    return s.m[key]
}

func (s *SyncMap) Put(key, value string) {
    s.mux.Lock()
    defer s.mux.Unlock()
    s.m[key] = value
}
