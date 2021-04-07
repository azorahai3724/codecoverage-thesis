package main

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_newApp(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			newApp(tArgs.w, tArgs.r)

		})
	}
}

func Test_parseCoverageFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      float64
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := parseCoverageFile(tArgs.s)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseCoverageFile got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("parseCoverageFile error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_enableCORS(t *testing.T) {
	type args struct {
		w *http.ResponseWriter
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			enableCORS(tArgs.w)

		})
	}
}

func Test_getOneApp(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			getOneApp(tArgs.w, tArgs.r)

		})
	}
}

func Test_getAllApps(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			getAllApps(tArgs.w, tArgs.r)

		})
	}
}
