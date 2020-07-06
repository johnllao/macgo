package storage

type MapStore map[[32]byte][]string

func (s MapStore) Exists(key [32]byte) bool {
	var ok bool
	_, ok = s[key]
	return ok
}

func (s MapStore) Add(key [32]byte, val string) {

}