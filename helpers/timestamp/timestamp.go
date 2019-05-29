package helpers

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time.Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	f, _ := strconv.ParseFloat(string(b), 64)
	sec, _ := math.Modf(f)

	t.Time = time.Unix(int64(sec), 0)

	return nil
}