set terminal pngcairo
set output '205IQ.png'
set grid
set title '205IQ by Eydou and Xelium'
plot '.data.csv' w lp title 'IQ' lt rgb "blue"