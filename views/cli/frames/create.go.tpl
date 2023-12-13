{{ define "frames/create" -}}
{{ with .Payload -}}
Successfully added frame {{ .Frame }} to workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
