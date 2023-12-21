{{ define "configurations/configuration" -}}
{{ if eq (len .Info) 0 -}}
<none>
{{ else -}}
{{ include "configurations/settings" .SettingsMaps }}
{{ end -}}
{{ end -}}
