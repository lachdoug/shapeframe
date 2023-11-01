{{ with index . "Result" -}}
Created frame {{ index . "Name" }} in workspace {{ index . "Workspace" }}
{{ end -}}