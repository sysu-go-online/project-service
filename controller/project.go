package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sysu-go-online/public-service/tools"

	"github.com/gorilla/mux"

	projectModel "github.com/sysu-go-online/project-service/model"
	userModel "github.com/sysu-go-online/user-service/model"
)

// ProjectController is controller for user
type ProjectController struct {
	projectModel.Project
	userModel.User
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	GitPath     string `json:"git_path"`
	IsClone     bool   `json:"isClone"`
}

// ListProjectsResponse is response for list projects
type ListProjectsResponse struct {
	Name string `json:"name"`
}

// CreateProjectHandler create project
// TODO: Check if the same name exists
func CreateProjectHandler(w http.ResponseWriter, r *http.Request) error {
	// Check token
	if ok, err := tools.CheckJWT(r.Header.Get("Authorization"), AuthRedisClient); !(ok && err == nil) {
		w.WriteHeader(401)
		return nil
	}
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	project := ProjectController{}
	if err := json.Unmarshal(body, &project); err != nil {
		return err
	}
	// TODO: check content
	project.Project.Name = project.Name
	project.Project.Description = project.Description
	project.Project.GitPath = project.GitPath
	project.Project.Path = project.Path

	session := MysqlEngine.NewSession()
	project.User.Username = mux.Vars(r)["username"]
	has, err := project.User.GetWithUsername(session)
	if err != nil {
		session.Rollback()
		return err
	}
	if !has {
		w.WriteHeader(401)
		return nil
	}
	project.Project.UserID = project.User.ID
	affected, err := project.Project.Insert(session)
	if err != nil {
		session.Rollback()
		return err
	}
	// create project root
	err = project.Project.CreateProjectRoot(project.User.Username)
	if err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	// clone from git path
	// TODO: use mq to decrease waitting time
	if project.IsClone {
		err := project.Project.CloneFromGitPath(project.User.Username)
		if err != nil {
			fmt.Println(err)
		}
	}
	if affected == 0 {
		w.WriteHeader(400)
		return nil
	}
	return nil
}

// ListProjectsHandler list projects
func ListProjectsHandler(w http.ResponseWriter, r *http.Request) error {
	project := ProjectController{}
	session := MysqlEngine.NewSession()
	project.User.Username = mux.Vars(r)["username"]
	has, err := project.User.GetWithUsername(session)
	if err != nil {
		session.Rollback()
		return err
	}
	if !has {
		w.WriteHeader(401)
		return nil
	}
	project.Project.UserID = project.User.ID
	ps, err := project.Project.GetWithUserID(session)
	if err != nil {
		session.Rollback()
		return err
	}
	if len(ps) == 0 {
		w.WriteHeader(204)
		return nil
	}

	ret := make([]ListProjectsResponse, 0)
	for _, v := range ps {
		tmp := ListProjectsResponse{v.Name}
		ret = append(ret, tmp)
	}
	body, err := json.Marshal(ret)
	if err != nil {
		return err
	}
	w.Write(body)
	return nil
}
