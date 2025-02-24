package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const DIFFICULTY = "000"

func mySHA256(s string) string {
	h := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}

func concatStringAndInt(s string, n int) string {
	return s + strconv.Itoa(n)
}

func calculateHash(s string, difficulty string) int {
	nonce := 0
	for {
		hash := mySHA256(concatStringAndInt(s, nonce))
		if strings.HasPrefix(hash, difficulty) {
			fmt.Println()
			fmt.Println("FIND !!!", hash, "for", s, "string and with", nonce, "nonce")
			return nonce
		}
		log.Println(nonce, ":", hash)
		nonce++
	}
}

func checkPoW(s string, nonce int) string {
	return mySHA256(concatStringAndInt(s, nonce))
}

func PoW() {
	s := "Exo"

	start := time.Now()
	nonce := calculateHash(s, DIFFICULTY)
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time : %s\n", elapsed)
	fmt.Println("Check :", s, "+", nonce, "=", checkPoW(s, nonce))
}
