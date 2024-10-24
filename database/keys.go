package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Keys struct {
	GUID       uuid.UUID `gorm:"primaryKey" json:"guid"`
	BusinessId string    `json:"business_id"`
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"public_key"`
	Timestamp  int64     `json:"timestamp"`
}

type KeysView interface {
	QueryKeysByBusId(string, uint64, uint64) ([]Keys, error)
}

type KeysDB interface {
	KeysView

	StoreKeys([]Keys, uint64) error
}

type addressDB struct {
	gorm *gorm.DB
}

func (db *addressDB) QueryKeysByBusId(s string, u uint64, u2 uint64) ([]Keys, error) {
	//TODO implement me
	panic("implement me")
}

func NewKeysDB(db *gorm.DB) KeysDB {
	return &addressDB{gorm: db}
}

func (db *addressDB) StoreKeys(keyList []Keys, keyLength uint64) error {
	result := db.gorm.CreateInBatches(&keyList, int(keyLength))
	return result.Error
}
