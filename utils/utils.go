package utils

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func RandomString(letterBytes string) string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}


func Tofloat(s string) float32 {

	complieRegex := regexp.MustCompile("(.*?)万")
	s2 := complieRegex.FindString(s)
	s3 := strings.Split(s2, "万")
	i2, err := strconv.ParseFloat(s3[0], 32)
	if err != nil {
		log.Fatal(err)
	}

	return float32(i2)
}


//CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
