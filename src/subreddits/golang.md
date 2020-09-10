# golang
## [1][We are the Go Time podcast. Ask us anything. (AMA)](https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/)
- url: https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/
---
Hi everyone! I'm Jon Calhoun, one of the panelists on the Go Time podcast. For those of you unfamiliar - it is a Go podcast that we record live every Tuesday at 3pm ET. We usually have a guest or two on each episode, and discuss a specific topic (defer, testing, databases, infra, etc). You can check it out here: &lt;https://changelog.com/gotime&gt;

This coming episode we want to try something a little different - we want to make a Q&amp;A episode. There are two reasons for doing this:

1. We are hoping this inspires topics for future episodes.
2. We want a venue to discuss questions that don't really fit into an entire episode on their own.

To make this happen we would like everyone here to post any Go-related questions that you would like us to discuss on air, and we will try to get to as many as possible. I'll also try to type up answers here while we discuss them on the episode.

We will be answering questions live tomorrow, Tuesday, Sep 8. We will repeat questions on air, and since we record live you can join in on the Gophers Slack to ask follow-up questions or to elaborate on questions.
## [2][Go Modules have a v2+ Problem](https://www.reddit.com/r/golang/comments/ipwea6/go_modules_have_a_v2_problem/)
- url: https://donatstudios.com/Go-v2-Modules
---

## [3][Golang 1.15.2 released](https://www.reddit.com/r/golang/comments/ipwctk/golang_1152_released/)
- url: https://www.reddit.com/r/golang/comments/ipwctk/golang_1152_released/
---
[https://github.com/golang/go/issues?q=milestone%3AGo1.15.2+label%3ACherryPickApproved](https://github.com/golang/go/issues?q=milestone%3AGo1.15.2+label%3ACherryPickApproved)
## [4][🔐 prvt 0.5 is out. Personal, E2E encrypted storage, accessible through your browser](https://www.reddit.com/r/golang/comments/ipue2p/prvt_05_is_out_personal_e2e_encrypted_storage/)
- url: https://github.com/ItalyPaleAle/prvt
---

## [5][I tried to build a minimal load balancer](https://www.reddit.com/r/golang/comments/ipzyu6/i_tried_to_build_a_minimal_load_balancer/)
- url: https://youtu.be/4i7_5NE6tlM
---

## [6][Are design patterns important in Go ?](https://www.reddit.com/r/golang/comments/iprdy9/are_design_patterns_important_in_go/)
- url: https://www.reddit.com/r/golang/comments/iprdy9/are_design_patterns_important_in_go/
---
Hi all! I know that certain concepts like dependency injection are necessary to write a well testable code in Go , but what about the design patterns mentioned in the Gang of Four book like visitor pattern , observer pattern etc . Are these design patterns language agnostic ? Do these patterns make sense in the context of Go ?

Are these patterns important to write a well maintainable code in Go ?
## [7][Guide to Environment variables in Go](https://www.reddit.com/r/golang/comments/ipzna1/guide_to_environment_variables_in_go/)
- url: https://www.reddit.com/r/golang/comments/ipzna1/guide_to_environment_variables_in_go/
---
Hey Gophers 👋

Do you store your application configs like DB Passwords, API secrets, etc in your code?   
In this article, I try to cover ways to manage them in better ways.

Understand what are environment variables/configs and how to manage them in Go.

Check this out here -  
[https://www.mohitkhare.com/blog/environment-variable-golang](https://www.mohitkhare.com/blog/environment-variable-golang)

Do suggest feedbacks and improvements!

[environment-variable-golang](https://preview.redd.it/yrsem12baam51.png?width=673&amp;format=png&amp;auto=webp&amp;s=d00a78bba75c74a3a41b91c3b831750b19cf0fe0)
## [8][RD Gateway Client](https://www.reddit.com/r/golang/comments/iq07ph/rd_gateway_client/)
- url: https://www.reddit.com/r/golang/comments/iq07ph/rd_gateway_client/
---
Hi,

At the moment I'm trying to implement a RD Gateway client in go. Is someone of you familiar with the MS RD Gateway protocol? Currently I'm stuck at the initial handshake after opening the 'in' and 'out' channels.

I uploaded my code here: https://github.com/develerik/rdtunnel

It would be great if you could help me :)
## [9][go template for whole folder (tree) structure](https://www.reddit.com/r/golang/comments/iq2j69/go_template_for_whole_folder_tree_structure/)
- url: https://www.reddit.com/r/golang/comments/iq2j69/go_template_for_whole_folder_tree_structure/
---
As I couldn't find a suitable folder structure generator, I've created a simple one for my needs. I hope you like ti too: [https://github.com/lorands/tymlate](https://github.com/lorands/tymlate)
## [10][Implementing of python's threading.Event using golang channels](https://www.reddit.com/r/golang/comments/ipuvdp/implementing_of_pythons_threadingevent_using/)
- url: https://github.com/trivigy/event
---

## [11][[CGO] How to create a struct with nested array of struct?](https://www.reddit.com/r/golang/comments/iq0q5s/cgo_how_to_create_a_struct_with_nested_array_of/)
- url: https://www.reddit.com/r/golang/comments/iq0q5s/cgo_how_to_create_a_struct_with_nested_array_of/
---
I am trying to create a struct, in Go, where one of it's children is an array of struct and then translate it to C using CGO.

I tried something like this in Go

```
*/
typedef struct FileInfo{
	int64_t Size;
	char *Name;
}FileInfo;

typedef struct Result{
	FileInfo **files;
}Result;

int64_t GetResult(void **presult, FileInfo **files) {
	Result *result = (Result *)malloc(sizeof(Result));
	result-&gt;files=files;

	*presult = result;

	int64_t ptr = (int64_t)result;

	return ptr;
}
*/
import "C"
```

```
func Run() {
	var arr []*C.struct_FileInfo

	ai := C.struct_FileInfo{
		Size: C.int64_t(1234),
		Name: C.CString("some name"),
	}

	arr = append(arr, &amp;ai)

	var presult unsafe.Pointer
	ptr := C.GetResult(&amp;presult, &amp;arr[0])

	println("\nResult struct pointer: %v", ptr)
}
```


It threw an `panic: runtime error: cgo argument has Go pointer to Go pointer` error.


How to I fix this error?
