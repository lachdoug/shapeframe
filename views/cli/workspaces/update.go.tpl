{{ define "workspaces/update" -}}
{{ with .Payload -}}
Successfully updated workspace
From:
{{ include "labels/label" .From 2 }}
To: 
{{ include "labels/label" .To 2 }}
{{ end -}}
{{ end -}}
