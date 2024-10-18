package utils

import (
	"fmt"
	"strconv"
	"time"
)


func GetInt(param string) (uint, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value: %v", err)
	}
	return (uint(id)), nil
}

func UintToString(i uint) string  {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatTime(t time.Time) string  {
	return t.Format("Jan 2, 2006 at 3:04pm")
}
