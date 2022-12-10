package hi

// Must panics if err is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must0 panics if err is not nil.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 returns v1 if err is nil, otherwise it panics.
func Must1[T1 any](v1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return v1
}

// Must2 returns (v1, v2) if err is nil, otherwise it panics.
func Must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// Must3 returns (v1, v2, v3) if err is nil, otherwise it panics.
func Must3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}

// Must4 returns (v1, v2, v3, v4) if err is nil, otherwise it panics.
func Must4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4
}

// Must5 returns (v1, v2, v3, v4, v5) if err is nil, otherwise it panics.
func Must5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4, v5
}

// Must6 returns (v1, v2, v3, v4, v5, v6) if err is nil, otherwise it panics.
func Must6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4, v5, v6
}

// Must7 returns (v1, v2, v3, v4, v5, v6, v7) if err is nil, otherwise it panics.
func Must7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4, v5, v6, v7
}

// Must8 returns (v1, v2, v3, v4, v5, v6, v7, v8) if err is nil, otherwise it panics.
func Must8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4, v5, v6, v7, v8
}
