#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course_types () {
   course_desplay_name=$1
   course_icon_path=$2
   course_json_output=$(echo $1 | sed 's/\//_/g')

   echo "$course_desplay_name $course_icon_path $course_json_output"
   curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure \
      --form "description=$course_desplay_name" \
      --form "icon_file=@$course_icon_path" \
      --output "$course_json_output.json"
}

create_course_types "Coding Minecraft" "$USERPROFILE/Downloads/stemex-curriculum/Coding Minecraft/Level 1-min.png"
create_course_types "Coding Roblox" "$USERPROFILE/Downloads/stemex-curriculum/Coding Roblox/Level 1-min.png"
create_course_types "Coding Python" "$USERPROFILE/Downloads/stemex-curriculum/Coding Python/Level 1-min.png"
create_course_types "Coding Scratch" "$USERPROFILE/Downloads/stemex-curriculum/Coding Scratch/Level 1-min.png"
create_course_types "Cyber Virtual Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Cyber Virtual Robotics/Level 1.png"
create_course_types "Lego Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Lego Robotics/Level 1-min.png"
create_course_types "VEX Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Vex Robotics/Level 1-min.png"

create_course_types "AppInventor Mobile Apps" "$USERPROFILE/Downloads/stemex-curriculum/AppInventor/Level 1-min.png"
create_course_types "A.I.& Machine Learning" "$USERPROFILE/Downloads/stemex-curriculum/A.I.& Machine Learning/Level 1-min.png"
create_course_types "3D Design & Printing" "$USERPROFILE/Downloads/stemex-curriculum/3D_Design_Printing/Level 1-min.png"
create_course_types "Micro:bits"  "$USERPROFILE/Downloads/stemex-curriculum/Micro_bit/Level 1-min.png"
create_course_types "AR/VR" "$USERPROFILE/Downloads/stemex-curriculum/AR_VR/Level 1-min.png"
create_course_types "Smart City" "$USERPROFILE/Downloads/stemex-curriculum/Smart City/Level 1-min.png"

read -p "done"