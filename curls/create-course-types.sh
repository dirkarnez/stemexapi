#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course_type () {
   course_desplay_name=$1
   course_icon_path=$2
   course_type_json_output=$(echo $1 | sed 's/\//_/g' | sed 's/:/_/g')

   echo "$course_desplay_name $course_icon_path $course_type_json_output"

   curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure \
      --form "description=$course_desplay_name" \
      --form "icon_file=@$course_icon_path" \
      --output "$course_type_json_output.json"
}

create_course_type "Coding Minecraft" "$USERPROFILE/Downloads/stemex-curriculum/Coding Minecraft/Level 1-min.png"
create_course_type "Coding Roblox" "$USERPROFILE/Downloads/stemex-curriculum/Coding Roblox/Level 1-min.png"
create_course_type "Coding Python" "$USERPROFILE/Downloads/stemex-curriculum/Coding Python/Level 1-min.png"
create_course_type "Coding Scratch" "$USERPROFILE/Downloads/stemex-curriculum/Coding Scratch/Level 1-min.png"
create_course_type "Cyber Virtual Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Cyber Virtual Robotics/Level 1.png"
create_course_type "Lego Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Lego Robotics/Level 1-min.png"
create_course_type "VEX Robotics" "$USERPROFILE/Downloads/stemex-curriculum/Vex Robotics/Level 1-min.png"
create_course_type "AppInventor Mobile Apps" "$USERPROFILE/Downloads/stemex-curriculum/AppInventor/Level 1-min.png"
create_course_type "A.I.& Machine Learning" "$USERPROFILE/Downloads/stemex-curriculum/A.I.& Machine Learning/Level 1-min.png"
create_course_type "3D Design & Printing" "$USERPROFILE/Downloads/stemex-curriculum/3D_Design_Printing/Level 1-min.png"
create_course_type "Micro:bits"  "$USERPROFILE/Downloads/stemex-curriculum/Micro_bit/Level 1-min.png"
create_course_type "AR/VR" "$USERPROFILE/Downloads/stemex-curriculum/AR_VR/Level 1-min.png"
create_course_type "Smart City" "$USERPROFILE/Downloads/stemex-curriculum/Smart City/Level 1-min.png"

read -p "done"