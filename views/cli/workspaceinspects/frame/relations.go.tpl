{{ define "workspaceinspects/frame/relations" -}}
Parent: {{ .Parent }}
Children: {{ include "workspaceinspects/frame/children" .Children 2 0 }}
Ancestors: {{ include "workspaceinspects/frame/ancestors" .Ancestors 2 0 }}
Descendents: {{ include "workspaceinspects/frame/descendents" .Descendents 2 0 }}
{{ end -}}


