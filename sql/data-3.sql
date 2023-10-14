-- Coding Minecraft
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Medieval Machinations Redstone', 'This course will introduce students to use Redstone, electrical circuitry, in a Medieval Theme. Students make mine carts to gather resources, collaborate to build their kingdom and to defend their castle. They are going to experience a lot of creation, adventure and exploration.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Elementary';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Theme Park', 'Everyone loves amusement theme parks. Students will have to navigate a number of engineering and teamwork challenges. They draft blueprints and plan for their parks build and create it collaboratively. They will play around and make it as much like the process of designing a real amusement park.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Elementary';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Travelling Into the Future', 'Minecraft is a versatile and fantasy game in which players are immersed in a world made up of various kinds of blocks. In order to use blocks, players must gather resources from the world they are in and can use them to craft new materials, tools or potions. In this lesson, students will be introduced to Minecraft in a future world that will teach them the basics of playing Minecraft and will teach them to work as a team to overcome obstacles and build a survival area in a new area.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Elementary';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Sheep City', 'Changing arrow’s explosive power, bounciness of golden block, game mode and difficulties…from basic programming components to more complex changes, students will have fun changing/programming the Minecraft worlds to their like.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Heist and Seek', 'Be the player to mine the most blocks in this timed hunt for stolen goods! Watch out, there’s a bank robber in town and they’ve hidden their stolen goods all over the map! It’s your job to go head to head against the other players and find the most boxes to win! Use loops, conditionals, and timers to add players into different teams and add different rounds into a treasure hunt game.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Spartan School', 'Build an infinite mob arena game, to fight alone or with friends! Oh, and did we mention you’ll be fighting blazes whilst you do it? Want to be the greatest Minecraft Spartan warrior of all time? Learn programming basics while creating wave after wave of mobs to fight in a Spartan training arena. Learn programming basics such as loops, methods and variables to create this mob fighting mini-game. Battle increasingly harder waves of enemies that multiply every round.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-EggWars', 'Battle your friends to have the last egg still intact and become champion of the server! Details to tend to include notification on eggs being placed and eggs being broken, signals on game start, building base for eggs, etc. This course will allow you to expand your knowledge of adding rules to a PVP game, as well as learn more about structure generation and for loops', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Rise Of The Werebunnies', 'Create a multiplayer fight for survival, complete with exploding eggs! Beware! The moon is full and the werebunnies are out…This course shows you how to split players into different teams and add a scoring system. This course uses loops and conditionals to split players into different teams with different characteristics, then gives players scores when they defeat their enemy.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'E-Hungary Games', 'Build your own hungry games style server to make the game how you want it to be!', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Intermediate';



INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Haunted Mansion', 'Troll your friend with three spooky mode to scare them in a haunted house!', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Flappy Block', 'Rack up a high score and try not touch the lava in our 3D version of Floppy Bird in Minecraft.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Advanced';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Dances to the Moosic', 'Create a cow themed nightclub where your players have to dance in time to the moo-sic!', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Swoop De Loop', 'Learn how to create rings using mathematical knowledge which give players a superboost when gliding.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Minecraft Master';

-- Coding Roblox

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Prison Escape', 'Duck through the maze, avoid the flashing lasers and break out of prison armed with your trusty dynamite slingshot – just try not to get blown up! This course is great for beginners. Get to grips with the code editor and learn how to make your Roblox game from scratch using functions and conditionals.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Box Car Racer', 'Create a box racing game where the fastest racer wins! Students learn how to code checkpoints, a finish line and write clean code in this fast-paced box racing game! This Roblox course shows you how to structure code well and add a finish line to the track using inheritance.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Wrestleball', 'Players must face off in the wrestling arena, firing balls to knock each other out of the centre. The player who stays in the centre the longest the winner! In this course you’ll learn how to make an arena-based PVP game using vectors, loops and a score system. It’s perfect for confident Roblox fans who are familiar with the Code Editor.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Ninja Obby', 'Work on your particular skills while navigating this tricky obstacle course!', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Riddle Ruin', 'Crack the codes inside the pyramid to set yourself free!', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'E-Platform Game Design', 'Avoid red obstacles and collect coins in this 2D platform Roblox game.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Roblox Intermediate';

