package controller

import (
	"net/http"
	"testing"
)

func TestCreateProjectHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := CreateProjectHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListProjectsHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
			if err := ListProjectsHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ListProjectsHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
