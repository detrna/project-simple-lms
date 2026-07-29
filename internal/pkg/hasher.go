package pkg

type Hasher interface {
	Compare(hashed string, literal string) error
	Hash(literal string) ([]byte, error)
}
