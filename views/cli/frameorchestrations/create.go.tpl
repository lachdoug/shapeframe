{{ define "frameorchestrations/create" -}}
{{ with .Result -}}
Successfully orchestrated frame {{ .Frame }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

