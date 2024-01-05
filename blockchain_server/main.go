package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

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

func HelloWord(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWord)
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
