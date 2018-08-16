package main

import (
	"flag"
	"fmt"

	"github.com/skycoin/skycoin/src/cipher"
	bip39 "github.com/skycoin/skycoin/src/cipher/go-bip39"
)

func main() {
	n := flag.Int64("n", 1, "number of address need to generate")
	entropy := flag.Int("e", 128, "bip39 entropy, can be 128, 256")
	flag.Parse()

	for i := int64(0); i < *n; i++ {
		seed, err := bip39Seed(*entropy)
		if err != nil {
			panic(err)
		}
		_, seckey := cipher.GenerateDeterministicKeyPair([]byte(seed))
		addr := cipher.AddressFromSecKey(seckey)
		fmt.Printf("\"%s\" \"%s\"\n", addr, seed)
	}
}

func bip39Seed(n int) (string, error) {
	entropy, err := bip39.NewEntropy(n)
	if err != nil {
		return "", fmt.Errorf("generate bip39 entropy failed, err:%v", err)
	}

	seed, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("generate bip39 seed failed, err:%v", err)
	}

	return seed, nil
}
