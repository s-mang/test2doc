#!/bin/bash

set -e

# find all *.apib files (after tests have run, generated files)
files=`find ./**/ -type f -name "*.apib"`

# copy template file to new apiary.apib file
cp apib.tmpl apiary.apib

# copy contents of each generated apib file into apiary.apib
# and delete the apib file
for f in ${files[@]}; do
	cat $f >> apiary.apib
	rm $f
done