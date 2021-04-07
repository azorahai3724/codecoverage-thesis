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

		want1      *mongo.Client
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := getDbClient()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getDbClient got1 = %v, want1: %v", got1, tt.want1)
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
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
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

		want1      *mongo.Collection
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := getDbCollection(tArgs.CollectionName, tArgs.DbName)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getDbCollection got1 = %v, want1: %v", got1, tt.want1)
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
