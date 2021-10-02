// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

//nolint: nlreturn
package flags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mkungla/vars/v5"
	"pkg.howijd.network/howi/v8/namespace"
)

var (
	ErrFlag              = errors.New("flag error")
	ErrMissingOption     = errors.New("need to specify atleast one option")
	ErrInvalidValue      = errors.New("invalid value for flag")
	ErrFlagAlreadyParsed = errors.New("flag is already parsed")
)

type (
	// Flag is howi/cli/flags.Flags interface.
	Flag interface {
		// Parse value for the flag from given string. It returns true if flag
		// was found in provided args string and false if not.
		// error is returned when flag was set but had invalid value.
		Parse(*[]string) (bool, error)
		// Get primary name for the flag. Usually that is long option
		Name() string
		// Usage returns a usage description for that flag
		Usage(...string) string
		// Flag returns flag with leading - or --
		// useful for help menus
		Flag() string
		// Hide flag from help menu.
		Hide()
		// Return flag aliases
		Aliases() []string
		// AliasesString returns string representing flag aliases.
		// e.g. used in help menu
		AliasesString() string
		// IsHidden reports whether to show that flag in help menu or not.
		IsHidden() bool
		// IsGlobal reports whether this flag was global and was set before any command or arg
		IsGlobal() bool
		// Pos returns flags position after command. Case of global since app name
		// min value 1 which means first global flag or first flag after command
		Pos() int
		// Unset unsets the value for the flag if it was parsed, handy for cases where
		// one flag cancels another like --debug cancels --verbose
		Unset()
		// Present reports whether flag was set in commandline
		Present() bool
		// Variable returns vars.Variable for this flag.
		// where key is flag and Value flags value.
		Variable() vars.Variable
		// Value returns vars.Value for given flag
		Value() vars.Value
		// Required sets this flag as required
		Required()
		// IsRequired returns true if this flag is required
		IsRequired() bool
		// Set flag default value
		Default(def ...interface{}) vars.Value
		// String calls Value().String()
		String() string
	}

	// Common is default string flag. Common flag ccan be used to
	// as base for custom flags by owerriding .Parse func.
	Common struct {
		// name of this flag
		name string
		// aliases for this flag
		aliases []string
		// hide from help menu
		hidden bool
		// global is set to true if value was parsed before any command or arg occurred
		global bool
		// position in os args how many commands where before that flag
		pos int
		// usage string
		usage string
		// isPresent enables to mock removal and .Unset the flag it reports whether flag was "present"
		isPresent bool
		// value for this flag
		variable vars.Variable
		// is this flag required
		required bool
		// default value
		defval vars.Value
		// flag already parsed
		parsed bool
	}

	// OptionFlag is string flag type which can have value of one of the options.
	OptionFlag struct {
		Common
		opts map[string]bool
	}

	// NumFlag is numeric flag type with default value 0.
	NumFlag struct {
		Common
	}

	// BoolFlag is boolean flag type with default value "false".
	BoolFlag struct {
		Common
	}
)

// New returns new common string flag. Argument "a" can be any nr of aliases.
func New(name string, aliases ...string) (Flag, error) {
	f, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f.variable = vars.New(name, "")
	return f, err
}

// NewNumFlag returns new numeric flag. Argument "a" can be any nr of aliases.
func NewNumFlag(name string, aliases ...string) (*NumFlag, error) {
	c, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f := &NumFlag{*c}
	f.variable, _ = vars.NewTyped(name, "0", vars.TypeFloat64)
	return f, nil
}

// NewBoolFlag returns new bool flag. Argument "a" can be any nr of aliases.
func NewBoolFlag(name string, aliases ...string) (*BoolFlag, error) {
	c, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f := &BoolFlag{*c}
	f.variable, _ = vars.NewTyped(name, "false", vars.TypeBool)
	return f, nil
}

// NewOptionFlag returns new string flag. Argument "opts" is string slice
// of options this flag accepts.
func NewOptionFlag(name string, opts []string, aliases ...string) (*OptionFlag, error) {
	c, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f := &OptionFlag{Common: *c}
	f.opts = make(map[string]bool, len(opts))
	for _, o := range opts {
		f.opts[o] = false
	}

	f.variable = vars.New(name, "")
	return f, nil
}

func newCommon(name string, aliases ...string) (*Common, error) {
	if !namespace.ValidSlug(name) {
		return nil, fmt.Errorf("%w: flag name %q is not valid", ErrFlag, name)
	}

	f := &Common{}
	f.name = strings.TrimLeft(name, "-")
	for _, alias := range aliases {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	return f, nil
}
