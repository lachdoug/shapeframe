{{ define "shapes/update" -}}
{{ with .Result -}}
Successfully updated shape in frame {{ .Frame }} workspace {{ .Workspace }}
From:
{{ include "labels/label" (.From) 2 }}
To:
{{ include "labels/label" (.To) 2 }}
{{ end -}}
{{ end -}}
