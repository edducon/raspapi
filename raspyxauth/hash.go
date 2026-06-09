package main

import (
"crypto/sha256"
"fmt"
"golang.org/x/crypto/bcrypt"
)

func main() {
password := "s5j6h6ers56sef6s23"
h := sha256.Sum256([]byte(password))
hash, err := bcrypt.GenerateFromPassword(h[:], bcrypt.DefaultCost)
if err != nil {
panic(err)
}
fmt.Println(string(hash))
}
