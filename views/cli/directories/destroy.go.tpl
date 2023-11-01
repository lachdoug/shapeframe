{{ with index . "Result" -}}
Removed directory {{ index . "Path" }} from workspace {{ index . "Workspace" }}
{{ end -}}
