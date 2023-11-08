{{ define "frames/destroy" -}}
{{ with .Result -}}
Successfully removed frame {{ .Frame }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
