{{ define "workspaceinspects/repository" -}}
{{ include "gitrepos/gitrepo" (.GitRepo) }}
{{ end -}}