{{ define "repositorybranches/update" -}}
{{ with .Result -}}
Successfully checked-out branch {{ .Branch }} for repository {{ .URI }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

