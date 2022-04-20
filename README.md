# flatuuid

Makes old kind UUIDs more compact.

## Why?!
Because I use UUID as a part of models and URLs in my APIs. But I don't like generated unique string which contains unuseful data like "-".

Also, I'm tired of copying from project to project the code for removing these pieces of string.

## Install
```
go get github.com/ixalender/flatuuid
```

## Usage

```go
compactUUID, err := flatuuid.Compact("f47ac10b-58cc-4372-0567-0e02b2c3d479")
// f47ac10b58cc437205670e02b2c3d479

rfc4122UUID, err := flatuuid.Compact("urn:uuid:f47ac10b-58cc-4372-0567-0e02b2c3d479")
// f47ac10b58cc437205670e02b2c3d479

microsoftGUID, err := flatuuid.Compact("{78B83322-84F6-46E8-A0BE-238FB5703A09}")
// 78b8332284f646e8a0be238fb5703a09
```

## License

MIT