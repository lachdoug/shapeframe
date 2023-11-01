{{ with index . "Result" -}}
Added repository {{ index . "URI" }} to workspace {{ index . "Workspace" }}
Git: {{ if not (index . "GitRepo") }}none{{ end }} 
{{ with index . "GitRepo" }}  URI: {{ index . "URI"}}
  URL: {{ index . "URL"}}
  Branch: {{ index . "Branch"}}
  Shapers: {{ $shapers := index . "Shapers" -}}
{{ if eq (len $shapers) 0 -}}
none
{{ else }}
{{ range $shapers }}  - URI: {{ index . "URI"}}
    Name: {{ index . "Name"}}
    About: {{ index . "About"}}
{{ end -}}
{{ end }}  Framers: {{ $Framers := index . "Framers" -}}
{{ if eq (len $Framers) 0 -}}
none
{{ else }}
{{ range $Framers }}  - URI: {{ index . "URI"}}
    Name: {{ index . "Name"}}
    About: {{ index . "About"}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
