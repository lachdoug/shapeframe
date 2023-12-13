{{ define "workspaceinspects/read" -}}
{{ with .Payload -}}
Name: {{ .Name }}
About: {{ .About }}
Directories: {{ include "workspaceinspects/directories" .Directories 2 0 }}
Repositories: {{ include "workspaceinspects/repositories" .Repositories 2 0 }}
Frames: {{ include "workspaceinspects/frames" .Frames 2 0 }}
{{ end -}}
{{ end -}}
