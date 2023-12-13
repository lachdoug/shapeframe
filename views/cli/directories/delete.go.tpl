{{ define "directories/delete" -}}
{{ with .Payload -}}
Successfully removed directory {{ .Path }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}