-- Coding Python

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Python Turtle', 'What is AI? How does a machine learn? Let your kids start to know more about Python. This course is suitable for students who have a little experience in coding. They will learn and understand the Python turtle library and graphics. They will be challenged to animate a clock and to control a spinner enhancing their creativity.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Python Game Design', 'Students will step into simple game design using Python. They will learn how to interact with the computer on screen or using keyboard. They are going to apply coding to design racing, word guess game plus other challenges! Sharing amongst friends and further exploration are encouraged.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Python Simon Says', 'Students are going to have a comprehensive understanding on the data structure in Python. They will advance their skills in using Python editor. Apply their understanding in data structure and engineering design process to create games as challenges.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Python Flappy & Snake Games', 'Students learn how to utilize other python resources. They will concentrate on coding the movement of objects using vectors and control. A lot of logic training will be involved in understanding the conditions in gaming and how to solve them in coding.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Python Ping Pong', 'In this course, students learn complicated interaction in game design. They are going to apply knowledge in python to solve geniune challenges and stunning geometric dancing figures. At the end they simulate the classic game of life revealing special patterns from simple rules.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Python Hangman and Pacman', 'In this course, students learn complicated interaction in game design. They are going to apply knowledge in python to solve geninue challenges and stunning geometric dancing figures. At the end they simulate the classic game of life revealing speical patterns from simple rules.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Intermediate';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Python Game Concept', 'With their deeper understanding of Python and its various structures, students will take a deep dive into the pygame module and other modules, as well as a deeper understanding into data types, to learn more about how they can be used to create more complex program without having to write everything from scratch.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Python Discord Bot', 'Students will learn to creating a discord bot and how to do with language processing. In the course, students will have a chance to do some works related to the concept of neural network.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Python Complex Game Design', 'Students will learn the most popular python module pygame. Through designing the classic space invader game, they will consolidate the skills in game design. A lot of revisions of the python coding structure and logic deductions in solving problems involved. Use of different platforms in coding, how to utilize resources in machine learning.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Python Machine Learning', 'Students will learn how to use tensflow and the use of the module keraus. They are going to learn how to build up a machine learning module to identify different clothing from thousands of images.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Advanced';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Yolo', 'We will learn how to use the yolov3 model to detect objects present in an image. It will help differentiate different objects.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Tic tac toe', 'We will learn about graphical user interface in this project, learn about Tkinter and use it to make a game GUI of tic tac toe and we will learn logic of the game design .', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Brick Breaker game', 'The kids will learn about game design theory. We will start with a simple Brick breaker game in which there is a ball that bounces off a platform to break a brick wall and the player has to keep the ball going by making sure the paddle is always there to bounce off the ball back.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Photo Editing App', 'We will learn about the pillow library which is used for image processing, we will learn about how we can edit images using python.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Python Master';

-- Coding Scratch
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Cartoon & mini-Game Creation', 'Learn by doing and problem solving. Make characters move, jump and sing.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Jr';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Rock n Roll Animation', 'This course is suitable for first time coders. Enjoy coding with interesting challenges that foster student’s creativity and imagination. Animate their favorite names, program a rock and roll band like creating a band using different instruments, with sound effect plus animation. Finally, students will create their own unique story using scratch.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Animate an Adventure Game', 'With Scratch, you can program your own interactive stories, games, and animations — and share your creations with others in the online community. Scratch helps young people learn to think creatively, reason systematically, and work collaboratively — essential skills for life in the 21st century.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Animate the Crab / Maze Starter', 'With Scratch, you can program your own interactive stories, games, and animations — and share your creations with others in the online community. Here you will use coordinates, random number and forever loop to make a crab with different costumes and move around randomly.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-The Pico Show / Hide and Seek', 'With Scratch, you can program your own interactive stories, games, and animations — and share your creations with others in the online community. Here you will use function and variables to create a dance party with music and background.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Balloon Pop / Catching Game', 'With Scratch, you can program your own interactive stories, games, and animations — and share your creations with others in the online community. Use video sensing and conditional statement to set up a score system that keeps track of score when balloons move around randomly and being pop with hand touching them.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-Dance Party / Pong Starter', 'With Scratch, you can program your own interactive stories, games, and animations — and share your creations with others in the online community. Use wait, loop and conditional statement to create your own sprites with different costumes and music played for sprites to dance.', `id`
FROM curriculum_entries
WHERE `description` = 'Coding Scratch Intermediate';



