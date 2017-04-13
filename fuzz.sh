#!/bin/bash
shopt -s extglob
declare -A functionNames
functionNames=(["FuzzJourney"]="journey")

echo "Functions that will be built:"
printf "\t- %s\n" "${!functionNames[@]}"

workdirPath="/tmp/fuzz"
echo "Creating global workdir: $workdirPath"
mkdir $workdirPath

for i in "${!functionNames[@]}"
do
	path="$workdirPath/$i"
	mkdir $path
	mkdir "$path/corpus"
	binpath="$path/$i.zip"
	
	corpuspath="./testdata/${functionNames[$i]}"
	echo "Copying corpus for $i from $corpuspath"
	cp $corpuspath/known/* "$path/corpus/"
	cp $corpuspath/corpus/* "$path/corpus/"
	
	echo "Building $i"
	go-fuzz-build -func $i -o $binpath github.com/aabizri/gonavitia/types
	
	echo "Running $i ($binpath)"
	go-fuzz -bin $binpath -workdir=$path
	
	echo "Copying back corpus from $corpuspath"
	destination="./testdata/${functionNames[$i]}/corpus"
	rsync --exclude="*.json" $path/corpus/* $destination/
	
	echo "Copying crashers"
	destination="./testdata/${functionNames[$i]}/crasher"
	commit=`git rev-parse HEAD`
	for j in $path/crashers/*
	do
		filename=$commit-${j##*/}
		cp $j $destination/$filename
	done
done
