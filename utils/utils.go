package utils

import "slices"

func Contains(list []string, key string) bool {
	return slices.Contains(list, key)
}
