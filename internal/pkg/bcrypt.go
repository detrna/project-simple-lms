package pkg

type BcryptHasher interface {
	CompareHashAndPassword(string, string) error
}
