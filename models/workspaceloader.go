package models

import (
	"path/filepath"
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"slices"
)

type WorkspaceLoader struct {
	Workspace       *Workspace
	Loads           []string
	Shapers         bool
	Framers         bool
	Repositories    bool
	ShaperLoads     []string
	FramerLoads     []string
	FrameLoads      []string
	RepositoryLoads []string
	DirectoryLoads  []string
	Preloads        []string
}

func NewWorkspaceLoader(w *Workspace, loads []string) (wl *WorkspaceLoader) {
	wl = &WorkspaceLoader{
		Workspace: w,
		Loads:     loads,
	}
	return
}

func (wl *WorkspaceLoader) load() (err error) {
	wl.dependencies()
	wl.settle()
	wl.query()
	err = wl.assign()
	return
}

func (wl *WorkspaceLoader) dependencies() {
	primaries := primaryLoads(wl.Loads)
	if slices.Contains(primaries, "Shapers") {
		wl.Loads = append(wl.Loads,
			"Directories.Shapers.Workspace",
			"Repositories.Shapers",
		)
	}
	if slices.Contains(primaries, "Framers") {
		wl.Loads = append(wl.Loads,
			"Directories.Framers.Workspace",
			"Repositories.Framers",
		)
	}
}

func (wl *WorkspaceLoader) settle() {
	utils.UniqStrings(&wl.Loads)
	for _, load := range wl.Loads {
		switch primaryLoad(load) {
		case "Frames":
			databaseAssociation(load, "Frames", &wl.Preloads, &wl.FrameLoads)
		case "Shapers":
			abstractAssociation(load, "Shapers", &wl.Shapers, &wl.ShaperLoads)
		case "Framers":
			abstractAssociation(load, "Framers", &wl.Framers, &wl.FramerLoads)
		case "Repositories":
			abstractAssociation(load, "Repositories", &wl.Repositories, &wl.RepositoryLoads)
		case "Directories":
			databaseAssociation(load, "Directories", &wl.Preloads, &wl.DirectoryLoads)
		default:
			wl.Preloads = append(wl.Preloads, load)
		}
	}
}

func (wl *WorkspaceLoader) query() {
	if len(wl.Preloads) > 0 {
		utils.UniqStrings(&wl.Preloads)
		queries.Load(wl.Workspace, wl.Workspace.ID, wl.Preloads...)
	}
}

func (wl *WorkspaceLoader) assign() (err error) {
	if err = wl.loadRepositories(); err != nil {
		return
	}
	if err = wl.loadDirectories(); err != nil {
		return
	}
	if err = wl.loadShapers(); err != nil {
		return
	}
	if err = wl.loadFramers(); err != nil {
		return
	}
	if err = wl.loadFrames(); err != nil {
		return
	}
	return
}

func (wl *WorkspaceLoader) loadRepositories() (err error) {
	if wl.Repositories {
		if err = wl.setRepositories(); err != nil {
			return
		}
		for _, r := range wl.Workspace.Repositories {
			if err = r.Load(wl.RepositoryLoads...); err != nil {
				return
			}
		}
	}
	return
}

func (wl *WorkspaceLoader) loadDirectories() (err error) {
	if len(wl.DirectoryLoads) > 0 {
		for _, d := range wl.Workspace.Directories {
			if err = d.Load(wl.DirectoryLoads...); err != nil {
				return
			}
		}
	}
	return
}

func (wl *WorkspaceLoader) loadShapers() (err error) {
	if wl.Shapers {
		for _, d := range wl.Workspace.Directories {
			wl.Workspace.Shapers = append(wl.Workspace.Shapers, d.Shapers...)
		}
		for _, r := range wl.Workspace.Repositories {
			wl.Workspace.Shapers = append(wl.Workspace.Shapers, r.Shapers...)
		}
	}
	return
}

func (wl *WorkspaceLoader) loadFramers() (err error) {
	if wl.Framers {
		for _, d := range wl.Workspace.Directories {
			wl.Workspace.Framers = append(wl.Workspace.Framers, d.Framers...)
		}
		for _, r := range wl.Workspace.Repositories {
			wl.Workspace.Framers = append(wl.Workspace.Framers, r.Framers...)
		}
	}
	return
}

func (wl *WorkspaceLoader) loadFrames() (err error) {
	if len(wl.FrameLoads) > 0 {
		for _, d := range wl.Workspace.Frames {
			if err = d.Load(wl.FrameLoads...); err != nil {
				return
			}
		}
	}
	return
}

func (wl *WorkspaceLoader) setRepositories() (err error) {
	var rus []string
	if rus, err = wl.repositoryURIs(); err != nil {
		return
	}
	for _, ru := range rus {
		r := NewRepository(wl.Workspace, ru)
		wl.Workspace.Repositories = append(wl.Workspace.Repositories, r)
	}
	return
}

func (wl *WorkspaceLoader) repositoryURIs() (dirPaths []string, err error) {
	if dirPaths, err = utils.GitURIs(filepath.Join(wl.Workspace.directory(), "repos")); err != nil {
		err = app.ErrorWrapf(err, "repository URIs")
		return
	}
	return
}
