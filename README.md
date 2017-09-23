# binquiry


[![Travis CI Build Status](https://travis-ci.org/marthjod/binquiry.svg?branch=master)](https://travis-ci.org/marthjod/binquiry)

[BÍN](http://bin.arnastofnun.is) "AJAX" parser/converter.

Also see [binquiry-web](https://github.com/marthjod/binquiry-web).

## Recognized inputs

- [x] Unambiguous nouns in nominative singular
- [x] Ambiguous nouns in nominative singular
- [x] (experimental) Unambiguous adjectives in nominative singular

## Examples

### JSON output

```bash
./bingo -q orð -f json | jq '.[] | .cases[].form'
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
./bingo -q kona -f json | jq '.[] | .cases[] | select(.number == "Plural" and .case == "Accusative")'
```

yields

```json
{
  "case": "Accusative",
  "number": "Plural",
  "form": "konur"
}
```

Ambiguous input fetches and returns all suggested paradigms.

```bash
$ ./bingo -q ár | jq '. | length'
3
```

## Dependencies

- https://github.com/go-xmlpath/xmlpath
- https://github.com/campoy/jsonenums

