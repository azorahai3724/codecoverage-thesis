package main

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func Test_getDbClient(t *testing.T) {
	tests := []struct {
		name string

		want       *mongo.Client
		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name:    "False connection if db is not started",
			want:    nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDbClient()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDbClient got = %v, want: %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("getDbClient error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_testDbConnection(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *mongo.Client
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name: "Invalid connection if db is not started",
			args: func(t *testing.T) args {
				return args{
					ctx: nil,
					c:   &mongo.Client{},
				}
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := testDbConnection(tArgs.ctx, tArgs.c)

			if (err != nil) != tt.wantErr {
				t.Fatalf("testDbConnection error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_getDbCollection(t *testing.T) {
	type args struct {
		CollectionName string
		DbName         string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want       *mongo.Collection
		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name: "Invalid collection if db is not started",
			args: func(t *testing.T) args {
				return args{
					CollectionName: "",
					DbName:         "",
				}
			},
			want:    nil,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got, err := getDbCollection(tArgs.CollectionName, tArgs.DbName)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDbCollection got = %v, want: %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("getDbCollection error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
