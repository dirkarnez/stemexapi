[stemexapi](https://ec2-43-198-151-195.ap-east-1.compute.amazonaws.com/)
========================================================================
### Endpoints
- original
  - https://ec2-43-198-50-165.ap-east-1.compute.amazonaws.com/
  - https://ec2-43-198-50-165.ap-east-1.compute.amazonaws.com:3000
- new
  - https://ec2-43-198-151-195.ap-east-1.compute.amazonaws.com/api/

### `main.go`
- `go run cmd/gormplayground/main.go`
- `go run cmd/codegen/main.go`

### TODOs
```
curr
- add genesis

courses
- id
- cirr id
- normal price courses_prices

courses_prices
- effective start 
- effective end

course_lessons
- id
- courses-id
- lesson date
- lesson instructr


last minute offer
- courses
- offer price courses_prices

- 100vh
- parent
- upload component

```
### Mail
- [app password](https://myaccount.google.com/apppasswords): `cyqb rllp qhep glnx`

### Google Drive
- https://console.cloud.google.com/apis/enableflow?apiid=drive.googleapis.com&pli=1&project=stemex-academy&supportedpurview=project
- https://console.cloud.google.com/apis/library/drive.googleapis.com?project=stemex-academy&supportedpurview=project
- https://developers.google.com/drive/api/reference/rest/v3/files/list
- https://pkg.go.dev/google.golang.org/api/drive/v3

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
- [ ] https://github.com/gin-gonic/gin/blob/master/binding/multipart_form_mapping_test.go
- [ ] do not create parent table, but create a mapping table for hubsapot key
- [ ] [dirkarnez/queueserver](https://github.com/dirkarnez/queueserver)
- [ ] oauth?
- [ ] no server side redirects
- [ ] [dirkarnez/go-json-customized-marshall](https://github.com/dirkarnez/go-json-customized-marshall)
- [ ] [dirkarnez/stemexfaker](https://github.com/dirkarnez/stemexfaker)
- [ ] Build variant
- [ ] [dirkarnez/stemexdocs](https://github.com/dirkarnez/stemexdocs)
- [ ] test cases (instead of postman)
- [ ] [iris/_examples/request-ratelimit/rate-middleware/main.go at main · kataras/iris](https://github.com/kataras/iris/blob/main/_examples/request-ratelimit/rate-middleware/main.go)
- [ ] [iris/_examples/auth/recaptcha at main · kataras/iris · GitHub](https://github.com/kataras/iris/tree/main/_examples/auth/recaptcha)
- [ ] [iris/_examples/websocket/socketio/main.go at main · kataras/iris · GitHub](https://github.com/kataras/iris/blob/main/_examples/websocket/socketio/main.go)
- [ ] gorm wrapper
  - ```go
    import (
        "reflect"
        "github.com/jinzhu/gorm"
    )
    
    type User struct {
        gorm.Model
        Name  string `gorm:"column:user_name"`
        Email string `gorm:"column:user_email"`
    }
    
    func main() {
        user := User{}
        reflectType := reflect.TypeOf(user)
    
        for i := 0; i < reflectType.NumField(); i++ {
            field := reflectType.Field(i)
            gormTag := field.Tag.Get("gorm")
    
            if gormTag != "" {
                // Do something with the GORM tag
                // In this example, we'll just print it
                fmt.Printf("Field: %s, GORM Tag: %s\n", field.Name, gormTag)
            }
        }
    }
    ```
  - mapper
    - ```go
      package main

      import "fmt"
      
      type VertexA struct {
      	X int
      	Y int
      }
      
      type VertexB struct {
      	X int
      	Y int
      }
      
      func mapper(a *VertexA, b *VertexB) {
      	b.Y = a.Y
      	return
      }
      
      func main() {
      	a := VertexA{1, 2}
      	b := VertexB{}
      	mapper(&a, &b)
      
      	fmt.Println(b.Y)
      }

      ```
  - status code
    ```go
    	ctx.StopWithStatus(iris.StatusForbidden)
    ``` 
### Notes
- `	Map(user, func(item User) UUIDEx {return item.ID})`
- `Curriculm plan`
- public files
- lesson sorting
- lesson deleting bugs
