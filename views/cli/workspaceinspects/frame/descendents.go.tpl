{{ define "workspaceinspects/frame/descendents" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
