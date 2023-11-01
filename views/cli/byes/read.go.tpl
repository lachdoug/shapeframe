{{ with index . "Result" -}}
{{ if eq (index . "Tone") "sad" -}}
Thank you {{ index . "Name" }} for the life we shared
{{ else -}}
Catch you later {{ index . "Name" }}!
{{ end -}}
{{ end -}}