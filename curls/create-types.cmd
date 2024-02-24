@echo off

set "description=Coding Minecraft"
set "file=upcoming-schedule/codingMinecraft/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=Coding Roblox"
set "file=upcoming-schedule/codingRoblox/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=Coding Python"
set "file=upcoming-schedule/Coding_Python/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"


set "description=Coding Scratch"
set "file=upcoming-schedule/Coding_Scratch/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=Cyber Virtual Robotics"
set "file=upcoming-schedule/cyberVirtualRobotics/Level 1.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=Lego Robotics"
set "file=upcoming-schedule/LEGO_Robotics/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=VEX Robotics"
set "file=upcoming-schedule/Vex Robotics/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=AppInventor Mobile Apps"
set "file=upcoming-schedule/AppInventor Mobile Apps/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

set "description=A.I. ^& Machine Learning"
set "file=upcoming-schedule/A.I. _ Machine Learning/Level 1-min.png"
set "output=A.I._Machine_Learning"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

set "description=3D Design ^& Printing"
set "file=upcoming-schedule/3D_Design_Printing/Level 1-min.png"
set "output=3D_Design_Printing"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

set "description=Micro:bits"
set "file=upcoming-schedule/Micro_bit/Level 1-min.png"
set "output=Micro_bit"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

set "description=AR/VR"
set "file=upcoming-schedule/AR_VR/Level 1-min.png"
set "output=AR_VR"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

set "description=Smart City"
set "file=upcoming-schedule/Smart City/Level 1-min.png"
curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
--form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

