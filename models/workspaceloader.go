package models

import (
	"path/filepath"
	"sf/app"
	"sf/database/queries"
	"sf/utils"
	"strings"

	"golang.org/x/exp/slices"
)

type WorkspaceLoader struct {
	Workspace       *Workspace
	Loads           []string
	Repositories    bool
	Shapers         bool
	Framers         bool
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
	if slices.Contains(wl.Loads, "Framers") {
		wl.Loads = append(wl.Loads,
			"Directories.Workspace",
			"Directories.Framers",
			"Repositories.Framers",
		)
	}
	if slices.Contains(wl.Loads, "Shapers") {
		wl.Loads = append(wl.Loads,
			"Directories.Workspace",
			"Directories.Shapers",
			"Repositories.Shapers",
		)
	}
	utils.UniqStrings(&wl.Loads)
}

func (wl *WorkspaceLoader) settle() {
	for _, load := range wl.Loads {
		elem := strings.SplitN(load, ".", 2)
		switch elem[0] {
		case "Frames":
			wl.Preloads = append(wl.Preloads, "Frames")
			if len(elem) > 1 {
				elem1 := strings.SplitN(elem[1], ".", 2)
				switch elem1[0] {
				case "Shapes":
					wl.Preloads = append(wl.Preloads, "Frames.Shapes")
					if len(elem1) > 1 {
						switch elem1[1] {
						case "Configuration":
							wl.FrameLoads = append(wl.FrameLoads, "Shapes.Configuration")
						default:
							wl.Preloads = append(wl.Preloads, load)
						}
					}
				case "Configuration":
					wl.FrameLoads = append(wl.FrameLoads, elem1[0])
				default:
					wl.Preloads = append(wl.Preloads, load)
				}
			}
		case "Shapers":
			wl.Shapers = true
		case "Framers":
			wl.Framers = true
		case "Repositories":
			wl.Repositories = true
			if len(elem) > 1 {
				wl.RepositoryLoads = append(wl.RepositoryLoads, elem[1])
			}
		case "Directories":
			wl.Preloads = append(wl.Preloads, "Directories")
			if len(elem) > 1 {
				switch elem[1] {
				case "Shapers", "Framers":
					wl.DirectoryLoads = append(wl.DirectoryLoads, elem[1])
				default:
					wl.Preloads = append(wl.Preloads, load)
				}
			}
		default:
			wl.Preloads = append(wl.Preloads, load)
		}
	}
	utils.UniqStrings(&wl.Preloads)
}

func (wl *WorkspaceLoader) query() {
	queries.Load(wl.Workspace, wl.Workspace.ID, wl.Preloads...)
}

func (wl *WorkspaceLoader) assign() (err error) {
	if wl.Repositories {
		if err = wl.SetRepositories(); err != nil {
			return
		}
		if err = wl.LoadRepositories(); err != nil {
			return
		}
	}
	if len(wl.DirectoryLoads) > 0 {
		if err = wl.LoadDirectories(); err != nil {
			return
		}
	}
	if len(wl.FrameLoads) > 0 {
		if err = wl.LoadFrames(); err != nil {
			return
		}
	}
	if wl.Shapers {
		wl.SetShapers()
	}
	if wl.Framers {
		wl.SetFramers()
	}
	return
}

func (wl *WorkspaceLoader) LoadRepositories() (err error) {
	for _, r := range wl.Workspace.Repositories {
		if err = r.Load(wl.RepositoryLoads...); err != nil {
			return
		}
	}
	return
}

func (wl *WorkspaceLoader) LoadDirectories() (err error) {
	for _, d := range wl.Workspace.Directories {
		if err = d.Load(wl.DirectoryLoads...); err != nil {
			return
		}
	}
	return
}

func (wl *WorkspaceLoader) LoadFrames() (err error) {
	for _, d := range wl.Workspace.Frames {
		if err = d.Load(wl.FrameLoads...); err != nil {
			return
		}
	}
	return
}

func (wl *WorkspaceLoader) RepositoryURIs() (dirPaths []string, err error) {
	if dirPaths, err = utils.GitRepoURIs(filepath.Join(wl.Workspace.directory(), "repos")); err != nil {
		err = app.ErrorWith(err, "repository URIs")
		return
	}
	return
}

func (wl *WorkspaceLoader) SetRepositories() (err error) {
	var rus []string
	if rus, err = wl.RepositoryURIs(); err != nil {
		return
	}
	for _, ru := range rus {
		var protocol string
		if ru[0:4] == "https" {
			protocol = "HTTPS"
		} else {
			protocol = "SSH"

		}
		r := NewRepository(wl.Workspace, ru, protocol)
		wl.Workspace.Repositories = append(wl.Workspace.Repositories, r)
	}
	return
}

func (wl *WorkspaceLoader) SetFramers() {
	for _, d := range wl.Workspace.Directories {
		wl.Workspace.Framers = append(wl.Workspace.Framers, d.Framers...)
	}
	for _, r := range wl.Workspace.Repositories {
		wl.Workspace.Framers = append(wl.Workspace.Framers, r.Framers...)
	}
}

func (wl *WorkspaceLoader) SetShapers() {
	for _, d := range wl.Workspace.Directories {
		wl.Workspace.Shapers = append(wl.Workspace.Shapers, d.Shapers...)
	}
	for _, r := range wl.Workspace.Repositories {
		wl.Workspace.Shapers = append(wl.Workspace.Shapers, r.Shapers...)
	}
}
