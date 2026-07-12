package pkg

type Packages struct {
	JWTProvider  JWTProvider
	BcryptHasher BcryptHasher
	Logger       Logger
	// Pagination   Pagination
	RedisClient  RedisClient
	ResendClient ResendClient
}
