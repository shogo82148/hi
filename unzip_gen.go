// Code generated by generate_unzip.pl; DO NOT EDIT.
package hi

import (
	"github.com/shogo82148/hi/tuple"
)

// Unzip2 converts a slice of 2-tuple to slices of each elements.
func Unzip2[T1, T2 any](s []tuple.Tuple2[T1, T2]) ([]T1, []T2) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
	}
	return s1, s2
}

// Unzip3 converts a slice of 3-tuple to slices of each elements.
func Unzip3[T1, T2, T3 any](s []tuple.Tuple3[T1, T2, T3]) ([]T1, []T2, []T3) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
	}
	return s1, s2, s3
}

// Unzip4 converts a slice of 4-tuple to slices of each elements.
func Unzip4[T1, T2, T3, T4 any](s []tuple.Tuple4[T1, T2, T3, T4]) ([]T1, []T2, []T3, []T4) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
	}
	return s1, s2, s3, s4
}

// Unzip5 converts a slice of 5-tuple to slices of each elements.
func Unzip5[T1, T2, T3, T4, T5 any](s []tuple.Tuple5[T1, T2, T3, T4, T5]) ([]T1, []T2, []T3, []T4, []T5) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
	}
	return s1, s2, s3, s4, s5
}

// Unzip6 converts a slice of 6-tuple to slices of each elements.
func Unzip6[T1, T2, T3, T4, T5, T6 any](s []tuple.Tuple6[T1, T2, T3, T4, T5, T6]) ([]T1, []T2, []T3, []T4, []T5, []T6) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
	}
	return s1, s2, s3, s4, s5, s6
}

// Unzip7 converts a slice of 7-tuple to slices of each elements.
func Unzip7[T1, T2, T3, T4, T5, T6, T7 any](s []tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
	}
	return s1, s2, s3, s4, s5, s6, s7
}

// Unzip8 converts a slice of 8-tuple to slices of each elements.
func Unzip8[T1, T2, T3, T4, T5, T6, T7, T8 any](s []tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
	}
	return s1, s2, s3, s4, s5, s6, s7, s8
}

// Unzip9 converts a slice of 9-tuple to slices of each elements.
func Unzip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](s []tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9
}

// Unzip10 converts a slice of 10-tuple to slices of each elements.
func Unzip10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](s []tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10
}

// Unzip11 converts a slice of 11-tuple to slices of each elements.
func Unzip11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](s []tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11
}

// Unzip12 converts a slice of 12-tuple to slices of each elements.
func Unzip12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](s []tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11, []T12) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	s12 := make([]T12, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
		s12[i] = t.V12
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12
}

// Unzip13 converts a slice of 13-tuple to slices of each elements.
func Unzip13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](s []tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11, []T12, []T13) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	s12 := make([]T12, len(s))
	s13 := make([]T13, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
		s12[i] = t.V12
		s13[i] = t.V13
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13
}

// Unzip14 converts a slice of 14-tuple to slices of each elements.
func Unzip14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](s []tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11, []T12, []T13, []T14) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	s12 := make([]T12, len(s))
	s13 := make([]T13, len(s))
	s14 := make([]T14, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
		s12[i] = t.V12
		s13[i] = t.V13
		s14[i] = t.V14
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14
}

// Unzip15 converts a slice of 15-tuple to slices of each elements.
func Unzip15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](s []tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11, []T12, []T13, []T14, []T15) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	s12 := make([]T12, len(s))
	s13 := make([]T13, len(s))
	s14 := make([]T14, len(s))
	s15 := make([]T15, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
		s12[i] = t.V12
		s13[i] = t.V13
		s14[i] = t.V14
		s15[i] = t.V15
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15
}

// Unzip16 converts a slice of 16-tuple to slices of each elements.
func Unzip16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](s []tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11, []T12, []T13, []T14, []T15, []T16) {
	s1 := make([]T1, len(s))
	s2 := make([]T2, len(s))
	s3 := make([]T3, len(s))
	s4 := make([]T4, len(s))
	s5 := make([]T5, len(s))
	s6 := make([]T6, len(s))
	s7 := make([]T7, len(s))
	s8 := make([]T8, len(s))
	s9 := make([]T9, len(s))
	s10 := make([]T10, len(s))
	s11 := make([]T11, len(s))
	s12 := make([]T12, len(s))
	s13 := make([]T13, len(s))
	s14 := make([]T14, len(s))
	s15 := make([]T15, len(s))
	s16 := make([]T16, len(s))
	for i, t := range s {
		s1[i] = t.V1
		s2[i] = t.V2
		s3[i] = t.V3
		s4[i] = t.V4
		s5[i] = t.V5
		s6[i] = t.V6
		s7[i] = t.V7
		s8[i] = t.V8
		s9[i] = t.V9
		s10[i] = t.V10
		s11[i] = t.V11
		s12[i] = t.V12
		s13[i] = t.V13
		s14[i] = t.V14
		s15[i] = t.V15
		s16[i] = t.V16
	}
	return s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16
}
