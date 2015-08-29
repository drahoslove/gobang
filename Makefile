# Game Go in lang Go
#
# author: Drahoslav Bednář

NAME = gobang

build: $(NAME).exe

run: rung

runc: $(NAME).exe
	./$^ --cli

rung: $(NAME).exe
	./$^ --gui

$(NAME).exe: *.go logic/*.go
	go build

%.syso: %.rc
	windres -o $@ $^ 
