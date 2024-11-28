// Code generated by girgen. DO NOT EDIT.

package spelling

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libspelling.h>
import "C"

// GType values.
var (
	GTypeDictionary = coreglib.Type(C.spelling_dictionary_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDictionary, F: marshalDictionary},
	})
}

// Dictionary: abstract base class for spellchecking dictionaries.
type Dictionary struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Dictionary)(nil)
)

// Dictionarier describes types inherited from class Dictionary.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Dictionarier interface {
	coreglib.Objector
	baseDictionary() *Dictionary
}

var _ Dictionarier = (*Dictionary)(nil)

func wrapDictionary(obj *coreglib.Object) *Dictionary {
	return &Dictionary{
		Object: obj,
	}
}

func marshalDictionary(p uintptr) (interface{}, error) {
	return wrapDictionary(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *Dictionary) baseDictionary() *Dictionary {
	return self
}

// BaseDictionary returns the underlying base object.
func BaseDictionary(obj Dictionarier) *Dictionary {
	return obj.baseDictionary()
}

// AddWord adds word to the dictionary.
//
// The function takes the following parameters:
//
//   - word to be added.
func (self *Dictionary) AddWord(word string) {
	var _arg0 *C.SpellingDictionary // out
	var _arg1 *C.char               // out

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(word)))
	defer C.free(unsafe.Pointer(_arg1))

	C.spelling_dictionary_add_word(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)
}

// ContainsWord checks if the dictionary contains word.
//
// The function takes the following parameters:
//
//   - word to be checked.
//
// The function returns the following values:
//
//   - ok: TRUE if the dictionary contains the word.
func (self *Dictionary) ContainsWord(word string) bool {
	var _arg0 *C.SpellingDictionary // out
	var _arg1 *C.char               // out
	var _arg2 C.gssize
	var _cret C.gboolean // in

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg2 = (C.gssize)(len(word))
	_arg1 = (*C.char)(C.calloc(C.size_t((len(word) + 1)), C.size_t(C.sizeof_char)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), len(word)), word)
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.spelling_dictionary_contains_word(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Code gets the language code of the dictionary, or NULL if undefined.
//
// The function returns the following values:
//
//   - utf8 (optional): language code of the dictionary.
func (self *Dictionary) Code() string {
	var _arg0 *C.SpellingDictionary // out
	var _cret *C.char               // in

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_dictionary_get_code(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ExtraWordChars gets the extra word characters of the dictionary.
//
// The function returns the following values:
//
//   - utf8: extra word characters.
func (self *Dictionary) ExtraWordChars() string {
	var _arg0 *C.SpellingDictionary // out
	var _cret *C.char               // in

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_dictionary_get_extra_word_chars(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IgnoreWord requests the dictionary to ignore word.
//
// The function takes the following parameters:
//
//   - word to be ignored.
func (self *Dictionary) IgnoreWord(word string) {
	var _arg0 *C.SpellingDictionary // out
	var _arg1 *C.char               // out

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(word)))
	defer C.free(unsafe.Pointer(_arg1))

	C.spelling_dictionary_ignore_word(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)
}

// ListCorrections retrieves a list of possible corrections for word.
//
// The function takes the following parameters:
//
//   - word to be checked.
//
// The function returns the following values:
//
//   - utf8s (optional): A list of possible corrections, or NULL.
func (self *Dictionary) ListCorrections(word string) []string {
	var _arg0 *C.SpellingDictionary // out
	var _arg1 *C.char               // out
	var _arg2 C.gssize
	var _cret **C.char // in

	_arg0 = (*C.SpellingDictionary)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg2 = (C.gssize)(len(word))
	_arg1 = (*C.char)(C.calloc(C.size_t((len(word) + 1)), C.size_t(C.sizeof_char)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), len(word)), word)
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.spelling_dictionary_list_corrections(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.char
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}
