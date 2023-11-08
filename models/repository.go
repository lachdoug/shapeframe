package models

import "sf/app"

type Repository struct {
	Workspace *Workspace
	Path      string
	GitRepo   *GitRepo
	Framers   []*Framer
	Shapers   []*Shaper
}

type RepositoryInspector struct {
	Path    string
	GitRepo *GitRepoInspector
}

// Construction

func NewRepository(w *Workspace, path string) (r *Repository) {
	r = &Repository{
		Workspace: w,
		Path:      path,
	}
	return
}

func CreateRepository(w *Workspace, uri string, ssh bool, st *Stream) (r *Repository, err error) {
	path := w.GitRepoDirectoryFor(uri)
	r = NewRepository(w, path)
	if err = r.Load("GitRepo"); err != nil {
		return
	}
	if r.IsExists() {
		err = app.Error("repository %s already exists in workspace %s", uri, w.Name)
		return
	}
	r.Create(uri, ssh, st)
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
	path := w.GitRepoDirectoryFor(uri)
	r = w.FindRepository(path)
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

func (r *Repository) Inspect() (ri *RepositoryInspector) {
	ri = &RepositoryInspector{
		Path:    r.Path,
		GitRepo: r.GitRepoInspect(),
	}
	return
}

func (r *Repository) GitRepoInspect() (gri *GitRepoInspector) {
	if r.GitRepo.isExists() {
		gri = r.GitRepo.Inspect()
	}
	return
}

// Data

func (r *Repository) IsExists() (is bool) {
	is = r.GitRepo.isExists()
	return
}

func (r *Repository) Load(loads ...string) (err error) {
	dl := NewRepositoryLoader(r, loads)
	err = dl.load()
	return
}

func (r *Repository) Create(uri string, ssh bool, st *Stream) {
	go r.GitRepo.clone(uri, ssh, st)
}

func (r *Repository) Update(st *Stream) {
	go r.GitRepo.pull(st)
}

func (r *Repository) Destroy() {
	r.GitRepo.remove()
}
