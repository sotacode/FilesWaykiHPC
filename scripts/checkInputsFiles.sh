#!/bin/bash



nInputs=$(ssh pgallardo@152.74.52.21 "cd /var/www/html/server/uploads/instances/;ls $1-* | wc -l")

if [ -e /home/nelson/wayki/inputChallenges/$1 ]; then
	cd /home/nelson/wayki/inputChallenges/$1
else
	mkdir /home/nelson/wayki/inputChallenges/$1
	cd /home/nelson/wayki/inputChallenges/$1
fi

nInputsCurrent=$(ls | wc -l)

if [ $nInputs -ne $nInputsCurrent ]; then
	scp pgallardo@152.74.52.21:/var/www/html/server/uploads/instances/$1-* /home/nelson/wayki/inputChallenges/$1/
else
	echo "no copie ningun archivo porque ya estan todos"
fi