-- Cyber Virtual Robotics
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-CVR Jr Adventure', 'Explore different worlds and terrain with LEGO Education SPIKE Prime. This course is ideal for kids with no experience in coding robots. Through navigating through unique maps, kids will learn how to precisely instruct the robot through specific maneuver and basic coding logic.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Elementary';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Cyber Robotics 101', 'Want to experience virtual robotics but scared it is too difficult? In this course you will learn all the basic topics of robotics to jumpstart your virtual robotic journey.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Physics with Ruby', 'Want to experience virtual robotics but scared it is too difficult? In this course you will learn all the basic topics of robotics to jumpstart your virtual robotic journey.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Tour with Ruby', 'With ruby''s excellent navigation tools, help Ruby explore through thin and dangerous road. This course will encourage students to plan for the most efficient route through using different tools such as color sensors, precision manuevers and more', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Road Safety with Ruby', 'Our roads serve different kinds of vehicles, from buses, trucks and taxis. With all these vehicles on the road, we must ensure the safety of everyone travelling. In this course, Ruby will help you navigate the roads and how we can make the roads safer and vehicles smarter', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Intermediate';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-CVR Python Gym - Code Robot Pioneer', 'This course suits students who have some experience in CoderZ 102 to pursue coding robot using python. In this course they learn the basic of using python to control their Ruby to explore different terrains. They need to apply some science idea and calculation to solve the problems to accomplish the tasks to be the Robot Pioneer!', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-CVR Python Gym – Code Robot Pioneer', 'In Smart Robot, students learn Python to let Ruby to perform amazing art skills. With Ruby''s sensitive sensors plus more complex coding skills, they learn how to utilise for cool tricks, make complex movement. Finally, to draw their unique name as the milestone to proceed to the next course.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-CVR Python Gym – Code Action Robot', 'Ruby is going take action and arm up! With the introduction of magnetic arm and other sensors Ruby is going to maneuver objects here and there. To meet the challenges, students need to learn more complex python and do a lot of testing and improvement.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Master';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D-CVR Python Gym – Code Intelligent Robot', 'Ruby is getting more and more intelligent. In this course, students will use a number of sensors and other their previous coding skills to code a self driving Ruby without hitting obstacles. They will find it very challenging as the problem will get harder and more complex.', `id`
FROM curriculum_entries
WHERE `description` = 'Cyber Virtual Robotics Master';

-- Lego Robotics 
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Let’s Get Moving', 'Learn about the basics of mechanical engineering all based around the idea of moving, and moving very quickly. In this course, students will build various models, such as ships and race cars, in order to learn about how motors and gears function.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Elementary - Wedo';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Wild Animals', 'Learn about mechanical engineering based on the animal kingdom. In this course, students will build various models, such as lions and birds, in order to learn about how motors and gears function, as well as little facts about the animals themselves.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Elementary - Wedo';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Rescue Heroes', 'Learn about mechanical engineering based on natural disasters. In this course, students will build various models, such as a helicopter, in order to learn about how motors and gears function, as well as how natural disasters can be prevented and how people can be rescued.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Elementary - Wedo';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Robotics 101', 'A robot is a device that is designed and programmed to perform a task either autonomously or with human input. Robots typically come in two forms: those resembling humans or role-specific robots, such as NASA space probes and Mars Rovers. Robots are generally used to perform either dangerous or monotonous tasks. The challenge facing robotics engineers is that the robot knows only what is written into the program. The design of the robot must also be capable of performing the task at hand. In this unit, students will experience both the designing and programming roles of being a robotics engineer.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Introductory - EV3';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Olympics', 'The Olympics curriculum is designed to introduce students to the world of not only building, but also programming basic robots. Throughout this course, the students will explore different ways in which a robot could be utilized to engage in various challenges related to the Olympics.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Introductory - EV3';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Missions to Mars', 'This course is designed to introduce students to the world of not only building, but also programming basic robots. Throughout this course the students will explore different ways in which a robot could be utilized to explore a distant planet.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Intermediate - EV3';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Envirobots', 'In Rescue EnviroBots, students will design and programme robots to help accomplish environmentally friendly tasks. By creating robots that can transfer nuclear waste, mine raw minerals, and deliver food and goods more efficiently, they will be sure to contribute to a more sustainable environment.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Intermediate - EV3';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Ocean Missions', 'In the Ocean Missions curriculum, students will be introduced to the world of robotics in an interesting and engaging way. The goal is to teach students about the building and programming aspects of robotics as it relates to real-world issues in ocean exploration.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Advanced - EV3';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Robot Sergeons', 'The Robot Surgeons curriculum is designed to introduce students to the world of not only building, but also programming basic robots. Throughout this course, the students will explore different ways in which a robot could be utilized in the medical field.', `id`
FROM curriculum_entries
WHERE `description` = 'Lego Robotics Advanced - EV3';

