package goschemagen

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/pablor21/gonnotation/annotations"
	"github.com/pablor21/gonnotation/parser"
)

// FieldProcessor handles field processing
type FieldProcessor struct {
	resolver *parser.TypeResolver
	config   *Config
}

// NewFieldProcessor creates field processor
func NewFieldProcessor(resolver *parser.TypeResolver, config *Config) *FieldProcessor {
	return &FieldProcessor{
		resolver: resolver,
		config:   config,
	}
}

// ProcessedField contains processed field info
type ProcessedField struct {
	GoName       string
	SchemaName   string
	GoType       ast.Expr
	ResolvedType string
	IsPointer    bool
	IsSlice      bool
	IsMap        bool
	IsEmbedded   bool
	Description  string
	Annotations  []annotations.Annotation
	Tags         annotations.StructTags
}

// ProcessField processes a field
func (fp *FieldProcessor) ProcessField(field *parser.FieldInfo) *ProcessedField {
	pf := &ProcessedField{
		GoName:      field.GoName,
		GoType:      field.Type,
		IsEmbedded:  field.IsEmbedded,
		Annotations: field.Annotations,
	}

	pf.IsPointer = fp.resolver.IsPointer(field.Type)
	pf.IsSlice = fp.resolver.IsSlice(field.Type)
	pf.IsMap = fp.resolver.IsMap(field.Type)
	pf.ResolvedType = fp.resolver.GetTypeName(field.Type)

	if field.Tag != nil {
		pf.Tags = fp.parseTags(field.Tag.Value)
	}

	pf.SchemaName = fp.resolveFieldName(pf.GoName, pf.Tags)

	return pf
}

// resolveFieldName resolves schema field name
func (fp *FieldProcessor) resolveFieldName(goName string, tags annotations.StructTags) string {
	structTagName := parser.DerefPtr(fp.config.StructTagName, "")
	if structTagName != "" {
		if tagValue := tags[structTagName]; tagValue != "" {
			parts := strings.Split(tagValue, ",")
			if parts[0] != "" && parts[0] != "-" {
				return parts[0]
			}
		}
	}

	if parser.DerefPtr(fp.config.UseJsonTag, true) {
		if jsonTag := tags["json"]; jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" && parts[0] != "-" {
				return parts[0]
			}
		}
	}

	return TransformFieldName(goName, parser.DerefPtr(fp.config.FieldCase, FieldCaseCamel))
}

// parseTags parses struct tags
func (fp *FieldProcessor) parseTags(tagStr string) annotations.StructTags {
	tagStr = strings.Trim(tagStr, "`")
	tag := reflect.StructTag(tagStr)
	tags := make(annotations.StructTags)

	for _, key := range []string{"json", "yaml", "xml", "gql", "graphql", "openapi", "description"} {
		if value := tag.Get(key); value != "" {
			tags[key] = value
		}
	}

	return tags
}
