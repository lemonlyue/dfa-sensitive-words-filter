# dfa-sensitive-words-filter
> Sensitive word filtering based on DFA algorithm Golang language.
> This project has 100% unit test coverage.

## Install
> go get github.com/lemon_lyue/dfa-sensitive-words-filter

## Example
```go
// get sensitive words filter instance
filter := SensitiveWordsFilter.GetInstance()

// build
filter.Build([]string{"丑"})

// set level
filter.SetLevel(SensitiveWordsFilter.FilterLevelHight)

// set skip distance
filter.SetSkipDistance(1)

// set replace text
filter.SetReplaceText("帅")

before := "你好丑"
// filter
after, hasSensitiveWords := filter.Filter(before)

fmt.Println("filter before: ", before)
fmt.Println("filter after: ", after)
fmt.Println("has sensitive words: ", hasSensitiveWords)
```

## Unit Test
1. Run Test
```shell script
go test ./ -cover
```
