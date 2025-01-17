package java

const enumConstTpl = `{{ $ctx := . }}{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In }}
			{{ javaTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn }}
			{{ javaTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}`

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			io.nonaxanon.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, 
				{{ javaTypeFor . }}.forNumber({{ $r.GetConst }}));
{{- end -}}
{{- if $r.GetDefinedOnly }}
			io.nonaxanon.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.In }}
			io.nonaxanon.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			io.nonaxanon.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
`
