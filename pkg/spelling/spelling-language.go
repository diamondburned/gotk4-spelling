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
	GTypeLanguage = coreglib.Type(C.spelling_language_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLanguage, F: marshalLanguage},
	})
}

// LanguageOverrides contains methods that are overridable.
type LanguageOverrides struct {
}

func defaultLanguageOverrides(v *Language) LanguageOverrides {
	return LanguageOverrides{}
}

// Language represents a spellchecking language.
type Language struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Language)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Language, *LanguageClass, LanguageOverrides](
		GTypeLanguage,
		initLanguageClass,
		wrapLanguage,
		defaultLanguageOverrides,
	)
}

func initLanguageClass(gclass unsafe.Pointer, overrides LanguageOverrides, classInitFunc func(*LanguageClass)) {
	if classInitFunc != nil {
		class := (*LanguageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLanguage(obj *coreglib.Object) *Language {
	return &Language{
		Object: obj,
	}
}

func marshalLanguage(p uintptr) (interface{}, error) {
	return wrapLanguage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Code gets the code of the language, or NULL if undefined.
//
// The function returns the following values:
//
//   - utf8 (optional): code of the language.
func (self *Language) Code() string {
	var _arg0 *C.SpellingLanguage // out
	var _cret *C.char             // in

	_arg0 = (*C.SpellingLanguage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_language_get_code(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Group gets the group of the language, or NULL if undefined.
//
// The function returns the following values:
//
//   - utf8 (optional): group of the language.
func (self *Language) Group() string {
	var _arg0 *C.SpellingLanguage // out
	var _cret *C.char             // in

	_arg0 = (*C.SpellingLanguage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_language_get_group(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of the language, or NULL if undefined.
//
// The function returns the following values:
//
//   - utf8 (optional): name of the language.
func (self *Language) Name() string {
	var _arg0 *C.SpellingLanguage // out
	var _cret *C.char             // in

	_arg0 = (*C.SpellingLanguage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.spelling_language_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LanguageClass: instance of this type is always passed by reference.
type LanguageClass struct {
	*languageClass
}

// languageClass is the struct that's finalized.
type languageClass struct {
	native *C.SpellingLanguageClass
}
