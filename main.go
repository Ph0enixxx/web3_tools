package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	for {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		if strings.HasPrefix(address, "0x00000") {
			fmt.Printf("----------------------------\n")
			fmt.Printf("Time: %v\n", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("Private: %v\n", hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
			fmt.Printf("Public: %v\n", hexutil.Encode(publicKeyBytes)[4:])   // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
			fmt.Printf("Address: %v\n", address)                             // 0x96216849c49358B10257cb55b28eA603c874b05E
			fmt.Printf("----------------------------\n")
		}
	}
}
