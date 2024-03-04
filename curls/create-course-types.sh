#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course_types () {
   course_desplay_name = $1
   course_icon_path = $2
   course_json_output = $(echo $1 | sed 's/\//-/g')

   echo "$course_desplay_name $course_icon_path $course_json_output"
   # curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure \
   
   # --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"
   # set "description="
# set "file="
# curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure ^
# --form "description=\"%description%\"" --form "icon_file=@\"../uploads/%file%\"" --output "%description%.json"
}

# create_course_types "Coding Minecraft" "upcoming-schedule/codingMinecraft/Level 1-min.png"
# create_course_types "Coding Roblox" "upcoming-schedule/codingRoblox/Level 1-min.png"
# create_course_types "Coding Python" "upcoming-schedule/Coding_Python/Level 1-min.png"
# create_course_types "Coding Scratch" "upcoming-schedule/Coding_Scratch/Level 1-min.png"
# create_course_types "Cyber Virtual Robotics" "upcoming-schedule/cyberVirtualRobotics/Level 1.png"
# create_course_types "Lego Robotics" "upcoming-schedule/LEGO_Robotics/Level 1-min.png"
# create_course_types "VEX Robotics" "upcoming-schedule/Vex Robotics/Level 1-min.png"
# create_course_types "AppInventor Mobile Apps" "$USERPROFILE/Downloads/stemex-curriculum/AppInventor/Level 1-min.png"
# create_course_types "A.I.& Machine Learning" "$USERPROFILE/Downloads/stemex-curriculum/A.I.& Machine Learning/Level 1-min.png"
# create_course_types "3D Design & Printing" "upcoming-schedule/3D_Design_Printing/Level 1-min.png"
# create_course_types "Micro:bits"  "upcoming-schedule/Micro_bit/Level 1-min.png"
create_course_types "AR/VR" "upcoming-schedule/AR_VR/Level 1-min.png"
# create_course_types "Smart City" "upcoming-schedule/Smart City/Level 1-min.png"

read -p "done"