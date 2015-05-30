package raftkv

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// KVstore interface has get and set function to be implemented
type KVstore interface {
	Get(key []byte) ([]byte, error)
	Put(key, val []byte) error
	Delete(key []byte) error
}

// LeveldbKvstore implements KVstore interface
type LeveldbKvstore struct {
	db *leveldb.DB
}

// NewLevelDB returns a LeveldbKvstore
func NewLevelDB(filename string) (*LeveldbKvstore, error) {
	// TODO: add path
	ldb, err := leveldb.OpenFile(filename, nil)
	if err != nil {
		return nil, err
	}
	kvs := &LeveldbKvstore{
		db: ldb,
	}
	return kvs, nil
}

// Get implements KVstore's Get function to get a value by key.
// It returns ErrNotFound if the goleveldb does not contains the key.
func (l *LeveldbKvstore) Get(key []byte) ([]byte, error) {
	return l.db.Get(key, nil)
}

// Put implements KVstore's Put function to put a key-value pair.
// It It overwrites any previous value for that key
func (l *LeveldbKvstore) Put(key, val []byte) error {
	return l.db.Put(key, val, nil)
}

// Delete implements KVstore's delete function. It deletes a key
func (l *LeveldbKvstore) Delete(key []byte) error {
	return l.db.Delete(key, nil)
}