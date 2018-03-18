package rlog

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/apex/log/handlers/discard"
)

func BenchmarkNoCtx(b *testing.B) {
	buf := bytes.Buffer{}
	std.Handler = NewHandler(&buf)

	for i := 0; i < b.N; i++ {
		Info("test message")
	}
}

func BenchmarkDiscard(b *testing.B) {
	std.Handler = discard.Default

	for i := 0; i < b.N; i++ {
		Info("test message")
	}
}

var errExample = errors.New("fail")

type user struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var _jane = user{
	Name:      "Jane Doe",
	Email:     "jane@test.com",
	CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
}

func BenchmarkWithFields(b *testing.B) {
	std.Handler = discard.Default
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			WithFields(Fields{
				"int":               1,
				"int64":             int64(1),
				"float":             3.0,
				"string":            "four!",
				"bool":              true,
				"time":              time.Unix(0, 0),
				"error":             errExample.Error(),
				"duration":          time.Second,
				"user-defined type": _jane,
				"another string":    "done!",
			}).Info("test message")
		}
	})
}
