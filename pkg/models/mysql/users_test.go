package mysql

import (
	"github.com/axxyhtrx/snippetbox/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestUserModelGet(t *testing.T) {
	// Skip the test if the `-short` flag is provided when running the test.
	// We'll talk more about this in a moment.
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}
	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Jones",
				Email:   "Jones@email.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()
			m := UserModel{db}
			user, err := m.Get(tt.userID)
			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}
