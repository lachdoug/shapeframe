{{ define "frameconfigurations/update" -}}
{{ with .Payload -}}
Successfully configured settings for frame {{ .Frame }} in workspace {{ .Workspace }}
From: {{ include "configurations/configuration" .From }}
To: {{ include "configurations/configuration" .To }}
{{ end -}}
{{ end -}}

