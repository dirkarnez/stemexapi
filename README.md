stemexapi
=========
### Endpoints
- original
  - https://ec2-43-198-50-165.ap-east-1.compute.amazonaws.com/
  - https://ec2-43-198-50-165.ap-east-1.compute.amazonaws.com:3000
- new
  - https://ec2-43-198-151-195.ap-east-1.compute.amazonaws.com/api/

### Service
- https://raw.githubusercontent.com/containerd/containerd/main/containerd.service
  ```bash
  systemctl daemon-reload
  systemctl enable --now containerd
  ```
### Prerequisites
- Go
  - [All releases - The Go Programming Language](https://go.dev/dl/)
- Code editor (preferably)
  - [Visual Studio Code - Code Editing. Redefined](https://code.visualstudio.com/)
    - Extensions
      - [Go - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- Git
  - [Git - Downloads](https://git-scm.com/downloads)

### Testing tool
- API test
  - [Postman API Platform | Sign Up for Free](https://www.postman.com/)


### API test cases
```
curl \
-X POST https://api.hubapi.com/crm/v3/objects/deals/search \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d" \
-d '{"filterGroups":[{"filters":[{"propertyName":"student_id","operator":"EQ","value":"20220014.stemex"}]}],"properties":["dealname","student_id","new_course_name","course_dates","zoom_link"],"sorts":[{"propertyName":"createdate","direction":"DESCENDING"}]}'
```

```
curl \
-X GET https://api.hubapi.com/deals/v1/deal/13953328933 \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d"
```

```
curl \
-X GET https://api.hubapi.com/crm/v4/objects/deal/8914581675/associations/note \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d"

curl \
-X GET https://api.hubapi.com/crm/v3/objects/notes/30398162463?properties=hs_attachment_ids \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d"

curl \
-X GET https://api.hubapi.com/files/v3/files/99697390981 \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d"
```

### AWS db
- setup
  - **https://webdock.io/en/docs/how-guides/database-guides/how-enable-remote-access-your-mariadbmysql-database**
  - **https://cloudinfrastructureservices.co.uk/how-to-install-mariadb-on-ubuntu-22-04/**
  - https://mariadb.com/kb/en/configuring-mariadb-for-remote-client-access/



### Local dev-
- https://localhost/api/login
- ```
  {
    "user_name": "joe",
    "password": "stemex"
  }
  ```

### TODOs
- [ ] do not create parent table, but create a mapping table for hubsapot key
