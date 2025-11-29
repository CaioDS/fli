package context

import (
	"time"

	bolt "go.etcd.io/bbolt"
	boltErrors "go.etcd.io/bbolt/errors"
)

type LocalDbContext struct {
	DB *bolt.DB
}

func CreateLocalDBContext(source string) *LocalDbContext {
	db, err := bolt.Open(source, 0666, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		panic("Failed to create local db context")
	}

	return &LocalDbContext{
		DB: db,
	}
}

func (db *LocalDbContext) Close() {
	db.DB.Close()
}

func (d *LocalDbContext) CreateBucket(name string) error {
	return d.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

// Put stores a key/value pair
func (d *LocalDbContext) Put(bucket string, key, value []byte) error {
	return d.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return boltErrors.ErrBucketNotFound
		}
		
		return b.Put(key, value)
	})
}

// Get retrieves a value
func (d *LocalDbContext) Get(bucket string, key []byte) ([]byte, error) {
	var val []byte
	err := d.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return boltErrors.ErrBucketNotFound
		}
		
		v := b.Get(key)
		if v != nil {
			val = append([]byte(nil), v...)
		}
		
		return nil
	})
	return val, err
}

// Delete removes a key
func (d *LocalDbContext) Delete(bucket string, key []byte) error {
	return d.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return boltErrors.ErrBucketNotFound
		}
		
		return b.Delete(key)
	})
}
