package goschemagen

import (
	"strings"

	"github.com/pablor21/gonnotation/annotations"
	"github.com/pablor21/gonnotation/parser"
)

// GetDescription extracts description from various sources based on configuration
func (c *Config) GetDescription(typeInfo interface{}) string {
	if c.UseCommentsAsDescription == nil || !*c.UseCommentsAsDescription {
		return ""
	}

	var comment string
	var typeAnnotations []annotations.Annotation

	switch ti := typeInfo.(type) {
	case *parser.StructInfo:
		comment = ti.Comment
		typeAnnotations = ti.Annotations
	case *parser.FieldInfo:
		comment = ti.Comment
		typeAnnotations = ti.Annotations
	case *parser.EnumInfo:
		comment = ti.Comment
		typeAnnotations = ti.Annotations
	case *parser.EnumValue:
		comment = ti.Comment
		typeAnnotations = ti.Annotations
	case *parser.InterfaceInfo:
		comment = ti.Comment
		typeAnnotations = ti.Annotations
	default:
		return ""
	}

	// First, try to get description from annotations
	if desc := getDescriptionFromAnnotations(typeAnnotations); desc != "" {
		return desc
	}

	// Fallback to extracted comment
	return strings.TrimSpace(comment)
}

// GetStructDescription extracts description for a struct
func (c *Config) GetStructDescription(structInfo *parser.StructInfo) string {
	return c.GetDescription(structInfo)
}

// GetFieldDescription extracts description for a field
func (c *Config) GetFieldDescription(fieldInfo *parser.FieldInfo) string {
	return c.GetDescription(fieldInfo)
}

// GetEnumDescription extracts description for an enum
func (c *Config) GetEnumDescription(enumInfo *parser.EnumInfo) string {
	return c.GetDescription(enumInfo)
}

// GetEnumValueDescription extracts description for an enum value
func (c *Config) GetEnumValueDescription(enumValue *parser.EnumValue) string {
	return c.GetDescription(enumValue)
}

// GetInterfaceDescription extracts description for an interface
func (c *Config) GetInterfaceDescription(interfaceInfo *parser.InterfaceInfo) string {
	return c.GetDescription(interfaceInfo)
}

// getDescriptionFromAnnotations searches for description in annotations
func getDescriptionFromAnnotations(annotations []annotations.Annotation) string {
	for _, ann := range annotations {
		switch ann.Name {
		case "description", "desc", "comment":
			if desc := ann.Params["description"]; desc != "" {
				return strings.TrimSpace(desc)
			}
			if desc := ann.Params["desc"]; desc != "" {
				return strings.TrimSpace(desc)
			}
			if desc := ann.Params["comment"]; desc != "" {
				return strings.TrimSpace(desc)
			}
		}
	}
	return ""
}
