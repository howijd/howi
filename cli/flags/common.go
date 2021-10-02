// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"strings"

	"github.com/mkungla/vars/v5"
)

// Name returns primary name for the flag usually that is long option.
func (f *Common) Name() string {
	return f.name
}

// Default sets flag default.
func (f *Common) Default(def ...interface{}) vars.Value {
	if len(def) > 0 && f.defval.Empty() {
		f.defval = vars.NewValue(def[0])
	}
	return f.defval
}

// Usage returns a usage description for that flag.
func (f *Common) Usage(usage ...string) string {
	if len(usage) > 0 {
		f.usage = strings.TrimSpace(usage[0])
	}
	if !f.defval.Empty() {
		return fmt.Sprintf("%s default: %q", f.usage, f.defval.String())
	}
	return f.usage
}

// Flag returns flag with leading - or --.
func (f *Common) Flag() string {
	if len(f.name) == 1 {
		return "-" + f.name
	}
	return "--" + f.name
}

// GetAliases Returns all aliases for the flag together with primary "name".
func (f *Common) Aliases() []string {
	return f.aliases
}

// AliasesString returns string representing flag aliases.
// e.g. used in help menu.
func (f *Common) AliasesString() string {
	if len(f.aliases) <= 1 {
		return ""
	}
	aliases := []string{}

	for _, a := range f.aliases {
		if len(a) == 1 {
			aliases = append(aliases, "-%s"+a)
			continue
		}
		aliases = append(aliases, "--%s"+a)
	}
	return strings.Join(aliases, ",")
}

// Hidden reports whether flag should be visible in help menu.
func (f *Common) IsHidden() bool {
	return f.hidden
}

// Hide flag from help menu.
func (f *Common) Hide() {
	f.hidden = true
}

// IsGlobal reports whether this flag is global.
func (f *Common) IsGlobal() bool {
	return f.global
}

// Pos returns flags position after command and case of global since app name
// min value is 1 which means first global flag or first flag after command.
func (f *Common) Pos() int {
	return f.pos
}

// Unset the value.
func (f *Common) Unset() {
	if !f.defval.Empty() {
		f.variable = vars.New(f.name, f.defval)
	} else {
		f.variable = vars.New(f.name, "")
	}
	f.isPresent = false
}

// Present reports whether flag was set in commandline.
func (f *Common) Present() bool {
	return f.isPresent
}

// Variable returns vars.Variable for this flag.
// where key is flag and Value flags value.
func (f *Common) Variable() vars.Variable {
	return f.variable
}

func (f *Common) Value() vars.Value {
	return f.variable.Value()
}

// Required sets this flag as required.
func (f *Common) Required() {
	f.required = true
}

// IsRequired returns true if this flag is required.
func (f *Common) IsRequired() bool {
	return f.required
}

// Parse the StringFlag.
func (f *Common) Parse(args *[]string) (bool, error) {
	return f.parse(args, func(v vars.Value) (err error) {
		f.variable = vars.New(f.name, v)
		return err
	})
}

// String calls Value().String().
func (f *Common) String() string {
	return f.Value().String()
}

// Parse value for the flag from given string.
// It returns true if flag has been parsed
// and error if flag has been already parsed.
func (f *Common) parse(args *[]string, read func(vars.Value) error) (bool, error) {
	if f.parsed {
		return false, fmt.Errorf("%w: %s", ErrFlagAlreadyParsed, f.name)
	}
	if f.isPresent {
		return f.isPresent, nil
	}

	if args == nil || len(*args) == 0 {
		return f.isPresent, nil
	}
	return f.parseArgs(args, read, false)
}

func (f *Common) parseAll(args *[]string, read func(vars.Value) error) (ok bool, err error) {
	f.isPresent, err = f.parseArgs(args, read, true)

	// search more
	if f.isPresent {
		for isPresent, err := f.parseArgs(args, read, true); isPresent && err != nil; {
		}
	}

	return f.isPresent, err
}

//nolint: funlen,gocognit,cyclop
func (f *Common) parseArgs(args *[]string, read func(vars.Value) error, all bool) (bool, error) {
	seek := false
	var flag string
	value := f.defval
	for i, arg := range *args {
		if len(arg) == 0 {
			return f.isPresent, nil
		}

		if arg[0] != '-' && !seek {
			f.pos++
			continue
		}

		if seek { //nolint:nestif,gocritic
			seek = false

			value = vars.NewValue((*args)[0])
			*args = append((*args)[:i-1], (*args)[i:]...)

			if err := read(value); err != nil {
				return f.isPresent, err
			}

			goto validate
		} else if strings.Contains(arg, "=") {
			kv, _ := vars.NewFromKeyVal(strings.TrimLeft(arg, "-"))
			flag = kv.Key()
			value = vars.NewValue(kv.String())
		} else {
			flag = strings.TrimLeft(arg, "-")
			if f.variable.Type() == vars.TypeBool {
				value = vars.NewValue("true")
			} else {
				// seek value from next arg
				seek = true
			}
		}

		if flag == f.name {
			f.isPresent = true
			*args = append((*args)[:i], (*args)[i+1:]...)
			if seek {
				continue
			}
			if err := read(value); err != nil {
				return f.isPresent, err
			}
			goto validate
		}

		hasAlias := false
		// if we got so far lets search alias then
		for _, alias := range f.aliases {
			if flag == alias {
				f.isPresent = true
				*args = append((*args)[:i], (*args)[i+1:]...)
				if seek {
					hasAlias = true
					break
				}
				if err := read(value); err != nil {
					return f.isPresent, err
				}
				goto validate
			}
		}
		seek = hasAlias
	}

validate:
	// was it global
	if !f.variable.Empty() && f.pos == 0 {
		f.global = true
	}
	if seek {
		return f.isPresent, fmt.Errorf("%w: did not find value for flag %q", ErrFlag, f.name)
	}
	// set default
	if !f.isPresent && !f.defval.Empty() && !all {
		return false, read(f.defval)
	}

	return f.isPresent, nil
}
