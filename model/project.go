package model

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-xorm/xorm"
	git "gopkg.in/src-d/go-git.v4"
)

// Project corresponds to project table in db
type Project struct {
	ID          int        `xorm:"pk autoincr 'id'"`
	UserID      int        `xorm:"'user_id'"`
	Name        string     `xorm:"notnull"`
	CreateTime  *time.Time `xorm:"created"`
	Description string
	GitPath     string
	Path        string
}

// TableName defines table name
func (p Project) TableName() string {
	return "project"
}

// Insert insert a project to db
func (p *Project) Insert(session *xorm.Session) (int, error) {
	affected, err := session.Insert(p)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(affected), nil
}

// CreateProjectRoot create project root in the user home
func (p *Project) CreateProjectRoot(username string) error {
	userHome := filepath.Join("/home", username, "projects")
	path := filepath.Join(userHome, p.Path, p.Name)
	return os.MkdirAll(path, os.ModeDir)
}

// GetWithUserID returns projects with given user id
func (p *Project) GetWithUserID(session *xorm.Session) ([]Project, error) {
	var ps []Project
	err := session.Where("user_id = ?", p.UserID).Find(&ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// GetWithUserIDAndName returns project with given user id and project name
func (p *Project) GetWithUserIDAndName(session *xorm.Session) (bool, error) {
	return session.Where("user_id = ?", p.UserID).And("name = ?", p.Name).Get(p)
}

// GetWithID returns project with given project id
func (p *Project) GetWithID(session *xorm.Session) (bool, error) {
	return session.Where("id = ?", p.ID).Get(p)
}

// CloneFromGitPath clone project form given git path
func (p *Project) CloneFromGitPath(username string) error {
	userHome := filepath.Join("/home", username, "projects")
	path := filepath.Join(userHome, p.Path, p.Name)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: p.GitPath,
	})
	return err
}
