package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	ens "github.com/wealdtech/go-ens/v3"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("API_Key"))
	if err != nil {
		panic(err)
	}

	// Resolve a name to an address.
	domain := "ethrereum.eth"
	address, err := ens.Resolve(client, domain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

	// Reverse resolve an address to a name.
	reverse, err := ens.ReverseResolve(client, address)
	if err != nil {
		panic(err)
	}
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
}
