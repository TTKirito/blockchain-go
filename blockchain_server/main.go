package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/blockchain-go/block"
	"github.com/blockchain-go/wallet"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

func init() {
	log.SetPrefix("Blockchain: ")
}

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minersWallet.BlockchainAddress())

	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func HelloWord(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	e := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil)
	fmt.Println(e)
	log.Fatal(e)
}

// fix bad getwate sudo ufw allow 3000
func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
