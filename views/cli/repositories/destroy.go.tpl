{{ with index . "Result" -}}
Removed repository {{ index . "URI" }} from workspace {{ index . "Workspace" }}
{{ end -}}
