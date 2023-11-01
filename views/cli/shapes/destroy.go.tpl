{{ with index . "Result" -}}
Destroyed shape {{ index . "Name" }} in frame {{ index . "Frame" }}
{{ end -}}