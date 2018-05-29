package domain

type (
	// Flag represents a command line flag.
	Flag struct {
		Description *string `yaml:"description"`
		Name        *string `yaml:"name"`
		Short       *string `yaml:"short"`
		Bind        *bool   `yaml:"bind"`
	}
)

// IsBind returns whether the flag is bound.
// func (f Flag) IsBind(def *bool) bool {
// 	b := f.Bind
// 	if b == nil {
// 		if f.Name != nil || f.Description != nil || f.Short != nil {
// 			return true
// 		}
// 		if def == nil {
// 			return false
// 		}
// 		return *def
// 	}
// 	return *b
// }
