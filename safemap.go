package safemap

import (
    "fmt"
    "io"
    "sort"
    "sync"
)

type SafeMap struct {
    m map[string] string
    sync.RWMutex
}

func New() *SafeMap {
    return &SafeMap{
        m: make(map[string] string),
    }
}

func (s *SafeMap) Dump(w io.Writer) {
    s.Lock()
    defer s.Unlock()

    keys := make([]string, len(s.m))

    i := 0
    for k, _ := range s.m {
        keys[i] = k
        i = i + 1
    }

    sort.Strings(keys)

    for _, k := range keys {
        fmt.Fprintf(w, "%s = %s\n", k, s.m[k])
    }
}

func (s *SafeMap) Get(key string) string {
    s.Lock()
    defer s.Unlock()
    return s.m[key]
}

func (s *SafeMap) Set(key string, value string) {
    s.Lock()
    defer s.Unlock()
    s.m[key] = value
}

func (s *SafeMap) Test(key string) (string, bool) {
    s.Lock()
    defer s.Unlock()
    value, ok := s.m[key]
    return value, ok
}

