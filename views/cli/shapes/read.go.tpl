{{ define "shapes/read" -}}
{{ with .Payload -}}
Name: {{ .Name }}
About: {{ .About }}
Workspace: {{ .Workspace }}
Frame: {{ .Frame }}
Configuration:
  Shape: {{ include "configurations/configuration" .ShapeSettings 4 0 }}
  FrameShape: {{ include "configurations/configuration" .FrameShapeSettings 4 0 }}
{{ end -}}
{{ end -}}

