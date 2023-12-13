package models

import (
	"sf/app/errors"
	"sf/app/repos"
	"sf/app/streams"
	"sf/utils"
)

type GitRepo struct {
	Workspace *Workspace
	Path      string
	Shapers   []*Shaper
	Framers   []*Framer
}

type GitRepoInspector struct {
	Path     string
	URI      string
	URL      string
	Branch   string
	Branches []string
	Shapers  []*ShaperInspector
	Framers  []*FramerInspector
}

// Construct

func NewGitRepo(w *Workspace, path string) (g *GitRepo) {
	g = &GitRepo{
		Workspace: w,
		Path:      path,
	}
	return
}

// Inspection

func (g *GitRepo) Inspect() (gri *GitRepoInspector, err error) {
	var uri string
	var url string
	var branch string
	var branches []string
	if uri, err = g.URI(); err != nil {
		return
	}
	if url, err = g.URL(); err != nil {
		return
	}
	if branch, err = g.Branch(); err != nil {
		return
	}
	if branches, err = g.Branches(); err != nil {
		return
	}
	gri = &GitRepoInspector{
		Path:     g.Path,
		URI:      uri,
		URL:      url,
		Branch:   branch,
		Branches: branches,
		Shapers:  g.ShapersInspect(),
		Framers:  g.FramersInspect(),
	}
	return
}

func (g *GitRepo) ShapersInspect() (sris []*ShaperInspector) {
	sris = []*ShaperInspector{}
	for _, sr := range g.Shapers {
		sri := sr.Inspect()
		sris = append(sris, sri)
	}
	return
}

func (g *GitRepo) FramersInspect() (fris []*FramerInspector) {
	fris = []*FramerInspector{}
	for _, fr := range g.Framers {
		fri := fr.Inspect()
		fris = append(fris, fri)
	}
	return
}

// Data

func (g *GitRepo) isExists() (is bool) {
	is = utils.IsGitRepoDir(g.Path)
	return
}

func (g *GitRepo) Load(loads ...string) (err error) {
	gl := NewGitRepoLoader(g, loads)
	err = gl.load()
	return
}

func (g *GitRepo) remove() {
	utils.RemoveDir(g.Path)
}

func (g *GitRepo) clone(url string, username string, password string, st *streams.Stream) {
	repos.GitClone(g.Path, url, username, password, st)
}

func (g *GitRepo) pull(username string, password string, st *streams.Stream) {
	repos.GitPull(g.Path, username, password, st)
}

func (g *GitRepo) checkout(branch string) (err error) {
	if err = repos.GitCheckout(g.Path, branch); err != nil {
		err = errors.ErrorWrapf(err, "checkout %s", branch)
	}
	return
}

// Repo

func (g *GitRepo) URI() (url string, err error) {
	if url, err = utils.GitURI(g.Path); err != nil {
		err = errors.ErrorWrap(err, "gitrepo uri")
	}
	return
}

func (g *GitRepo) URL() (url string, err error) {
	if url, err = utils.GitURL(g.Path); err != nil {
		err = errors.ErrorWrap(err, "gitrepo url")
	}
	return
}

func (g *GitRepo) Branch() (branch string, err error) {
	if branch, err = repos.GitBranch(g.Path); err != nil {
		err = errors.ErrorWrap(err, "gitrepo branch")
	}
	return
}

func (g *GitRepo) Branches() (branches []string, err error) {
	if branches, err = repos.GitBranches(g.Path); err != nil {
		err = errors.ErrorWrap(err, "gitrepo branches")
	}
	return
}

func (g *GitRepo) FramerNames() (ns []string) {
	ns = []string{}
	for _, fr := range g.Framers {
		ns = append(ns, fr.Name)
	}
	return
}

func (g *GitRepo) ShaperNames() (ns []string) {
	ns = []string{}
	for _, sr := range g.Shapers {
		ns = append(ns, sr.Name)
	}
	return
}
