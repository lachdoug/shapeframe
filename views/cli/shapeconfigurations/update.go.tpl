{{ define "shapeconfigurations/update" -}}
{{ with .Result -}}
Successfully configured shape {{ .Shape }} in frame {{ .Frame }} workspace {{ .Workspace }}
{{ with .From -}}
From:
{{ include "configurations/configuration" (.Configuration) 2 }}
{{ end -}}
{{ with .To -}}
To:
{{ include "configurations/configuration" (.Configuration) 2 }}
{{ end -}}
{{ end -}}
{{ end -}}
