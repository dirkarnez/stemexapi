@echo off
set PATH=%USERPROFILE%\Downloads\putty

pscp -i %USERPROFILE%\Downloads\alex.ppk -r ubuntu@ec2-43-198-151-195.ap-east-1.compute.amazonaws.com:/home/ubuntu/uploads/ ./uploads/

pause