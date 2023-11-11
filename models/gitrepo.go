package models

import (
	"sf/app"
	"sf/utils"

	"github.com/go-git/go-git/v5"
)

type GitRepo struct {
	Workspace *Workspace
	Path      string
	Shapers   []*Shaper
	Framers   []*Framer
}

type GitRepoInspector struct {
	URI     string
	URL     string
	Branch  string
	Shapers []*ShaperInspector
	Framers []*FramerInspector
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
	if uri, err = g.URI(); err != nil {
		return
	}
	if url, err = g.URL(); err != nil {
		return
	}
	if branch, err = g.Branch(); err != nil {
		return
	}
	gri = &GitRepoInspector{
		URI:     uri,
		URL:     url,
		Branch:  branch,
		Shapers: g.ShapersInspect(),
		Framers: g.FramersInspect(),
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

func (g *GitRepo) clone(url string, st *utils.Stream) {
	utils.GitRemoteClone(g.Path, url, st)
	var err error
	tmp := utils.TempDir("clone")
	d := g.Path
	utils.RemoveDir(tmp)
	utils.MakeDir(tmp)
	o := &git.CloneOptions{
		URL:      url,
		Depth:    1,
		Progress: st,
	}
	if _, err = git.PlainClone(tmp, false, o); err != nil {
		st.Error(err)
	}
	utils.MoveDir(tmp, d)
	st.Close()
}

func (g *GitRepo) pull(st *utils.Stream) {
	var gr *git.Repository
	var gw *git.Worktree
	var err error
	d := g.Path
	o := &git.PullOptions{
		RemoteName: "origin",
		Depth:      1,
		Progress:   st,
	}
	if gr, err = git.PlainOpen(d); err != nil {
		st.Error(err)
	} else if gw, err = gr.Worktree(); err != nil {
		st.Error(err)
	} else if err = gw.Pull(o); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			st.Writef("Already up to date\n")
		} else {
			st.Error(err)
		}
	}
	st.Close()
}

// Repo

func (g *GitRepo) URI() (url string, err error) {
	if url, err = utils.GitRepoURI(g.Path); err != nil {
		err = app.ErrorWith(err, "gitrepo uri")
	}
	return
}

func (g *GitRepo) URL() (url string, err error) {
	if url, err = utils.GitRepoURL(g.Path); err != nil {
		err = app.ErrorWith(err, "gitrepo url")
	}
	return
}

func (g *GitRepo) Branch() (branch string, err error) {
	if branch, err = utils.GitRepoBranch(g.Path); err != nil {
		err = app.ErrorWith(err, "gitrepo branch")
	}
	return
}
