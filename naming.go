package goschemagen

import (
	"strings"

	"github.com/pablor21/gonnotation/parser"
)

// NamingStrategy handles naming transformations
type NamingStrategy struct {
	config *Config
}

// NewNamingStrategy creates naming strategy
func NewNamingStrategy(config *Config) *NamingStrategy {
	return &NamingStrategy{config: config}
}

// TransformFieldName transforms a field name
func (ns *NamingStrategy) TransformFieldName(name string) string {
	return TransformFieldName(name, parser.DerefPtr(ns.config.FieldCase, FieldCaseCamel))
}

// TransformTypeName applies type naming transformations (strip/add prefix/suffix)
// If customName is provided (from annotation), it is used instead
func (ns *NamingStrategy) TransformTypeName(name string, customName string) string {
	if customName != "" {
		return customName
	}
	return ns.applyTypeTransformations(name)
}

// TransformEnumName applies enum naming transformations (strip/add prefix/suffix)
// If customName is provided (from annotation), it is used instead
func (ns *NamingStrategy) TransformEnumName(name string, customName string) string {
	if customName != "" {
		return customName
	}
	return ns.applyEnumTransformations(name)
}

// TransformEnumValue transforms enum value name based on configured style
func (ns *NamingStrategy) TransformEnumValue(value string, isIota bool) string {
	var caseStyle FieldCase
	enumValueCase := parser.DerefPtr(ns.config.EnumValueCase, FieldCaseScreamingSnake)
	if isIota && parser.DerefPtr(ns.config.IotaEnumValueStyle, EnumValueStyleName) == EnumValueStyleName {
		caseStyle = enumValueCase
	} else if !isIota && parser.DerefPtr(ns.config.EnumValueStyle, EnumValueStyleValue) == EnumValueStyleName {
		caseStyle = enumValueCase
	} else {
		return value // Use original value
	}

	return TransformFieldName(value, caseStyle)
}

// applyTypeTransformations applies strip/add prefix/suffix for types
func (ns *NamingStrategy) applyTypeTransformations(name string) string {
	result := name

	// Strip prefixes
	stripPrefix := parser.DerefPtr(ns.config.StripPrefix, "")
	if stripPrefix != "" {
		prefixes := strings.Split(stripPrefix, ",")
		for _, prefix := range prefixes {
			prefix = strings.TrimSpace(prefix)
			if strings.HasPrefix(result, prefix) {
				result = strings.TrimPrefix(result, prefix)
				break
			}
		}
	}

	// Strip suffixes
	stripSuffix := parser.DerefPtr(ns.config.StripSuffix, "")
	if stripSuffix != "" {
		suffixes := strings.Split(stripSuffix, ",")
		for _, suffix := range suffixes {
			suffix = strings.TrimSpace(suffix)
			if strings.HasSuffix(result, suffix) {
				result = strings.TrimSuffix(result, suffix)
				break
			}
		}
	}

	// Add prefix
	addTypePrefix := parser.DerefPtr(ns.config.AddTypePrefix, "")
	if addTypePrefix != "" {
		result = addTypePrefix + result
	}

	// Add suffix
	addTypeSuffix := parser.DerefPtr(ns.config.AddTypeSuffix, "")
	if addTypeSuffix != "" {
		result = result + addTypeSuffix
	}

	return result
}

