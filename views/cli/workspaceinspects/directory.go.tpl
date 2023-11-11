{{ define "workspaceinspects/directory" -}}
{{ .Path }}
{{ include "gitrepos/gitrepo" (.GitRepo) }}
{{ end -}}
