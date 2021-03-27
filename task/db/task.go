package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Connect(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: time.Second * 1})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

func CreateTask(t string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		return b.Put(key, []byte(t))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func ReadTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DoTask(id int) error {
	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(tasksBucket)

		b.Delete(itob(id))

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int {
	return int(binary.BigEndian.Uint64(v))
}
