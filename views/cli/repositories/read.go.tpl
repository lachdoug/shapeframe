{{ define "repositories/read" -}}
{{ with .Payload -}}
Workspace: {{ .Workspace }}
URI: {{ .URI }}
Branch: {{ .Branch }}
Branches: {{ include "repositories/branches" .Branches }}
Shapers: {{ include "repositories/shapers" .Shapers }}
Framers: {{ include "repositories/framers" .Framers }}
{{ end -}}
{{ end -}}
