# Quicktype to generate W3C Thing Description Models for your Language

## References

* [W3C WoT Thing Description](https://www.w3.org/TR/wot-thing-description11/)
* [W3C WoT Modbus Binding](https://w3c.github.io/wot-binding-templates/bindings/protocols/modbus/)

## Prerequisits

* NodeJS: v18.16.0
* npm: 9.5.1
* ~/.npm-packages/bin should be in your PATH
* Toolchain for your language of choice

## Installation

```
git clone git@github.com:hadjian/quicktype-td.git
npm install -g quicktype
chmod 755 ~/.npm-packages/bin/quicktype
cd quicktype-td
mkdir <name-of-programming-language>
```

In the last command, you should use the name of the programming language you want to use, e.g. "go", "C++", "JavaScript" etc. See the official [quicktype project](https://github.com/quicktype/quicktype) for a list of supported languages.

In the quicktype-td project you will already find three subdirectories:

```
- go      // Example project for the go programming language
- schemas // offical schemas for TD and Modbus binding
- things  // Example things
```

Now we can use ```quicktype``` to generate model files for a range of supported languages as follows:

```
quicktype --src-lang schema --lang go --out go/td-parser.go schemas/td-schema.json
```

The above command will generate model files for the go programming language.

```
quicktype --src-lang schema --lang go --out go/modbus-parser.go schemas/modbus.schema.json
```

Now you can add these model files to your project and make use of them to parse a TD with Modbus bindings. The example ```go``` directory contains a few lines of code to demonstrate this. But before you can compile your project, we need to employ a small manual hack, which we will try to automate in the future. Read on.

## Manual Fix

Now you have your model files in your language sub-dicrectory, but the Modbus form class and the TD form class are separate. Use your programming language abstraction to inherit the TD form class from the Modbus Form class. In go for example we can use struct embeddings like so:

```go
type FormElementProperty struct {
	ModbusForm
	Op                  *FormElementPropertyOp          `json:"op"`
	AdditionalResponses []AdditionalResponsesDefinition `json:"additionalResponses,omitempty"`
	ContentCoding       *string                         `json:"contentCoding,omitempty"`
	ContentType         *string                         `json:"contentType,omitempty"`
	Href                string                          `json:"href"`
	Response            *ExpectedResponse               `json:"response,omitempty"`
	Scopes              *TypeDeclaration                `json:"scopes"`
	Security            *TypeDeclaration                `json:"security"`
	Subprotocol         *string                         `json:"subprotocol,omitempty"`
}
```

Now you can use the the generated ```UnmarshalTdParser``` function to parse raw bytes read from the ```things/pac4200.json``` file into an instance of the TdParser class/struct. See ```go/main.go``` for reference.
