{{ define "gitrepos/gitrepo" -}}
URI: {{ .URI }}
URL: {{ .URL }}
Branch: {{ .Branch }}
Shapers: {{ if eq (len (.Shapers)) 0 }}<none>
{{ else -}}
{{ include "gitrepos/framers" (.Shapers) }}
{{ end -}}
Framers: {{ if eq (len (.Framers)) 0 }}<none>
{{ else -}}
{{ include "gitrepos/framers" (.Framers) }}
{{ end -}}
{{ end -}}


