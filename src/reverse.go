package polynomials

import (
    "sort"
)

// SOURCE: https://stackoverflow.com/a/71904070


func Reverse[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}