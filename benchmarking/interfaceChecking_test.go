package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"case2", func(userID string) ([]Pet, error) {
			return nil, errors.New("invalid id: 3")
		}, "3", nil, "invalid id: 3"},
	}
	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets}
			petNames, err := l.GetPetNames(d.userID)
			if diff := cmp.Diff(petNames, d.petNames); diff != "" {
				t.Error(diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
