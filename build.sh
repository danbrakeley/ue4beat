#!/bin/sh
rm -f ue4beat.gz
docker build -t ue4beat-builder .
echo
echo Extracting ue4beat.gz...
docker container create --name ue4beat-extract ue4beat-builder > /dev/null
docker container cp ue4beat-extract:/app/ue4beat.gz ./ue4beat.gz
docker container rm -f ue4beat-extract > /dev/null
echo Done.
