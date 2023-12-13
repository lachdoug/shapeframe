{{ define "shapes/update" -}}
{{ with .Payload -}}
Successfully updated shape in frame {{ .Frame }} workspace {{ .Workspace }}
From: {{ include "labels/label" (.From) }}
To: {{ include "labels/label" (.To) }}
{{ end -}}
{{ end -}}
