package rand_test

import (
	"testing"
	"time"

	"github.com/goexl/gox/rand"
)

func TestTime(t *testing.T) {
	tests := []struct {
		from time.Time
		to time.Time
	}{
		{from: time.Now(), to: time.Now().Add(time.Hour)},
	}

	for _, test := range tests {
		got := rand.New().Time().Between(test.from, test.to).Build().Generate()
		if got.Before(test.from) || got.After(test.to) {
			t.Fatalf("期望：从%v到%v，实际：%v", test.from, test.to, got)
		}
	}
}
