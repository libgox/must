package must

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		err     error
		want    any
		wantErr bool
	}{
		{
			name: "NoError",
			val:  "val",
			err:  nil,
			want: "val",
		},
		{
			name:    "Error",
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() {
					Must(tt.val, tt.err)
				}, "expected panic, but did not panic")
			} else {
				got := Must(tt.val, tt.err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMust2(t *testing.T) {
	tests := []struct {
		name    string
		val1    any
		val2    any
		err     error
		want1   any
		want2   any
		wantErr bool
	}{
		{
			name:  "NoError",
			val1:  "val1",
			val2:  "val2",
			err:   nil,
			want1: "val1",
			want2: "val2",
		},
		{
			name:    "Error",
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() {
					Must2(tt.val1, tt.val2, tt.err)
				}, "expected panic, but did not panic")
			} else {
				got1, got2 := Must2(tt.val1, tt.val2, tt.err)
				assert.Equal(t, tt.want1, got1)
				assert.Equal(t, tt.want2, got2)
			}
		})
	}
}

func TestMust3(t *testing.T) {
	tests := []struct {
		name    string
		val1    any
		val2    any
		val3    any
		err     error
		want1   any
		want2   any
		want3   any
		wantErr bool
	}{
		{
			name:  "NoError",
			val1:  "val1",
			val2:  "val2",
			val3:  "val3",
			err:   nil,
			want1: "val1",
			want2: "val2",
			want3: "val3",
		},
		{
			name:    "Error",
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() {
					Must3(tt.val1, tt.val2, tt.val3, tt.err)
				}, "expected panic, but did not panic")
			} else {
				got1, got2, got3 := Must3(tt.val1, tt.val2, tt.val3, tt.err)
				assert.Equal(t, tt.want1, got1)
				assert.Equal(t, tt.want2, got2)
				assert.Equal(t, tt.want3, got3)
			}
		})
	}
}
