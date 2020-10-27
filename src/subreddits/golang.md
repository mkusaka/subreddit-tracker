# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][TIL Go support continue statements with labels](https://www.reddit.com/r/golang/comments/jiveet/til_go_support_continue_statements_with_labels/)
- url: https://relistan.com/continue-statement-with-labels-in-go
---

## [3][From HTTP to HTTPS with Go (with mTLS)](https://www.reddit.com/r/golang/comments/jiy3eq/from_http_to_https_with_go_with_mtls/)
- url: https://www.prakharsrivastav.com/posts/from-http-to-https-using-go/
---

## [4][Creating your own GitHub Action using Go](https://www.reddit.com/r/golang/comments/jijwqb/creating_your_own_github_action_using_go/)
- url: https://www.youtube.com/watch?v=8IgNY8QT3vk
---

## [5][Tiny progress bar package](https://www.reddit.com/r/golang/comments/jiz7s4/tiny_progress_bar_package/)
- url: https://github.com/vardius/progress-go
---

## [6][awsinventory 0.4.0: Big batch of new AWS services!](https://www.reddit.com/r/golang/comments/jizfbj/awsinventory_040_big_batch_of_new_aws_services/)
- url: https://www.reddit.com/r/golang/comments/jizfbj/awsinventory_040_big_batch_of_new_aws_services/
---
[https://github.com/manywho/awsinventory/releases/tag/0.4.0](https://github.com/manywho/awsinventory/releases/tag/0.4.0)

AWS Inventory is a command line tool written in Go to fetch data from  AWS and use it to generate a FedRAMP compliant inventory of your assets

This release adds a load of new AWS services as well as other fixes and improvements.

The following new services can now be included:

* CloudFront
* CodeCommit
* DynamoDB
* ECR
* ECS
* ELBv2 (ALB's &amp; NLB's)
* ElastiCache
* Elasticsearch
* KMS
* Lambda
* RDS
* SQS
## [7][What do you think about : Qmgo - The MongoDB driver for Go .](https://www.reddit.com/r/golang/comments/jivcm9/what_do_you_think_about_qmgo_the_mongodb_driver/)
- url: https://www.reddit.com/r/golang/comments/jivcm9/what_do_you_think_about_qmgo_the_mongodb_driver/
---
Qmgo - The MongoDB driver for Go . It‘s based on official mongo-go-driver but easier to use like Mgo.  [https://github.com/qiniu/qmgo](https://github.com/qiniu/qmgo)

* Qmgo allow user to use the new features of MongoDBin a more elegant way.
* Qmgo is the first choice for migrating from mgoto the new MongoDB driverwith minimal code changes.

Current supported features:

* CRUD to documents
* Sort、limit、count、select、distinct
* Transactions
* Hooks
* Automatically update default and custom fields
* Predefine operator keys
## [8][Do more with GitLab from your CLI with the newly released version of glab (v1.11.0). You can now view merge request diffs, view issue boards, and do more with glab from your CLI](https://www.reddit.com/r/golang/comments/jimf12/do_more_with_gitlab_from_your_cli_with_the_newly/)
- url: https://github.com/profclems/glab/releases/latest
---

## [9][distribyted/distribyted: Torrent client with on-demand file downloading as a filesystem in Go](https://www.reddit.com/r/golang/comments/ji9zcb/distribyteddistribyted_torrent_client_with/)
- url: https://github.com/distribyted/distribyted
---

## [10][kboard: a terminal game to practice keyboard typing](https://www.reddit.com/r/golang/comments/jid5l4/kboard_a_terminal_game_to_practice_keyboard_typing/)
- url: https://github.com/CamiloGarciaLaRotta/kboard
---

## [11][Why is my Middleware being called multiple times?](https://www.reddit.com/r/golang/comments/jitcb2/why_is_my_middleware_being_called_multiple_times/)
- url: https://www.reddit.com/r/golang/comments/jitcb2/why_is_my_middleware_being_called_multiple_times/
---
I am writing a middleware for graphql resolvers and I am using the Chi router. My middleware is running into an issue where it keeps getting called after a request has been made to a resolver. I would expect it to be called only once each time a resolver function is called. What is causing it to keep calling itself indefinitely? 

`type authResponseWriter struct {`  
`http.ResponseWriter`  
`}`  
`func (w *authResponseWriter) Write(b []byte) (int, error) {`  
 `// some logic to generate token`  
`http.SetCookie(w, &amp;http.Cookie{`  
`Name:     "token",`  
`Value:    string(token),`  
`HttpOnly: true,`  
`Path:     "/",`  
`SameSite: http.SameSiteStrictMode,`  
`})`  
 `return w.ResponseWriter.Write(b)`  
`}`  
`// Middleware returns the handler for authenticating the user`  
`func Middleware(next http.Handler) func(http.Handler) http.Handler {`  
 `return func(next http.Handler) http.Handler {`  
 `return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {`  
 `authResponseWriter := &amp;authResponseWriter{w}`  
`next.ServeHTTP(authResponseWriter, r)`  
`})`  
`}`  
`}`
