package goschemagen

type KnownTypeMapping struct {
	Model []string // Go type patterns
	Type  string   // Protobuf type
}
