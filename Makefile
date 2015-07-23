# Game Go in lang Go
#
# author: Drahoslav Bednář

NAME = gobang

build: $(NAME).exe

run: $(NAME).exe
	$^

$(NAME).exe: $(NAME).go 
	go build -o $@ $^

