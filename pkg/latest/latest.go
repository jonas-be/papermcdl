package latest

import (
	"strconv"
	"strings"
)

type NoStableItemFoundError struct {
	msg string
}

func (e NoStableItemFoundError) Error() string {
	return e.msg
}

// GetLatestItem items should have the latest item as first index
func GetLatestItem(items []string) (int, string, error) {
	for i, item := range items {
		if !strings.Contains(strings.ToLower(item), "pre") && !strings.Contains(strings.ToLower(item), "snapshot") {
			return i, item, nil
		}
	}
	return 0, "", NoStableItemFoundError{}
}

// Like getLatestItem but for snapshots
func GetLatestItemSnapshot(items []string) (int, string, error) {
	for i, item := range items {
		if strings.Contains(strings.ToLower(item), "snapshot") || strings.Contains(strings.ToLower(item), "pre") {
			return i, item, nil
		}
	}
	return 0, "", NoStableItemFoundError{}
}

func ConvertIntArrayToStringArray(intArr []int) []string {
	strArr := make([]string, len(intArr))
	for i, v := range intArr {
		strArr[i] = strconv.Itoa(v)
	}
	return strArr
}
