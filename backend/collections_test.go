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
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newApp(tt.args.w, tt.args.r)
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

		want       float64
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		{
			name: "100 percent",
			args: func(t *testing.T) args {
				return args{s: `github.com/bitrise-io/stack-service/internal/stack/handler.go:207:	execute		100.0%
					total:									(statements)	100.0%`}
			},
			want:    100.0,
			wantErr: false,
			inspectErr: func(err error, t *testing.T) {
			},
		},
		{
			name: "Invalid input",
			args: func(t *testing.T) args {
				return args{s: `sehv 1111234%`}
			},
			want:    0.0,
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got, err := parseCoverageFile(tArgs.s)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCoverageFile got = %v, want: %v", got, tt.want)
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
