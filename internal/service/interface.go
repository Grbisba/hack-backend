package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	CreateConsumer(ctx context.Context, req dto.CreateConsumer) (*dto.Token, error)
	UpdateConsumerPassword(ctx context.Context, req dto.UpdatePassword) error
	DeleteConsumerByID(ctx context.Context, token string) error
	GetConsumerByID(ctx context.Context, token string) (*dto.Consumer, error)
	Login(ctx context.Context, req dto.Login) (*dto.Token, error)
	RefreshToken(_ context.Context, token string) (*dto.Token, error)

	SaveResult(ctx context.Context, token string, req dto.ResultCreation) error // todo
	GetResultByResolvedID(ctx context.Context, token string, resultID uuid.UUID) (*dto.Result, error)
	GetResultsByUserID(ctx context.Context, token string) ([]dto.Result, error)
	CreateResultByFirstTest(ctx context.Context, token string, req dto.Resolved) (*dto.Result, error)
	CreateResultBySecondTest(ctx context.Context, token string, req dto.Resolved) (*dto.Result, error)
	GetResultByTestID(ctx context.Context, token string, testID uuid.UUID) (*dto.Result, error) // todo
	GetAllResultsByAI(ctx context.Context, token string) ([]string, error)

	GetAllTests(_ context.Context, token string) ([]dto.Test, error)
	GetTestByID(_ context.Context, token string, testID uuid.UUID) (*dto.Test, error)

	CreateResolved(ctx context.Context, token string, req dto.ResolvedRequest) (*dto.Resolved, error)
	GetResolvedByID(ctx context.Context, token string, resolvedID uuid.UUID) (*dto.Resolved, error)

	GetQuestionsForThirdTest(_ context.Context, token string, questions dto.ThirdTestAnswers) (*dto.ThirdTestQuestions, error)
	CreateResultByThirdTest(ctx context.Context, token string, questions dto.ThirdTestAnswers) (*dto.Result, error)

	SetImageToResult(ctx context.Context, token string, image dto.ImageCreation) (bool, error) // todo

	SendResultOnEmail(_ context.Context, token string, req dto.MailSending) (bool, error)
}
