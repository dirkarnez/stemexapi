INSERT INTO curriculum_course_youtube_video_entries (`url`, `title`, `entry_id`)
SELECT 'https://www.youtube.com/embed/0SLnKsFWwFA', '', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Elementary';

