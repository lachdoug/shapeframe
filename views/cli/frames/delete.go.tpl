{{ define "frames/delete" -}}
{{ with .Payload -}}
Successfully removed frame {{ .Frame }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
