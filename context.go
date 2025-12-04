package goschemagen

import "github.com/pablor21/gonnotation/parser"

type GenerationContext struct {
	Config         *Config
	FieldProcessor *FieldProcessor
	NamingStrategy *NamingStrategy
	parser.GenerationContext
}

// func NewGenerationContext() *GenerationContext {
// 	return &GenerationContext{
// 		Config:         NewConfig(),
// 		GenerationContext: *parser.NewGenerationContext(),
// 	}
// }
