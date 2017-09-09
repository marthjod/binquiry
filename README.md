# bingo


[![Travis CI Build Status](https://travis-ci.org/marthjod/bingo.svg?branch=master)](https://travis-ci.org/marthjod/bingo)

[BÍN](http://bin.arnastofnun.is) parser/converter.

## Recognized inputs

- [x] Unambiguous nouns in nominative singular

## Examples

### JSON output

```bash
./bingo -q orð -f json | jq '.cases[].form'
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
./bingo -q kona -f json | jq '.cases[] | select(.number == "Plural" and .name == "Accusative")'
```

yields

```json
{
  "name": "Accusative",
  "number": "Plural",
  "form": "konur"
}
```

### List output


```bash
./bingo -q penni -f list
[penni penna penna penna pennar penna pennum penna]
```

## Dependencies

- https://github.com/campoy/jsonenums

