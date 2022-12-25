package hi

// If is a 1 line if/else expression.
func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	} else {
		return falseValue
	}
}

type SwitchCase[T any, P comparable] struct {
	predicate P
	result    T
	done      bool
}

func Switch[T any, P comparable](predicate P) SwitchCase[T, P] {
	var result T
	return SwitchCase[T, P]{
		predicate: predicate,
		result:    result,
		done:      false,
	}
}

func (s SwitchCase[T, P]) Case(val P, result T) SwitchCase[T, P] {
	if s.done {
		return s
	}
	if s.predicate == val {
		return SwitchCase[T, P]{
			predicate: s.predicate,
			result:    result,
			done:      true,
		}
	}
	return s
}

func (s SwitchCase[T, P]) Default(result T) T {
	if s.done {
		return s.result
	}
	return result
}
