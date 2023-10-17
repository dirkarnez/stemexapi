INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingRoblox/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Scratch/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/cyberVirtualRobotics/Level 1.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Vex Robotics/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/3D_Design_Printing/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Micro_bit/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AR_VR/Level 1-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Smart City/Level 1-min.png'); 

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Elementary-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Introductory-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Intermediate-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Advanced-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Master-min.png'); 

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingRoblox/Level 2- Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingRoblox/Level 2-Intermediate-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Advanced-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Master-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Scratch/Level 2-Scratch Jr-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Scratch/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Scratch/Level 2-Intermediate-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/cyberVirtualRobotics/Level 2-Elementary.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/cyberVirtualRobotics/Level 2-Introductory.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/cyberVirtualRobotics/Level 2-Intermediate.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/cyberVirtualRobotics/Level 2-Master.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 2-Elementary-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 2-Advanced-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Advanced-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Advanced-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/3D_Design_Printing/Level 2-Introductory-min.jpg');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/3D_Design_Printing/Level 2-Intermediate-min.jpg');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Micro_bit/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Micro_bit/Level 2-Intermediate-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AR_VR/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AR_VR/Level 2-Intermediate-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Smart City/Level 2-Elementary-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Smart City/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Smart City/Level 2-Intermediate-min.png');

-- TRUNCATE TABLE curriculum_entries;

INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Coding Minecraft', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Coding Roblox', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/codingRoblox/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Coding Python', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/Coding_Python/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Coding Scratch', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/Coding_Scratch/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Cyber Virtual Robotics', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/cyberVirtualRobotics/Level 1.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Lego Robotics', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/LEGO_Robotics/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'VEX Robotics', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/Vex Robotics/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'AppInventor Mobile Apps', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/AppInventor Mobile Apps/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'A.I. & Machine Learning', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/A.I. _ Machine Learning/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT '3D Design & Printing', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/3D_Design_Printing/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Micro:bits', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/Micro_bit/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'AR/VR', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/AR_VR/Level 1-min.png';
INSERT INTO curriculum_entries(`description`, `icon_id`) SELECT 'Smart City', `id` FROM files WHERE `physical_file_name` = 'upcoming-schedule/Smart City/Level 1-min.png';


-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Minecraft Elementary', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Minecraft'
AND files.`physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Elementary-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Minecraft Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Minecraft'
AND files.`physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Introductory-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Minecraft Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Minecraft'
AND files.`physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Intermediate-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Minecraft Advanced', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Minecraft'
AND files.`physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Advanced-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Minecraft Master', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Minecraft'
AND files.`physical_file_name` = 'upcoming-schedule/codingMinecraft/Level 2-Master-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Roblox Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Roblox'
AND files.`physical_file_name` = 'upcoming-schedule/codingRoblox/Level 2- Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Roblox Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Roblox'
AND files.`physical_file_name` = 'upcoming-schedule/codingRoblox/Level 2-Intermediate-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Python Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Python'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Python/Level 2-Introductory-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Python Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Python'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Python/Level 2-Intermediate-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Python Advanced', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Python'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Python/Level 2-Advanced-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Python Master', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Python'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Python/Level 2-Master-min.png';



-- 1

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Scratch Jr', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Scratch'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Scratch/Level 2-Scratch Jr-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Scratch Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Scratch'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Scratch/Level 2-Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Coding Scratch Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Coding Scratch'
AND files.`physical_file_name` = 'upcoming-schedule/Coding_Scratch/Level 2-Intermediate-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Cyber Virtual Robotics Elementary', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Cyber Virtual Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/cyberVirtualRobotics/Level 2-Elementary.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Cyber Virtual Robotics Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Cyber Virtual Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/cyberVirtualRobotics/Level 2-Introductory.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Cyber Virtual Robotics Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Cyber Virtual Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/cyberVirtualRobotics/Level 2-Intermediate.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Cyber Virtual Robotics Master', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Cyber Virtual Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/cyberVirtualRobotics/Level 2-Master.png';


-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Lego Robotics Elementary - Wedo', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Lego Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/LEGO_Robotics/Level 2-Elementary-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Lego Robotics Introductory - EV3', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Lego Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/LEGO_Robotics/Level 2-Introductory-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Lego Robotics Intermediate - EV3', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Lego Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/LEGO_Robotics/Level 2-Intermediate-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Lego Robotics Advanced - EV3', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Lego Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/LEGO_Robotics/Level 2-Advanced-min.png';


-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'VEX Robotics Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'VEX Robotics'
AND files.`physical_file_name` = 'upcoming-schedule/Vex Robotics/Level 1-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'AppInventor Mobile Apps Development Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'AppInventor Mobile Apps'
AND files.`physical_file_name` = 'upcoming-schedule/AppInventor Mobile Apps/Level 2-Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'AppInventor Mobile Apps Development Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'AppInventor Mobile Apps'
AND files.`physical_file_name` = 'upcoming-schedule/AppInventor Mobile Apps/Level 2-Intermediate-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'AppInventor Mobile Apps Development Advanced', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'AppInventor Mobile Apps'
AND files.`physical_file_name` = 'upcoming-schedule/AppInventor Mobile Apps/Level 2-Advanced-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'A.I. & Machine Learning Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'A.I. & Machine Learning'
AND files.`physical_file_name` = 'upcoming-schedule/A.I. _ Machine Learning/Level 2-Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'A.I. & Machine Learning Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'A.I. & Machine Learning'
AND files.`physical_file_name` = 'upcoming-schedule/A.I. _ Machine Learning/Level 2-Intermediate-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'A.I. & Machine Learning Advanced', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'A.I. & Machine Learning'
AND files.`physical_file_name` = 'upcoming-schedule/A.I. _ Machine Learning/Level 2-Advanced-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT '3D Design & Printing Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = '3D Design & Printing'
AND files.`physical_file_name` = 'upcoming-schedule/3D_Design_Printing/Level 2-Introductory-min.jpg';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT '3D Design & Printing Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = '3D Design & Printing'
AND files.`physical_file_name` = 'upcoming-schedule/3D_Design_Printing/Level 2-Intermediate-min.jpg';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Micro:bit Digital Making Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Micro:bits'
AND files.`physical_file_name` = 'upcoming-schedule/Micro_bit/Level 2-Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Micro:bit Digital Making Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Micro:bits'
AND files.`physical_file_name` = 'upcoming-schedule/Micro_bit/Level 2-Intermediate-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'AR/VR Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'AR/VR'
AND files.`physical_file_name` = 'upcoming-schedule/AR_VR/Level 2-Introductory-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'AR/VR Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'AR/VR'
AND files.`physical_file_name` = 'upcoming-schedule/AR_VR/Level 2-Intermediate-min.png';

-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Smart City Elementary', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Smart City'
AND files.`physical_file_name` = 'upcoming-schedule/Smart City/Level 2-Elementary-min.png';

INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Smart City Introductory', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Smart City'
AND files.`physical_file_name` = 'upcoming-schedule/Smart City/Level 2-Introductory-min.png';


INSERT INTO curriculum_entries (`description`, `parent_id`, `icon_id`)
SELECT 'Smart City Intermediate', curriculum_entries.`id`, files.`id`
FROM curriculum_entries, files 
WHERE curriculum_entries.`description` = 'Smart City'
AND files.`physical_file_name` = 'upcoming-schedule/Smart City/Level 2-Intermediate-min.png';