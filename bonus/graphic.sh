#!/usr/bin/env bash
##
## EPITECH PROJECT, 2020
## 205IQ_2019
## File description:
## graphic
##

main() {
    if [ "$#" -ne 2 ]; then
        echo "error need file 2 numbers"
        exit 84
    fi
    cd .. && make re &>/dev/null && cd bonus
    .././205IQ $1 $2 > .data.csv
    echo -e "\e[92mOk\033[0m Convert in png"
    cat drawer.gnu | gnuplot
    rm .data.csv
    echo -e "\e[92mCompleted !\033[0m"
}

main "$@"