-- VEX Robotics
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-High Rise Challenges', 'Meet VEX GO. An affordable construction system that teaches the fundamentals of STEM through fun, hands-on activities that help young students perceive coding and engineering in a fun and positive way!', `id`
FROM curriculum_entries
WHERE `description` = 'VEX Robotics Introductory';

-- AppInventor Mobile Apps Development
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-HelloPurr: Tap the Kitty, Hear Him Meow', 'HelloPurr is a simple app that you can build in a very fun way. You will create a button that has a picture of your favorite cat on it, and then program the button so that when it is clicked a "meow" sound plays with some vibrations.', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Piccall', 'PicCall shows you how you can use App Inventor to make apps that do actual things, like calling friends. We will learn about how real-life applications work and are programmed.', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-QuizMe', 'Youll design the quiz game so that the user proceeds from question to question by clicking a Next button and receives simple correct/incorrect feedback on each answer', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Snow Globe', 'In this project you will create a virtual “Snow Globe” that displays snowflakes falling randomly on New York City at night whenever you shake your Android device. You will be introduced to the “Any Component” advanced feature in App Inventor which is used to give collective behaviors to components of the same type', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Intermediate';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Android, Where''s My Car?', 'Time to use the advanced features for app inventors to remember where you parked your car in case you go to a new location and are not familiar with it. With your very own app and your mobile device we can pinpoint and remember it using the sensors in our devices.', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Firebase Authentication in App Inventor Using Javascript', 'The kids will learn what firebase is and set up for it and how we use it for authentication purposes in google and update any number of apps with fresh data, how data is managed in it.', `id`
FROM curriculum_entries
WHERE `description` = 'AppInventor Mobile Apps Development Advanced';

-- A.I. & Machine Learning

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Chameleon', 'In this project kids will make a chameleon that changes color to match its background using a webcam to take pictures of different colors, then use machine learning with those examples to train the chameleon to recognize colors.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Shoot your goal', 'In this project you will train a computer to play a simple arcade game. The game is based on shooting balls at a target. You can’t aim at the target directly because there is a wall in the way, so you need to bounce the ball off a wall to do it. You will teach the computer to be able to play this game by collecting examples of shots that hit and miss, so that it can learn to make predictions about the shots it can take.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A- Chatbots', 'In this project you will make a chatbot that can answer questions about a topic of your choice.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B- Zombie Escape!', 'You are trapped in a huge hotel that has been overrun by zombies! To help you escape, you have a small remote-controlled robot.There’s no point trying to use it to memorize where the zombies are –there are too many rooms and too many zombies, and they’re all moving around the hotel too much anyway. You need to make your robot learn.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C- Secret Code', 'In this project you will train the computer to understand secret code Words. You’ll use that to say commands to a spy to guide it around a town.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'D- Laser Eyes', 'In this project you will make voice-powered laser eyes! This is a game where you will see laser beams shoot from your eyes in your computer’s webcam. You will use these to shoot at bottles. Your laser eyes will be voice-activated, so you will have to shout “laser eyes” to make them shoot. You will be using two kinds of machine learning model. Speech recognition to activate the lasers and face detection to aim them!', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Intermediate';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A- Top Trumps', 'In this project you will train a computer to play a card game. For some values on the cards, you win by having the highest number. For others, you win by having the lowest. The range of numbers for different values will vary. The aim will be for the computer to learn how to play the game well without you having to give it a list of all the cards or tell it the rules.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Advanced';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B- Phishing', 'People are sent links to these fake phishing websites in emails or instant messages. How can they know if a link is safe to click on? In this project, you will learn about the research that is being done to train machine learning systems to predict if a link is to a phishing website or a legitimate website.', `id`
FROM curriculum_entries
WHERE `description` = 'A.I. & Machine Learning Advanced';

-- 3D Design & Printing
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-3D Design & Printing Introductory', 'Learn to create your very own 3D structures that can be printed in the future. At the same time, students will be able to learn about how to use TinkerCAD and its various tools, such as alignment tools and hole generation.', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Superhero', 'In this topic, student will create their superhero with their designed costume, wings and decoration through learning mirror tool, rotation, wokrplane. Student will also revisiting Boolean addition, duplication, scaling and grouping', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Architecture', 'In this topic of 3D printing, students will learn more about the functions of TinkerCAD, such as scaling and aligning objects in architecture design. Through learning different style of famous architectures. student will have an opportunity to create a Japanese style architecture and Roman Dome with columns architecture for their masterpieces.', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Transport', 'In this topic, students will consolidate their knowledge gained in Tickercad to create their own car using the balloon connector and to design the best boat by exploring buoyancy designed Sea Craft. Student will unleash their creativities from planning, designing and to the building process.', `id`
FROM curriculum_entries
WHERE `description` = '3D Design & Printing Intermediate';

