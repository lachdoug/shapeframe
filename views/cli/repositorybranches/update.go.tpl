{{ define "repositorybranches/update" -}}
{{ with .Payload -}}
Successfully checked-out branch {{ .Branch }} for repository {{ .URI }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

