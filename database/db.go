package database

import (
	"context"
	"fmt"
	"github.com/zhulida1234/go-rpc-service/common/retry"
	"github.com/zhulida1234/go-rpc-service/config"
	"gorm.io/gorm"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
)

type DB struct {
	gorm *gorm.DB

	Keys KeysDB
}

func NewDB(ctx context.Context, dbConf config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Name)
	if dbConf.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConf.Port)
	}
	if dbConf.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConf.User)
	}
	if dbConf.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConf.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm: gorm,
		Keys: NewKeysDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm: tx,
			Keys: NewKeysDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		fmt.Println("第一步执行")
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		fmt.Println("第二步执行")
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		fmt.Println("第三步执行")
		return nil
	})
	fmt.Println("执行不成功", err)
	return err
}
