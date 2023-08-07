package block

import (
	"log"
	"time"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// ==============================
// Blockchain Neighbor Management Methods
// ==============================

// SetNeighbors discovers and sets the neighbors for the blockchain instance.
func (bc *Blockchain) SetNeighbors() {
	bc.neighbors = utils.FindNeighbors(
		utils.GetHost(), bc.port,
		NEIGHBOR_IP_RANGE_START, NEIGHBOR_IP_RANGE_END,
		BLOCKCHAIN_PORT_RANGE_START, BLOCKCHAIN_PORT_RANGE_END)
	log.Printf("%v", bc.neighbors)
}

// SyncNeighbors synchronizes the neighbors ensuring thread safety.
func (bc *Blockchain) SyncNeighbors() {
	bc.muxNeighbors.Lock()
	defer bc.muxNeighbors.Unlock()
	bc.SetNeighbors()
}

// StartSyncNeighbors initiates the synchronization process and schedules it to run periodically.
func (bc *Blockchain) StartSyncNeighbors() {
	bc.SyncNeighbors()
	_ = time.AfterFunc(time.Second*BLOCKCHIN_NEIGHBOR_SYNC_TIME_SEC, bc.StartSyncNeighbors)
}
