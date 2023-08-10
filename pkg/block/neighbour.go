package block

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// ==============================
// Blockchain Neighbor Management Methods
// ==============================

// SetNeighbors discovers and sets the neighbors for the blockchain instance.
func (bc *Blockchain) SetNeighbors() {
	if host := os.Getenv("MINER_HOST"); host != "" {
		bc.neighbors = []string{
			"http://" + host + "-1:5001",
			"http://" + host + "-2:5002",
			"http://" + host + "-3:5003",
		}
	} else {
		bc.neighbors = utils.FindNeighbors(
			utils.GetHost(), bc.port,
			NEIGHBOR_IP_RANGE_START, NEIGHBOR_IP_RANGE_END,
			BLOCKCHAIN_PORT_RANGE_START, BLOCKCHAIN_PORT_RANGE_END)
	}

	// Filter out the neighbors with the same port as the current instance
	bc.neighbors = filterOutSelfPort(bc.neighbors, strconv.Itoa(int(bc.port)))

	log.Printf("%v", bc.neighbors)
}

//* This is a debug method, until blockchain broadcasting is implemented
// filterOutSelfPort removes neighbors with the same port as the current instance
func filterOutSelfPort(neighbors []string, port string) []string {
	var filtered []string
	for _, neighbor := range neighbors {
		if !strings.HasSuffix(neighbor, ":"+port) {
			filtered = append(filtered, neighbor)
		}
	}
	return filtered
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
