package utils

import "encoding/base64"

func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func Base64UrlDecode(data string) ([]byte, error) {
	return base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(data)
}