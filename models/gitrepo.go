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
}

type GitRepoInspector struct {
	URI     string
	URL     string
	Branch  string
	Shapers []*ShaperInspector
	Framers []*FramerInspector
}

// Construct

func GitRepoNew(w *Workspace, path string) (g *GitRepo) {
	if w == nil {
		panic("GitRepo Workspace is <nil>")
	}
	g = &GitRepo{Workspace: w, Path: path}
	return
}

// Inspection

func (g *GitRepo) Inspect() (gri *GitRepoInspector, err error) {
	var sris []*ShaperInspector
	var fris []*FramerInspector
	if sris, err = g.ShapersInspect(); err != nil {
		return
	}
	if fris, err = g.FramersInspect(); err != nil {
		return
	}
	gri = &GitRepoInspector{
		URI:     g.URI,
		URL:     g.URL,
		Branch:  g.Branch,
		Shapers: sris,
		Framers: fris,
	}
	return
}

func (g *GitRepo) ShapersInspect() (sris []*ShaperInspector, err error) {
	var srs []*Shaper
	sris = []*ShaperInspector{}
	if srs, err = g.Shapers(); err != nil {
		return
	}
	for _, sr := range srs {
		sris = append(sris, sr.Inspect())
	}
	return
}

func (g *GitRepo) FramersInspect() (fris []*FramerInspector, err error) {
	var frs []*Framer
	fris = []*FramerInspector{}
	if frs, err = g.Framers(); err != nil {
		return
	}
	for _, fr := range frs {
		fris = append(fris, fr.Inspect())
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

func (g *GitRepo) load() {
	fmt.Println(g.Path)

	g.URI = utils.GitRepoURI(g.Path)
	g.URL = utils.GitRepoURL(g.Path)
	g.Branch = utils.GitRepoBranch(g.Path)
}

func (g *GitRepo) remove() (err error) {
	utils.RemoveDir(g.Path)
	return
}

func (g *GitRepo) clone(s *Stream, uri string, ssh bool) {
	var err error
	d := g.Path
	utils.MakeDir(d)
	o := &git.CloneOptions{URL: g.OriginURL(uri, ssh), Depth: 1, Progress: s}
	if _, err = git.PlainClone(d, false, o); err != nil {
		s.Error(err)
		g.remove()
	}
	s.Close()
}

func (g *GitRepo) pull(s *Stream) {
	var gr *git.Repository
	var gw *git.Worktree
	var err error
	d := g.Path
	o := &git.PullOptions{RemoteName: "origin", Depth: 1, Progress: s}
	if gr, err = git.PlainOpen(d); err != nil {
		s.Error(err)
	} else if gw, err = gr.Worktree(); err != nil {
		s.Error(err)
	} else if err = gw.Pull(o); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			s.Write([]byte("Already up to date\n"))
		} else {
			s.Error(err)
		}
	}
	s.Close()
}

// Associations

func (g *GitRepo) Framers() (frs []*Framer, err error) {
	var fr *Framer
	for _, n := range g.FramerNames() {
		if fr, err = g.Framer(n); err != nil {
			return
		}
		frs = append(frs, fr)
	}
	return
}

func (g *GitRepo) Shapers() (srs []*Shaper, err error) {
	var sr *Shaper
	for _, n := range g.ShaperNames() {
		if sr, err = g.Shaper(n); err != nil {
			return
		}
		srs = append(srs, sr)
	}
	return
}

func (g *GitRepo) FramerNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(g.Path, "framers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(g.Path, "framers", dir, "framer.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (g *GitRepo) ShaperNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(g.Path, "shapers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(g.Path, "shapers", dir, "shaper.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (g *GitRepo) Framer(name string) (fr *Framer, err error) {
	fr = FramerNew(
		g.Workspace,
		g.Path,
		name,
	)
	err = fr.Load()
	return
}

func (g *GitRepo) Shaper(name string) (sr *Shaper, err error) {
	sr = ShaperNew(
		g.Workspace,
		g.Path,
		name,
	)
	err = sr.Load()
	return
}
