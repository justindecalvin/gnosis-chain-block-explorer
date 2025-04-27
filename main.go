package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to the Gnosis Chain RPC
    client, err := ethclient.Dial("https://rpc.gnosischain.com")
    if err != nil {
        log.Fatalf("Failed to connect to Gnosis Chain RPC: %v", err)
    }
    fmt.Println("✅ Successfully connected to Gnosis Chain")

    // Fetch the latest block
    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatalf("Failed to fetch latest block header: %v", err)
    }

    fmt.Printf("Latest Block Number: %d\n", header.Number.Uint64())
    fmt.Printf("Block Hash: %s\n", header.Hash().Hex())
}
