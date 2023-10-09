TRUNCATE TABLE curriculum_entries;

INSERT INTO curriculum_entries(`description`) VALUES('Coding Minecraft');
INSERT INTO curriculum_entries(`description`) VALUES('Coding Roblox');
INSERT INTO curriculum_entries(`description`) VALUES('Coding Python');
INSERT INTO curriculum_entries(`description`) VALUES('Coding Scratch');
INSERT INTO curriculum_entries(`description`) VALUES('Cyber Virtual Robotics');
INSERT INTO curriculum_entries(`description`) VALUES('Lego Robotics');
INSERT INTO curriculum_entries(`description`) VALUES('VEX Robotics');
INSERT INTO curriculum_entries(`description`) VALUES('AppInventor Mobile Apps');
INSERT INTO curriculum_entries(`description`) VALUES('A.I. & Machine Learning');
INSERT INTO curriculum_entries(`description`) VALUES('3D Design & Printing');
INSERT INTO curriculum_entries(`description`) VALUES('Micro:bits');
INSERT INTO curriculum_entries(`description`) VALUES('AR / VR');
INSERT INTO curriculum_entries(`description`) VALUES('Smart City'); 

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
SELECT 'AppInventor Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AppInventor Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AppInventor Advanced', `id`
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
SELECT 'AR / VR Introductory', `id`
FROM curriculum_entries
WHERE `description` = 'AR / VR';

INSERT INTO curriculum_entries (`description`, `parent_id`)
SELECT 'AR / VR Intermediate', `id`
FROM curriculum_entries
WHERE `description` = 'AR / VR';
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