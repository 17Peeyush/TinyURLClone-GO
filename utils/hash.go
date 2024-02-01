package utils

import (
	"hash/fnv"
	"math"
	"math/rand"
)
const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LinkHash(link string) string {
	// Create a new FNV-1a hash object
	hashObj := fnv.New32()

	// Write the bytes of the input string to the hash object
	hashObj.Write([]byte(link))
	hashNum := int(math.Abs((float64(hashObj.Sum32()))))
	//fmt.Printf("the hashNum: %d",hashNum)
	var hashVal []rune
	for i := 0; i < 7; i++ {
		// Ensure a non-negative index
		randNum := int(rand.Int31n(64))
		if randNum < 2{
			i--
			continue
		}
		idx := base62Digits[(hashNum)%randNum]
		hashVal = append(hashVal, rune(idx))
	}

	//return string(strconv.FormatFloat(),'f',0,64)
	return string(hashVal)
}