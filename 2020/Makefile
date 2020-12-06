


$(shell touch .env)
include .env
export

.DEFAULT:
		mkdir $@
		touch $@/$@.go
		echo "cGFja2FnZSBtYWluCgppbXBvcnQgImdpdGh1Yi5jb20vbWF4bWNkL2FvYy8yMDIwL2xpYiIKCmZ1bmMgbWFpbigpIHsKCWxpYi5JbnB1dExpbmVzKDIpCn0K" | base64 -d > $@/$@.go
		curl -H "Cookie: session=$${SESSION}" https://adventofcode.com/2020/day/$@/input > $@/$@.txt
