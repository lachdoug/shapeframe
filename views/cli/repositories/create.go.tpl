{{ define "repositories/create" -}}
{{ with .Result -}}
Successfully added repository {{ .URI }} to workspace {{ .Workspace }}
Repository:
{{ include "gitrepos/gitrepo" (.GitRepo) 2 }}
{{ end -}}
{{ end -}}
