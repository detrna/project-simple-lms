package template

type UseCase struct {
	repo IRepository
}

type IUseCase interface {
	Ping() (string, error)
}

func NewUseCase(repo IRepository) *UseCase {
	return (&UseCase{repo: repo})
}

func (usecase UseCase) Ping() (string, error) {
	result, err := usecase.repo.Ping()
	return result, err
}
