ifeq ($(OS),Windows_NT)
	SHELL := pwsh.exe
else
	SHELL := pwsh
endif

.SHELLFLAGS := -NoProfile -Command


.PHONY: all
all: run

precommit_install:
	pre-commit install
precommit_checkall: precommit_install
	pre-commit run --all-files
fyne_test: # Use this to test your fyne install
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest

get_deps: # install os packages first: https://developer.fyne.io/started/
	go get fyne.io/fyne/v2@latest
	go install fyne.io/fyne/v2/cmd/fyne@latest
run: get_deps
	go run main.go