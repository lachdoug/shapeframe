{{ with index . "Result" -}}
Destroyed frame {{ index . "Name" }} in workspace {{ index . "Workspace" }}
{{ end -}}