package models

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
	user := User{Password: "testpassword"}
	err := user.HashPassword()
	assert.NoError(t, err)
	assert.NotEqual(t, "testpassword", user.Password)
}

func TestUser_GetAge(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name           string
		birthday       *time.Time
		expectedYears  int
		expectedMonths int
		expectedDays   int
	}{
		{
			name:           "No birthday",
			birthday:       nil,
			expectedYears:  0,
			expectedMonths: 0,
			expectedDays:   0,
		},
		{
			name:           "Same day birthday",
			birthday:       &now,
			expectedYears:  0,
			expectedMonths: 0,
			expectedDays:   0,
		},
		{
			name: "One year ago",
			birthday: func() *time.Time {
				past := now.AddDate(-1, 0, 0)
				return &past
			}(),
			expectedYears:  1,
			expectedMonths: 0,
			expectedDays:   0,
		},
		{
			name: "One month ago",
			birthday: func() *time.Time {
				past := now.AddDate(0, -1, 0)
				return &past
			}(),
			expectedYears:  0,
			expectedMonths: 1,
			expectedDays:   0,
		},
		{
			name: "One day ago",
			birthday: func() *time.Time {
				past := now.AddDate(0, 0, -1)
				return &past
			}(),
			expectedYears:  0,
			expectedMonths: 0,
			expectedDays:   1,
		},
		{
			name: "Future date with negative days",
			birthday: func() *time.Time {
				past := now.AddDate(-1, 1, -1)
				return &past
			}(),
			expectedYears:  0,
			expectedMonths: 11,
			expectedDays:   30, // Assuming 30 days in month
		},
		{
			name: "Future date with negative months",
			birthday: func() *time.Time {
				past := now.AddDate(-2, 1, 0)
				return &past
			}(),
			expectedYears:  0,
			expectedMonths: 11,
			expectedDays:   0, // Assuming 30 days in month
		},

		{
			name: "Full complex date",
			birthday: func() *time.Time {
				t, _ := time.Parse(time.RFC3339, "1995-03-27T00:00:00Z")
				return &t
			}(),
			expectedYears: func() int {
				t, _ := time.Parse(time.RFC3339, "1995-03-27T00:00:00Z")
				return time.Now().Year() - t.Year()
			}(),
			expectedMonths: func() int {
				t, _ := time.Parse(time.RFC3339, "1995-03-27T00:00:00Z")
				months := time.Now().Month() - t.Month()

				if time.Now().Day() < t.Day() {
					months--
				}
				if months < 0 {
					months += 12
				}
				return int(months)

			}(),
			expectedDays: func() int {
				t, _ := time.Parse(time.RFC3339, "1995-03-27T00:00:00Z")
				days := time.Now().Day() - t.Day()
				if days < 0 {
					days += time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC).Day()
				}
				return days
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := User{Birthday: tc.birthday}
			years, months, days := user.GetAge()
			assert.Equal(t, tc.expectedYears, years)
			assert.Equal(t, tc.expectedMonths, months)
			assert.Equal(t, tc.expectedDays, days)
		})
	}
}

func TestUser_CheckPassword(t *testing.T) {
	user := User{Password: "testpassword"}
	user.HashPassword()
	err := user.CheckPassword("testpassword")
	assert.NoError(t, err)

	err = user.CheckPassword("wrongpassword")
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid credentials")
}

func TestUser_HasUpcomingBirthday(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name     string
		birthday *time.Time
		expected bool
	}{
		{
			name:     "No birthday",
			birthday: nil,
			expected: false,
		},
		{
			name: "Birthday in 5 days",
			birthday: func() *time.Time {
				future := now.AddDate(0, 0, 5)
				return &future
			}(),
			expected: true,
		},
		{
			name: "Birthday in 15 days",
			birthday: func() *time.Time {
				future := now.AddDate(0, 0, 15)
				return &future
			}(),
			expected: false,
		},
		{
			name: "Birthday in 10 days",
			birthday: func() *time.Time {
				future := now.AddDate(0, 0, 10)
				return &future
			}(),
			expected: true,
		},
		{
			name: "Birthday in 9 days",
			birthday: func() *time.Time {
				future := now.AddDate(0, 0, 9)
				return &future
			}(),
			expected: true,
		},
		{
			name: "Birthday yesterday",
			birthday: func() *time.Time {
				past := now.AddDate(0, 0, -1)
				return &past
			}(),
			expected: false,
		},
		{
			name: "Birthday next year but same day",
			birthday: func() *time.Time {
				next := now.AddDate(1, 0, 0)
				return &next
			}(),
			expected: false,
		},
		{
			name: "Birthday was last year in 5 days",
			birthday: func() *time.Time {
				next := now.AddDate(-1, 0, 5)
				return &next
			}(),
			expected: true,
		},
		{
			name: "Birthday is next month",
			birthday: func() *time.Time {
				next := now.AddDate(0, 1, 0)
				return &next
			}(),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := User{Birthday: tc.birthday}
			result := user.HasUpcomingBirthday()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestUser_DefaultFields(t *testing.T) {
	user := User{}

	assert.Empty(t, user.Name)
	assert.Empty(t, user.Email)
	assert.Empty(t, user.Password)
	assert.Equal(t, uint(0), user.ID)
	assert.Nil(t, user.Birthday)
	assert.Equal(t, sql.NullTime{}.Valid, user.ActivatedAt.Valid)
	assert.True(t, user.CreatedAt.IsZero())
	assert.True(t, user.UpdatedAt.IsZero())
}
