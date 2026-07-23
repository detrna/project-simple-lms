package pkg

type BcryptHasher interface {
	Compare(hashed string, literal string) error
	Hash(literal string) ([]byte, error)
}
