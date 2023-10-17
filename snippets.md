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

```sql
INSERT INTO testing.files (id,created_at,updated_at,deleted_at,physical_file_name) VALUES
	 (0x0AF749296CF211EE9AA006C3BC34E27E,NULL,NULL,NULL,'upcoming-schedule/Coding_Python/Level 3-Introductory B-min.png'),
```