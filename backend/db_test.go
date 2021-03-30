package main

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func Test_getDbClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *mongo.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDbClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("getDbClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDbClient() = %v, want %v", got, tt.want)
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
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDbConnection(tt.args.ctx, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("testDbConnection() error = %v, wantErr %v", err, tt.wantErr)
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
		name    string
		args    args
		want    *mongo.Collection
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDbCollection(tt.args.CollectionName, tt.args.DbName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDbCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDbCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
