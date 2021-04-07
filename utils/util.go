package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func GenerateRandomNumber(length int) int64 {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	num, _:=strconv.ParseInt(b.String(),10,64)
	return num
}
