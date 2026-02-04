package service

import (
	"context"
	"database/sql"
	db "todo_list_roadmap/db/genarated"
	"todo_list_roadmap/dto"
	"todo_list_roadmap/handle/response"
	"todo_list_roadmap/util"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TodoService struct {
	sqlDB *sql.DB
}

func (s *TodoService) Save(ctx context.Context, req dto.TodoCreateRequest) (dto.TodoCreateResponse, error) {
	session := db.New(s.sqlDB)
	ID := uuid.New().String()
	_, err := session.CreateTodo(
		ctx,
		db.CreateTodoParams{
			ID:          ID,
			Title:       req.Title,
			Description: sql.NullString{Valid: false},
		},
	)
	if err != nil {
		zap.L().Error("Failed Save Todo", zap.Error(err))
		return dto.TodoCreateResponse{}, response.ErrInternalServer
	}
	return dto.TodoCreateResponse{
		ID:          ID,
		Title:       req.Title,
		Description: req.Description,
	}, nil
}

func (s *TodoService) GetTodos(ctx context.Context, req dto.TodoListRequest) ([]dto.TodoDetailResponse, int64, error) {
	session := db.New(s.sqlDB)

	total, err := session.CountTodo(ctx)
	if err != nil {
		zap.L().Error("Failed CountTodo", zap.Error(err))
		return nil, 0, response.ErrInternalServer
	}

	todos, err := session.GetListTodo(
		ctx, db.GetListTodoParams{Limit: int32(req.Limit), Offset: int32(util.Offset(req.PageRequest))},
	)
	if err != nil {
		zap.L().Error("Failed GetListTodo", zap.Error(err))
		return nil, 0, response.ErrInternalServer
	}

	return s.toTodoDetailResponse(todos), total, nil
}

func (s *TodoService) toTodoDetailResponse(todos []db.Todo) []dto.TodoDetailResponse {
	listResponse := make([]dto.TodoDetailResponse, len(todos))
	for index, todo := range todos {
		listResponse[index] = dto.TodoDetailResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description.String,
		}
	}
	return listResponse
}

func (s *TodoService) UpdateTodo(ctx context.Context, req dto.TodoUpdateRequest, ID string) (dto.TodoUpdateResponse, error) {
	session := db.New(s.sqlDB)
	description := sql.NullString{Valid: false}
	if req.Description == "" {
		description.String = req.Description
	}
	err := session.UpdateTodoByID(ctx, db.UpdateTodoByIDParams{
		Title:       req.Title,
		Description: description,
		ID:          ID,
	})
	if err != nil {
		zap.L().Error("Failed UpdateTodo", zap.Error(err), zap.String("ID", ID))
		return dto.TodoUpdateResponse{}, response.ErrInternalServer
	}
	return dto.TodoUpdateResponse{
		Title:       req.Title,
		Description: description.String,
		ID:          ID,
	}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, ID string) error {
	session := db.New(s.sqlDB)
	err := session.DeleteTodoByID(ctx, ID)
	if err != nil {
		zap.L().Error("Failed DeleteTodo", zap.Error(err), zap.String("ID", ID))
		return err
	}
	return nil
}

func NewTodoService(sqlDB *sql.DB) *TodoService {
	return &TodoService{sqlDB: sqlDB}
}
