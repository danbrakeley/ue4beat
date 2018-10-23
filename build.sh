#!/bin/sh
rm -f ue4beat.gz
docker build -t ue4beat-builder .
echo
FILENAME_TO_EXTRACT=ue4beat.tar.gz
echo Extracting $FILENAME_TO_EXTRACT...
docker container create --name ue4beat-extract ue4beat-builder > /dev/null
docker container cp ue4beat-extract:/app/$FILENAME_TO_EXTRACT ./$FILENAME_TO_EXTRACT
docker container rm -f ue4beat-extract > /dev/null
echo Done.
