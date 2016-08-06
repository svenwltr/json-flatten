json-flatten
============

Flatten JSON data.

Why?!
-----

```bash
$ json-flatten generated-json | grep -i "interesting-value"
a.b=interesting-value
foo[0].bar=interesting-value

$ generate-json | jq 'foo[] | .bar' # for use in scripts
```

Usage
-----

```bash
json-flatten file.json
```

or

```bash
cat file.json | json-flatten
```
