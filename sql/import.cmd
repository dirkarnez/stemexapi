set PATH=%USERPROFILE%\Downloads\mariadb-11.1.2-winx64\mariadb-11.1.2-winx64\bin
mysql --host=ec2-43-198-151-195.ap-east-1.compute.amazonaws.com --port=3306 --user=webadmin --password=password --database=testing < data-1.sql
mysql --host=ec2-43-198-151-195.ap-east-1.compute.amazonaws.com --port=3306 --user=webadmin --password=password --database=testing < data-2.sql
mysql --host=ec2-43-198-151-195.ap-east-1.compute.amazonaws.com --port=3306 --user=webadmin --password=password --database=testing < data-3.sql
mysql --host=ec2-43-198-151-195.ap-east-1.compute.amazonaws.com --port=3306 --user=webadmin --password=password --database=testing < data-4.sql

pause
