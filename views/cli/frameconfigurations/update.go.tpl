{{ define "frameconfigurations/update" -}}
{{ with .Result -}}
Successfully configured frame {{ .Frame }} in workspace {{ .Workspace }}
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

