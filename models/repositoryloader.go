package models

import (
	"sf/utils"

	"golang.org/x/exp/slices"
)

type RepositoryLoader struct {
	Repository   *Repository
	Loads        []string
	GitRepo      bool
	Shapers      bool
	Framers      bool
	GitRepoLoads []string
}

func NewRepositoryLoader(w *Repository, loads []string) (rl *RepositoryLoader) {
	rl = &RepositoryLoader{
		Repository: w,
		Loads:      loads,
	}
	return
}

func (rl *RepositoryLoader) load() (err error) {
	rl.dependencies()
	rl.settle()
	err = rl.assign()
	return
}

func (rl *RepositoryLoader) dependencies() {
	primaries := primaryLoads(rl.Loads)
	if slices.Contains(primaries, "Shapers") {
		rl.Loads = append(rl.Loads,
			"GitRepo.Shapers",
		)
	}
	if slices.Contains(rl.Loads, "Framers") {
		rl.Loads = append(rl.Loads,
			"GitRepo.Framers",
		)
	}
}

func (rl *RepositoryLoader) settle() {
	utils.UniqStrings(&rl.Loads)
	for _, load := range rl.Loads {
		switch primaryLoad(load) {
		case "GitRepo":
			abstractAssociation(load, "GitRepo", &rl.GitRepo, &rl.GitRepoLoads)
		case "Shapers":
			abstractAssociation(load, "Shapers", &rl.Shapers)
		case "Framers":
			abstractAssociation(load, "Framers", &rl.Framers)
		}
	}
}

func (rl *RepositoryLoader) assign() (err error) {
	if err = rl.loadGitRepo(); err != nil {
		return
	}
	rl.loadShapers()
	rl.loadFramers()
	return
}

func (rl *RepositoryLoader) loadGitRepo() (err error) {
	if rl.GitRepo {
		rl.setGitRepo()
		if err = rl.Repository.GitRepo.Load(rl.GitRepoLoads...); err != nil {
			return
		}
	}
	return
}

func (rl *RepositoryLoader) loadShapers() {
	if rl.Shapers {
		rl.setShapers()
	}
}

func (rl *RepositoryLoader) loadFramers() {
	if rl.Framers {
		rl.setFramers()
	}
}

func (rl *RepositoryLoader) setGitRepo() {
	rl.Repository.GitRepo = NewGitRepo(rl.Repository.Workspace, rl.Repository.directory())
}

func (rl *RepositoryLoader) setFramers() {
	rl.Repository.Framers = rl.Repository.GitRepo.Framers
}

func (rl *RepositoryLoader) setShapers() {
	rl.Repository.Shapers = rl.Repository.GitRepo.Shapers
}
