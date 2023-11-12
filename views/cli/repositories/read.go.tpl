{{ define "repositories/read" -}}
{{ with .Result -}}
Workspace: {{ .Workspace }}
URI: {{ .URI }}
Branch: {{ .Branch }}
{{ include "repositories/branches" (.Branches) }}
{{ include "repositories/shapers" (.Shapers) }}
{{ include "repositories/framers" (.Framers) }}
{{ end -}}
{{ end -}}
