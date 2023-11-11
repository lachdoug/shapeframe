{{ define "framejoins/create" -}}
{{ with .Result -}}
Successfully joined frame {{ .Workspace }} to frame {{ .Parent }}
{{ end -}}
{{ end -}}
