{{ with index . "Result" -}}
CONTEXT
{{ with index . "Exit" -}}
From:
{{ if index . "Workspace" }}  Workspace: {{ index . "Workspace" }}
{{ else }}  No context
{{ end -}}
{{ if index . "Frame" }}  Frame: {{ index . "Frame" }}
{{ end -}}
{{ if index . "Shape" }}  Shape: {{ index . "Shape" }}
{{ end -}}
{{ end -}}

{{ with index . "Enter" -}}
To:
{{ if index . "Workspace" }}  Workspace: {{ index . "Workspace" }}
{{ else }}  No context
{{ end -}}
{{ if index . "Frame" }}  Frame: {{ index . "Frame" }}
{{ end -}}
{{ if index . "Shape" }}  Shape: {{ index . "Shape" }}
{{ end -}}
{{ end -}}

{{ end -}}
