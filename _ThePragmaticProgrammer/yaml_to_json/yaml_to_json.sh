#!/bin/bash

# Ref: https://stackoverflow.com/questions/965053/extract-filename-and-extension-in-bash

DEST_DIR=$(pwd)/json_files/

find ./yaml_files -type f \( -iname '*.yaml' -o -name '*.yml' \) | while read FilePath; do
	echo "process: $FilePath"

	FilenameWithExt=$(basename -- "$FilePath")
	Filename="${FilenameWithExt%.*}"

	yq eval $FilePath -j > $DEST_DIR$Filename.json
done
