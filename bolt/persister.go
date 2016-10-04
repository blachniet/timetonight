package bolt

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/blachniet/timetonight"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	// Ensure Persister fulfills the interface
	_ timetonight.Persister = &Persister{}

	bktSettings      = []byte("Settings")
	keyTogglAPIToken = []byte("Toggl.APIToken")
	keyTimePerDay    = []byte("TimePerDay")
)

// Persister provides persistence methods saved via a bolt database.
type Persister struct {
	// Path to the bolt database file
	Path string

	db *bolt.DB
}

// NewPersister creates a new persister with the path to the bolt database file.
func NewPersister(path string) *Persister {
	p := &Persister{}
	p.Path = path
	return p
}

// Open opens the persister's underlying bolt database.
func (p *Persister) Open() error {
	db, err := bolt.Open(p.Path, 0600, nil)
	if err != nil {
		return errors.Wrap(err, "Err opening bolt database")
	}

	p.db = db

	// Create root buckets
	err = db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(bktSettings)
		return e
	})

	return errors.Wrap(err, "Error initializing buckets")
}

// Close closes the underlying bolt database connection
func (p *Persister) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *Persister) TogglAPIToken() (string, error) {
	var token string
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bktSettings)
		v := b.Get(keyTogglAPIToken)
		if v != nil {
			token = string(v)
		}
		return nil
	})
	return token, errors.Wrap(err, "Err getting toggl API token")
}

func (p *Persister) SetTogglAPIToken(token string) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bktSettings)
		return b.Put(keyTogglAPIToken, []byte(token))
	})
	return errors.Wrap(err, "Err setting toggl API token")
}

func (p *Persister) TimePerDay() (time.Duration, error) {
	var tpd time.Duration
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bktSettings)
		v := b.Get(keyTimePerDay)
		if v != nil {
			buf := bytes.NewReader(v)
			return binary.Read(buf, binary.LittleEndian, &tpd)
		}
		return nil
	})
	return tpd, errors.Wrap(err, "Err getting toggl API token")
}

func (p *Persister) SetTimePerDay(dur time.Duration) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bktSettings)
		return b.Put(keyTimePerDay, durtob(dur))
	})
	return errors.Wrap(err, "Err setting time per day")
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return b
}

func i64tob(v int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return b
}

func durtob(v time.Duration) []byte {
	return i64tob(int64(v))
}
