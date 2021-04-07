package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	return strBuilder(chars,length).String()
}

func GenerateRandomNumber(length int) int64 {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	num, _:=strconv.ParseInt(strBuilder(chars,length).String(),10,64)
	return num
}

func strBuilder(chars []rune, length int) *strings.Builder {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return &b
}