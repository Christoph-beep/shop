IDS=$(ps ax | grep "/var/folders" | grep "b001/exe/main" | grep -v "grep" | awk '{print $1}')
    if [ ! -z "$IDS" ]
    then
        kill $IDS;
    fi
sleep 0.3
go run main.go &
