// Code generated by generate_zip.pl; DO NOT EDIT.

package optional

import (
	"github.com/shogo82148/hi/tuple"
)

// Zip2 returns an optional of valid 2-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip2[T1, T2 any](o1 Optional[T1], o2 Optional[T2]) Optional[tuple.Tuple2[T1, T2]] {
	if o1.valid && o2.valid {
		return Optional[tuple.Tuple2[T1, T2]]{valid: true, value: tuple.Tuple2[T1, T2]{V1: o1.value, V2: o2.value}}
	} else {
		return Optional[tuple.Tuple2[T1, T2]]{}
	}
}

// Zip3 returns an optional of valid 3-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip3[T1, T2, T3 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3]) Optional[tuple.Tuple3[T1, T2, T3]] {
	if o1.valid && o2.valid && o3.valid {
		return Optional[tuple.Tuple3[T1, T2, T3]]{valid: true, value: tuple.Tuple3[T1, T2, T3]{V1: o1.value, V2: o2.value, V3: o3.value}}
	} else {
		return Optional[tuple.Tuple3[T1, T2, T3]]{}
	}
}

// Zip4 returns an optional of valid 4-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip4[T1, T2, T3, T4 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4]) Optional[tuple.Tuple4[T1, T2, T3, T4]] {
	if o1.valid && o2.valid && o3.valid && o4.valid {
		return Optional[tuple.Tuple4[T1, T2, T3, T4]]{valid: true, value: tuple.Tuple4[T1, T2, T3, T4]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value}}
	} else {
		return Optional[tuple.Tuple4[T1, T2, T3, T4]]{}
	}
}

// Zip5 returns an optional of valid 5-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip5[T1, T2, T3, T4, T5 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5]) Optional[tuple.Tuple5[T1, T2, T3, T4, T5]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid {
		return Optional[tuple.Tuple5[T1, T2, T3, T4, T5]]{valid: true, value: tuple.Tuple5[T1, T2, T3, T4, T5]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value}}
	} else {
		return Optional[tuple.Tuple5[T1, T2, T3, T4, T5]]{}
	}
}

// Zip6 returns an optional of valid 6-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip6[T1, T2, T3, T4, T5, T6 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6]) Optional[tuple.Tuple6[T1, T2, T3, T4, T5, T6]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid {
		return Optional[tuple.Tuple6[T1, T2, T3, T4, T5, T6]]{valid: true, value: tuple.Tuple6[T1, T2, T3, T4, T5, T6]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value}}
	} else {
		return Optional[tuple.Tuple6[T1, T2, T3, T4, T5, T6]]{}
	}
}

// Zip7 returns an optional of valid 7-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip7[T1, T2, T3, T4, T5, T6, T7 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7]) Optional[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid {
		return Optional[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]]{valid: true, value: tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value}}
	} else {
		return Optional[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]]{}
	}
}

// Zip8 returns an optional of valid 8-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip8[T1, T2, T3, T4, T5, T6, T7, T8 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8]) Optional[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid {
		return Optional[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]{valid: true, value: tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value}}
	} else {
		return Optional[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]{}
	}
}

// Zip9 returns an optional of valid 9-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9]) Optional[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid {
		return Optional[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]{valid: true, value: tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value}}
	} else {
		return Optional[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]{}
	}
}

// Zip10 returns an optional of valid 10-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10]) Optional[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid {
		return Optional[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]{valid: true, value: tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value}}
	} else {
		return Optional[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]{}
	}
}

// Zip11 returns an optional of valid 11-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11]) Optional[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid {
		return Optional[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]]{valid: true, value: tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value}}
	} else {
		return Optional[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]]{}
	}
}

// Zip12 returns an optional of valid 12-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11], o12 Optional[T12]) Optional[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid && o12.valid {
		return Optional[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]]{valid: true, value: tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value, V12: o12.value}}
	} else {
		return Optional[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]]{}
	}
}

// Zip13 returns an optional of valid 13-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11], o12 Optional[T12], o13 Optional[T13]) Optional[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid && o12.valid && o13.valid {
		return Optional[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]]{valid: true, value: tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value, V12: o12.value, V13: o13.value}}
	} else {
		return Optional[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]]{}
	}
}

// Zip14 returns an optional of valid 14-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11], o12 Optional[T12], o13 Optional[T13], o14 Optional[T14]) Optional[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid && o12.valid && o13.valid && o14.valid {
		return Optional[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]]{valid: true, value: tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value, V12: o12.value, V13: o13.value, V14: o14.value}}
	} else {
		return Optional[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]]{}
	}
}

// Zip15 returns an optional of valid 15-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11], o12 Optional[T12], o13 Optional[T13], o14 Optional[T14], o15 Optional[T15]) Optional[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid && o12.valid && o13.valid && o14.valid && o15.valid {
		return Optional[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]]{valid: true, value: tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value, V12: o12.value, V13: o13.value, V14: o14.value, V15: o15.value}}
	} else {
		return Optional[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]]{}
	}
}

// Zip16 returns an optional of valid 16-tuples, if all of arguments are valid.
// Otherwise, one or more arguments are not valid, it returns none
func Zip16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](o1 Optional[T1], o2 Optional[T2], o3 Optional[T3], o4 Optional[T4], o5 Optional[T5], o6 Optional[T6], o7 Optional[T7], o8 Optional[T8], o9 Optional[T9], o10 Optional[T10], o11 Optional[T11], o12 Optional[T12], o13 Optional[T13], o14 Optional[T14], o15 Optional[T15], o16 Optional[T16]) Optional[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]] {
	if o1.valid && o2.valid && o3.valid && o4.valid && o5.valid && o6.valid && o7.valid && o8.valid && o9.valid && o10.valid && o11.valid && o12.valid && o13.valid && o14.valid && o15.valid && o16.valid {
		return Optional[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]]{valid: true, value: tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{V1: o1.value, V2: o2.value, V3: o3.value, V4: o4.value, V5: o5.value, V6: o6.value, V7: o7.value, V8: o8.value, V9: o9.value, V10: o10.value, V11: o11.value, V12: o12.value, V13: o13.value, V14: o14.value, V15: o15.value, V16: o16.value}}
	} else {
		return Optional[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]]{}
	}
}
