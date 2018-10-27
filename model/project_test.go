package model

import (
	"reflect"
	"testing"

	"github.com/go-xorm/xorm"
)

func TestProject_TableName(t *testing.T) {
	tests := []struct {
		name string
		p    Project
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.TableName(); got != tt.want {
				t.Errorf("Project.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_Insert(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Insert(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Project.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_CreateProjectRoot(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.CreateProjectRoot(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Project.CreateProjectRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProject_GetWithUserID(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		want    []Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetWithUserID(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.GetWithUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Project.GetWithUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_GetWithUserIDAndName(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetWithUserIDAndName(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.GetWithUserIDAndName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Project.GetWithUserIDAndName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_GetWithID(t *testing.T) {
	type args struct {
		session *xorm.Session
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetWithID(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("Project.GetWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Project.GetWithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_CloneFromGitPath(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		p       *Project
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.CloneFromGitPath(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Project.CloneFromGitPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
