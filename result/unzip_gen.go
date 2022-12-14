// Code generated by generate_zip.pl; DO NOT EDIT.

package result

import (
	"github.com/shogo82148/hi/tuple"
)

// Unzip2 converts an optional of 2-tuple to optionals of each elements.
func Unzip2[T1, T2 any](o Result[tuple.Tuple2[T1, T2]]) (Result[T1], Result[T2]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}
}

// Unzip3 converts an optional of 3-tuple to optionals of each elements.
func Unzip3[T1, T2, T3 any](o Result[tuple.Tuple3[T1, T2, T3]]) (Result[T1], Result[T2], Result[T3]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}
}

// Unzip4 converts an optional of 4-tuple to optionals of each elements.
func Unzip4[T1, T2, T3, T4 any](o Result[tuple.Tuple4[T1, T2, T3, T4]]) (Result[T1], Result[T2], Result[T3], Result[T4]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}
}

// Unzip5 converts an optional of 5-tuple to optionals of each elements.
func Unzip5[T1, T2, T3, T4, T5 any](o Result[tuple.Tuple5[T1, T2, T3, T4, T5]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}
}

// Unzip6 converts an optional of 6-tuple to optionals of each elements.
func Unzip6[T1, T2, T3, T4, T5, T6 any](o Result[tuple.Tuple6[T1, T2, T3, T4, T5, T6]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}
}

// Unzip7 converts an optional of 7-tuple to optionals of each elements.
func Unzip7[T1, T2, T3, T4, T5, T6, T7 any](o Result[tuple.Tuple7[T1, T2, T3, T4, T5, T6, T7]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}
}

// Unzip8 converts an optional of 8-tuple to optionals of each elements.
func Unzip8[T1, T2, T3, T4, T5, T6, T7, T8 any](o Result[tuple.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}
}

// Unzip9 converts an optional of 9-tuple to optionals of each elements.
func Unzip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](o Result[tuple.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}
}

// Unzip10 converts an optional of 10-tuple to optionals of each elements.
func Unzip10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](o Result[tuple.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}
}

// Unzip11 converts an optional of 11-tuple to optionals of each elements.
func Unzip11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](o Result[tuple.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}
}

// Unzip12 converts an optional of 12-tuple to optionals of each elements.
func Unzip12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](o Result[tuple.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11], Result[T12]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}, Result[T12]{err: o.err, value: o.value.V12}
}

// Unzip13 converts an optional of 13-tuple to optionals of each elements.
func Unzip13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](o Result[tuple.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11], Result[T12], Result[T13]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}, Result[T12]{err: o.err, value: o.value.V12}, Result[T13]{err: o.err, value: o.value.V13}
}

// Unzip14 converts an optional of 14-tuple to optionals of each elements.
func Unzip14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](o Result[tuple.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11], Result[T12], Result[T13], Result[T14]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}, Result[T12]{err: o.err, value: o.value.V12}, Result[T13]{err: o.err, value: o.value.V13}, Result[T14]{err: o.err, value: o.value.V14}
}

// Unzip15 converts an optional of 15-tuple to optionals of each elements.
func Unzip15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](o Result[tuple.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11], Result[T12], Result[T13], Result[T14], Result[T15]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}, Result[T12]{err: o.err, value: o.value.V12}, Result[T13]{err: o.err, value: o.value.V13}, Result[T14]{err: o.err, value: o.value.V14}, Result[T15]{err: o.err, value: o.value.V15}
}

// Unzip16 converts an optional of 16-tuple to optionals of each elements.
func Unzip16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](o Result[tuple.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]]) (Result[T1], Result[T2], Result[T3], Result[T4], Result[T5], Result[T6], Result[T7], Result[T8], Result[T9], Result[T10], Result[T11], Result[T12], Result[T13], Result[T14], Result[T15], Result[T16]) {
	return Result[T1]{err: o.err, value: o.value.V1}, Result[T2]{err: o.err, value: o.value.V2}, Result[T3]{err: o.err, value: o.value.V3}, Result[T4]{err: o.err, value: o.value.V4}, Result[T5]{err: o.err, value: o.value.V5}, Result[T6]{err: o.err, value: o.value.V6}, Result[T7]{err: o.err, value: o.value.V7}, Result[T8]{err: o.err, value: o.value.V8}, Result[T9]{err: o.err, value: o.value.V9}, Result[T10]{err: o.err, value: o.value.V10}, Result[T11]{err: o.err, value: o.value.V11}, Result[T12]{err: o.err, value: o.value.V12}, Result[T13]{err: o.err, value: o.value.V13}, Result[T14]{err: o.err, value: o.value.V14}, Result[T15]{err: o.err, value: o.value.V15}, Result[T16]{err: o.err, value: o.value.V16}
}
