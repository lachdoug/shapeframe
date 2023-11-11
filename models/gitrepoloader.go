package models

import (
	"path/filepath"
	"sf/utils"

	"golang.org/x/exp/slices"
)

type GitRepoLoader struct {
	GitRepo *GitRepo
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
	err = gl.associations()
	return
}

func (gl *GitRepoLoader) associations() (err error) {
	if slices.Contains(gl.Loads, "Shapers") {
		gl.SetShapers()
		for _, sr := range gl.GitRepo.Shapers {
			if err = sr.Load(); err != nil {
				return
			}
		}
	}
	if slices.Contains(gl.Loads, "Framers") {
		gl.SetFramers()
		for _, sr := range gl.GitRepo.Framers {
			if err = sr.Load(); err != nil {
				return
			}
		}
	}
	return
}

func (gl *GitRepoLoader) SetFramers() {
	for _, n := range gl.FramerNames() {
		gl.GitRepo.Framers = append(gl.GitRepo.Framers, gl.Framer(n))
	}
}

func (gl *GitRepoLoader) SetShapers() {
	for _, n := range gl.ShaperNames() {
		gl.GitRepo.Shapers = append(gl.GitRepo.Shapers, gl.Shaper(n))
	}
}

func (gl *GitRepoLoader) FramerNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(gl.GitRepo.Path, "framers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(gl.GitRepo.Path, "framers", dir, "framer.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (gl *GitRepoLoader) ShaperNames() (ns []string) {
	dirs := utils.SubDirs(filepath.Join(gl.GitRepo.Path, "shapers"))
	for _, dir := range dirs {
		if utils.IsFile(filepath.Join(gl.GitRepo.Path, "shapers", dir, "shaper.yaml")) {
			ns = append(ns, dir)
		}
	}
	return
}

func (gl *GitRepoLoader) Framer(name string) (fr *Framer) {
	fr = NewFramer(
		gl.GitRepo.Workspace,
		gl.GitRepo.Path,
		name,
	)
	return
}

func (gl *GitRepoLoader) Shaper(name string) (sr *Shaper) {
	sr = NewShaper(
		gl.GitRepo.Workspace,
		gl.GitRepo.Path,
		name,
	)
	return
}
