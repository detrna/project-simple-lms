package pkg

type BcryptHasher interface {
	CompareHashAndPassword(string, string) error
	Hash(string) ([]byte, error)
}
