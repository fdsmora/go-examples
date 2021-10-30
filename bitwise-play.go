package main

import (
	"fmt"
)

var (
	MaxInt uint     = 0x8000 << (uint(^uint(0)>>63) << 4)
	/*
		0x8000 is hexadecimal 2^15
                uint depends on the machine, can be either 32 bits or 64 bits
                uint(0) is 32 0's or 64 0's
                ^uint(0) is the complement: 32 1's or 64 1's, i.e. 2^32 or 2^64
                (^uint(0))>>63 is 1
                ((^uint(0))>>63) << 4 is 16
                0x8000 << 16 = 2^31
	*/
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
}

