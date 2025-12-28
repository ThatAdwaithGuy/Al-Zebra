package utils

type Stack[T any] []T

func (s *Stack[T]) PushBack(item T) {
	*s = append([]T{item}, *s...)
}

func (s *Stack[T]) PushFront(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) PeekBack() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	return &(*s)[0]
}

func (s *Stack[T]) PeekFront() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	return &(*s)[length-1]
}

func (s *Stack[T]) PopBack() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[0]
	*s = (*s)[1:]

	return &last
}

func (s *Stack[T]) PopFront() *T {
	length := len(*s)

	if length == 0 {
		return nil
	}
	last := (*s)[length-1]
	*s = (*s)[:length-1]

	return &last
}
