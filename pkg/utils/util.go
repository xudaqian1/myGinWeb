package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func GetRandomString(len int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<len;i++{
		result = append(result,bytes[r.Intn(62)])
	}
	return  string(result)
}