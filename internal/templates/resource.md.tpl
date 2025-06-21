{{`---
# Resource **umbrella_{{ .ResourceName }}**
---

## Example Usage

```hcl
resource "umbrella_{{ .ResourceName }}" "example" {
{{- range .Fields }}
  {{ .Name }} = "..."
{{- end }}
}
```

## Argument Reference

{{- range .Fields }}* `{{ .Name }}` – {{ if .Required }}(Required){{ else }}(Optional){{ end }} {{ .Description }}
{{- end }}

## Attributes Reference

* `id` – Umbrella ID of the resource.
`}}
