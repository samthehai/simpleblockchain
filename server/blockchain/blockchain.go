package blockchain

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/samthehai/simpleblockchain/block"
	"github.com/samthehai/simpleblockchain/server"
	"github.com/samthehai/simpleblockchain/wallet"
)

type blockchainServer struct {
	port   uint64
	chain  *block.Blockchain
	server *http.Server
}

func NewBlockChainServer(port uint64) server.Server {
	return &blockchainServer{
		port: port,
	}
}

func (s *blockchainServer) getBlockchain() *block.Blockchain {
	if s.chain == nil {
		minerWallet := wallet.NewWallet()
		s.chain = block.NewBlockChain(minerWallet.BlockchainAddress)
		log.Printf("private_key %v", minerWallet.PrivateKey)
		log.Printf("publick_key %v", minerWallet.PublicKey)
		log.Printf("blockchain_address %v", minerWallet.BlockchainAddress)
	}
	return s.chain
}

func (s *blockchainServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bc := s.getBlockchain()
			m, _ := json.Marshal(bc)
			io.WriteString(w, string(m[:]))
		default:
			log.Printf("invalid http method")
		}
	})

	s.server = &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(int(s.port)),
		Handler: r,
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *blockchainServer) Close() error {
	if s.server == nil {
		return nil
	}

	return s.server.Close()
}
