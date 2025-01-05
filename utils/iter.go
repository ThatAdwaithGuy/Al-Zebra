package utils

func Map[T, R any](slice []T, predicate func(T) R) []R {
  var result []R
  for _, ele := range slice {
    result = append(result, predicate(ele))
  }
  return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
  var result []T
  for _, ele := range slice {
    if predicate(ele) {
      result = append(result, ele)
    }
  }
  return result
}
