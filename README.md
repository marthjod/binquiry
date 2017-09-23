# binquiry


[![Travis CI Build Status](https://travis-ci.org/marthjod/binquiry.svg?branch=master)](https://travis-ci.org/marthjod/binquiry)

[BÍN](http://bin.arnastofnun.is) "AJAX" parser/converter.

Also see [binquiry-web](https://github.com/marthjod/binquiry-web).

## Recognized inputs

- [x] Unambiguous nouns in nominative singular
- [x] Ambiguous nouns in nominative singular
- [x] Unambiguous adjectives in nominative singular (experimental)

## Examples

### JSON output

```bash
./binquiry | jq '.[] | .cases[].form'
```

yields

```json
"orð"
"orð"
"orði"
"orðs"
"orð"
"orð"
"orðum"
"orða"
```

```bash
./binquiry -q kona | \
jq '.[]
    | .cases[]
    | select(
        .number == "Plural" and
        .case == "Accusative"
    )'
```

yields

```json
{
  "case": "Accusative",
  "number": "Plural",
  "form": "konur"
}
```

```bash
./binquiry -q gamall | \
jq '.[]
    | .cases[]
    | select(
        .gender == "Neuter" and
        .declension == "Weak" and
        .number == "Plural" and
        .degree == "Superlative"
    )
    | .form'
```

yields

```json
"elstu"
"elstu"
"elstu"
"elstu"
```

Ambiguous input fetches and returns all suggested paradigms.

```bash
$ ./binquiry -q ár | jq '. | length'
3
```

## Dependencies

- https://github.com/go-xmlpath/xmlpath
- https://github.com/campoy/jsonenums

