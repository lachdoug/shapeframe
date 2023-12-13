{{ define "workspaces/read" -}}
{{ with .Payload -}}
Name: {{ .Name }}
About: {{ .About }}
Directories: {{ include "workspaces/directories" .Directories }}
Repositories: {{ include "workspaces/repositories" .Repositories }}
Frames: {{ include "workspaces/frames" .Frames }}
{{ end -}}
{{ end -}}
