{{ define "workspaceinspects/read" -}}
{{ with .Result -}}
Name: {{ .Name }}
About: {{ .About }}
{{ include "workspaceinspects/directories" (.Directories) }}
{{ include "workspaceinspects/repositories" (.Repositories) }}
{{ include "workspaceinspects/frames" (.Frames) }}
{{ end -}}
{{ end -}}
