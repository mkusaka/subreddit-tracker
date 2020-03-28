# golang
## [1][My girlfriend’s evening project](https://www.reddit.com/r/golang/comments/fq9gf2/my_girlfriends_evening_project/)
- url: https://i.imgur.com/mjfzcOu.jpg
---

## [2][GoCache - Cache Server in Go](https://www.reddit.com/r/golang/comments/fqjrat/gocache_cache_server_in_go/)
- url: https://github.com/kadnan/gocache
---

## [3][A form deocder that decode request body of any types(xml, json, form, multipart form...) into a scturct by same codebase.](https://www.reddit.com/r/golang/comments/fqbrza/a_form_deocder_that_decode_request_body_of_any/)
- url: https://github.com/clevergo/form
---

## [4][Help, noob, trying to understand](https://www.reddit.com/r/golang/comments/fqg1xq/help_noob_trying_to_understand/)
- url: https://www.reddit.com/r/golang/comments/fqg1xq/help_noob_trying_to_understand/
---
I only know a little bit of python and I'm trying to learn Go. I'm following this book: The Go Programming Language by Alan A. A. Donovan and Brian W. Kernighan. In chapter 1, page 11, a program to find duplicate lines in files is discussed. [This is the code](https://pastebin.com/vAjJhWr0) . Can someone explain why we are using pointers while mentioning the type for "f" in this line.

func countLines(f \*os.File, counts map\[string\]int)

This is the first instance I came across such thing.  Help please.
## [5][To people who came from python programming](https://www.reddit.com/r/golang/comments/fqatyp/to_people_who_came_from_python_programming/)
- url: https://www.reddit.com/r/golang/comments/fqatyp/to_people_who_came_from_python_programming/
---
How did you adjust to life with Golang. I know Golang's model is simplicity is complicated, but when I'm building, say an api applications using http sessions its just simpler for me to go back to requests.session package in python, then figure it out with net/http and gorilla mux; I would think for a language that prides itself in simplicity that http session managing would be an internal package and pointers, I'm still trying to get the hang of pointers. I understand them, but when I see a pointer on an Http request type it confuses me.
## [6][A path to learn Go?](https://www.reddit.com/r/golang/comments/fpzcrc/a_path_to_learn_go/)
- url: https://www.reddit.com/r/golang/comments/fpzcrc/a_path_to_learn_go/
---
Does anyone have any valuable resources that they want to share with me?
## [7][Json object passing to a function fails](https://www.reddit.com/r/golang/comments/fqik2y/json_object_passing_to_a_function_fails/)
- url: https://www.reddit.com/r/golang/comments/fqik2y/json_object_passing_to_a_function_fails/
---
In this program, I try pass a JSON object to a function and then to a REST API with `POST`. To do that I have following.

```
type resposeObj struct{
	uid 	string `json :"uid"`
	service string `json : "service"`
}

response := resposeObj{uid:validatemailID.String(),service:"API Gateway"}

validateResponse,err := json.Marshal(response)

		if err != nil{
			log.Println("Error in marshaling data",err)
		}

		log.Println("Validate response byte slice: ",validateResponse)

err = sendResponse(validatemailResponse) 

```
```

func sendResponse(req []byte)error{
	log.Println("Received response : ",string(req))
    return nil
}

```

And I get `Validate response byte slice:  [123 125]` as output. Which should be the byte array of Json object.

What am I missing here? Why its not converting to byte slice?

https://stackoverflow.com/questions/60899550/json-object-passing-to-a-function-fails-in-golang
## [8][Coworker made a nice tool to manage SSH hosts](https://www.reddit.com/r/golang/comments/fq2nyw/coworker_made_a_nice_tool_to_manage_ssh_hosts/)
- url: https://github.com/aalbacetef/ssh-helper
---

## [9][Coding bud](https://www.reddit.com/r/golang/comments/fqeat1/coding_bud/)
- url: https://www.reddit.com/r/golang/comments/fqeat1/coding_bud/
---
I am re-starting my journey with Go and I am looking for someone to bug about basics reasoning in Go programming. Anyone is generous enough to help me learn? I have 3 years of experience in programming (Python, JavaScript), so I understand tech talk, just need someone to discuss Go specifics with.
## [10][Starter Projects in Go](https://www.reddit.com/r/golang/comments/fqbda4/starter_projects_in_go/)
- url: https://www.reddit.com/r/golang/comments/fqbda4/starter_projects_in_go/
---
I am from a Java/JS/PHP/C background. I m learning Go, having a hard time with the syntax but I really want to learn. 

What starter projects (preferably web based) I can work on  in Go to learn it decently and put on my resume as I will be graduating in 2 months and looking for a job which has some Go involved?
