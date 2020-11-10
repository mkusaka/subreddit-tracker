# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][How Go helped save HealthCare.gov featuring Paul Smith, CTO of Ad Hoc (Go Time #154)](https://www.reddit.com/r/golang/comments/jr7fp0/how_go_helped_save_healthcaregov_featuring_paul/)
- url: https://changelog.com/gotime/154
---

## [3][I have rewritten "Spinning Donut" to Golang](https://www.reddit.com/r/golang/comments/jqxm6d/i_have_rewritten_spinning_donut_to_golang/)
- url: https://www.reddit.com/r/golang/comments/jqxm6d/i_have_rewritten_spinning_donut_to_golang/
---
Implementation of "Spinning Donut" on Golang, not much different but cool

Left: Go  
Right: C  
Source: [https://github.com/Vlad-Shevliakov/ASCII-render](https://github.com/Vlad-Shevliakov/ASCII-render)

https://reddit.com/link/jqxm6d/video/9m2dr0o5v7y51/player
## [4][What is your success story converting your team/mates to code with Golang?](https://www.reddit.com/r/golang/comments/jrg9sa/what_is_your_success_story_converting_your/)
- url: https://www.reddit.com/r/golang/comments/jrg9sa/what_is_your_success_story_converting_your/
---
Hi Gophers, I inspired by the post [here](https://www.reddit.com/r/golang/comments/jnhzm5/today_i_presented_go_to_my_team/) to share also my experience in evangelizing Go at my workplace.

I have created a narration of Go benefits, I POC-ing some compatibility with internal tech stack. I have created internal tools to do automation, I create an internal community to share Go news and meetup announcement. However, those efforts don't seem to be fruitful. I am in a level of Manager right now, but basically, it doesn't provide any influence to the adoption. Lol.

So, my fellow gophers in Reddit. What is your success story in influencing your team for the Go adoption? It could provide learnings and reflection to me so I could hustle in evangelizing Go more.

Thanks.
## [5][Happy Birthday Go!](https://www.reddit.com/r/golang/comments/jrkn2m/happy_birthday_go/)
- url: https://www.reddit.com/r/golang/comments/jrkn2m/happy_birthday_go/
---
Was going through tour of golang  [here](https://tour.golang.org/welcome/4) and found that it's Go's birthday today!

Long live Go!
## [6][Can Java microservices be as fast as Go?](https://www.reddit.com/r/golang/comments/jrhkkl/can_java_microservices_be_as_fast_as_go/)
- url: https://www.reddit.com/r/golang/comments/jrhkkl/can_java_microservices_be_as_fast_as_go/
---
https://medium.com/helidon/can-java-microservices-be-as-fast-as-go-5ceb9a45d673

If I'm gonna leave my opinion on the article, take it with a pinch of salt because I dont know sometimes what I'm talking about, just a friendly opinion.

To me at least, it seems like even though sometimes Java outperformed in response times for a small 5-10% margin, it's still not worth the cost of Java's x2-10 times more memory consumption which costs very much in the hosting market. Like, most of the tests have as big margins as Go=100mb Java=1.5GB memory, that's a huge difference. The cost is too big for the advantage, which is not even certain, in contrast with RAM cost which is very expensive.

Other than that, it just seems to me that Golang's optimization is still maturing and has a long way to go. But good job at Java native images for succeeding outperforming it. It was honestly surprising to me.
What's your thoughts?
## [7][What Makes Go So Different? | Better Programming on Medium](https://www.reddit.com/r/golang/comments/jrdher/what_makes_go_so_different_better_programming_on/)
- url: https://medium.com/better-programming/what-makes-go-so-different-eb0648498ce0
---

## [8][Commercial applications with golang](https://www.reddit.com/r/golang/comments/jrj4zt/commercial_applications_with_golang/)
- url: https://www.reddit.com/r/golang/comments/jrj4zt/commercial_applications_with_golang/
---
What has been some of the compliance challenges of using golang for developing commercial applications.

I am taking steps in this direction but I have to get the licenses reviewed by our lawyers. I am concerned about the modules I will be using and their licenses. 

Community's insights and advice welcome.

Thanks, srini
## [9][Makeless - SaaS Framework just released](https://www.reddit.com/r/golang/comments/jr00ox/makeless_saas_framework_just_released/)
- url: https://www.reddit.com/r/golang/comments/jr00ox/makeless_saas_framework_just_released/
---
Hey guys,

today we are released **the beta** of **Makeless - a SaaS Framework**:

It's split up into a modular Golang backend and a Typescript + Vue.js frontend where everything is ready to go.

https://preview.redd.it/wt1dn7urj8y51.png?width=2374&amp;format=png&amp;auto=webp&amp;s=66a8a290e62be6a23cdd47815c2b54585b2f72a0

**Features:**

* Based on Golang ([gin](https://github.com/gin-gonic/gin) &amp; [gorm](https://github.com/go-gorm/gorm))
* Concurrency safe &amp; scalable
* Super clean and small
* Fully customizable and configurable
* Multilingual
* State of the art Authentication with JWT HttpOnly Cookies
* User management
* Team management
* Token management for users and teams
* Realtime events
* Subscriptions and Per-Seat Payments out of the box (coming soon)

**Setup:** 

* Backend: [https://github.com/makeless/makeless-demo/blob/master/backend/main.go](https://github.com/makeless/makeless-demo/blob/master/backend/main.go)
* Frontend: [https://github.com/makeless/makeless-demo/blob/master/frontend/src/main.ts](https://github.com/makeless/makeless-demo/blob/master/frontend/src/main.ts)
* Config: [https://github.com/makeless/makeless-demo/blob/master/makeless.json](https://github.com/makeless/makeless-demo/blob/master/makeless.json)

**Links:** 

* Go Backend: [https://github.com/makeless/makeless-go](https://github.com/makeless/makeless-go)
* Typescript + Vue.js Frontend: [https://github.com/makeless/makeless-ui](https://github.com/makeless/makeless-ui)
* Production ready Docker Demo: [https://github.com/makeless/makeless-demo](https://github.com/makeless/makeless-demo)
* Scalable redis event package: [https://github.com/makeless/makeless-go-event-redis](https://github.com/makeless/makeless-go-event-redis)
* Scalable mailgun mailer package: [https://github.com/makeless/makeless-go-mailer-mailgun](https://github.com/makeless/makeless-go-mailer-mailgun)
* In-Memory Authenticator: [https://github.com/makeless/makeless-go-authenticator-in-memory](https://github.com/makeless/makeless-go-authenticator-in-memory) 

**Next steps:**  


* Fair and usable License
* Documentation

**Would love to just get some feedback**

Have a good day guys and stay safe!
## [10][TinyGo, what’s your experience with it?](https://www.reddit.com/r/golang/comments/jr40ej/tinygo_whats_your_experience_with_it/)
- url: https://www.reddit.com/r/golang/comments/jr40ej/tinygo_whats_your_experience_with_it/
---
I’m interested in learning more about automation  digitally and physically. 

Recently I’ve been researching arduinos and esp32. I came across tinygo and see that it works with microcontrollers.

Would Go be a good language to learn for this use case. I’m interested in learning Go because I’m also into web development too. (so it might a two birds one stone situation 😅)

It’s more hobby at the moment but don’t know if c &amp; c++ would be easier but since I’m only doing it for hobby would Go be more interesting. 

What is your opinion?
## [11][Flexible MQTT-telegram-bot for IoT](https://www.reddit.com/r/golang/comments/jr6fls/flexible_mqtttelegrambot_for_iot/)
- url: https://xd-wart.medium.com/flexible-mqtt-telegram-bot-for-iot-70d567edfb2e
---