// applyEnumTransformations applies strip/add prefix/suffix for enums
func (ns *NamingStrategy) applyEnumTransformations(name string) string {
	result := name

	// Strip enum prefixes
	stripEnumPrefix := parser.DerefPtr(ns.config.StripEnumPrefix, "")
	if stripEnumPrefix != "" {
		prefixes := strings.Split(stripEnumPrefix, ",")
		for _, prefix := range prefixes {
			prefix = strings.TrimSpace(prefix)
			if strings.HasPrefix(result, prefix) {
				result = strings.TrimPrefix(result, prefix)
				break
			}
		}
	}

	// Strip enum suffixes
	stripEnumSuffix := parser.DerefPtr(ns.config.StripEnumSuffix, "")
	if stripEnumSuffix != "" {
		suffixes := strings.Split(stripEnumSuffix, ",")
		for _, suffix := range suffixes {
			suffix = strings.TrimSpace(suffix)
			if strings.HasSuffix(result, suffix) {
				result = strings.TrimSuffix(result, suffix)
				break
			}
		}
	}

	// Add enum prefix
	addEnumPrefix := parser.DerefPtr(ns.config.AddEnumPrefix, "")
	if addEnumPrefix != "" {
		result = addEnumPrefix + result
	}

	// Add enum suffix
	addEnumSuffix := parser.DerefPtr(ns.config.AddEnumSuffix, "")
	if addEnumSuffix != "" {
		result = result + addEnumSuffix
	}

	return result
}

// TransformFieldName transforms field name based on case
func TransformFieldName(name string, fieldCase FieldCase) string {
	switch fieldCase {
	case FieldCaseSnake:
		return ToSnakeCase(name)
	case FieldCasePascal:
		return name
	case FieldCaseOriginal:
		return name
	case FieldCaseScreamingSnake:
		return strings.ToUpper(ToSnakeCase(name))
	case FieldCaseKebab:
		return strings.ReplaceAll(ToSnakeCase(name), "_", "-")
	case FieldCaseLower:
		return strings.ToLower(name)
	case FieldCaseUpper:
		return strings.ToUpper(name)
	case FieldCaseCamel:
		fallthrough
	default:
		if len(name) == 0 {
			return name
		}
		// If the whole name is uppercase (acronym), just lowercase it
		allUpper := true
		for _, r := range name {
			if r >= 'a' && r <= 'z' {
				allUpper = false
				break
			}
		}
		if allUpper {
			return strings.ToLower(name)
		}
		// Detect leading acronym (sequence of capitals) preceding a capital+lowercase boundary
		// e.g. ID -> id, URLValue -> urlValue, HTTPServer -> httpServer
		if len(name) > 1 && isUpper(name[0]) && isUpper(name[1]) {
			boundary := -1
			for i := 1; i < len(name); i++ {
				if isUpper(name[i]) {
					// If next char exists and is lowercase, current position starts next word, boundary at i
					if i+1 < len(name) && !isUpper(name[i+1]) && isLetter(name[i+1]) && strings.ToLower(string(name[i+1])) != string(name[i+1]) {
						boundary = i
						break
					}
				} else {
					// first lowercase encountered inside acronym run -> boundary at previous char
					boundary = i - 1
					break
				}
			}
			if boundary == -1 {
				// Entire string was uppercase or no boundary found; already handled allUpper above, fallback to standard
				return strings.ToLower(name[:1]) + name[1:]
			}
			// Lower leading acronym segment
			return strings.ToLower(name[:boundary]) + name[boundary:]
		}
		// Standard lowerCamel: lowercase first character only
		return strings.ToLower(name[:1]) + name[1:]
	}
}

func isUpper(b byte) bool  { return b >= 'A' && b <= 'Z' }
func isLetter(b byte) bool { return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') }

// ToSnakeCase converts to snake_case
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	var out strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		isUpper := r >= 'A' && r <= 'Z'
		var prev rune
		var next rune
		if i > 0 {
			prev = runes[i-1]
		}
		if i+1 < len(runes) {
			next = runes[i+1]
		}
		prevUpper := prev >= 'A' && prev <= 'Z'
		prevLower := prev >= 'a' && prev <= 'z'
		nextLower := next >= 'a' && next <= 'z'

		// Insert underscore on transitions:
		// - lower/digit to upper
		// - acronym boundary: upper followed by upper then next lower (URLValue -> url_value)
		if i > 0 && isUpper && (prevLower || (prevUpper && nextLower)) {
			out.WriteRune('_')
		}
		out.WriteRune(r)
	}
	return strings.ToLower(out.String())
}
