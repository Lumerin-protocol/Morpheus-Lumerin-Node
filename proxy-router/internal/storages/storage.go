package storages

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	badger "github.com/dgraph-io/badger/v4"
)

type Storage struct {
	db *badger.DB
}

func NewStorage(log lib.ILogger, path string) *Storage {
	opts := badger.DefaultOptions(path)
	opts.Logger = NewBadgerLogger(log)

	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}
	return &Storage{db}
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	return valCopy, err
}

func (s *Storage) Set(key, val []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (s *Storage) Delete(key []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}
