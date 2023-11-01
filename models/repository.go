package models

type Repository struct {
	Workspace *Workspace
	Path      string
	GitRepo   *GitRepo
}

type RepositoryInspector struct {
	Path    string
	GitRepo *GitRepoInspector
}

// Construction

func RepositoryNew(w *Workspace, path string) (r *Repository) {
	if w == nil {
		panic("Repository Workspace is <nil>")
	}
	r = &Repository{
		Workspace: w,
		Path:      path,
		GitRepo:   GitRepoNew(w, path),
	}
	return
}

// Inspection

func (r *Repository) Inspect() (ri *RepositoryInspector, err error) {
	var gri *GitRepoInspector
	if gri, err = r.GitRepoInspect(); err != nil {
		return
	}
	ri = &RepositoryInspector{
		Path:    r.Path,
		GitRepo: gri,
	}
	return
}

func (r *Repository) GitRepoInspect() (gri *GitRepoInspector, err error) {
	if r.GitRepo.isExists() {
		gri, err = r.GitRepo.Inspect()
	}
	return
}

// Data

func (r *Repository) IsExists() (is bool) {
	is = r.GitRepo.isExists()
	return
}

func (r *Repository) Load() {
	// gr := GitRepoNew(r.Workspace, r.Path)
	// gr.load()
	r.GitRepo.load()
}

func (r *Repository) Create(s *Stream, uri string, ssh bool) {
	go r.GitRepo.clone(s, uri, ssh)
}

func (r *Repository) Update(s *Stream) {
	go r.GitRepo.pull(s)
}

func (r *Repository) Destroy() {
	r.GitRepo.remove()
}

// Associations

func (r *Repository) Framers() (frs []*Framer, err error) {
	if r.GitRepo != nil {
		frs, err = r.GitRepo.Framers()
	}
	return
}

func (r *Repository) Shapers() (frs []*Shaper, err error) {
	if r.GitRepo != nil {
		frs, err = r.GitRepo.Shapers()
	}
	return
}
