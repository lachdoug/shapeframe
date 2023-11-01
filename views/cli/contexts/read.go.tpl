{{ with index . "Result" -}}
{{ if index . "Workspace" -}}
CONTEXT
  Workspace: {{ index . "Workspace" }}
{{ else -}}
No context
{{ end -}}
{{ if index . "Frame" }}  Frame: {{ index . "Frame" }}
{{ end -}}
{{ if index . "Shape" }}  Shape: {{ index . "Shape" }}
{{ end -}}
{{ end -}}
