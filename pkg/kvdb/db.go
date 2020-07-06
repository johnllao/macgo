package kvdb

import (
	"bytes"
	"os"
	"os/signal"
	"syscall"

	"github.com/boltdb/bolt"
)

type DB struct {
	db *bolt.DB
}

func OpenDB(path string) (*DB, error) {
	var err error
	var bdb *bolt.DB
	bdb, err =  bolt.Open(path, 0664, nil)
	if err != nil {
		return nil, err
	}

	var sigc = make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		bdb.Close()
		os.Exit(0)
	}()

	var db = &DB{
		db: bdb,
	}

	return db, nil
}

func (d *DB) Close() {
	d.Close()
}

func (d *DB) Save(bucket string, key string, value []byte) error {

	return d.db.Update(func (tx *bolt.Tx) error {
		var err error

		var bucketb = []byte(bucket)
		var b *bolt.Bucket
		b, err = tx.CreateBucketIfNotExists(bucketb)
		if err != nil {
			return err
		}

		var keyb = []byte(key)
		err = b.Put(keyb, value)
		if err != nil {
			return err
		}

		return nil
	})

}

func (d *DB) Get(bucket string, key string) ([]byte, error) {

	var err error
	var value []byte
	err = d.db.View(func (tx *bolt.Tx) error {

		var bucketb = []byte(bucket)
		var b = tx.Bucket(bucketb)

		var keyb = []byte(key)
		value = b.Get(keyb)

		return nil
	})

	return value, err
}

func (d *DB) GetAll(bucket string) ([]string, [][]byte, error) {

	var err error
	var keys = make([]string, 0)
	var values = make([][]byte, 0)
	err = d.db.View(func(tx *bolt.Tx) error {

		var bucketb = []byte(bucket)
		var b = tx.Bucket(bucketb)

		var c = b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			keys = append(keys, string(k))
			values = append(values, v)
		}

		return nil
	})

	return keys, values, err
}

func (d *DB) GetByPrefix(bucket string, prefix string) ([]string, [][]byte, error) {

	var err error
	var keys = make([]string, 0)
	var values = make([][]byte, 0)
	var prefixb = []byte(prefix)
	err = d.db.View(func(tx *bolt.Tx) error {

		var bucketb = []byte(bucket)
		var b = tx.Bucket(bucketb)

		var c = b.Cursor()
		for k, v := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, prefixb); k, v = c.Next() {
			keys = append(keys, string(k))
			values = append(values, v)
		}

		return nil
	})

	return keys, values, err
}

func (d *DB) GetByFilter(bucket string, fn func(string, []byte) bool) ([]string, [][]byte, error) {

	var err error
	var keys = make([]string, 0)
	var values = make([][]byte, 0)
	err = d.db.View(func(tx *bolt.Tx) error {

		var bucketb = []byte(bucket)
		var b = tx.Bucket(bucketb)

		var c = b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if !fn(string(k), v) {
				continue
			}
			keys = append(keys, string(k))
			values = append(values, v)
		}

		return nil
	})

	return keys, values, err
}

func (d *DB) GetByFilterAndPrefix(bucket string, prefix string, fn func(string, []byte) bool) ([]string, [][]byte, error) {

	var err error
	var keys = make([]string, 0)
	var values = make([][]byte, 0)
	var prefixb = []byte(prefix)
	err = d.db.View(func(tx *bolt.Tx) error {

		var bucketb = []byte(bucket)
		var b = tx.Bucket(bucketb)

		var c = b.Cursor()
		for k, v := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, prefixb); k, v = c.Next() {
			if !fn(string(k), v) {
				continue
			}

			keys = append(keys, string(k))
			values = append(values, v)
		}

		return nil
	})

	return keys, values, err
}



