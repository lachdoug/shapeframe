{{ define "workspaces/repository" -}}
{{ include "gitrepos/gitrepo" (.GitRepo) }}
{{ end -}}