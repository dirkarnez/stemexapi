#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course_types () {
    $(echo $PATH | sed 's/\//-/g')

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
create_course_types "Coding Python" "upcoming-schedule/Coding_Python/Level 1-min.png"
create_course_types "Coding Scratch" "upcoming-schedule/Coding_Scratch/Level 1-min.png"
create_course_types "Cyber Virtual Robotics" "upcoming-schedule/cyberVirtualRobotics/Level 1.png"
create_course_types "Lego Robotics" "upcoming-schedule/LEGO_Robotics/Level 1-min.png"
create_course_types "VEX Robotics" "upcoming-schedule/Vex Robotics/Level 1-min.png"
create_course_types "AppInventor Mobile Apps" "$USERPROFILE/Downloads/stemex-curriculum/AppInventor/Level 1-min.png"
create_course_types "A.I.& Machine Learning" "$USERPROFILE/Downloads/stemex-curriculum/A.I.& Machine Learning/Level 1-min.png"
create_course_types "3D Design & Printing" "upcoming-schedule/3D_Design_Printing/Level 1-min.png"

# set "Micro:bits"
# set "upcoming-schedule/Micro_bit/Level 1-min.png"
# set "output=Micro_bit"
# set "AR/VR"
# set "upcoming-schedule/AR_VR/Level 1-min.png"
# set "output=AR_VR"
# set "Smart City"
# set "upcoming-schedule/Smart City/Level 1-min.png"

read -p "done"