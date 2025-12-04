package goschemagen

import (
	"github.com/pablor21/gonnotation/parser"
)

// FieldCase defines how field names should be transformed
type FieldCase string

const (
	FieldCaseCamel          FieldCase = "camel"
	FieldCaseSnake          FieldCase = "snake"
	FieldCasePascal         FieldCase = "pascal"
	FieldCaseOriginal       FieldCase = "original"
	FieldCaseNone           FieldCase = "none"
	FieldCaseScreamingSnake FieldCase = "screaming_snake"
	FieldCaseKebab          FieldCase = "kebab"
	FieldCaseLower          FieldCase = "lower"
	FieldCaseUpper          FieldCase = "upper"
)

// EnumValueStyle defines how enum values should be transformed
type EnumValueStyle string

const (
	EnumValueStyleName  EnumValueStyle = "name"
	EnumValueStyleValue EnumValueStyle = "value"
)

type Config struct {
	// Field naming
	FieldCase  *FieldCase `yaml:"field_case,omitempty"`
	UseJsonTag *bool      `yaml:"use_json_tag,omitempty"`

	// General naming transformations
	UseCommentsAsDescription *bool   `yaml:"use_comments_as_description,omitempty"`
	StripPrefix              *string `yaml:"strip_prefix,omitempty"`
	StripSuffix              *string `yaml:"strip_suffix,omitempty"`
	AddTypePrefix            *string `yaml:"add_type_prefix,omitempty"`
	AddTypeSuffix            *string `yaml:"add_type_suffix,omitempty"`

	// Enum naming transformations
	AddEnumPrefix   *string `yaml:"add_enum_prefix,omitempty"`
	AddEnumSuffix   *string `yaml:"add_enum_suffix,omitempty"`
	StripEnumPrefix *string `yaml:"strip_enum_prefix,omitempty"`
	StripEnumSuffix *string `yaml:"strip_enum_suffix,omitempty"`

	// Enum value transformations
	EnumValueCase      *FieldCase      `yaml:"enum_value_case,omitempty"`
	EnumValueStyle     *EnumValueStyle `yaml:"enum_value_style,omitempty"`
	IotaEnumValueStyle *EnumValueStyle `yaml:"iota_enum_value_style,omitempty"`

	// Type mappings
	TypeMappings map[string]string `yaml:"type_mappings"` // Go type -> Protobuf type

	// Imports
	KnownTypes map[string]KnownTypeMapping `yaml:"known_types,omitempty"`

	// Embed common configuration that can be overridden at plugin level
	parser.CommonConfig `yaml:",inline"`
}

func NewConfig() *Config {
	// // Parse the nested configuration structure
	// type ConfigFile struct {
	// 	Plugins struct {
	// 		Protobuf Config `yaml:"protobuf"`
	// 	} `yaml:"plugins"`
	// }

	cfg := &Config{
		CommonConfig: parser.NewConfigWithDefaults().CommonConfig,
	}
	// err := yaml.Unmarshal(DefaultConfigData, cfg)
	// if err != nil {
	// 	panic("failed to parse default config: " + err.Error())
	// }
	return cfg
}
