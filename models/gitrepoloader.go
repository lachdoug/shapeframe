package models

import (
	"path/filepath"
	"sf/utils"

	"golang.org/x/exp/slices"
)

type GitRepoLoader struct {
	GitRepo *GitRepo
	Shapers bool
	Framers bool
	Loads   []string
}

func NewGitRepoLoader(g *GitRepo, loads []string) (gl *GitRepoLoader) {
	gl = &GitRepoLoader{
		GitRepo: g,
		Loads:   loads,
	}
	return
}

func (gl *GitRepoLoader) load() (err error) {
	gl.settle()
	err = gl.assign()
	return
}

func (gl *GitRepoLoader) settle() {
	if slices.Contains(gl.Loads, "Shapers") {
		gl.Shapers = true
	}
	if slices.Contains(gl.Loads, "Framers") {
		gl.Framers = true
	}
	return
}

func (gl *GitRepoLoader) assign() (err error) {
	if err = gl.loadShapers(); err != nil {
		return
	}
	if err = gl.loadFramers(); err != nil {
		return
	}
	return
}

func (gl *GitRepoLoader) loadShapers() (err error) {
	if gl.Shapers {
		gl.setShapers()
		for _, sr := range gl.GitRepo.Shapers {
			if err = sr.Load(); err != nil {
				return
			}
		}
	}
	return
}

func (gl *GitRepoLoader) loadFramers() (err error) {
	if gl.Framers {
		gl.setFramers()
		for _, fr := range gl.GitRepo.Framers {
			if err = fr.Load(); err != nil {
				return
			}
		}
	}
	return
}

func (gl *GitRepoLoader) setShapers() {
	for _, n := range gl.shaperNames() {
		gl.GitRepo.Shapers = append(gl.GitRepo.Shapers, gl.shaper(n))
	}
}

func (gl *GitRepoLoader) setFramers() {
	for _, n := range gl.framerNames() {
		gl.GitRepo.Framers = append(gl.GitRepo.Framers, gl.framer(n))
	}
}

func (gl *GitRepoLoader) shaperNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(gl.GitRepo.Path, "shapers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(gl.GitRepo.Path, "shapers", dir, "shaper.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (gl *GitRepoLoader) framerNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(gl.GitRepo.Path, "framers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(gl.GitRepo.Path, "framers", dir, "framer.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (gl *GitRepoLoader) shaper(name string) (sr *Shaper) {
	sr = NewShaper(
		gl.GitRepo.Workspace,
		gl.GitRepo.Path,
		name,
	)
	return
}

func (gl *GitRepoLoader) framer(name string) (fr *Framer) {
	fr = NewFramer(
		gl.GitRepo.Workspace,
		gl.GitRepo.Path,
		name,
	)
	return
}
