package controllers

// import (
// 	"sf/app/errors"
// 	"sf/models"
// )

// type FrameJoinsCreateParams struct {
// 	Workspace string
// 	Frame     string
// 	Parent    string
// }

// type FrameJoinsCreateResult struct {
// 	Workspace string
// 	Frame     string
// 	Parent    string
// }

// func FrameJoinsCreate(jparams []byte) (jbody []byte, err error) {
// 	var w *models.Workspace
// 	var fc *models.Frame
// 	var fp *models.Frame
// 	var isCircular bool
// 	params := ParamsFor[FrameJoinsCreateParams](jparams)

// 	uc := models.ResolveUserContext("Workspaces")
// 	if w, err = models.ResolveWorkspace(uc, params.Workspace,
// 		"Frames",
// 	); err != nil {
// 		return
// 	}
// 	if fc, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
// 		return
// 	}
// 	if fp, err = models.ResolveFrame(uc, w, params.Parent); err != nil {
// 		return
// 	}
// 	fc.Parent = fp
// 	if isCircular, err = fc.IsCircular(); err != nil {
// 		return
// 	}
// 	if isCircular {
// 		err = errors.Error("circular references are not permitted")
// 	}
// 	fc.Save()

// 	result := &FrameJoinsCreateResult{
// 		Workspace: w.Name,
// 		Frame:     fc.Name,
// 		Parent:    fp.Name,
// 	}

// 	jbody = jbodyFor(result)
// 	return
// }
