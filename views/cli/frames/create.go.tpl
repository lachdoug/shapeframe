{{ define "frames/create" -}}
{{ with .Result -}}
Successfully added frame {{ .Frame }} to workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
