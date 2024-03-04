#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course_types () {
   echo "Parameter #1 is $1"
   curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure \
   
   # --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"
   # set "description="
# set "file="
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"
}

create_course_types "Coding Minecraft" "upcoming-schedule/codingMinecraft/Level 1-min.png"
create_course_types "Coding Roblox" "upcoming-schedule/codingRoblox/Level 1-min.png"
create_course_types Coding Python" "upcoming-schedule/Coding_Python/Level 1-min.png"
create_course_types "Coding Scratch" "upcoming-schedule/Coding_Scratch/Level 1-min.png"
create_course_types "Cyber Virtual Robotics" "upcoming-schedule/cyberVirtualRobotics/Level 1.png"


# set "Lego Robotics"
# set "upcoming-schedule/LEGO_Robotics/Level 1-min.png"

# set "description=VEX Robotics"
# set "file=upcoming-schedule/Vex Robotics/Level 1-min.png"

# set "description=AppInventor Mobile Apps"
# set "file=%USERPROFILE%/Downloads/stemex-curriculum/AppInventor/Level 1-min.png"

# set "description=A.I. ^& Machine Learning"
# set "file=upcoming-schedule/A.I. _ Machine Learning/Level 1-min.png"
# set "output=A.I._Machine_Learning"

# set "description=3D Design ^& Printing"
# set "file=upcoming-schedule/3D_Design_Printing/Level 1-min.png"
# set "output=3D_Design_Printing"
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

# set "description=Micro:bits"
# set "file=upcoming-schedule/Micro_bit/Level 1-min.png"
# set "output=Micro_bit"
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

# set "description=AR/VR"
# set "file=upcoming-schedule/AR_VR/Level 1-min.png"
# set "output=AR_VR"
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%output%.json"

# set "description=Smart City"
# set "file=upcoming-schedule/Smart City/Level 1-min.png"
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"

read -p "done"