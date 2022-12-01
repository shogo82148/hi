// Code generated by generate-tuples.pl; DO NOT EDIT.
package hi

import "fmt"

// Tuple2 is a 2-tuple.
type Tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

// New2 returns a new 2-tuple.
func New2[T1, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{v1, v2}
}

// New2 returns a new 2-tuple.
func (t Tuple2[T1, T2]) Get() (T1, T2) {
	return t.V1, t.V2
}

func (t Tuple2[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", t.V1, t.V2)
}

// Tuple3 is a 3-tuple.
type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// New3 returns a new 3-tuple.
func New3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{v1, v2, v3}
}

// New3 returns a new 3-tuple.
func (t Tuple3[T1, T2, T3]) Get() (T1, T2, T3) {
	return t.V1, t.V2, t.V3
}

func (t Tuple3[T1, T2, T3]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.V1, t.V2, t.V3)
}

// Tuple4 is a 4-tuple.
type Tuple4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

// New4 returns a new 4-tuple.
func New4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// New4 returns a new 4-tuple.
func (t Tuple4[T1, T2, T3, T4]) Get() (T1, T2, T3, T4) {
	return t.V1, t.V2, t.V3, t.V4
}

func (t Tuple4[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4)
}

// Tuple5 is a 5-tuple.
type Tuple5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// New5 returns a new 5-tuple.
func New5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// New5 returns a new 5-tuple.
func (t Tuple5[T1, T2, T3, T4, T5]) Get() (T1, T2, T3, T4, T5) {
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t Tuple5[T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5)
}

// Tuple6 is a 6-tuple.
type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

// New6 returns a new 6-tuple.
func New6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}

// New6 returns a new 6-tuple.
func (t Tuple6[T1, T2, T3, T4, T5, T6]) Get() (T1, T2, T3, T4, T5, T6) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6)
}

// Tuple7 is a 7-tuple.
type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

// New7 returns a new 7-tuple.
func New7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}
}

// New7 returns a new 7-tuple.
func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get() (T1, T2, T3, T4, T5, T6, T7) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7)
}

// Tuple8 is a 8-tuple.
type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

// New8 returns a new 8-tuple.
func New8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}
}

// New8 returns a new 8-tuple.
func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8)
}

// Tuple9 is a 9-tuple.
type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
}

// New9 returns a new 9-tuple.
func New9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

// New9 returns a new 9-tuple.
func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9)
}

// Tuple10 is a 10-tuple.
type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
}

// New10 returns a new 10-tuple.
func New10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}
}

// New10 returns a new 10-tuple.
func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10)
}

// Tuple11 is a 11-tuple.
type Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
}

// New11 returns a new 11-tuple.
func New11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11) Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11}
}

// New11 returns a new 11-tuple.
func (t Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11
}

func (t Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11)
}

// Tuple12 is a 12-tuple.
type Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
	V12 T12
}

// New12 returns a new 12-tuple.
func New12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12) Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12}
}

// New12 returns a new 12-tuple.
func (t Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12
}

func (t Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12)
}

// Tuple13 is a 13-tuple.
type Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
}

// New13 returns a new 13-tuple.
func New13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13) Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13}
}

// New13 returns a new 13-tuple.
func (t Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13
}

func (t Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13)
}

// Tuple14 is a 14-tuple.
type Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
}

// New14 returns a new 14-tuple.
func New14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14) Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14}
}

// New14 returns a new 14-tuple.
func (t Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14
}

func (t Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14)
}

// Tuple15 is a 15-tuple.
type Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
}

// New15 returns a new 15-tuple.
func New15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15) Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
}

// New15 returns a new 15-tuple.
func (t Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14, t.V15
}

func (t Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14, t.V15)
}

// Tuple16 is a 16-tuple.
type Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
}

// New16 returns a new 16-tuple.
func New16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16) Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16}
}

// New16 returns a new 16-tuple.
func (t Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Get() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14, t.V15, t.V16
}

func (t Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10, t.V11, t.V12, t.V13, t.V14, t.V15, t.V16)
}
