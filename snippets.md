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

### deployment
delete all inside public
put new

sudo pkill stemexapi
delete 
put new
sudo chmod +x ./stemexapi
sudo nohup ./stemexapi &