# golang
## [1][This is a JavaScript bundler and minifier. It packages up JavaScript and TypeScript code for distribution on the web.](https://www.reddit.com/r/golang/comments/iw9kfv/this_is_a_javascript_bundler_and_minifier_it/)
- url: https://github.com/evanw/esbuild
---

## [2][Typesafe database access for Go](https://www.reddit.com/r/golang/comments/iw02hp/typesafe_database_access_for_go/)
- url: https://github.com/prisma/prisma-client-go
---

## [3][YoMo is an open-source project for building your own IoT edge computing applications. With YoMo, you can speed up the development of microservices-based applications, and your industrial IoT platform will take full advantage of the low latency and high bandwidth brought by 5G.](https://www.reddit.com/r/golang/comments/iw1b3z/yomo_is_an_opensource_project_for_building_your/)
- url: https://github.com/yomorun/yomo
---

## [4][A load testing tool with a real-time analyzer, written in Go](https://www.reddit.com/r/golang/comments/iwdx5v/a_load_testing_tool_with_a_realtime_analyzer/)
- url: https://i.redd.it/bf5f5nb42bo51.gif
---

## [5][When should you use(or not use) a blank identifier in Go?](https://www.reddit.com/r/golang/comments/ivuzfr/when_should_you_useor_not_use_a_blank_identifier/)
- url: https://www.youtube.com/watch?v=FKaMiSGmQoE
---

## [6][Go implementation of a WebRTC Selective Forwarding Unit](https://www.reddit.com/r/golang/comments/iw5pk0/go_implementation_of_a_webrtc_selective/)
- url: https://github.com/pion/ion-sfu
---

## [7][[Student Research Project] Collaborated Code Reviews using Customised Static Code Analysis (Go, C#, PHP, Python)](https://www.reddit.com/r/golang/comments/iw58wn/student_research_project_collaborated_code/)
- url: https://www.reddit.com/r/golang/comments/iw58wn/student_research_project_collaborated_code/
---
This year I have developed a growing interest in coding and software development. When it came time to choose a topic for my final research project at my university, I decided to conduct research about a local and interesting startup in New Zealand.

They collaboratively automize code reviews by enabling teams to discuss and code their own rules using a sophisticated engine for static code analysis.

Alongside interviews, I have also compiled a survey to explore how such prototype creates value for software development teams. I would like to invite you to participate in my research. Your professional input will be highly valuable for the research report and further development of this prototype.

If you are interested and want to know more, it takes ten minutes to complete my survey which provides further details. Your support would be immensely appreciated. [You can find the survey here.](https://otagopolytechnic.au1.qualtrics.com/jfe/form/SV_b15riYH1K3pDkb3)

Happy coding!
## [8][Cool Stuff with Go's AST Package Pt 2: Comments and Token Positions](https://www.reddit.com/r/golang/comments/iw7zrk/cool_stuff_with_gos_ast_package_pt_2_comments_and/)
- url: https://medium.com/@riptidedata/cool-stuff-with-gos-ast-package-pt-2-e4d39ab7e9db
---

## [9][Which of OpenTracing, OpenCensus and OpenTelemetry is better and why?](https://www.reddit.com/r/golang/comments/ivzzks/which_of_opentracing_opencensus_and_opentelemetry/)
- url: https://www.reddit.com/r/golang/comments/ivzzks/which_of_opentracing_opencensus_and_opentelemetry/
---
They all looks similar and all of them are big and complicated, so it's hard to compare and choose.

Is anyone using any of this instrumentations tools or any other?

Why did you choose it?  
What are pros and cons?
## [10][Have issue with Golang structure](https://www.reddit.com/r/golang/comments/iw3bux/have_issue_with_golang_structure/)
- url: https://www.reddit.com/r/golang/comments/iw3bux/have_issue_with_golang_structure/
---
I created a Database on mysql and create an Sql file to add my table to my database, i am trying to Build a CRUD app, so in my struct, the the naming is different from what is in my struct, 

// Golang struct

type Client struct {  
gorm.Model  
clientID    uint \`gorm:"primary\_key"\`  
ActiveAt    time.Time  
CreatedAt   time.Time  
UpdatedAt   time.Time  
Client\_name string  
}

&amp;#x200B;

//Mysql table

CREATE TABLE IF NOT EXISTS \`client\` (  
 \`client\_id\` BIGINT UNSIGNED AUTO\_INCREMENT,  
 \`created\_at\` DATETIME NOT NULL,  
 \`updated\_at\` DATETIME DEFAULT CURRENT\_TIMESTAMP ON UPDATE CURRENT\_TIMESTAMP,  
 \`active\` BOOLEAN,  
 \`client\_name\` VARCHAR(255) NOT NULL,  
 PRIMARY KEY (\`client\_id\`)  
)

could this be the issue why my code is not working?
