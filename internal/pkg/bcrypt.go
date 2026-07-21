package pkg

type BcryptHasher interface {
	CompareHashAndPassword(hashed string, literal string) error
	Hash(literal string) ([]byte, error)
}
