package models

import (
	"fmt"
	"net/url"
	"path/filepath"
	"sf/app/errors"
	"sf/app/streams"
	"sf/app/validations"
	"sf/utils"
	"strings"
)

type Repository struct {
	Workspace *Workspace
	URI       string
	Protocol  string
	Token     string
	GitRepo   *GitRepo
	Framers   []*Framer
	Shapers   []*Shaper
}

type RepositoryInspector struct {
	GitRepo *GitRepoInspector
}

type RepositoryReader struct {
	Workspace string
	URI       string
	Branch    string
	Branches  []string
	Shapers   []string
	Framers   []string
}

// Construction

func NewRepository(w *Workspace, uri string) (r *Repository) {
	r = &Repository{
		Workspace: w,
		URI:       uri,
	}
	return
}

func CreateRepository(w *Workspace, uri string, protocol string, username string, token string, st *streams.Stream) (r *Repository, vn *validations.Validation, err error) {
	r = NewRepository(w, uri)
	if err = r.Load("GitRepo"); err != nil {
		return
	}
	if vn = r.Validation(); vn.IsValid() {
		if r.IsExists() {
			err = errors.Errorf("repository %s already exists in workspace %s", uri, w.Name)
			return
		}
		r.Create(protocol, username, token, st)
	}
	return
}

func ResolveRepository(w *Workspace, uri string, loads ...string) (r *Repository, err error) {
	if w == nil {
		err = errors.Error("no workspace")
		return
	}
	if len(w.Repositories) == 0 {
		err = errors.Errorf("no repositories exist in workspace %s", w.Name)
		return
	}
	r = w.FindRepository(uri)
	if r == nil {
		err = errors.Errorf("repository %s does not exist in workspace %s", uri, w.Name)
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
	ri = &RepositoryInspector{
		GitRepo: gri,
	}
	return
}

// Read

func (r *Repository) Read() (rr *RepositoryReader, err error) {
	var branch string
	var branches []string
	if branch, err = r.GitRepo.Branch(); err != nil {
		return
	}
	if branches, err = r.GitRepo.Branches(); err != nil {
		return
	}
	rr = &RepositoryReader{
		Workspace: r.Workspace.Name,
		URI:       r.URI,
		Branch:    branch,
		Branches:  branches,
		Shapers:   r.GitRepo.ShaperNames(),
		Framers:   r.GitRepo.FramerNames(),
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

func (r *Repository) Validation() (vn *validations.Validation) {
	vn = validations.NewValidation()
	if r.URI == "" {
		vn.Add("URI", "must not be blank")
	}
	return
}

func (r *Repository) Create(protocol string, username string, password string, st *streams.Stream) {
	go r.GitRepo.clone(r.OriginURL(protocol), username, password, st)
}

func (r *Repository) Update(username string, password string, st *streams.Stream) {
	go r.GitRepo.pull(username, password, st)
}

func (r *Repository) Checkout(branch string) (err error) {
	err = r.GitRepo.checkout(branch)
	return
}

func (r *Repository) Delete() {
	utils.RemoveDir(r.directory())
	r.GitRepo.remove()
}

// URL

func (r *Repository) OriginURL(protocol string) (u string) {
	uUrl, _ := url.Parse("https://" + r.URI)
	host := uUrl.Host
	path := filepath.Join(strings.Split(uUrl.Path, "/")...)
	if protocol == "HTTPS" {
		u = fmt.Sprintf("https://%s/%s", host, path)
	} else {
		u = fmt.Sprintf("git@%s:%s.git", host, path)
	}
	return
}
