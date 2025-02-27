package common

import (
	"encoding/base32"
	"encoding/base64"
)

func EncodeToBase32(data ...string) string {
	var tmpStr string
	for _, v := range data {
		if len(tmpStr) == 0 {
			tmpStr = v
		} else {
			tmpStr = tmpStr + "@" + v
		}
	}
	base32String := base32.StdEncoding.EncodeToString([]byte(tmpStr))
	return base32String
}

func DecodeFromBase32(base32String string) string {
	tmpStr, _ := base32.StdEncoding.DecodeString(base32String)
	return string(tmpStr)
}

func EncodeToBase64(data ...string) string {
	var tmpStr string
	for _, v := range data {
		if len(tmpStr) == 0 {
			tmpStr = v
		} else {
			tmpStr = tmpStr + "@" + v
		}
	}
	base64String := base64.StdEncoding.EncodeToString([]byte(tmpStr))
	return base64String
}

func DecodeFromBase64(base64String string) string {
	tmpStr, _ := base64.StdEncoding.DecodeString(base64String)
	return string(tmpStr)
}
