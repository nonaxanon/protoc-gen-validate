package java

const wrapperConstTpl = `{{ $f := .Field }}{{ $r := .Rules }}			
			{{- renderConstants (unwrap .) }}`

const wrapperTpl = `{{ $f := .Field }}{{ $r := .Rules }}			
			if ({{ hasAccessor . }}) {
				{{- render (unwrap .) }}
			}
			{{ if .MessageRules.GetRequired }} else {
				throw new io.nonaxanon.pgv.ValidationException("{{ $f }}", "null", "is required");
			} {{ end }}`
