.PHONY: all

GOPATH := ${PWD}/.workspace/:${GOPATH}
export GOPATH

all:
	@mkdir -p .workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go test ./... -cover
