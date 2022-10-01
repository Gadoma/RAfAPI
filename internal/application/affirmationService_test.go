package application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gadoma/rafapi/internal/application"
	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gadoma/rafapi/test/mock"
)

func prepareAffirmationServiceTest() (
	repositoryMock mock.AffirmationRepository,
	affirmationStub domain.Affirmation,
	affirmationUpdateStub domain.AffirmationUpdate,
	ctx context.Context) {
	repositoryMock = mock.AffirmationRepository{}

	affirmationStub = domain.Affirmation{
		Id:         1,
		CategoryId: 1,
		Text:       "I am a stub.",
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	affirmationUpdateStub = domain.AffirmationUpdate{
		CategoryId: 2,
		Text:       "I am another stub.",
	}

	ctx = context.Background()

	return
}

func TestAffirmationServiceGetAffirmations(t *testing.T) {
	repositoryMock, affirmationStub, _, ctx := prepareAffirmationServiceTest()

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
	repositoryMock, _, _, ctx := prepareAffirmationServiceTest()

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
	repositoryMock, affirmationStub, _, ctx := prepareAffirmationServiceTest()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id int) (*domain.Affirmation, error) {
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
	repositoryMock, affirmationStub, _, ctx := prepareAffirmationServiceTest()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id int) (*domain.Affirmation, error) {
		return nil, nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.GetAffirmation(ctx, affirmationStub.Id)

	if err != domain.ErrorResourceNotFound {
		t.Errorf("error=%q, want domain.ErrorResourceNotFound", err)
	}
}

func TestAffirmationServiceGetAffirmationError(t *testing.T) {
	repositoryMock, affirmationStub, _, ctx := prepareAffirmationServiceTest()

	repositoryMock.GetAffirmationFn = func(ctx context.Context, id int) (*domain.Affirmation, error) {
		return nil, errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.GetAffirmation(ctx, affirmationStub.Id)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceCreateAffirmation(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	repositoryMock.CreateAffirmationFn = func(ctx context.Context, au *domain.AffirmationUpdate) (int, error) {
		return 1, nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	result, err := sut.CreateAffirmation(ctx, &affirmationUpdateStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	if got, want := result, 1; got != want {
		t.Errorf("result=%v, want %v", got, want)
	}
}

func TestAffirmationServiceCreateAffirmationValidationError(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	affirmationUpdateStub.Text = ""

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.CreateAffirmation(ctx, &affirmationUpdateStub)

	if err != domain.ErrorAffirmationUpdateInvalidText {
		t.Errorf("error=%q, want domain.ErrorAffirmationUpdateInvalidText", err)
	}
}

func TestAffirmationServiceCreateAffirmationError(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	repositoryMock.CreateAffirmationFn = func(ctx context.Context, au *domain.AffirmationUpdate) (int, error) {
		return 0, errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	_, err := sut.CreateAffirmation(ctx, &affirmationUpdateStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceUpdateAffirmation(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	repositoryMock.UpdateAffirmationFn = func(ctx context.Context, id int, au *domain.AffirmationUpdate) error {
		return nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, 1, &affirmationUpdateStub)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestAffirmationServiceUpdateAffirmationValidationError(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	affirmationUpdateStub.Text = ""

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, 1, &affirmationUpdateStub)

	if err != domain.ErrorAffirmationUpdateInvalidText {
		t.Errorf("error=%q, want domain.ErrorAffirmationUpdateInvalidText", err)
	}
}

func TestAffirmationServiceUpdateAffirmationError(t *testing.T) {
	repositoryMock, _, affirmationUpdateStub, ctx := prepareAffirmationServiceTest()

	repositoryMock.UpdateAffirmationFn = func(ctx context.Context, id int, au *domain.AffirmationUpdate) error {
		return errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.UpdateAffirmation(ctx, 1, &affirmationUpdateStub)

	if err == nil {
		t.Error("an error was expected")
	}
}

func TestAffirmationServiceDeleteAffirmation(t *testing.T) {
	repositoryMock, _, _, ctx := prepareAffirmationServiceTest()

	repositoryMock.DeleteAffirmationFn = func(ctx context.Context, id int) error {
		return nil
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.DeleteAffirmation(ctx, 1)

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}
}

func TestAffirmationServiceDeleteAffirmationError(t *testing.T) {
	repositoryMock, _, _, ctx := prepareAffirmationServiceTest()

	repositoryMock.DeleteAffirmationFn = func(ctx context.Context, id int) error {
		return errors.New("something went wrong")
	}

	sut := application.NewAffirmationService(&repositoryMock)

	err := sut.DeleteAffirmation(ctx, 1)

	if err == nil {
		t.Error("an error was expected")
	}
}
