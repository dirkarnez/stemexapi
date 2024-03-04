#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

/C/Windows/System32/curl.exe -k -c cookie.txt --location "https://localhost:443/api/login" \
    --header "Content-Type: application/json" \
    --data "{\"user_name\": \"joe\",\"password\": \"stemex\"}"

read -p "done"
