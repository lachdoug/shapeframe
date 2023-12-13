{{ define "shapeframeconfigurations/update" -}}
{{ with .Payload -}}
Successfully configured frame settings for shape {{ .Shape }} in frame {{ .Frame }} workspace {{ .Workspace }}
From: {{ include "configurations/configuration" .From }}
To: {{ include "configurations/configuration" .To }}
{{ end -}}
{{ end -}}
