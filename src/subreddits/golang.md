# golang
## [1][Which framework to use for a multiplayer mobile game?](https://www.reddit.com/r/golang/comments/fuqwqm/which_framework_to_use_for_a_multiplayer_mobile/)
- url: https://www.reddit.com/r/golang/comments/fuqwqm/which_framework_to_use_for_a_multiplayer_mobile/
---
Hi, I have to develop the back-end of a multiplayer mobile game. I specify that the real time part of the game will be managed with a separate service, so in Go I should just develop the services (REST API) that interface with the database (for example user registration, writing matches completed on db etc.), In particular we will use Mongodb.

So far I've simply used Gorilla/Mux and net/http to create APIs. I wanted to ask if you recommend a framework (e.g. Gin Gonic, Chi etc.) that allows me to better structure the code and allows me to develop faster than using only Gorilla/mux and net/http.

Thanks in advance.
## [2][Understanding bytes in Go by building a TCP protocol](https://www.reddit.com/r/golang/comments/fucme1/understanding_bytes_in_go_by_building_a_tcp/)
- url: https://ieftimov.com/post/understanding-bytes-golang-build-tcp-protocol/
---

## [3][Recursive JSON Parser](https://www.reddit.com/r/golang/comments/ful672/recursive_json_parser/)
- url: https://www.reddit.com/r/golang/comments/ful672/recursive_json_parser/
---
Howdy!

Do y'all use any libraries that parse JSON into structs without needing to know the structure of the JSON beforehand?

I've found a couple around the interwebs but they don't fit my needs (too slow for hundreds of thousands of concurrent JSON requests).

Edit: We have a couple hundred thousand requests coming through our API hourly and we grab the data from each request. Most fit our API spec but ~5,000/hour do not. We want to capture these in a _struct way_ instead of simply logging and parsing a JSON string to potentially expand our API to better serve our users so we're looking for a lib that can recursively parse a JSON payload into a struct so we can use that in some of our functions. The goal is not to simply log, but use the payload data to call other functions based on the value payloads that we weren't expecting to receive (which are already built).
## [4][Access denied.](https://www.reddit.com/r/golang/comments/fuph9l/access_denied/)
- url: https://www.reddit.com/r/golang/comments/fuph9l/access_denied/
---
Hi. I am new to Golang. I use VScode as IDE. I created a new file, save it as main.go. I wrote a simple println method and tried to run it from vsvode trrminal. When i hit &gt;go run main.go I get an Access Denied error. 

Can someone help me?
## [5][A tiny Go library which can help you to keep your configurations structured. Currently you can init fields from: environment variables, flags, files(json and yaml) and 'default' tag.](https://www.reddit.com/r/golang/comments/fu7s9o/a_tiny_go_library_which_can_help_you_to_keep_your/)
- url: https://github.com/BoRuDar/configuration
---

## [6][Gobinaries.com](https://www.reddit.com/r/golang/comments/fuo5ez/gobinariescom/)
- url: https://gobinaries.com/
---

## [7][Anyone participating in Code Jam?](https://www.reddit.com/r/golang/comments/fukkqd/anyone_participating_in_code_jam/)
- url: https://www.reddit.com/r/golang/comments/fukkqd/anyone_participating_in_code_jam/
---
Has anyone been able to do indicium?
## [8][Quarantine has driven me crazy - I'm writing ArnoldC as a UNIX shell, in Go](https://www.reddit.com/r/golang/comments/ftxrwx/quarantine_has_driven_me_crazy_im_writing_arnoldc/)
- url: https://www.reddit.com/r/golang/comments/ftxrwx/quarantine_has_driven_me_crazy_im_writing_arnoldc/
---
For those unfamiliar: [https://esolangs.org/wiki/ArnoldC](https://esolangs.org/wiki/ArnoldC)

&amp;#x200B;

I'm a lowly CS student with no experience in Go and no business writing any systems stuff, but what else would I do with all my free time now? Anyway, here's the repo with only one command implemented. **Any criticism is welcome, because honestly I have no idea what I'm doing.**

&amp;#x200B;

[https://github.com/John123Allison/ArnoldShell](https://github.com/John123Allison/ArnoldShell)
## [9][time.Duration question](https://www.reddit.com/r/golang/comments/fugefd/timeduration_question/)
- url: https://www.reddit.com/r/golang/comments/fugefd/timeduration_question/
---
why does time stop at hour?  why does it not at least go to day?
## [10][JWT parse error](https://www.reddit.com/r/golang/comments/fuj0s3/jwt_parse_error/)
- url: https://www.reddit.com/r/golang/comments/fuj0s3/jwt_parse_error/
---
Does anyone has experience with this error?

&gt;cannot use (func(token \*jwt.Token) (interface{}, error) literal) (value of type func(token \*jwt.Token) (interface{}, error)) as jwt.Keyfunc value in argument to jwt.Parse

my code:

&amp;#x200B;

https://preview.redd.it/i6cq21xluoq41.png?width=776&amp;format=png&amp;auto=webp&amp;s=42b73d98fe3bdcf638b0caad0020ff3030f6020e

Edit: Reddit won't let me post the code properly:

`func GetLoggedUser(w http.ResponseWriter, r *http.Request) {`  
`w.Header().Set("Content-Type", "application/json")`  
`tokenString := r.Header.Get("Authorization")`  
`token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {`  
 `if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {`  
 `return nil, fmt.Errorf("Unexpected signing method")`  
 `}`  
 `return []byte("secret"), nil`  
 `})`  
 `var result models.User`  
 `var res models.Exception`  
 `if claims, ok := token.Claims.(jwt.MapClaims); ok &amp;&amp; token.Valid {`  
`result.Username = claims["Username"].(string)`  
`json.NewEncoder(w).Encode(result)`  
 `return`  
 `} else {`  
`res.Message = err.Error()`  
`json.NewEncoder(w).Encode(res)`  
 `return`  
 `}`  
`}`  

