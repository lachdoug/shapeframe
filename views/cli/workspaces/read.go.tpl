{{ define "workspaces/read" -}}
{{ with .Result -}}
Name: {{ .Name }}
About: {{ .About }}
{{ include "workspaces/directories" (.Directories) }}
{{ include "workspaces/repositories" (.Repositories) }}
{{ include "workspaces/frames" (.Frames) }}
{{ end -}}
{{ end -}}
