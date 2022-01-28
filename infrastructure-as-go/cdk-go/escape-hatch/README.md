# CDK Escape hatch in GO

This does not create useful resources - do NOT deploy! 
The generated CloudFormation just shows a concept.

Example from AWS [Docu](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#aws-properties-s3-bucket-analyticsconfiguration--examples) as unsupported property in a CDK Construct.

## Why typeconversion as in TypeScript does not work

### GO Conversions

See [conversions](https://go.dev/ref/spec#Conversions) in the GO doc.

TypeScript like cast as in the other CDK escape hatches are not possible in GO:

#### Conversions

A non-constant value x can be converted to type T in any of these cases:
- x is assignable to T.
...

#### Assignability
A value x is assignable to a variable of type T ("x is assignable to T") if one of the following conditions applies:

- x's type is identical to T.
- x's type V and T have identical underlying types and at least one of V or T is not a defined type.
- T is an interface type and x implements T.
- x is a bidirectional channel value, T is a channel type, x's type V and T have identical element types, and at least one of V or T is not a defined type.
- x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
- x is an untyped constant representable by a value of type T.

## Method 1

Convert json to GO struct.

See `escape-hatch.go` wintersoldier/cap

## Method 2

Use GO SDK types

See `escape-hatch.go` olaf/elsa

## GO on AWS

This is a code for the chapter Escape Hatches for https://www.go-on-aws.com/