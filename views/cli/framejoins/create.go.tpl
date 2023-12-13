{{ define "framejoins/create" -}}
{{ with .Payload -}}
Successfully joined frame {{ .Workspace }} to frame {{ .Parent }}
{{ end -}}
{{ end -}}
