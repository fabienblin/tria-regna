NAME=tria-regna

all:
	go build -o $(NAME) .

run:
	go run .

clean:
	@rm $(NAME)

re: clean all