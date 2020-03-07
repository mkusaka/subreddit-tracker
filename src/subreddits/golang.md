# golang
## [1][My experience with learning Golang](https://www.reddit.com/r/golang/comments/fesci5/my_experience_with_learning_golang/)
- url: https://pmihaylov.com/learning-golang-experience/
---

## [2][CertMagic v0.10.0 released - CertMAutomatic HTTPS for any Go program: fully-managed TLS certificate issuance and renewal.](https://www.reddit.com/r/golang/comments/fet7wz/certmagic_v0100_released_certmautomatic_https_for/)
- url: https://github.com/caddyserver/certmagic
---

## [3][How DRY?](https://www.reddit.com/r/golang/comments/femrpz/how_dry/)
- url: https://www.reddit.com/r/golang/comments/femrpz/how_dry/
---
I try to always follow this philosophy. And in Go, I tend to follow it more often than not. But I always wonder when people say "Enough. This is too much copying".

In this [gist](https://gist.github.com/justrudd/8517dfbea90afe9b4b64be11d7a58290) are two functions - populateOne and populateTwo. populateOne queries a table in chunks of 1000 pulling in an int32 and int64 column and then adding those columns to a cache. populateTwo queries a table in chunks of 1000 pulling in an int32 and string column and then adding those columns to a cache.

They are probably 90% similar. In the case of 2 functions, not a huge deal. but I'm repeating that pattern 12 times. 3 of the 12 definitely have more complicated queries and deal with null columns. But 9 of the 12 are structurally the same - query by id, retrieve 3 columns, update a map with one column mapping to the other.

I'm not asking for how you'd get rid of the duplication (unless you think you have a neat way to do it), but just curious as to how much you care. A little? A lot? Does writing the code in Go make you care less? In Ruby, I'd get rid of it. In Java/Kotlin, I'd consider removing as much of it as possible but lean toward duplication when it started looking contrived. But in Go, it doesn't really bother me...
## [4][Micro In Action, Coda: Distributed Cron Job](https://www.reddit.com/r/golang/comments/fetctd/micro_in_action_coda_distributed_cron_job/)
- url: https://itnext.io/micro-in-action-coda-distributed-cron-job-a2b577885b24
---

## [5][r8limiter - envoy pluggable rate limiting service - Feedback is needed :)](https://www.reddit.com/r/golang/comments/fet1qj/r8limiter_envoy_pluggable_rate_limiting_service/)
- url: https://github.com/samueltorres/r8limiter
---

## [6][Designed in ZPL !](https://www.reddit.com/r/golang/comments/fecwcd/designed_in_zpl/)
- url: https://i.redd.it/6ygxcb44s1l41.jpg
---

## [7][load testing with vegeta and postman](https://www.reddit.com/r/golang/comments/feqcz6/load_testing_with_vegeta_and_postman/)
- url: https://www.reddit.com/r/golang/comments/feqcz6/load_testing_with_vegeta_and_postman/
---
Hello fellow Gophers,

Been using a small library I wrote to streamline my load testing with vegeta. 

The use case really came about when running the usual postman tests, and not having a great way to utilize what was defined in my postman collection to run my load tests as well. I had already been using vegeta for load testing but for anything other than a simple API that code becomes verbose very quickly. So I really just wanted a way for vegeta to use my postman collection and environment and hammer away at the API.

It simply ingests your postman collection and environment to generate vegeta targets and then allows you to configure vegeta as normal. 

Figured I'd share as others might find it useful within their own workflows. 

Any feedback on the implementation is welcome! Hope you guys find this useful!

[https://github.com/dgparker/vegeta-powerup](https://github.com/dgparker/vegeta-powerup)
## [8][Finding a graphql client](https://www.reddit.com/r/golang/comments/feubr3/finding_a_graphql_client/)
- url: https://www.reddit.com/r/golang/comments/feubr3/finding_a_graphql_client/
---
I am trying to call graphql api by golang for few days, but most of libraries in GitHub are graphql server implementation not client.


Could anybody recommend some graphql client libraries ? 

machinebox/graphql or shurcool/graphql is the only one choices in GitHub?
## [9][Yet another go module challenge for me](https://www.reddit.com/r/golang/comments/feu2pl/yet_another_go_module_challenge_for_me/)
- url: https://www.reddit.com/r/golang/comments/feu2pl/yet_another_go_module_challenge_for_me/
---
In the app I am developing there is a library module and a couple of executables using the library. All maintained in the same git repo.

As the library (and the signature of the functions) evolves, have a hard time convincing go to build one app or the other. when i try to build one of the apps, the "latest" change to the library is not picked up.

Perhaps the library, each app should be in their own git repos.

Trying to avoid splitting the library and the execs into separate git repos.

Is there a go mod option to indicate some module to be always the latest? Or to tell go not to go to the repo for the library. after all it is local.

Hope the issue is clear. Looking for the "proper" approach. 

thanks for any advice. srini
## [10][XML parsing in Go](https://www.reddit.com/r/golang/comments/fekgx1/xml_parsing_in_go/)
- url: https://www.reddit.com/r/golang/comments/fekgx1/xml_parsing_in_go/
---
    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;service name=""&gt;
      &lt;tabset name="Syslog"&gt;
        &lt;tab name="Syslog Config"/&gt;
        &lt;tab-contents name="Syslog Configuration"&gt;
          &lt;string desc="SuperUserID" name="superID"&gt;&lt;![CDATA[admin]]&gt;&lt;/string&gt;
          &lt;password desc="SuperUserPwd" name="superPwd"&gt;&lt;![CDATA[superman321]]&gt;&lt;/password&gt;
          &lt;string desc="GuestID" name="guestID"&gt;&lt;![CDATA[guest]]&gt;&lt;/string&gt;
          &lt;string desc="GuestPermission" name="guestPerm"&gt;&lt;![CDATA[1,23,122]]&gt;&lt;/string&gt;
          &lt;password desc="GuestPwd" name="guestPwd"&gt;&lt;![CDATA[guest123]]&gt;&lt;/password&gt;
          &lt;flag desc="Enable Guest" name="enableGuest"&gt;&lt;![CDATA[false]]&gt;&lt;/flag&gt;
          &lt;flag desc="Enable Logging" name="enableLogging"&gt;&lt;![CDATA[true]]&gt;&lt;/flag&gt;
        &lt;/tab-contents&gt;
      &lt;/tabset&gt;
    &lt;/service&gt;

I really haven't worked on XML at all. I thought it'd be simple one, but seems very flexible yet confuses heck out of me when "tab-contents" takes bunch of different type of data inside.

    type Ex_XML struct {
      XMLName xml.Name `xml:"service"`
      Tabset  []tabset `xml:"tabset"`
    }
    type tabset struct {
      Name        string      `xml:"name,attr"`
      Tab         tab         `xml:"tab"`
      TabContents tabContents `xml:"tab-contents"`
    }
    type tab struct {
      Name string `xml:"name,attr"`
    }
    type tabContents struct {                  // &lt;===== I'M LOST HERE
      Name     string `xml:"name,attr"`
      String   string `xml:"string"`
      Password string `xml:"password"`
      Flag     string `xml:"flag"`
    }

How am I supposed to code tabContents here? I think attribute "name" should be used as a key, but I can't find any examples..

Thanks!
