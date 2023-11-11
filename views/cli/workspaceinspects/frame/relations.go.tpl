{{ define "workspaceinspects/frame/relations" -}}
Relations:
  Parent: {{ .Parent }}
{{ include "workspaceinspects/frame/children" (.Children) 2 }}
{{ include "workspaceinspects/frame/ancestors" (.Ancestors) 2 }}
{{ include "workspaceinspects/frame/descendents" (.Descendents) 2 }}
{{ end -}}


