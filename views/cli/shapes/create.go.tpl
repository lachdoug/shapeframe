{{ with index . "Result" -}}
Created shape {{ index . "Name" }} in frame {{ index . "Frame" }}
{{ end -}}