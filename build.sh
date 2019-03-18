

template_tool --template-file=set/set.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ]
}' > set/set.generate.go && gofmt -w set/set.generate.go

template_tool --template-file=math/math.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "float32",
        "float64"
    ]
}' > math/math.generate.go  && gofmt -w math/math.generate.go

template_tool --template-file=sorts/sorts.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "float32",
        "float64",
        "string",
        "rune"
    ]
}' > sorts/sorts.generate.go && gofmt -w sorts/sorts.generate.go

template_tool --template-file=ternary/triple_get.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ]
}' > ternary/tripleget.generate.go && gofmt -w ternary/tripleget.generate.go


template_tool --template-file=typeconv/ptr.gotmpl --template-map='{
    "Types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string",
        "bool"
    ]
}' > typeconv/ptr.generate.go && gofmt -w typeconv/ptr.generate.go


template_tool  --template-file=slice/slice.gotmpl --template-map='{
    "keyTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string",
        "float32",
        "float64",
        "bool",
        "interface"
    ],
    "valueTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string",
        "float32",
        "float64",
        "bool"
    ]
}' > slice/slice.generate.go && gofmt -w slice/slice.generate.go


template_tool --template-file=slice/slice.unique.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ]
}' > slice/slice.unique.generate.go && gofmt -w slice/slice.unique.generate.go

template_tool  --template-file=slice/slice.find.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ]
}' > slice/slice.find.generate.go && gofmt -w slice/slice.find.generate.go

build_tools --target=template --template-file=assert/assert.gotmpl --template-map='{
    "types": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "float32",
        "float64",
        "string"
    ],
    "numTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "float32",
        "float64"
    ]
}' > assert/assert.generate.go && gofmt -w assert/assert.generate.go

template_tool --template-file=mapdefault/default.gotmpl --template-map='{
    "keyTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ],
    "valueTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string",
        "float32",
        "float64",
        "interface{}",
        "bool",
        "struct{}"
    ]
}' > mapdefault/default.generate.go && gofmt -w mapdefault/default.generate.go

template_tool --template-file=ordermap/order_map.gotmpl --template-map='{
    "keyTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string"
    ],
    "valueTypes": [
        "int",
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "string",
        "float32",
        "float64",
        "interface{}",
        "bool",
        "struct{}"
    ]
}' > ordermap/ordered_map.generate.go && gofmt -w ordermap/ordered_map.generate.go