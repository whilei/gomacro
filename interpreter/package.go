// this file was generated by gomacro command: import "github.com/cosmos72/gomacro/interpreter"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package interpreter

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// inception: allow interpreted code to import "github.com/cosmos72/gomacro/interpreter"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/interpreter"] = imports.Package{
		Binds: map[string]r.Value{
			"DefaultImporter": r.ValueOf(DefaultImporter),
			"New":             r.ValueOf(New),
			"NewEnv":          r.ValueOf(NewEnv),
			"NewInterpreter":  r.ValueOf(NewInterpreter),
			"Nil":             r.ValueOf(&Nil).Elem(),
			"None":            r.ValueOf(&None).Elem(),
			"One":             r.ValueOf(&One).Elem(),
			"OptDebugMacroExpand":        r.ValueOf(OptDebugMacroExpand),
			"OptDebugQuasiquote":         r.ValueOf(OptDebugQuasiquote),
			"OptShowAfterEval":           r.ValueOf(OptShowAfterEval),
			"OptShowAfterMacroExpansion": r.ValueOf(OptShowAfterMacroExpansion),
			"OptShowAfterParse":          r.ValueOf(OptShowAfterParse),
			"OptShowEvalDuration":        r.ValueOf(OptShowEvalDuration),
			"OptShowPrompt":              r.ValueOf(OptShowPrompt),
			"OptTrapPanic":               r.ValueOf(OptTrapPanic),
			"Read":                       r.ValueOf(Read),
			"ReadMultiline":              r.ValueOf(ReadMultiline),
			"Unknown":                    r.ValueOf(&Unknown).Elem(),
		},
		Types: map[string]r.Type{
			"Builtin":     r.TypeOf((*Builtin)(nil)).Elem(),
			"Cmd":         r.TypeOf((*Cmd)(nil)).Elem(),
			"Env":         r.TypeOf((*Env)(nil)).Elem(),
			"Function":    r.TypeOf((*Function)(nil)).Elem(),
			"Importer":    r.TypeOf((*Importer)(nil)).Elem(),
			"Inspector":   r.TypeOf((*Inspector)(nil)).Elem(),
			"Interpreter": r.TypeOf((*Interpreter)(nil)).Elem(),
			"Macro":       r.TypeOf((*Macro)(nil)).Elem(),
			"Options":     r.TypeOf((*Options)(nil)).Elem(),
			"PackageRef":  r.TypeOf((*PackageRef)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{},
	}
}
