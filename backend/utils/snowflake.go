package utils

import (
    "sync"
    "time"
)

type Snowflake struct {
    mu            sync.Mutex
    lastTimestamp int64
    sequence      int64
}

func NewSnowflake() *Snowflake {
    return &Snowflake{}
}

func (s *Snowflake) GenerateID() int64 {
    s.mu.Lock()
    defer s.mu.Unlock()

    timestamp := time.Now().UnixNano() / 1000000

    if timestamp == s.lastTimestamp {
        s.sequence = (s.sequence + 1) & 4095
        if s.sequence == 0 {
            for timestamp <= s.lastTimestamp {
                timestamp = time.Now().UnixNano() / 1000000
            }
        }
    } else {
        s.sequence = 0
    }

    s.lastTimestamp = timestamp

    id := (timestamp << 22) | (1 << 17) | (1 << 12) | s.sequence

    return id
}