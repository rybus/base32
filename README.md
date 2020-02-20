# base32
Golang base32 tool


## Installation

```bash
go install github.com/rybus/base32
```

add `~/go/bin` to your `$PATH`.

## Usage

```bash
echo "foo" | base32
MZXW6===

echo "MZXW6===" | base32 --decode
"foo"
```
