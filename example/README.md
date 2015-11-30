## Generating the example apib files

### 1. run the tests

```bash
cd example
go test ./...
```

apib files are generated like so:

```
.
├── foos
│   ├── ...
│   └── foos.apib
└── widgets
    ├── ...
    └── widgets.apib
```

### 2. combine the doc files

```bash
./scripts/combine.sh
```

final apib doc file is generated at `example/apidoc.apib`