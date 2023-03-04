package rand_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/goexl/gox/rand"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		from time.Duration
		to time.Duration
	}{
		{from: time.Second, to: time.Hour},
	}

	for _, test := range tests {
		got := rand.New().Duration().Between(test.from, test.to).Build().Generate()
		fmt.Println(got)
		if got < test.from || got > test.to {
			t.Fatalf("期望：从%v到%v，实际：%v", test.from, test.to, got)
		}
	}
}
