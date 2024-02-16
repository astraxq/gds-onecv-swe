package controller

import (
	"database/sql"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetConnection(t *testing.T) {
	validGinContext := &gin.Context{}
	mockDB := &sql.DB{}
	validGinContext.Set("db", mockDB)

	type args struct {
		c *gin.Context
	}

	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				c: validGinContext,
			},
			want:    mockDB,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				c: &gin.Context{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				c: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConnection(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

