{{ define "gitrepos/gitrepo" -}}
URI: {{ .URI }}
URL: {{ .URL }}
Branch: {{ .Branch }}
Branches: {{ include "gitrepos/branches" .Branches 2 0 }}
Shapers: {{ include "gitrepos/shapers" .Shapers 2 0 }}
Framers: {{ include "gitrepos/framers" .Framers 2 0 }}
{{ end -}}


