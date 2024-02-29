package main

import (
	"fmt"

	"github.com/blockchain-go/utils"
)

func main() {
	fmt.Println(utils.FindNeighbors("0.0.0.0", 5000, 0, 3, 5000, 5003))
	fmt.Println(utils.GetHost())
}
