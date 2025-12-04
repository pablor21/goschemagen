# GOSCHEMAGEN

This package contains utilities to generate schema files (protobuf, openapi, gql, etc.) from Go structs.
It is designed to be extensible and customizable, allowing users to define how Go types are mapped to schema definitions.

It leverages the `parser` package to analyze Go code and extract type information, and provides a plugin architecture to support different schema formats.

Please refer to the documentation  in the parser package for details on how to use the type resolver and core configuration options.
[https://github.com/pablor21/gonnotation](https://github.com/pablor21/gonnotation)

## Examples of code generators built with this package:

- [gqlschemagen](https://github.com/pablor21/gqlschemagen)
- [oasgen](https://github.com/pablor21/oaschemagen)
- [protoschemagen](https://github.com/pablor21/protoschemagen)


## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.