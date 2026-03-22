package frequency

import (
	"testing"
)

func TestGetByteFrequency(t *testing.T) {
	tbl := GetByteFrequency(&[]byte{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})
	for i := 0; i < 4; i++ {
		if tbl[byte(i)] != uint64(i) {
			t.Errorf("expected %d, instead got %d", i, tbl[byte(i)])
		}

	}
}
