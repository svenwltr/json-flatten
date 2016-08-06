json-flatten
============

Flatten JSON data.

[![license](https://img.shields.io/github/license/svenwltr/json-flatten.svg?maxAge=2592000?style=flat)](https://github.com/svenwltr/json-flatten/blob/master/LICENSE)
[![Latest Stable Version](https://img.shields.io/github/release/svenwltr/json-flatten.svg?style=flat)](https://github.com/svenwltr/json-flatten/releases)
[![Travis](https://img.shields.io/travis/svenwltr/json-flatten.svg?maxAge=2592000?style=flat)](https://travis-ci.org/svenwltr/json-flatten)
[![GoDoc](https://godoc.org/github.com/svenwltr/json-flatten?status.png)](https://godoc.org/github.com/svenwltr/json-flatten)
[![Go Report Card](https://goreportcard.com/badge/svenwltr/json-flatten)](http://goreportcard.com/report/svenwltr/json-flatten)

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