-- Micro:bit Digital Making

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Micro:bit Introductory', 'The micro:bit is a small computer that is well suited for introducing how software and hardware work together to perform tasks. It has an LED light display, buttons, sensors, and many input/output features that can be coded and physically interacted with.', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bit Digital Making Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Micro:bit Rock, Paper Scissors', 'Rock, paper, scissors is a simple game that everyone has played at least once in their life. But can this simple game be created using Micro:bit, the answer is yes, yes it can.', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bit Digital Making Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Micro:bit Hot Potato', 'For this lesson, students will recreate the game Hot Potato using their Micro:bit. For this game, students will start a timer with a random countdown and when the timer goes off, the game is over and whoever is still holding the potato has lost.', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bit Digital Making Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Micro:bit Compass', 'A compass is an instrument with a magnetic pointer which shows the direction of the magnetic north and the bearings from it. The Micro:bit comes with a magnetometer that can be used to detect magnetic north in much the same way as a compass.', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bit Digital Making Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Micro:bit Guita', 'Guitars are musical instruments that typically has six strings with history dating back to 1200s in Spain. Modern electric guitars were introduced in the 1930s and use electronic pickups and loudspeakers to amplify its sound during performances.', `id`
FROM curriculum_entries
WHERE `description` = 'Micro:bit Digital Making Intermediate';

-- AR/VR

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Interesting Zoo', 'Our AR/VR course keeps student abreast of the latest technology in STEM learning. Students will have knowledges about the basic applications of AR/VR technologies. They will engage in creations of three-dimensional scenes and even games with Cospaces, in which they can develop their spatial sense and design thinking skills.', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Comfort Home', 'Students consolidate their knowledge in CoSpaces through creating a special meal in special occasion of their own choice. They will be thrilled to have their own AR project and share with their family members or friends.', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Creative Story Remix', 'This course is suitable for kids who have experinece in CoSpaces and would like to challenge themselves. They need to rewrite novel stories with creativity and illustrate in an immersive VR environment.', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Interactive Art (Starry Night)', 'For these course students should have previous experience in working in CoSpaces. They are going to further explore on the potentials of VR/AR in different areas.', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Interactive Art (Pumpkin)', 'Introduce some background about Japanese Artist Yayoi Kusama. Her iconic dots in every pieces of her works. She said that " Keep creating artworks make me happy"', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'C-Interactive Art (Self Portrait)', 'This is an extension for those who challenge themselves and have great interest in exploring the possibilities of using VR integrating into artworks. They need to review some self portraits of great artists and create a VR of their own.', `id`
FROM curriculum_entries
WHERE `description` = 'AR/VR Intermediate';

-- Smart City
INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Jr Mechanical Toys', 'Young engineers start with a basic understanding of energy, force and materials. They are going to make and explore different toys using daily materials available and have fun to play with. All gadgets can be brought home for further investigation and share with family members.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Elementary';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Jr Marine Adventure', 'What causes things to sink or float? How scientists help to explore the ocean below? In this course, students experiment with different attributes related to water, its buoyancy, pressure. How to navigate above water and help lifes below water.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Elementary';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-Environmental Pioneer', 'Environmental protection is vital for the future of our kids. Students will learn principles behind on how to harness renewable energy, the importance and how nature makes clear water for us. They design and test solution on how to remedy in case of human fault and contaminate land and sea.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Introductory';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B- Aerospace Journey', 'Humans dream to fly in air. In this journey, students apply Engineering Design Process to design, create, test and refine a variety of flying machines. Not only to fly against gravity but also think of ways to land safely to complete their dream of space journey.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Introductory';


INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'A-HK Build Up', 'How is a city built nowadays? What will a city be like in the future? To learn more about these questions and civil engineering, students can take up our HK Build-Up course. In this course, students will build various structures using everyday items and learn how engineers solve issues using the Engineering Design Process.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Intermediate';

INSERT INTO curriculum_course_information_entries (`image_src`, `title`, `content`, `entry_id`)
SELECT '', 'B-Chemical Exploration', 'In the Chemical Exploration course, students will use the Engineering Design Process to design, create, test, and refine various mixtures and solutions with different chemical properties. They develop solutions to clean up an oil spill, synthesize their own rocket fuel, and investigate the secrets behind color pigmentation.', `id`
FROM curriculum_entries
WHERE `description` = 'Smart City Intermediate';