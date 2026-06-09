#!/bin/bash

DATA_FILE="./test_data.txt"
USER="postgres"
DB_ADDR="localhost"
DB_PORT="5432"
DB="raspyxdb"
TEMP_FILE="./tmp/file"

apt -y install postgresql-client-*

mkdir -p ./tmp 2> /dev/null

while read -r line; do
    if [[ ${line::1} != "\"" ]]; then
        if ! [ -z "$line" ]; then
	        TABLE=$line
	fi
    else
    	echo $line | tr \" \' > $TEMP_FILE
    	sed -i "s/' '/','/g" $TEMP_FILE
    	echo "INSERT INTO $TABLE VALUES ($(cat $TEMP_FILE));" >> ./tables.sql
    fi
done < $DATA_FILE

psql -U $USER -d $DB -h $DB_ADDR -p $DB_PORT -f ./tables.sql
rm ./tables.sql
rm -r ./tmp
