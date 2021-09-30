package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) int {
	mu.Lock()
	balance = balance + amount
	fmt.Printf("Depositing: %v\n", balance)
	mu.Unlock()
	return balance
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func Balance(amount int) int {
	mu.Lock()
	b := balance + amount
	mu.Unlock()
	return b
}

func main() {
	var dep = flag.Int("deposit", 0, "Enter amount to deposit")
	flag.Parse()
	b := Balance(1700)
	d := Deposit(*dep)
	enc := encrypt([]byte("{Th1s 7s 3ee Fl@g}"), "CoolCoolCoolclll")
	var needed_total = 164473
	f := decrypt(enc, "CoolCoolCoolclll")
	if d+b == needed_total {
		fmt.Printf("Total in you account: %v\n", b+d)
		fmt.Printf("Decrypted %s/n", f)
	} else {
		fmt.Println("Try harder....")
	}
}
