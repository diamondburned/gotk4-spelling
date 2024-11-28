// Code generated by girgen. DO NOT EDIT.

package spelling

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libspelling.h>
import "C"

// GType values.
var (
	GTypeChecker = coreglib.Type(C.spelling_checker_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeChecker, F: marshalChecker},
	})
}

// CheckerOverrides contains methods that are overridable.
type CheckerOverrides struct {
}

func defaultCheckerOverrides(v *Checker) CheckerOverrides {
	return CheckerOverrides{}
}

// Checker: SpellingChecker is the core class of libspelling. It provides
// high-level APIs for spellchecking.
type Checker struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Checker)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Checker, *CheckerClass, CheckerOverrides](
		GTypeChecker,
		initCheckerClass,
		wrapChecker,
		defaultCheckerOverrides,
	)
}

func initCheckerClass(gclass unsafe.Pointer, overrides CheckerOverrides, classInitFunc func(*CheckerClass)) {
	if classInitFunc != nil {
		class := (*CheckerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapChecker(obj *coreglib.Object) *Checker {
	return &Checker{
		Object: obj,
	}
}

func marshalChecker(p uintptr) (interface{}, error) {
	return wrapChecker(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewChecker: create a new SpellingChecker.
//
// The function takes the following parameters:
//
//   - provider
//   - language
//
// The function returns the following values:
//
//   - checker: newly created SpellingChecker.
func NewChecker(provider Providerer, language string) *Checker {
	var _arg1 *C.SpellingProvider // out
	var _arg2 *C.char             // out
	var _cret *C.SpellingChecker  // in

	_arg1 = (*C.SpellingProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(language)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.spelling_checker_new(_arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(language)

	var _checker *Checker // out

	_checker = wrapChecker(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _checker
}

// AddWord adds word to the active dictionary.
//
// The function takes the following parameters:
//
//   - word to be added.
func (self *Checker) AddWord(word string) {
	var _arg0 *C.SpellingChecker // out
	var _arg1 *C.char            // out

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(word)))
	defer C.free(unsafe.Pointer(_arg1))

	C.spelling_checker_add_word(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)
}

// CheckWord checks if the active dictionary contains word.
//
// The function takes the following parameters:
//
//   - word to be checked.
//
// The function returns the following values:
//
//   - ok: TRUE if the dictionary contains the word.
func (self *Checker) CheckWord(word string) bool {
	var _arg0 *C.SpellingChecker // out
	var _arg1 *C.char            // out
	var _arg2 C.gssize
	var _cret C.gboolean // in

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg2 = (C.gssize)(len(word))
	_arg1 = (*C.char)(C.calloc(C.size_t((len(word) + 1)), C.size_t(C.sizeof_char)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), len(word)), word)
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.spelling_checker_check_word(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(word)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtraWordChars gets the extra word characters of the active dictionary.
//
// The function returns the following values:
//
//   - utf8: extra word characters.
func (self *Checker) ExtraWordChars() string {
	var _arg0 *C.SpellingChecker // out
	var _cret *C.char            // in

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_checker_get_extra_word_chars(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Language gets the language being used by the spell checker.
//
// The function returns the following values:
//
//   - utf8 (optional): string describing the current language.
func (self *Checker) Language() string {
	var _arg0 *C.SpellingChecker // out
	var _cret *C.char            // in

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_checker_get_language(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Provider gets the spell provider used by the spell checker.
//
// Currently, only Enchant-2 is supported.
//
// The function returns the following values:
//
//   - provider: SpellingProvider.
func (self *Checker) Provider() Providerer {
	var _arg0 *C.SpellingChecker  // out
	var _cret *C.SpellingProvider // in

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_checker_get_provider(_arg0)
	runtime.KeepAlive(self)

	var _provider Providerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type spelling.Providerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Providerer)
			return ok
		})
		rv, ok := casted.(Providerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching spelling.Providerer")
		}
		_provider = rv
	}

	return _provider
}

// IgnoreWord requests the active dictionary to ignore word.
//
// The function takes the following parameters:
//
//   - word to be ignored.
func (self *Checker) IgnoreWord(word string) {
	var _arg0 *C.SpellingChecker // out
	var _arg1 *C.char            // out

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(word)))
	defer C.free(unsafe.Pointer(_arg1))

	C.spelling_checker_ignore_word(_arg0, _arg1)
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
func (self *Checker) ListCorrections(word string) []string {
	var _arg0 *C.SpellingChecker // out
	var _arg1 *C.char            // out
	var _cret **C.char           // in

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(word)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.spelling_checker_list_corrections(_arg0, _arg1)
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

// SetLanguage sets the language code to use when communicating with the
// provider, such as en_US.
//
// The function takes the following parameters:
//
//   - language to use.
func (self *Checker) SetLanguage(language string) {
	var _arg0 *C.SpellingChecker // out
	var _arg1 *C.char            // out

	_arg0 = (*C.SpellingChecker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(language)))
	defer C.free(unsafe.Pointer(_arg1))

	C.spelling_checker_set_language(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(language)
}

// CheckerGetDefault gets a default SpellingChecker using the default provider
// and language.
//
// The function returns the following values:
//
//   - checker: SpellingChecker.
func CheckerGetDefault() *Checker {
	var _cret *C.SpellingChecker // in

	_cret = C.spelling_checker_get_default()

	var _checker *Checker // out

	_checker = wrapChecker(coreglib.Take(unsafe.Pointer(_cret)))

	return _checker
}

// CheckerClass: instance of this type is always passed by reference.
type CheckerClass struct {
	*checkerClass
}

// checkerClass is the struct that's finalized.
type checkerClass struct {
	native *C.SpellingCheckerClass
}
