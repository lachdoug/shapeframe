package models

import (
	"fmt"
	"net/url"
	"path/filepath"
	"sf/utils"
	"strings"

	"github.com/go-git/go-git/v5"
)

type GitRepo struct {
	Workspace *Workspace
	Path      string
	URI       string
	URL       string
	Branch    string
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

func (g *GitRepo) Inspect() (gri *GitRepoInspector) {
	gri = &GitRepoInspector{
		URI:     g.URI,
		URL:     g.URL,
		Branch:  g.Branch,
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

// URL/URI

func (g *GitRepo) OriginURL(uri string, ssh bool) (u string) {
	uUrl, _ := url.Parse("https://" + uri)
	host := uUrl.Host
	path := filepath.Join(strings.Split(uUrl.Path, "/")...)
	if ssh {
		u = fmt.Sprintf("git@%s:%s.git", host, path)
	} else {
		u = fmt.Sprintf("https://%s/%s", host, path)
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

func (g *GitRepo) clone(uri string, ssh bool, st *Stream) {
	var err error
	d := g.Path
	utils.MakeDir(d)
	o := &git.CloneOptions{
		URL:      g.OriginURL(uri, ssh),
		Depth:    1,
		Progress: st,
	}
	if _, err = git.PlainClone(d, false, o); err != nil {
		st.Error(err)
		g.remove()
	}
	st.Close()
}

func (g *GitRepo) pull(st *Stream) {
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
