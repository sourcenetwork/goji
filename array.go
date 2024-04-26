//go:build js

package goji

import "syscall/js"

func init() {
	Array = arrayJS(js.Global().Get("Array"))
}

type arrayJS js.Value

// Array is a wrapper for the Array global object.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
var Array arrayJS

// ArrayValue is an instance of an array.
type ArrayValue js.Value

// New wraps the array constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/Array
func (a arrayJS) New(lengthOrElements ...any) ArrayValue {
	res := js.Value(a).New(lengthOrElements...)
	return ArrayValue(res)
}

// From wraps the array from static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/from
func (a arrayJS) From(arrayLike js.Value, mapFn js.Value, thisArg js.Value) ArrayValue {
	res := js.Value(a).Call("from", arrayLike, mapFn, thisArg)
	return ArrayValue(res)
}

// FromAsync wraps the array fromAsync static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fromAsync
func (a arrayJS) FromAsync(arrayLike js.Value, mapFn js.Value, thisArg js.Value) PromiseValue {
	res := js.Value(a).Call("fromAsync", arrayLike, mapFn, thisArg)
	return PromiseValue(res)
}

// IsArray wraps the array isArray static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/isArray
func (a arrayJS) IsArray(value any) bool {
	return js.Value(a).Call("isArray", value).Bool()
}

// Of wraps the array of static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/of
func (a arrayJS) Of(elements ...any) ArrayValue {
	res := js.Value(a).Call("of", elements...)
	return ArrayValue(res)
}

// Length wraps the array length property.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/length
func (a ArrayValue) Length() int {
	return js.Value(a).Get("length").Int()
}

// At wraps the array at instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/at
func (a ArrayValue) At(index int) js.Value {
	return js.Value(a).Call("at", index)
}

// Concat wraps the array concat instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/concat
func (a ArrayValue) Concat(elements ...any) ArrayValue {
	res := js.Value(a).Call("concat", elements...)
	return ArrayValue(res)
}

// CopyWithin wraps the array copyWithin instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/copyWithin
func (a ArrayValue) CopyWithin(target, start, end int) ArrayValue {
	res := js.Value(a).Call("copyWithin", target, start, end)
	return ArrayValue(res)
}

// Entries wraps the array entries instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/entries
func (a ArrayValue) Entries() js.Value {
	return js.Value(a).Call("entries")
}

// Every wraps the array every instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/every
func (a ArrayValue) Every(callbackFn js.Value, thisArg js.Value) bool {
	return js.Value(a).Call("every", callbackFn, thisArg).Bool()
}

// Fill wraps the array fill instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fill
func (a ArrayValue) Fill(value any, start, end int) ArrayValue {
	res := js.Value(a).Call("fill", value, start, end)
	return ArrayValue(res)
}

// Filter wraps the array filter instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/filter
func (a ArrayValue) Filter(callbackFn js.Value, thisArg js.Value) ArrayValue {
	res := js.Value(a).Call("filter", callbackFn, thisArg)
	return ArrayValue(res)
}

// Find wraps the array find instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/find
func (a ArrayValue) Find(callbackFn js.Value, thisArg js.Value) js.Value {
	return js.Value(a).Call("find", callbackFn, thisArg)
}

// FindIndex wraps the array findIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findIndex
func (a ArrayValue) FindIndex(callbackFn js.Value, thisArg js.Value) int {
	return js.Value(a).Call("findIndex", callbackFn, thisArg).Int()
}

// FindLast wraps the array findLast instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findLast
func (a ArrayValue) FindLast(callbackFn js.Value, thisArg js.Value) js.Value {
	return js.Value(a).Call("findLast", callbackFn, thisArg)
}

// FindLastIndex wraps the array findLastIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findLastIndex
func (a ArrayValue) FindLastIndex(callbackFn js.Value, thisArg js.Value) int {
	return js.Value(a).Call("findLastIndex", callbackFn, thisArg).Int()
}

// Flat wraps the array flat instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/flat
func (a ArrayValue) Flat(depth int) ArrayValue {
	res := js.Value(a).Call("flat", depth)
	return ArrayValue(res)
}

// FlatMap wraps the array flatMap instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/flatMap
func (a ArrayValue) FlatMap(callbackFn js.Value, thisArg js.Value) ArrayValue {
	res := js.Value(a).Call("flatMap", callbackFn, thisArg)
	return ArrayValue(res)
}

// ForEach wraps the array forEach instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach
func (a ArrayValue) ForEach(callbackFn js.Value, thisArg js.Value) {
	js.Value(a).Call("forEach", callbackFn, thisArg)
}

// Includes wraps the array includes instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/includes
func (a ArrayValue) Includes(searchElement any, fromIndex int) bool {
	return js.Value(a).Call("includes", searchElement, fromIndex).Bool()
}

