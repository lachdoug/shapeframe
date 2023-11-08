{{ define "directories/create" -}}
{{ with .Result -}}
Successfully added directory {{ .Path }} to workspace {{ .Workspace }}
Directory:
{{ include "gitrepos/gitrepo" (.GitRepo) 2 }}
{{ end -}}
{{ end -}}