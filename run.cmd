set PATH=%USERPROFILE%\Downloads\PortableGit\bin

:end
cls &&^
cd %cd% &&^
C:\Windows\System32\timeout.exe /t 5 &&^
echo committing &&^
git config --global credential.helper "" &&^
git add * &&^
git commit -m "- upload files" &&^
git push -u origin main
goto end

pause
