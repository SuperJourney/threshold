package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLastMonthTime(t *testing.T) {
	tests := []struct {
		now   time.Time
		name  string
		want  int64
		want1 int64
	}{
		{
			now:   time.Date(2022, 12, 2, 22, 0, 0, 0, time.Local),
			name:  "",
			want:  time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
			want1: time.Date(2022, 12, 31, 23, 59, 59, 0, time.Local).Unix(),
		},
		{
			now:   time.Date(2022, 11, 2, 22, 0, 0, 0, time.Local),
			name:  "",
			want:  time.Date(2022, 11, 1, 0, 0, 0, 0, time.Local).Unix(),
			want1: time.Date(2022, 11, 30, 23, 59, 59, 0, time.Local).Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, x1 := GetLastMonthTime(tt.now)
			assert.Equal(t, tt.want, x)
			assert.Equal(t, tt.want1, x1)
		})
	}
}

func TestGetLastWeekTime(t *testing.T) {
	type args struct {
		now time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		{
			args: args{
				now: time.Date(2022, 6, 19, 4, 0, 0, 0, time.Local),
			},
			name:  "test1",
			want:  time.Date(2022, 6, 13, 0, 0, 0, 0, time.Local).Unix(),
			want1: time.Date(2022, 6, 19, 23, 59, 59, 0, time.Local).Unix(),
		},

		{
			args: args{
				now: time.Date(2024, 1, 29, 4, 0, 0, 0, time.Local),
			},
			name:  "test2",
			want:  time.Date(2024, 1, 29, 0, 0, 0, 0, time.Local).Unix(),
			want1: time.Date(2024, 2, 4, 23, 59, 59, 0, time.Local).Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, x1 := GetLastWeekTime(tt.args.now)
			assert.Equal(t, tt.want, x)
			assert.Equal(t, tt.want1, x1)
		})
	}
}
