// Code generated by generate_zip.pl; DO NOT EDIT.

package list

import (
	"github.com/shogo82148/hi/tuple"
)

// Zip2 returns a list of 2-tuples.
// The returned list have the length of the shortest list.
func Zip2[T1, T2 any](l1 *List[T1], l2 *List[T2]) *List[tuple.Tuple2[T1, T2]] {
	var ret List[tuple.Tuple2[T1, T2]]
	for e1, e2 := l1.Front(), l2.Front(); e1 != nil && e2 != nil; e1, e2 = e1.Next(), e2.Next() {
		ret.PushBack(tuple.Tuple2[T1, T2]{V1: e1.Value, V2: e2.Value})
	}
	return &ret
}

// Zip3 returns a list of 3-tuples.
// The returned list have the length of the shortest list.
func Zip3[T1, T2, T3 any](l1 *List[T1], l2 *List[T2], l3 *List[T3]) *List[tuple.Tuple3[T1, T2, T3]] {
	var ret List[tuple.Tuple3[T1, T2, T3]]
	for e1, e2, e3 := l1.Front(), l2.Front(), l3.Front(); e1 != nil && e2 != nil && e3 != nil; e1, e2, e3 = e1.Next(), e2.Next(), e3.Next() {
		ret.PushBack(tuple.Tuple3[T1, T2, T3]{V1: e1.Value, V2: e2.Value, V3: e3.Value})
	}
	return &ret
}

// Zip4 returns a list of 4-tuples.
// The returned list have the length of the shortest list.
func Zip4[T1, T2, T3, T4 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4]) *List[tuple.Tuple4[T1, T2, T3, T4]] {
	var ret List[tuple.Tuple4[T1, T2, T3, T4]]
	for e1, e2, e3, e4 := l1.Front(), l2.Front(), l3.Front(), l4.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil; e1, e2, e3, e4 = e1.Next(), e2.Next(), e3.Next(), e4.Next() {
		ret.PushBack(tuple.Tuple4[T1, T2, T3, T4]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value})
	}
	return &ret
}

// Zip5 returns a list of 5-tuples.
// The returned list have the length of the shortest list.
func Zip5[T1, T2, T3, T4, T5 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5]) *List[tuple.Tuple5[T1, T2, T3, T4, T5]] {
	var ret List[tuple.Tuple5[T1, T2, T3, T4, T5]]
	for e1, e2, e3, e4, e5 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil; e1, e2, e3, e4, e5 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next() {
		ret.PushBack(tuple.Tuple5[T1, T2, T3, T4, T5]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value})
	}
	return &ret
}

// Zip6 returns a list of 6-tuples.
// The returned list have the length of the shortest list.
func Zip6[T1, T2, T3, T4, T5, T6 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6]) *List[tuple.Tuple6[T1, T2, T3, T4, T5, T6]] {
	var ret List[tuple.Tuple6[T1, T2, T3, T4, T5, T6]]
	for e1, e2, e3, e4, e5, e6 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil; e1, e2, e3, e4, e5, e6 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next() {
		ret.PushBack(tuple.Tuple6[T1, T2, T3, T4, T5, T6]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value})
	}
	return &ret
}

// Zip7 returns a list of 7-tuples.
// The returned list have the length of the shortest list.
func Zip7[T1, T2, T3, T4, T5, T6, T7 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7]) *List[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	var ret List[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]]
	for e1, e2, e3, e4, e5, e6, e7 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil; e1, e2, e3, e4, e5, e6, e7 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next() {
		ret.PushBack(tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value})
	}
	return &ret
}

// Zip8 returns a list of 8-tuples.
// The returned list have the length of the shortest list.
func Zip8[T1, T2, T3, T4, T5, T6, T7, T8 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8]) *List[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	var ret List[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]
	for e1, e2, e3, e4, e5, e6, e7, e8 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil; e1, e2, e3, e4, e5, e6, e7, e8 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next() {
		ret.PushBack(tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value})
	}
	return &ret
}

// Zip9 returns a list of 9-tuples.
// The returned list have the length of the shortest list.
func Zip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9]) *List[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	var ret List[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next() {
		ret.PushBack(tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value})
	}
	return &ret
}

