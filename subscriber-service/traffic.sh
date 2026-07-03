#!/bin/bash

URL="http://172.16.27.65:31727"

COUNT=1

while true
do
    ACTION=$((RANDOM % 4))

    case $ACTION in

    0)
        echo "CREATE"

        curl -s -X POST $URL/subscribers \
        -H "Content-Type: application/json" \
        -d "{
        \"name\":\"User$COUNT\",
        \"imsi\":\"404100$COUNT\",
        \"msisdn\":\"98765$COUNT\"
        }" > /dev/null

        COUNT=$((COUNT+1))
        ;;

    1)
        echo "GET"

        curl -s $URL/subscribers > /dev/null
        ;;

    2)
        echo "UPDATE"

        ID=$(curl -s $URL/subscribers | \
        grep -o '"id":"[^"]*"' | \
        cut -d'"' -f4 | \
        shuf -n1)

        if [ ! -z "$ID" ]; then
            curl -s -X PUT $URL/subscribers/$ID \
            -H "Content-Type: application/json" \
            -d "{
            \"name\":\"Updated-$RANDOM\",
            \"imsi\":\"404999$RANDOM\",
            \"msisdn\":\"99999$RANDOM\",
            \"plan\":\"Premium\",
            \"status\":\"ACTIVE\"
            }" > /dev/null
        fi
        ;;

    3)
        echo "DELETE"

        ID=$(curl -s $URL/subscribers | \
        grep -o '"id":"[^"]*"' | \
        cut -d'"' -f4 | \
        shuf -n1)

        if [ ! -z "$ID" ]; then
            curl -s -X DELETE $URL/subscribers/$ID > /dev/null
        fi
        ;;

    esac

    sleep $(awk -v seed=$RANDOM 'BEGIN{srand(seed); printf "%.2f\n", rand()*2}')

done
