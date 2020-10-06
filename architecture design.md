# architecture design

教师机与学生机之间的通信采用C/S、B/S并用方式

在教师机上运行一个Http Server，用以简化教师机与学生机之间的通信过程。

大部分情况下使用Http协议进行沟通。在需要长连接时采用WebSocket协议。

教师机Http Server运行端口号及WebSocket沟通端口号均为`1472`



## 接口及过程设计

### 层次一

* 时序图

```mermaid
sequenceDiagram

    participant student
    participant teacher
    participant DataBase
	

    student ->> teacher: Http POST: Login
  
  teacher ->> DataBase: check password
    DataBase ->> teacher: isPasswordRight
    
    
    rect rgb(175, 255, 212)
     
    opt Wrong Password
  	teacher -->> student: return Author Failure
    student -->> teacher: Http POST: Login
    teacher -->> DataBase: check password
    DataBase -->> teacher: isPasswordRight
    end
    
    end
    
    Note left of DataBase: PasswordRight
    
    teacher ->> DataBase: setSignInDate
    
    teacher ->> student: return Author Success
    
    teacher -> student: start a Websocket
    
  
    activate student
    activate teacher
    
    rect rgb(255,255,0)
     student --> teacher: some infomation exchange
     end
     
    student -> teacher: close Websocket
    deactivate teacher
    deactivate student

    
```


* 学生：登录

  `POST			/student/{studentId}`

  ```json
  {
      "student_password": "(string)",
  }
  ```

  return module:

  string: Authorization Accept / Authorization Failure


  if Accept:	start a webSocket `GET		/student/{studentId}/keepAlive`  and set `SignInDate`

* 学生：登出

  断开 webSocket and set `SignOutDate`





* 教师：导入信息

  `func ImportClassInfo()   { }`

  导入结构：

  ```go
  type ClassInfo struct{
      // as the primery key
      ClassName string
      ClassStartDate int64
      ClassOverDate int64
      Students []StudentInfo
  }
  
  type StudentInfo struct{
      StudentId string
      StudentName string
      StudentPassword string
  }
  ```

  转化为数据库中的结构：

  ```go
  type StudentStatus{
      className string
      ClassStartDate int64
      ClassOverDate int64
      StudentId string
      StudentName string
      StudentPassword string
      
      // as unixtimestamp
      SignInDate int64
      SignOutDate int64
      
      // constant package define SignStatus associated with numbers, and it must be caculate by Sign in/out time
      // four statuses are 正常/迟到/早退/缺勤
      SignStatus int
      
      
      // homework status
      
  }
  
  type Homework struct{
      HomeworkTitle string
      // define HomeworkType associated with numbers
      HomeworkType int
      HomeworkAnswer string
      
  }
  ```

* 更改签到状态

  `func  ChangeStudentSignStatus(studentId, className, classStartDate, targetSignStatus) (bool, error) {}`

---------

### 层次二



```mermaid
sequenceDiagram

participant student
participant teacher

rect rgb(12,255,255)
student -> teacher: EXIST Websocket 
end

student ->> teacher: Http POST
note left of student: POST body contains these infos: <br> · StudentInfo <br> · ClassInfo<br>· SeatInfo <br>· Question
teacher ->> student: Http return OK(200)
teacher -->> student: Websocket writeMsg: <br> Questions' Answer
```

`POST		/student/{studentId}`



### 层次三

New a tables named "`questionInfo`" containing the whole questions infos

```mermaid
sequenceDiagram

participant student
participant teacher
participant database

rect rgb(75,255,255)
student -> teacher: EXIST Websocket 
end

teacher ->> teacher: publish questions
teacher ->> database: Save homeworkInfo

note right of database: Questions should <br> contains a DDL

rect rgb(101, 102, 485)
teacher ->> student: Websocket <br> Broadcase Questions
end


note left of student: answer the questions

student ->> teacher: Http POST 
note left of student: POST body contains these infos: <br> ·HomeworkInfo <br>·StudentInfo <br> ·ClassInfo <br> all of this infos is in an Array

teacher ->> teacher: Auto Correct homework

teacher ->> database: Save homeworkStatus
teacher ->> student: Http return OK(200)

opt text/file questions
	teacher -->> teacher: corrent it at anytime
end
```

When a student want to upload a file, the requst API is different with the normal homework upload API. 
And also, he can just upload a file once. Therefore, when many files shuold be uploaded, he should request the "file upload API" for many times which contains the params named "studentId", "className", "classStartDate", "questionTitle" and   "file"

