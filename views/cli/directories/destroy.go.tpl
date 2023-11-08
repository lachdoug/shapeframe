{{ define "directories/destroy" -}}
{{ with .Result -}}
Successfully removed directory {{ .Path }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}


