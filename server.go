package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nrotta/plainchain/blockchain"
)

var node *blockchain.Node

func addTransaction(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("sender")
	d := r.FormValue("receiver")
	v, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v is not a valid value for the tx: %s", v, err)))
		return
	}

	t := blockchain.NewTx([]byte(s), []byte(d), v)
	node.AddTx(t)
	m, err := json.MarshalIndent(t, "", "   ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Marshaling transaction %v failed: %s", t, err)))
		return
	}
	w.Write([]byte(m))
}

func getBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	h, err := strconv.ParseUint(v["height"], 10, 32)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v is not a valid height for the block: %s", h, err)))
		return
	}

	b := node.Blockchain.GetBlock(uint32(h))
	m, err := json.MarshalIndent(b, "", "   ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Marshaling block %v failed: %s", b, err)))
		return
	}
	w.Write([]byte(m))
}

func getTxsPool(w http.ResponseWriter, r *http.Request) {
	t := node.GetTxsPool()
	m, err := json.MarshalIndent(t, "", "   ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Marshaling txs pool %v failed: %s", t, err)))
		return
	}
	w.Write([]byte(m))
}

func getTx(w http.ResponseWriter, r *http.Request) {
	h := blockchain.Hash{}
	v := mux.Vars(r)["hash"]
	s, err := hex.DecodeString(v)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Decoding transaction hash %v failed: %s", v, err)))
		return
	}
	copy(h[:], s)

	t := node.Blockchain.GetTx(h)
	m, err := json.MarshalIndent(t, "", "   ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Marshaling transaction %v failed: %s", t, err)))
		return
	}
	w.Write([]byte(m))
}

func main() {
	host := flag.String("host", "localhost:3000", "name:port where to run the blockchain node")
	address := flag.String("address", "1AdgxM5BhcLyRz6qRn8QPPBGJFfcXD5oA6", "address to award the coinbase reward to")
	flag.Parse()

	node = blockchain.NewNode(blockchain.Address(*address))

	go node.Mine()

	r := mux.NewRouter()
	r.HandleFunc("/transactions", addTransaction).Methods("POST")
	r.HandleFunc("/blocks/{height}", getBlock).Methods("GET")
	r.HandleFunc("/txspool", getTxsPool).Methods("GET")
	r.HandleFunc("/transactions/{hash}", getTx).Methods("GET")
	http.Handle("/", r)

	fmt.Printf("Running blockchain node at: %s\n", *host)
	log.Fatal(http.ListenAndServe(*host, nil))
}
