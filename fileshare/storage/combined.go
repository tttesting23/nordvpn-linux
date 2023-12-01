package storage

import (
	"errors"
	"log"
	"os"

	"github.com/NordSecurity/nordvpn-linux/fileshare"
	"github.com/NordSecurity/nordvpn-linux/fileshare/pb"
)

// Combined combines transfers from two storages
// Originally we had our own storage implementation in JSON file. Later libDrop introduced an
// integrated storage solution, so we migrated to that. But to not lose transfer history when
// updating the app, we still load transfers from the original file storage.
type Combined struct {
	json    fileshare.Storage
	libdrop fileshare.Storage
}

func NewCombined(storagePath string, fileshare fileshare.Fileshare) *Combined {
	return &Combined{NewJsonFile(storagePath), NewLibdrop(fileshare)}
}

func (c *Combined) Load() (map[string]*pb.Transfer, error) {
	libdropTransfers, err := c.libdrop.Load()
	if err != nil {
		return nil, err
	}

	jsonTransfers, err := c.json.Load()
	if err == nil {
		for key, value := range jsonTransfers {
			libdropTransfers[key] = value
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		log.Printf("json history file corrupted: %s", err)
	}

	return libdropTransfers, nil
}
