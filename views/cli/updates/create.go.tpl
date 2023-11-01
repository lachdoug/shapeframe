{{ with index . "Result" -}}
Updated {{ index . "Kind"}}
FROM
{{ with index . "From" }}  Name: {{ index . "Name" }}
  About: {{ index . "About" }}
{{ end -}}
TO
{{ with index . "To" }}  Name: {{ index . "Name" }}
  About: {{ index . "About" }}
{{ end -}}
{{ end -}}