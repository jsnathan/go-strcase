// Copyright (c) 2017, A. Stoewer <adrian.jsnathan@rz.ifi.lmu.de>
// All rights reserved.

package strcase

// KebabCase converts a string into kebab case.
func KebabCase(s string) string {
	return lowerDelimiterCase(s, '-')
}
