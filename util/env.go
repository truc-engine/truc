package util

import "os"

func GetEnv(key ...string) string {
	res := ""
	for _, k := range key {
		res = os.Getenv(k)
		if res != "" {
			return res
		}
	}
	if res == "" && len(key) > 1 {
		return key[len(key)-1]
	}
	return res
}
