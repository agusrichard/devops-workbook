#!/bin/bash

GIRLS=(veni sekar saskia arifa)

for girl in ${GIRLS[@]};
	do echo -n $girl | wc -c;
done
