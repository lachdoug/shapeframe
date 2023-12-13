{{ define "workspaceinspects/frame/shape" -}}
Name: {{ .Name }}
About: {{ .About }}
Shaper: {{ .Shaper }}
Configuration:
  Shape: {{ include "configurations/configuration" .ShapeSettings 4 0 }}
  FrameShape: {{ include "configurations/configuration" .FrameShapeSettings 4 0 }}
{{ end -}}
