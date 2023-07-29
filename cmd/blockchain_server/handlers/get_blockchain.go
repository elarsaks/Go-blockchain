package handlers

// // Get the blockchain of the BlockchainServer
// func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
// 	bc, ok := cache["blockchain"]
// 	if !ok {
// 		minersWallet := wallet.NewWallet()
// 		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.GetPort())
// 		cache["blockchain"] = bc

// 		// Call RegisterNewWallet to register the provided wallet address
// 		success := bc.RegisterNewWallet(minersWallet.BlockchainAddress(), "Register Miner Wallet")
// 		if !success {
// 			log.Println("ERROR: Failed to register wallet")
// 			// TODO: Handle error
// 			return nil
// 		}

// 		// Setting the wallet in the BlockchainServer object
// 		bcs.Wallet = minersWallet

// 		log.Printf("privateKey %v", minersWallet.PrivateKeyStr())
// 		log.Printf("publicKey %v", minersWallet.PublicKeyStr())
// 		log.Printf("blockchainAddress %v", minersWallet.BlockchainAddress())
// 	}
// 	return bc
// }
