


$(shell touch .env)
include .env
export

.DEFAULT:
	mkdir $@
	touch $@/main.go
	cp template.go $@/main.go
	curl -H "Cookie: session=$${SESSION}" https://adventofcode.com/2022/day/$@/input > $@/input.txt
