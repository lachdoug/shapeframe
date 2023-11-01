{{ with index . "Result" -}}
WORKSPACE
Name: {{ index . "Name"}}
About: {{ index . "About"}}
Frames: {{ $frames := index . "Frames" -}}
{{ if eq (len $frames) 0 -}}
none
{{ else }}
{{ range $frames }}- Name:   {{ index . "Name"}}
  About:  {{ index . "About"}}
  Framer: {{ index . "Framer"}}
  Shapes: {{ $shapes := index . "Shapes" -}}
{{ if eq (len $shapes) 0 -}}
none
{{ else }}
{{ range $shapes }}  - Name: {{ index . "Name"}}
    About: {{ index . "About"}}
    Shaper: {{ index . "Shaper"}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
Repositories: {{ $repositories := index . "Repositories" -}}
{{ if eq (len $repositories) 0 -}}
none
{{ else }}
{{ range $repositories }}- Path: {{ index . "Path"}}
  Git: {{ if not (index . "GitRepo") }}none{{ end }}
{{ with index . "GitRepo" }}    URI: {{ index . "URI"}}
    URL: {{ index . "URL"}}
    Branch: {{ index . "Branch"}}
    Shapers: {{ $shapers := index . "Shapers" -}}
{{ if eq (len $shapers) 0 -}}
none
{{ else }}
{{ range $shapers }}    - URI: {{ index . "URI"}}
      Name: {{ index . "Name"}}
      About: {{ index . "About"}}
{{ end -}}
{{ end }}    Framers: {{ $Framers := index . "Framers" -}}
{{ if eq (len $Framers) 0 -}}
none
{{ else }}
{{ range $Framers }}    - URI: {{ index . "URI"}}
      Name: {{ index . "Name"}}
      About: {{ index . "About"}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
Directories: {{ $directories := index . "Directories" -}}
{{ if eq (len $directories) 0 -}}
none
{{ else }}
{{ range $directories }}- Path: {{ index . "Path"}}
  Git: {{ if not (index . "GitRepo") }}none{{ end }} 
{{ with index . "GitRepo" }}    URI: {{ index . "URI"}}
    URL: {{ index . "URL"}}
    Branch: {{ index . "Branch"}}
    Shapers: {{ $shapers := index . "Shapers" -}}
{{ if eq (len $shapers) 0 -}}
none
{{ else }}
{{ range $shapers }}    - URI: {{ index . "URI"}}
      Name: {{ index . "Name"}}
      About: {{ index . "About"}}
{{ end -}}
{{ end }}    Framers: {{ $Framers := index . "Framers" -}}
{{ if eq (len $Framers) 0 -}}
none
{{ else }}
{{ range $Framers }}    - URI: {{ index . "URI"}}
      Name: {{ index . "Name"}}
      About: {{ index . "About"}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
