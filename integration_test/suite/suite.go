package testsuite

import (
	"fmt"
	"log"

	"main/internal/app"
	"main/internal/config"
	"main/internal/infrastructure"
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Suite struct {
	Router *gin.Engine
	Infra  *pkg.Packages
	DB     *gorm.DB
	Config *config.Config
}

type IntegrationTest[T any] struct {
	Name               string
	Data               T
	ExpectedStatusCode int
	ExpectedResponse   any
}

func New() *Suite {
	fmt.Print("TEST")
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("couldn't load config")
	}

	infra, db, repo, err := infrastructure.Initialize(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err = TruncateDatabase(db); err != nil {
		log.Fatal("database error")
	}

	router := app.SetupRouter(cfg, infra, repo)

	return &Suite{
		Router: router,
		Infra:  infra,
		DB:     db,
		Config: cfg,
	}
}

func TruncateDatabase(db *gorm.DB) error {
	return db.Exec(`
        TRUNCATE TABLE
            enrollments,
            classes,
            courses,
            users
        RESTART IDENTITY
        CASCADE;
    `).Error
}
