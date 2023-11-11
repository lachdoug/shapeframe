package controllers

import (
	"sf/models"
)

type FramesIndexParams struct {
	Workspace string // Limit list to workspace
}

type FramesIndexItemResult struct {
	Workspace string
	Framer    string
	Name      string
	About     string
	IsContext bool
}

func FramesIndex(jparams []byte) (jbody []byte, err error) {
	var fs []*models.Frame
	params := ParamsFor[FramesIndexParams](jparams)

	uc := models.ResolveUserContext(
		"Workspace.Frames.Workspace",
		"Workspaces.Frames.Workspace",
	)
	if params.Workspace == "" {
		for _, w := range uc.Workspaces {
			fs = append(fs, w.Frames...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
			return
		}
		fs = w.Frames
	}

	result := []*FramesIndexItemResult{}
	for _, f := range fs {
		result = append(result, &FramesIndexItemResult{
			Workspace: f.Workspace.Name,
			Framer:    f.FramerName,
			Name:      f.Name,
			About:     f.About,
			IsContext: uc.FrameID == f.ID,
		})
	}

	jbody = jbodyFor(result)
	return
}
