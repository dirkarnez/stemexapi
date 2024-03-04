#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function create_course () {
#    course_desplay_name=$1
#    course_icon_path=$2
#    course_json_output=$(echo $1 | sed 's/\//_/g' | sed 's/:/_/g')

#    echo "$course_desplay_name $course_icon_path $course_json_output"

# parent_id=$(jq-windows-amd64.exe --raw-output '.id' "./Coding Roblox.json")

parent_id=$(jq-windows-amd64.exe --raw-output '.id' "./AppInventor Mobile Apps.json")
description=
echo $parent_id
#    curl -X POST --location "https://localhost/api/curriculum-course-type" -b cookie.txt --insecure \
#       --form "description=$course_desplay_name" \
#       --form "icon_file=@$course_icon_path" \
#       --output "$course_json_output.json"
}


# use this to parse a json file to array of objects and foreach(element => console.log(element.name))



# @REM @REM INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`, `seq_no_same_level`)
# @REM @REM SELECT '', curriculum_entries.`id`, files.`id`, 0
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft'
# @REM @REM AND files.`server_physical_file_name` = '';



# @REM @REM INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`, `seq_no_same_level`)
# @REM @REM SELECT 'Coding Minecraft Introductory', curriculum_entries.`id`, files.`id`, 1
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Introductory-min.png';

# @REM @REM INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`, `seq_no_same_level`)
# @REM @REM SELECT 'Coding Minecraft Intermediate', curriculum_entries.`id`, files.`id`, 2
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Intermediate-min.png';

# @REM @REM INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`, `seq_no_same_level`)
# @REM @REM SELECT 'Coding Minecraft Advanced', curriculum_entries.`id`, files.`id`, 3
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Advanced-min.png';


# @REM @REM INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`, `seq_no_same_level`)
# @REM @REM SELECT 'Coding Minecraft Master', curriculum_entries.`id`, files.`id`, 4
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Master-min.png';


# @REM @REM INSERT INTO curriculum_course_information_entries (`icon_id`, `title`, `content`, `entry_id`)
# @REM @REM SELECT files.`id`, 'A - Medieval Machinations Redstone', 'This course will introduce students to use Redstone, electrical circuitry, in a Medieval Theme. Students make mine carts to gather resources, collaborate to build their kingdom and to defend their castle. They are going to experience a lot of creation, adventure and exploration.', curriculum_entries.`id`
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE curriculum_entries.`description` = 'Coding Minecraft Elementary'
# @REM @REM AND files.`server_physical_file_name` = 'schedule-details/codingMineCraftElementry/elementry1.png';

# @REM @REM INSERT INTO curriculum_course_information_entries (`icon_id`, `title`, `content`, `entry_id`)
# @REM @REM SELECT files.`id`, 'B - Theme Park', 'Everyone loves amusement theme parks. Students will have to navigate a number of engineering and teamwork challenges. They draft blueprints and plan for their parks build and create it collaboratively. They will play around and make it as much like the process of designing a real amusement park.', curriculum_entries.`id`
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE `description` = 'Coding Minecraft Elementary'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 3-Elementary B.png';

# @REM @REM INSERT INTO curriculum_course_information_entries (`icon_id`, `title`, `content`, `entry_id`)
# @REM @REM SELECT files.`id`, 'C - Travelling Into the Future', 'Minecraft is a versatile and fantasy game in which players are immersed in a world made up of various kinds of blocks. In order to use blocks, players must gather resources from the world they are in and can use them to craft new materials, tools or potions. In this lesson, students will be introduced to Minecraft in a future world that will teach them the basics of playing Minecraft and will teach them to work as a team to overcome obstacles and build a survival area in a new area.', curriculum_entries.`id`
# @REM @REM FROM curriculum_entries, files 
# @REM @REM WHERE `description` = 'Coding Minecraft Elementary'
# @REM @REM AND files.`server_physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 3-Elementary C.png';



# icon_file="upcoming-schedule/codingMinecraft/Level 2-Elementary-min.png"
# description="Coding Minecraft Elementary"

# echo $parent_id
# --form "blog_entries.0.external_url=\"https://hk.stemex.org/java-your-way-through-minecraft/\"" \
# --form "blog_entries.0.title=\"JAVA YOUR WAY THROUGH MINECRAFT!\"" \
# --form "blog_entries.1.external_url=\"https://hk.stemex.org/minecraft-kids/\"" \
# --form "blog_entries.1.title=\"Minecraft 編程 - 怎樣提升孩子\"" \

# curl -X POST --location "https://localhost/api/curriculum-course" -b cookie.txt --insecure \
# --form "parent_id=\"$parent_id\"" \
# --form "icon_file=@\"../uploads/$icon_file\"" \
# --form "blog_entries.0.external_url=\"https://hk.stemex.org/java-your-way-through-minecraft/\"" \
# --form "blog_entries.0.title=\"JAVA YOUR WAY THROUGH MINECRAFT!\"" \
# --form "blog_entries.1.external_url=\"https://hk.stemex.org/minecraft-kids/\"" \
# --form "blog_entries.1.title=\"Minecraft 編程 - 怎樣提升孩子\"" \
# --form "description=\"%description%\"" \
# --output "%output%.json"


 "AppInventor Mobile Apps"  "AppInventor Mobile Apps Development Introductory"



curl -X POST --location "https://localhost/api/curriculum-course" -b cookie.txt --insecure \
--form "parent_id=$parent_id" \
--form "icon_file=@\"../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Level 2-Introductory-min.png\"" \
--form "description=$description" \
--form "curriculum_plan_file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/App Inventor Intro Curriculum Guide.pdf" \
--form "youtube_video_entries.0.url=https://www.youtube.com/watch?v=zbpzr_hYwtg" \
--form "blog_entries.0.external_url=https://hk.stemex.org/self-control-app/" \
--form "blog_entries.0.title=從小培養孩子的自控能力 3款提升自控能力的電子應用程式" \
--form "levels.0.name=A" \
--form "levels.0.icon_file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Level 3-Introductory A-min.png" \
--form "levels.0.description=HelloPurr: Tap the Kitty, Hear Him Meow', 'HelloPurr is a simple app that you can build in a very fun way. You will create a button that has a picture of your favorite cat on it, and then program the button so that when it is clicked a "meow" sound plays with some vibrations." \
--form "levels.0.lessons.0.presentation_notes.0.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Introductory [L1-HelloCodi].pptx" \
--form "levels.0.lessons.0.student_notes.0.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Intro _Lesson1_Student Notes.pdf" \
--form "levels.0.lessons.0.teacher_notes.0.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Intro _Lesson1_Teacher Notes.txt" \
--form "levels.0.lessons.0.misc_materials.0.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/Bee-Sound.mp3" \
--form "levels.0.lessons.0.misc_materials.1.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/codi.jpg" \
--form "levels.0.lessons.0.misc_materials.2.file=@../../stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/HelloCodi.aia" \
--output "OUTPUT.json"

# https://localhost/api/curriculum-course?id=6dd4a6d9d2fa11ee9aa006c3bc34e27e
read -p "done"
