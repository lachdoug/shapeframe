{{ define "workspaces/directory" -}}
{{ .Path }}
{{ include "gitrepos/gitrepo" (.GitRepo) }}
{{ end -}}
