package main

//go:generate python -m peachpy.x86_64 detect_unique_chars.py -S -o detect_unique_chars_amd64.s -mabi=goasm
func DetectUniqueChars(str string) int
