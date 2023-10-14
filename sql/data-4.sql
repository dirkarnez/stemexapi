INSERT INTO curriculum_course_you_tube_videos (`url`, `entry_id`)
SELECT 'https://www.youtube.com/embed/0SLnKsFWwFA', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Elementary';