// IndexOf wraps the array indexOf instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/indexOf
func (a ArrayValue) IndexOf(searchElement any, fromIndex int) int {
	return js.Value(a).Call("indexOf", searchElement, fromIndex).Int()
}

// Join wraps the array join instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
func (a ArrayValue) Join(separator string) string {
	return js.Value(a).Call("join", separator).String()
}

// Keys wraps the array keys instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/keys
func (a ArrayValue) Keys() js.Value {
	return js.Value(a).Call("keys")
}

// LastIndexOf wraps the array lastIndexOf instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/lastIndexOf
func (a ArrayValue) LastIndexOf(searchElement any, fromIndex int) int {
	return js.Value(a).Call("lastIndexOf", searchElement, fromIndex).Int()
}

// Map wraps the array map instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map
func (a ArrayValue) Map(callbackFn js.Value, thisArg js.Value) ArrayValue {
	res := js.Value(a).Call("map", callbackFn, thisArg)
	return ArrayValue(res)
}

// Pop wraps the array pop instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/pop
func (a ArrayValue) Pop() js.Value {
	return js.Value(a).Call("pop")
}

// Push wraps the array push instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/push
func (a ArrayValue) Push(elements ...any) int {
	return js.Value(a).Call("push", elements...).Int()
}

// Reduce wraps the array reduce instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce
func (a ArrayValue) Reduce(callbackFn js.Value, initialValue js.Value) js.Value {
	return js.Value(a).Call("reduce", callbackFn, initialValue)
}

// ReduceRight wraps the array reduceRight instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduceRight
func (a ArrayValue) ReduceRight(callbackFn js.Value, initialValue js.Value) js.Value {
	return js.Value(a).Call("reduceRight", callbackFn, initialValue)
}

// Reverse wraps the array reverse instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reverse
func (a ArrayValue) Reverse() ArrayValue {
	res := js.Value(a).Call("reverse")
	return ArrayValue(res)
}

// Shift wraps the array shift intance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/shift
func (a ArrayValue) Shift() js.Value {
	return js.Value(a).Call("shift")
}

// Slice wraps the array slice instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/slice
func (a ArrayValue) Slice(start, end int) ArrayValue {
	res := js.Value(a).Call("slice", start, end)
	return ArrayValue(res)
}

// Some wraps the array some instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/some
func (a ArrayValue) Some(callbackFn js.Value, thisArg js.Value) bool {
	return js.Value(a).Call("some", callbackFn, thisArg).Bool()
}

// Sort wraps the array sort instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/sort
func (a ArrayValue) Sort(compareFn js.Value) ArrayValue {
	res := js.Value(a).Call("sort", compareFn)
	return ArrayValue(res)
}

// Splice wraps the array splice instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/splice
func (a ArrayValue) Splice(start, deleteCount int, items ...any) ArrayValue {
	args := []any{start, deleteCount}
	args = append(args, items...)

	res := js.Value(a).Call("splice", args...)
	return ArrayValue(res)
}

// ToLocaleString wraps the array toLocaleString instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/toLocaleString
func (a ArrayValue) ToLocaleString(locales js.Value, options js.Value) string {
	return js.Value(a).Call("toLocaleString", locales, options).String()
}

// ToReversed wraps the array toReversed instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/toReversed
func (a ArrayValue) ToReversed() ArrayValue {
	res := js.Value(a).Call("toReversed")
	return ArrayValue(res)
}

// ToSorted wraps the array toSorted instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/toSorted
func (a ArrayValue) ToSorted(compareFn js.Value) ArrayValue {
	res := js.Value(a).Call("toSorted", compareFn)
	return ArrayValue(res)
}

// ToSpliced wraps the array toSpliced instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/toSpliced
func (a ArrayValue) ToSpliced(start, deleteCount int, items ...any) ArrayValue {
	args := []any{start, deleteCount}
	args = append(args, items...)

	res := js.Value(a).Call("toSpliced", args...)
	return ArrayValue(res)
}

// ToString wraps the array toString instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/toString
func (a ArrayValue) ToString() string {
	return js.Value(a).Call("toString").String()
}

// Unshift wraps the array unshift instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/unshift
func (a ArrayValue) Unshift(elements ...any) int {
	return js.Value(a).Call("unshift", elements...).Int()
}

// Values wraps the array values instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/values
func (a ArrayValue) Values() js.Value {
	return js.Value(a).Call("values")
}

// With wraps the array with instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/with
func (a ArrayValue) With(index int, value any) ArrayValue {
	res := js.Value(a).Call("with", index, value)
	return ArrayValue(res)
}
