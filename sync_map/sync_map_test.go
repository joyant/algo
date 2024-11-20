package sync_map

import (
    "sync"
    "testing"
)

func TestSyncMap(t *testing.T) {
    g := sync.WaitGroup{}
    data := map[string]string{
        "a": "1",
        "b": "2",
        "c": "3",
    }
    g.Add(len(data) * 2)
    sm := NewSyncMap()
    // write and read
    for k, v := range data {
        go func(k, v string) {
            defer g.Done()
            sm.Put(k, v)
        }(k, v)
    }
    for k := range data {
        go func(k string) {
            defer g.Done()
            sm.Get(k)
        }(k)
    }
    g.Wait()
    // verify value
    for k, v := range data {
        val := sm.Get(k)
        if v != val {
            t.Errorf("expected %s got %s", v, val)
        }
    }
}
