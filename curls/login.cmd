@echo off

set PATH=%USERPROFILE%\Downloads\PortableGit\mingw64\bin

curl -k  -c cookie.txt --location "https://localhost:443/api/login"  --header "Content-Type: application/json"  --data "  {\"user_name\": \"joe\",\"password\": \"stemex\"}"
pause