// Zip10 returns a list of 10-tuples.
// The returned list have the length of the shortest list.
func Zip10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10]) *List[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	var ret List[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next() {
		ret.PushBack(tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value})
	}
	return &ret
}

// Zip11 returns a list of 11-tuples.
// The returned list have the length of the shortest list.
func Zip11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11]) *List[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]] {
	var ret List[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next() {
		ret.PushBack(tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value})
	}
	return &ret
}

// Zip12 returns a list of 12-tuples.
// The returned list have the length of the shortest list.
func Zip12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11], l12 *List[T12]) *List[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]] {
	var ret List[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(), l12.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil && e12 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next(), e12.Next() {
		ret.PushBack(tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value, V12: e12.Value})
	}
	return &ret
}

// Zip13 returns a list of 13-tuples.
// The returned list have the length of the shortest list.
func Zip13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11], l12 *List[T12], l13 *List[T13]) *List[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]] {
	var ret List[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(), l12.Front(), l13.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil && e12 != nil && e13 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next(), e12.Next(), e13.Next() {
		ret.PushBack(tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value, V12: e12.Value, V13: e13.Value})
	}
	return &ret
}

// Zip14 returns a list of 14-tuples.
// The returned list have the length of the shortest list.
func Zip14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11], l12 *List[T12], l13 *List[T13], l14 *List[T14]) *List[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]] {
	var ret List[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(), l12.Front(), l13.Front(), l14.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil && e12 != nil && e13 != nil && e14 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next(), e12.Next(), e13.Next(), e14.Next() {
		ret.PushBack(tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value, V12: e12.Value, V13: e13.Value, V14: e14.Value})
	}
	return &ret
}

// Zip15 returns a list of 15-tuples.
// The returned list have the length of the shortest list.
func Zip15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11], l12 *List[T12], l13 *List[T13], l14 *List[T14], l15 *List[T15]) *List[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]] {
	var ret List[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(), l12.Front(), l13.Front(), l14.Front(), l15.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil && e12 != nil && e13 != nil && e14 != nil && e15 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next(), e12.Next(), e13.Next(), e14.Next(), e15.Next() {
		ret.PushBack(tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value, V12: e12.Value, V13: e13.Value, V14: e14.Value, V15: e15.Value})
	}
	return &ret
}

// Zip16 returns a list of 16-tuples.
// The returned list have the length of the shortest list.
func Zip16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](l1 *List[T1], l2 *List[T2], l3 *List[T3], l4 *List[T4], l5 *List[T5], l6 *List[T6], l7 *List[T7], l8 *List[T8], l9 *List[T9], l10 *List[T10], l11 *List[T11], l12 *List[T12], l13 *List[T13], l14 *List[T14], l15 *List[T15], l16 *List[T16]) *List[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]] {
	var ret List[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]]
	for e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16 := l1.Front(), l2.Front(), l3.Front(), l4.Front(), l5.Front(), l6.Front(), l7.Front(), l8.Front(), l9.Front(), l10.Front(), l11.Front(), l12.Front(), l13.Front(), l14.Front(), l15.Front(), l16.Front(); e1 != nil && e2 != nil && e3 != nil && e4 != nil && e5 != nil && e6 != nil && e7 != nil && e8 != nil && e9 != nil && e10 != nil && e11 != nil && e12 != nil && e13 != nil && e14 != nil && e15 != nil && e16 != nil; e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16 = e1.Next(), e2.Next(), e3.Next(), e4.Next(), e5.Next(), e6.Next(), e7.Next(), e8.Next(), e9.Next(), e10.Next(), e11.Next(), e12.Next(), e13.Next(), e14.Next(), e15.Next(), e16.Next() {
		ret.PushBack(tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{V1: e1.Value, V2: e2.Value, V3: e3.Value, V4: e4.Value, V5: e5.Value, V6: e6.Value, V7: e7.Value, V8: e8.Value, V9: e9.Value, V10: e10.Value, V11: e11.Value, V12: e12.Value, V13: e13.Value, V14: e14.Value, V15: e15.Value, V16: e16.Value})
	}
	return &ret
}
