package timex

import "time"

// Zone returns the offset of the local zone from UTC, as a Duration.
//
// It returns 0 if the local zone is UTC.
func Zone() Duration {
	_, s := time.Now().Zone()
	return Second * Duration(s)
}
