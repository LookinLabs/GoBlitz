package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func generateAPIKey() string {
    key := make([]byte, 16)
    rand.Read(key)
    return hex.EncodeToString(key)
}

func main() {
    fmt.Println(generateAPIKey())
}