package simtest

import "goplag/source"

func supportedExt(src *source.Source) bool {
	ext := src.Ext

	if ext == ".go" {
		return true
	}

	if ext == ".c" || ext == ".cpp" {
		return true
	}

	if ext == ".java" {
		return true
	}

	if ext == ".v" {
		return true
	}

	return false
}
