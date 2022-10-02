package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gadoma/rafapi/internal/application"
	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gadoma/rafapi/test/mock"
)

func prepareRandomAffirmationServiceTest() (
	repositoryMock mock.RandomAffirmationRepository,
	randomAffirmationStubs []domain.RandomAffirmation,
	ctx context.Context) {
	repositoryMock = mock.RandomAffirmationRepository{}

	randomAffirmationStubs = append(
		randomAffirmationStubs,
		domain.RandomAffirmation{Text: "I am a stub."},
		domain.RandomAffirmation{Text: "I am another stub."},
	)

	ctx = context.Background()

	return
}

func TestRandomAffirmationServiceGetRandomAffirmation(t *testing.T) {
	repositoryMock, randomAffirmationStubs, ctx := prepareRandomAffirmationServiceTest()

	repositoryMock.GetRandomAffirmationsFn = func(ctx context.Context, categoryIds []int) ([]*domain.RandomAffirmation, error) {
		return []*domain.RandomAffirmation{&randomAffirmationStubs[0], &randomAffirmationStubs[1]}, nil
	}

	sut := application.NewRandomAffirmationService(&repositoryMock)

	result, err := sut.GetRandomAffirmation(ctx, []int{1})

	if err != nil {
		t.Errorf("error=%q, want nil", err)
	}

	expected := randomAffirmationStubs[0].Text + " " + randomAffirmationStubs[1].Text

	if got, want := result.Text, expected; got != want {
		t.Errorf("result.Text=%v, want %v", got, want)
	}
}

func TestRandomAffirmationServiceGetRandomAffirmationError(t *testing.T) {
	repositoryMock, _, ctx := prepareRandomAffirmationServiceTest()

	repositoryMock.GetRandomAffirmationsFn = func(ctx context.Context, categoryIds []int) ([]*domain.RandomAffirmation, error) {
		return nil, errors.New("something went wrong")
	}

	sut := application.NewRandomAffirmationService(&repositoryMock)

	_, err := sut.GetRandomAffirmation(ctx, []int{1})

	if err == nil {
		t.Error("an error was expected")
	}
}
