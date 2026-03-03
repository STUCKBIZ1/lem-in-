#!/bin/sh
for i in example00 example01 example02 example03 example04 example05 example06 example07 badexample00 badexample01
do
go run . $i
done