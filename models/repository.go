package models

import (
	"fmt"
	"net/url"
	"path/filepath"
	"sf/app"
	"sf/utils"
	"strings"
)

type Repository struct {
	Workspace *Workspace
	URI       string
	Protocol  string
	GitRepo   *GitRepo
	Framers   []*Framer
	Shapers   []*Shaper
}

type RepositoryInspector struct {
	GitRepo *GitRepoInspector
}

// Construction

func NewRepository(w *Workspace, uri string, protocol string) (r *Repository) {
	r = &Repository{
		Workspace: w,
		URI:       uri,
		Protocol:  protocol,
	}
	return
}

func CreateRepository(w *Workspace, uri string, protocol string, st *utils.Stream) (r *Repository, vn *app.Validation, err error) {
	r = NewRepository(w, uri, protocol)
	if err = r.Load("GitRepo"); err != nil {
		return
	}
	if vn = r.Validation(); vn.IsValid() {
		if r.IsExists() {
			err = app.Error("repository %s already exists in workspace %s", uri, w.Name)
			return
		}
		r.Create(st)
	}
	return
}

func ResolveRepository(w *Workspace, uri string, loads ...string) (r *Repository, err error) {
	if w == nil {
		err = app.Error("no workspace")
		return
	}
	if len(w.Repositories) == 0 {
		err = app.Error("no repositories exist in workspace %s", w.Name)
		return
	}
	r = w.FindRepository(uri)
	if r == nil {
		err = app.Error("repository %s does not exist in workspace %s", uri, w.Name)
		return
	}
	if len(loads) > 0 {
		if err = r.Load(loads...); err != nil {
			return
		}
	}
	return
}

// Inspection

func (r *Repository) Inspect() (ri *RepositoryInspector, err error) {
	var gri *GitRepoInspector
	if gri, err = r.GitRepo.Inspect(); err != nil {
		return
	}
	fmt.Println("REPOSITORY URI", r.URI)
	ri = &RepositoryInspector{
		GitRepo: gri,
	}
	return
}

// Data

func (r *Repository) directory() (dirPath string) {
	dirPath = r.Workspace.GitRepoDirectoryFor(r.URI)
	return
}

func (r *Repository) IsExists() (is bool) {
	is = r.GitRepo.isExists()
	return
}

func (r *Repository) Load(loads ...string) (err error) {
	dl := NewRepositoryLoader(r, loads)
	err = dl.load()
	return
}

func (r *Repository) Validation() (vn *app.Validation) {
	vn = &app.Validation{}
	if r.URI == "" {
		vn.Add("URI", "must not be blank")
	}
	// if utils.IsGitRemote(r.OriginURL()) {
	// 	vn.Add("URI", "must be a URI for an accessible remote")
	// }
	return
}

func (r *Repository) Create(st *utils.Stream) {
	go r.GitRepo.clone(r.OriginURL(), st)
}

func (r *Repository) Update(st *utils.Stream) {
	go r.GitRepo.pull(st)
}

func (r *Repository) Destroy() {
	r.GitRepo.remove()
}

// URL

func (r *Repository) OriginURL() (u string) {
	uUrl, _ := url.Parse("https://" + r.URI)
	host := uUrl.Host
	path := filepath.Join(strings.Split(uUrl.Path, "/")...)
	if r.Protocol == "HTTPS" {
		u = fmt.Sprintf("https://%s/%s", host, path)
	} else {
		u = fmt.Sprintf("git@%s:%s.git", host, path)
	}
	return
}
