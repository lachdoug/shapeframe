{{ with index . "Result" -}}
FRAME
From:
{{ with index . "From" }}  Name: {{ index . "Name" }}
  About: {{ index . "About" }}
{{ end -}}
To:
{{ with index . "To" }}  Name: {{ index . "Name" }}
  About: {{ index . "About" }}
{{ end -}}
{{ end -}}