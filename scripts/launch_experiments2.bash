#!/bin/bash

#$1 es la direccion donde se encuentran el contenedor a ejecutar
#$2 es el nombre del directorio donde se encuentran los archivos de entrada
#$3 es el directorio donde se deben guardar los resultados

#cd $2
#echo $1 > salida.txt
#echo $2 >> salida.txt
#echo $3 >> salida.txt

/home/nelson/wayki/scripts/./checkInputsFiles.sh $8

pids=""
RESULT=0

for input in $(ls $2); do
	echo $input > texto.txt
	srun singularity exec $1 /home/nelson/wayki/scripts/./launch_container.sbash $2 $input $3 &
	pids="$pids $!"
done

for pid in $pids; do
    wait $pid || let "RESULT=1"
done

if [ "$RESULT" == "1" ];
    then
       exit 1
fi

cd $3
zip -r mysolv.zip *


echo "file=@$3/mysolv.zip" > $3/data.txt
echo "solver_id=$4" >> $3/data.txt
echo "os=Ubuntu 20.04.1"  >> $3/data.txt
echo "cpu_model=Intel i7 3820" >> $3/data.txt
echo "cpu_manufacturer=intel" >> $3/data.txt
echo "cpu_cores=Octa-core" >> $3/data.txt
echo "cpu_clock=3.60 GHz" >> $3/data.txt
echo "gpu_type=Dedicated graphics" >> $3/data.txt
echo "gpu_manufacturer=NVIDIA" >> $3/data.txt
echo "gpu_model=GT218" >> $3/data.txt
echo "ram=ddr3" >> $3/data.txt
echo "ram_model=HyperX Genesis 4GB x 4" >> $3/data.txt
echo "version=$5" >> $3/data.txt
echo "changelog=$6" >> $3/data.txt
echo "type=results" >> $3/data.txt
echo "email=$7" >> $3/data.txt
echo "-X POST https://wayki.net/api/solution/submissioncertificated/$8" >> $3/data.txt
echo "jwt: $9" >> $3/data.txt


if [ $? -eq ]; then
   echo "OK" >> $3/data.txt
else
   echo "FAIL" >> $3/data.txt
fi


#rm $3/*


