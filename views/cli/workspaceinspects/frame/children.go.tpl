{{ define "workspaceinspects/frame/children" -}}
Children:{{ if eq (len .) 0 }} <none>
{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
