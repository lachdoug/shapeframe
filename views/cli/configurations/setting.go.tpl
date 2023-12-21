{{ define "configurations/setting" -}}
{{ range $k, $v := . -}}
{{ $k }}: {{ $v }}
{{ end -}}
{{ end -}}
