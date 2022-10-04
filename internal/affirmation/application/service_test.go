package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gadoma/rafapi/internal/affirmation/application"
	"github.com/gadoma/rafapi/internal/affirmation/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
	"github.com/gadoma/rafapi/test/mock"
	"github.com/oklog/ulid/v2"
)

func getAffirmationEntityStub() domain.Affirmation {
	return domain.Affirmation{
		Id:         ulid.Make(),
		CategoryId: ulid.Make(),
		Text:       "I am a stub.",
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}
}

func getCreateAffirmationCommandStub() domain.CreateAffirmationCommand {
	return domain.CreateAffirmationCommand{
		Id:         ulid.Make(),
		CategoryId: ulid.Make(),
		Text:       "I am a create stub.",
	}
}

func getUpdateAffirmationCommandStub() domain.UpdateAffirmationCommand {
	return domain.UpdateAffirmationCommand{
		CategoryId: ulid.Make(),
		Text:       "I am an update stub.",
	}
}

func TestAffirmationServiceGetAffirmations(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	affirmationStub := getAffirmationEntityStub()

	repositoryMock.GetAffirmationsFn = func(ctx context.Context) ([]*domain.Affirmation, int, error) {
		return []*domain.Affirmation{&affirmationStub}, 1, nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	result, n, err := sut.GetAffirmations(ctx)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result[0].Id, affirmationStub.Id; got != want {
		t.Errorf("result[0].Id=%v, want %v", got, want)
	} else if got, want := result[0].Text, affirmationStub.Text; got != want {
		t.Errorf("result[0].Text=%v, want %v", got, want)
	} else if got, want := n, 1; got != want {
		t.Errorf("count=%v, want %v", got, want)
	}
}

func TestAffirmationServiceGetAffirmationsError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()

	repositoryMock.GetAffirmationsFn = func(ctx context.Context) ([]*domain.Affirmation, int, error) {
		return nil, 0, errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, _, err := sut.GetAffirmations(ctx)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceGetAffirmation(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	affirmationStub := getAffirmationEntityStub()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
		return &affirmationStub, nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	result, err := sut.GetAffirmation(ctx, affirmationStub.Id)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result.Id, affirmationStub.Id; got != want {
		t.Errorf("result.Id=%v, want %v", got, want)
	} else if got, want := result.Text, affirmationStub.Text; got != want {
		t.Errorf("result.Text=%v, want %v", got, want)
	}
}

func TestAffirmationServiceGetAffirmationNotFound(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	affirmationStub := getAffirmationEntityStub()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
		return nil, nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.GetAffirmation(ctx, affirmationStub.Id)

	if !errors.Is(err, common.ErrorResourceNotFound) {
		t.Errorf("error=%q, want domain.ErrorResourceNotFound", err)
	}
}

func TestAffirmationServiceGetAffirmationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	affirmationStub := getAffirmationEntityStub()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
		return nil, errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.GetAffirmation(ctx, affirmationStub.Id)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceCreateAffirmation(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	createAffirmationStub := getCreateAffirmationCommandStub()

	repositoryMock.CreateAffirmationFn = func(ctx context.Context, cac *domain.CreateAffirmationCommand) error {
		return nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	result, err := sut.CreateAffirmation(ctx, &createAffirmationStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result, &createAffirmationStub.Id; got != want {
		t.Errorf("result=%v, want %v", got, want)
	}
}

func TestAffirmationServiceCreateAffirmationValidationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	createAffirmationStub := getCreateAffirmationCommandStub()

	createAffirmationStub.Text = ""

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.CreateAffirmation(ctx, &createAffirmationStub)

	if !errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidText) {
		t.Errorf("error=%q, want domain.ErrorCreateAffirmationCommandInvalidText", err)
	}
}

func TestAffirmationServiceCreateAffirmationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	createAffirmationStub := getCreateAffirmationCommandStub()

	repositoryMock.CreateAffirmationFn = func(ctx context.Context, cac *domain.CreateAffirmationCommand) error {
		return errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.CreateAffirmation(ctx, &createAffirmationStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceUpdateAffirmation(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	updateAffirmationStub := getUpdateAffirmationCommandStub()

	repositoryMock.UpdateAffirmationFn = func(ctx context.Context, id ulid.ULID, uac *domain.UpdateAffirmationCommand) error {
		return nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, ulid.Make(), &updateAffirmationStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestAffirmationServiceUpdateAffirmationValidationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	updateAffirmationStub := getUpdateAffirmationCommandStub()

	updateAffirmationStub.Text = ""

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, ulid.Make(), &updateAffirmationStub)

	if !errors.Is(err, domain.ErrorUpdateAffirmationCommandInvalidText) {
		t.Errorf("error=%q, want domain.ErrorAffirmationUpdateInvalidText", err)
	}
}

func TestAffirmationServiceUpdateAffirmationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()
	updateAffirmationStub := getUpdateAffirmationCommandStub()

	repositoryMock.UpdateAffirmationFn = func(ctx context.Context, id ulid.ULID, uac *domain.UpdateAffirmationCommand) error {
		return errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, ulid.Make(), &updateAffirmationStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceDeleteAffirmation(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()

	repositoryMock.DeleteAffirmationFn = func(ctx context.Context, id ulid.ULID) error {
		return nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.DeleteAffirmation(ctx, ulid.Make())

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestAffirmationServiceDeleteAffirmationError(t *testing.T) {
	repositoryMock := mock.AffirmationRepository{}
	ctx := context.Background()

	repositoryMock.DeleteAffirmationFn = func(ctx context.Context, id ulid.ULID) error {
		return errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.DeleteAffirmation(ctx, ulid.Make())

	if err == nil {
		t.Error("an error was expected")
	}
}
