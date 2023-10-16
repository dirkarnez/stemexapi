```sql
select COUNT(ua.id), u.user_name, CAST(ua.created_at AS DATE) as d  from users u
left join roles r  
on u.role_id = r.id
left join user_activities ua
on ua.user_id = u.id
GROUP BY d, u.user_name
```

total
```sql
select COUNT(ua.id), CAST(ua.created_at AS DATE) as d  from user_activities ua
left join users u
on ua.user_id = u.id
left join roles r  
on u.role_id = r.id
where r.name = 'parent' OR r.name is NULL
GROUP BY d 
```

### Tree
```Go
	// return func(ctx iris.Context) {
	// 	var curriculumEntryList []model.CurriculumEntry
	// 	if err := dbInstance.Raw(`
	// 		WITH RECURSIVE curriculum_entries_nested AS (
	// 			SELECT id, description, parent_id FROM curriculum_entries WHERE description = 'Micro:bits'
	// 			UNION
	// 			SELECT curriculum_entries.id, curriculum_entries.description, curriculum_entries.parent_id FROM curriculum_entries JOIN curriculum_entries_nested ON curriculum_entries_nested.id = curriculum_entries.parent_id
	// 		)
	// 		SELECT * FROM curriculum_entries_nested
	// 	`).
	// 		Scan(&curriculumEntryList).Error; err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		return
	// 	} else {
	// 		ctx.JSON(curriculumEntryList)
	// 	}
	// }
```



```js

	// import codingMinecraftElementary from "../assets/images/upcoming-schedule/codingMinecraft/Level 2-Elementary-min.png";
	// import codingMinecraftIntroductory from "../assets/images/upcoming-schedule/codingMinecraft/Level 2-Introductory-min.png";
	// import codingMinecraftIntermediate from "../assets/images/upcoming-schedule/codingMinecraft/Level 2-Intermediate-min.png";
	// import codingMinecraftAdvanced from "../assets/images/upcoming-schedule/codingMinecraft/Level 2-Advanced-min.png";
	// import codingMinecraftMaster from "../assets/images/upcoming-schedule/codingMinecraft/Level 2-Master-min.png";

	// import codingRobloxIntroductory from "../assets/images/upcoming-schedule/codingRoblox/Level 2- Introductory-min.png";
	// import codingRobloxIntermediate from "../assets/images/upcoming-schedule/codingRoblox/Level 2-Intermediate-min.png";

	// import Coding_PythonAdvanced from "../assets/images/upcoming-schedule/Coding_Python/Level 2-Advanced-min.png";
	// import Coding_PythonIntermediate from "../assets/images/upcoming-schedule/Coding_Python/Level 2-Intermediate-min.png";
	// import Coding_PythonIntroductory from "../assets/images/upcoming-schedule/Coding_Python/Level 2-Introductory-min.png";
	// import Coding_PythonMaster from "../assets/images/upcoming-schedule/Coding_Python/Level 2-Master-min.png";


	// import Coding_ScratchIntermediate from "../assets/images/upcoming-schedule/Coding_Scratch/Level 2-Intermediate-min.png";
	// import Coding_ScratchIntroductory from "../assets/images/upcoming-schedule/Coding_Scratch/Level 2-Introductory-min.png";
	// import Coding_ScratchJr from "../assets/images/upcoming-schedule/Coding_Scratch/Level 2-Scratch Jr-min.png";

	// import cyberVirtualRoboticsElementary from "../assets/images/upcoming-schedule/cyberVirtualRobotics/Level 2-Elementary.png";
	// import cyberVirtualRoboticsIntermediate from "../assets/images/upcoming-schedule/cyberVirtualRobotics/Level 2-Intermediate.png";
	// import cyberVirtualRoboticsIntroductory from "../assets/images/upcoming-schedule/cyberVirtualRobotics/Level 2-Introductory.png";
	// import cyberVirtualRoboticsMaster from "../assets/images/upcoming-schedule/cyberVirtualRobotics/Level 2-Master.png";

	// import LEGO_RoboticsAdvanced from "../assets/images/upcoming-schedule/LEGO_Robotics/Level 2-Advanced-min.png";
	// import LEGO_RoboticsElementary from "../assets/images/upcoming-schedule/LEGO_Robotics/Level 2-Elementary-min.png";
	// import LEGO_RoboticsIntermediate from "../assets/images/upcoming-schedule/LEGO_Robotics/Level 2-Intermediate-min.png";
	// import LEGO_RoboticsIntroductory from "../assets/images/upcoming-schedule/LEGO_Robotics/Level 2-Introductory-min.png";

	// import Vex_RoboticsIntroductory from "../assets/images/upcoming-schedule/Vex Robotics/Level 1-min.png";

	// import AppInventorAdvanced from "../assets/images/upcoming-schedule/AppInventor Mobile Apps/Level 2-Advanced-min.png";
	// import AppInventorIntermediate from "../assets/images/upcoming-schedule/AppInventor Mobile Apps/Level 2-Intermediate-min.png";
	// import AppInventorIntroductory from "../assets/images/upcoming-schedule/AppInventor Mobile Apps/Level 2-Introductory-min.png";

	// import AIMLAdvanced from "../assets/images/upcoming-schedule/A.I. _ Machine Learning/Level 2-Advanced-min.png";
	// import AIMLIntermediate from "../assets/images/upcoming-schedule/A.I. _ Machine Learning/Level 2-Intermediate-min.png";
	// import AIMLIntroductory from "../assets/images/upcoming-schedule/A.I. _ Machine Learning/Level 2-Introductory-min.png";

	// import ThreeD_Design_PrintingIntermediate from "../assets/images/upcoming-schedule/3D_Design_Printing/Level 2-Intermediate-min.jpg";
	// import ThreeD_Design_PrintingIntroductory from "../assets/images/upcoming-schedule/3D_Design_Printing/Level 2-Introductory-min.jpg";

	// import Micro_bitIntermediate from "../assets/images/upcoming-schedule/Micro_bit/Level 2-Intermediate-min.png";
	// import Micro_bitIntroductory from "../assets/images/upcoming-schedule/Micro_bit/Level 2-Introductory-min.png";

	// import AR_VRIntermediate from "../assets/images/upcoming-schedule/AR_VR/Level 2-Intermediate-min.png";
	// import AR_VRIntroductory from "../assets/images/upcoming-schedule/AR_VR/Level 2-Introductory-min.png";

	// import SmartCityElementary from "../assets/images/upcoming-schedule/Smart City/Level 2-Elementary-min.png";
	// import SmartCityIntermediate from "../assets/images/upcoming-schedule/Smart City/Level 2-Intermediate-min.png";
	// import SmartCityIntroductory from "../assets/images/upcoming-schedule/Smart City/Level 2-Introductory-min.png";

```