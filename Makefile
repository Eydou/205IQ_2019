##
## EPITECH PROJECT, 2019
## minishell1
## File description:
## Makefile
##

SRC	=	src/main.go

OBJ	=	$(SRC:.c=.o)

TEST	=	$(shell find ./ -name '*.go' ! -name 'main.cpp')	\

CC	=	go build

NAME	=	205IQ

CPP_FLAGS	=	-W -Wall -Werror -Wextra -std=c++11

all:	$(NAME)

tests_run:	$(TEST)
	go test -v src/functions/math.go src/functions/math_test.go

$(NAME):	$(OBJ)
	@$(CC) -o $(NAME) $(OBJ)
	@echo -e " -> \e[96mCompilation oklmzer\033[0m"
	@echo -e "\e[96mUSAGE\n   ./205IQ u s [IQ1] [IQ2]\033[0m"
clean:
	@rm -f $(shell find $(SOURCEDIR) -name '*.o')
	@rm -f $(shell find $(SOURCEDIR) -name '*~')
	@rm -f $(shell find $(SOURCEDIR) -name '*#')
	@rm -f $(shell find $(SOURCEDIR) -name '*vg*')
	@rm -f $(shell find $(SOURCEDIR) -name '*.gc*')
	@echo -e "\e[96mIs Clean\033[0m"


fclean: clean
	@rm -f $(NAME)

re:	fclean all

.PHONY: all tests_run clean fclean re
