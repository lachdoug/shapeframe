{{ define "directories/create" -}}
{{ with .Payload -}}
Successfully added directory {{ .Path }} to workspace {{ .Workspace }}
{{ end -}}
{{ end -}}