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

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", addTx).Methods("POST")
	r.HandleFunc("/blocks/{height}", getBlock).Methods("GET")
	r.HandleFunc("/txspool", getTxsPool).Methods("GET")
	r.HandleFunc("/transactions/{hash}", getTx).Methods("GET")
	return r
}

func createResponse(w http.ResponseWriter, data interface{}, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(m))
}

func addTx(w http.ResponseWriter, r *http.Request) {
	type tx struct {
		TxsIn []struct {
			PrevOutput string `json:"prev_tx_hash"`
			Index      int    `json:"index"`
		} `json:"txs_in"`
		TxsOut []struct {
			Address string `json:"recipient"`
			Value   int64  `json:"value"`
		} `json:"txs_out"`
	}
	t := &tx{}
	json.NewDecoder(r.Body).Decode(t)
	defer r.Body.Close()

	var txsIns []*blockchain.TxIn
	var txsOuts []*blockchain.TxOut

	for _, txIn := range t.TxsIn {
		b, err := hex.DecodeString(txIn.PrevOutput)
		if err != nil {
			createResponse(w, nil, http.StatusBadRequest, fmt.Errorf("Error decoding Prev Tx Output for Tx %v: %v", txIn, err))
			return
		}

		h := blockchain.Hash{}
		copy(h[:], b)
		txsIns = append(txsIns, &blockchain.TxIn{Index: txIn.Index, PrevOutput: h})
	}

	for _, txOut := range t.TxsOut {
		txsOuts = append(txsOuts, &blockchain.TxOut{Address: blockchain.Address(txOut.Address), Value: txOut.Value})
	}

	n := blockchain.NewTx(txsIns, txsOuts)
	node.AddTx(n)
	createResponse(w, n, http.StatusCreated, nil)
}

func getBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	h, err := strconv.ParseUint(v["height"], 10, 32)
	if err != nil {
		createResponse(w, nil, http.StatusBadRequest, fmt.Errorf("%v is not a valid height: %v", h, err))
		return
	}

	b := node.Blockchain.GetBlock(uint32(h))
	createResponse(w, b, http.StatusOK, nil)
}

func getTxsPool(w http.ResponseWriter, r *http.Request) {
	t := node.GetTxsPool()
	createResponse(w, t, http.StatusOK, nil)
}

func getTx(w http.ResponseWriter, r *http.Request) {
	h := mux.Vars(r)["hash"]
	b, err := hex.DecodeString(h)
	if err != nil {
		createResponse(w, nil, http.StatusBadRequest, fmt.Errorf("Error decoding transaction %s: %v", h, err))
		return
	}

	hash := blockchain.Hash{}
	copy(hash[:], b)
	t := node.Blockchain.GetTx(hash)
	createResponse(w, t, http.StatusOK, nil)
}

func main() {
	host := flag.String("host", "localhost:3000", "name:port where to run the blockchain node")
	address := flag.String("address", "1AdgxM5BhcLyRz6qRn8QPPBGJFfcXD5oA6", "address to award the coinbase reward to")
	flag.Parse()

	node = blockchain.NewNode(blockchain.Address(*address))
	go node.Mine()

	fmt.Printf("Running blockchain node at: %s\n", *host)
	http.Handle("/", newRouter())
	log.Fatal(http.ListenAndServe(*host, nil))
}
