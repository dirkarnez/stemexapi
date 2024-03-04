#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

# https://localhost/api/curriculum-course?id=
function create_course () {
    course_type_json_output=$(echo $1 | sed 's/\//_/g' | sed 's/:/_/g')
    parent_id=$(jq-windows-amd64.exe --raw-output '.id' "./$course_type_json_output.json")
    course_json_output=$(echo $description | sed 's/\//_/g' | sed 's/:/_/g')
    # $USERPROFILE/Downloads/curl/curl.exe
    
    /C/Windows/System32/curl.exe -X POST --location "https://localhost/api/curriculum-course" -b cookie.txt --insecure \
        --form "parent_id=$parent_id" \
        --form "description=$description" \
        --form "icon_file=@$icon_file" \
        --form "curriculum_plan_file=@$curriculum_plan_file" \
        --form "blog_entries.0.external_url=$blog_entries_0_external_url" \
        --form "blog_entries.0.title=$blog_entries_0_title" \
        --form "levels.0.name=$levels_0_name" \
        --form "levels.0.icon_file=@$levels_0_icon_file" \
        --form "levels.0.description=$levels_0_description" \
        --form "levels.0.lessons.0.lesson_number=1" \
        --form "levels.0.lessons.0.presentation_notes.0.id=" \
        --form "levels.0.lessons.0.presentation_notes.0.file=@$levels_0_lessons_0_presentation_notes_0_file" \
        --form "levels.0.lessons.0.student_notes.0.id=" \
        --form "levels.0.lessons.0.student_notes.0.file=@$levels_0_lessons_0_student_notes_0_file" \
        --form "levels.0.lessons.0.teacher_notes.0.id=" \
        --form "levels.0.lessons.0.teacher_notes.0.file=@$levels_0_lessons_0_teacher_notes_0_file" \
        --form "levels.0.lessons.0.misc_materials.0.id=" \
        --form "levels.0.lessons.0.misc_materials.0.file=@$levels_0_lessons_0_misc_materials_0_file" \
        --form "levels.0.lessons.0.misc_materials.1.id=" \
        --form "levels.0.lessons.0.misc_materials.1.file=@$levels_0_lessons_0_misc_materials_1_file" \
        --form "levels.0.lessons.0.misc_materials.2.id=" \
        --form "levels.0.lessons.0.misc_materials.2.file=@$levels_0_lessons_0_misc_materials_2_file" \
        --output "$course_json_output.json"
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


export icon_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Level 2-Introductory-min.png"
export description="AppInventor Mobile Apps Development Introductory"
export curriculum_plan_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/App Inventor Intro Curriculum Guide.pdf"
export blog_entries_0_external_url="https://hk.stemex.org/self-control-app/" 
export blog_entries_0_title="從小培養孩子的自控能力 3款提升自控能力的電子應用程式" 
export youtube_video_entries_0_url="https://www.youtube.com/watch?v=zbpzr_hYwtg"
export levels_0_name="A"
export levels_0_icon_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Level 3-Introductory A-min.png" 
export levels_0_description="HelloPurr: Tap the Kitty, Hear Him Meow', 'HelloPurr is a simple app that you can build in a very fun way. You will create a button that has a picture of your favorite cat on it, and then program the button so that when it is clicked a "meow" sound plays with some vibrations." 
export levels_0_lessons_0_presentation_notes_0_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Introductory [L1-HelloCodi].pptx" 
export levels_0_lessons_0_student_notes_0_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Intro _Lesson1_Student Notes.pdf" 
export levels_0_lessons_0_teacher_notes_0_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/App Inventor Intro _Lesson1_Teacher Notes.txt" 
export levels_0_lessons_0_misc_materials_0_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/Bee-Sound.mp3" 
export levels_0_lessons_0_misc_materials_1_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/codi.jpg" 
export levels_0_lessons_0_misc_materials_2_file="$USERPROFILE/Downloads/stemex-curriculum/AppInventor/STEMex_AppInventor_Introductory/Lesson 1/HelloCodi.aia" 

create_course "AppInventor Mobile Apps"

# https://localhost/api/curriculum-course?id=6dd4a6d9d2fa11ee9aa006c3bc34e27e
read -p "done"
