{{ define "frames/update" -}}
{{ with .Result -}}
Successfully updated frame in workspace {{ .Workspace }}
From:
{{ include "labels/label" (.From) 2 }}
To:
{{ include "labels/label" (.To) 2 }}
{{ end -}}
{{ end -}}
