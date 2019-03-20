package todos

import (
	"apiii/entities"
	"apiii/usecases/ports"
	"apiii/usecases/repositories"
	"apiii/infrastructure/db"
)

type TodoUsecase struct {
	TodoRepository repositories.TodoRepository
	DB *db.DB
}

func(usecase *TodoUsecase) GetAllTodo() ([]entities.Todo, error){
	todos, err := usecase.TodoRepository.FindAll(usecase.DB.GormDB)
	if err != nil {
		panic(err.Error())
	}
	return todos, nil
}

func (usecase *TodoUsecase) CreateTodo(input *ports.TodoInputPort) (*ports.TodoOutputPort, error) {
	todo := &entities.Todo{
		Title: input.Title,
		Text: input.Text,
	}

	if err := usecase.TodoRepository.Insert(usecase.DB.GormDB, todo); err != nil {
		return nil, err
	}

	return createOutputPort(todo), nil
}

func createOutputPort(todo *entities.Todo) *ports.TodoOutputPort {
	return &ports.TodoOutputPort{
		Title: todo.Title,
		Text: todo.Text,
	}
}
