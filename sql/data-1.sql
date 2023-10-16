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
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Advance-min.png'); 
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingMinecraft/Level 2-Master-min.png'); 

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingRoblox/Level 2- Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/codingRoblox/Level 2-Intermediate-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Coding_Python/Level 2-Advance-min.png');
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
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/LEGO_Robotics/Level 2-Advance-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/Vex Robotics/Level 1-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/AppInventor Mobile Apps/Level 2-Advance-min.png');

INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Introductory-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Intermediate-min.png');
INSERT INTO files(`physical_file_name`) VALUES('upcoming-schedule/A.I. _ Machine Learning/Level 2-Advance-min.png');

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
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Minecraft Elementary', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Minecraft Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Minecraft Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Minecraft Advance', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Minecraft Master', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft';
-- 1

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Micro:bit Digital Making Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bits';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Micro:bit Digital Making Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bits';


-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Roblox Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Roblox Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox';


-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Python Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Python Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Python Advanced', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Python Master', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Scratch Jr', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Scratch Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Coding Scratch Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Cyber Virtual Robotics Elementary', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Cyber Virtual Robotics Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Cyber Virtual Robotics Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Cyber Virtual Robotics Master', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Lego Robotics Elementary - Wedo', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Lego Robotics Introductory - EV3', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Lego Robotics Intermediate - EV3', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Lego Robotics Advanced - EV3', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'VEX Robotics Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'VEX Robotics';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AppInventor Mobile Apps Development Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AppInventor Mobile Apps Development Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AppInventor Mobile Apps Development Advanced', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'A.I. & Machine Learning Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'A.I. & Machine Learning Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'A.I. & Machine Learning Advanced', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT '3D Design & Printing Introductory', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT '3D Design & Printing Intermediate', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Micro:bit Digital Making Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bits';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Micro:bit Digital Making Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bits';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AR/VR Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AR/VR Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR';
-- 1
INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Smart City Elementary', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Smart City Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'Smart City Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City';