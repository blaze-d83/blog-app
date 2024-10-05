package utils

import (
	"fmt"
	"strconv"
)


func GetInt(param string) (uint, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value: %v", err)
	}
	return (uint(id)), nil
}
