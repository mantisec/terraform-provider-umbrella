{{`---
# Data Source **umbrella_{{ .ResourceName }}**
---

## Example Usage

```hcl
data "umbrella_{{ .ResourceName }}" "this" {
  id = var.{{ .ResourceName }}_id
}
```

## Attributes Reference

{{- range .Fields }}* `{{ .Name }}` – {{ .Description }}
{{- end }}
`}}
