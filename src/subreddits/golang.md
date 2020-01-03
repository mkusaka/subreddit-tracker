# golang
## [1][Chaos-mesh: A Chaos Engineering Platform for Kubernetes](https://www.reddit.com/r/golang/comments/ejbgmn/chaosmesh_a_chaos_engineering_platform_for/)
- url: https://github.com/pingcap/chaos-mesh
---

## [2][Retrieving cookies from a webview in Go](https://www.reddit.com/r/golang/comments/ejg2qd/retrieving_cookies_from_a_webview_in_go/)
- url: https://www.reddit.com/r/golang/comments/ejg2qd/retrieving_cookies_from_a_webview_in_go/
---
I'm writing a Go component, and on startup I want to provide personal user login before the program proceeds. This should ideally be done through a webview that connects to an identity server (not served by the program I am writing in Go; I simply want to use our pre-existing identity server in order to authorize users). Only the login part of the Go component should be webview based, as the rest of what it does is not GUI related.

There seems to be a few libraries that deal with webviews, notably [https://github.com/zserge/webview](https://github.com/zserge/webview). However, I have not found a way to extract cookies using this. Does anyone have experience with doing something similar?

I asked this same question on stackoverflow two weeks ago ( [https://stackoverflow.com/questions/59425959/how-do-i-retrieve-cookies-from-a-go-webview](https://stackoverflow.com/questions/59425959/how-do-i-retrieve-cookies-from-a-go-webview) ), but it seems that either my question is bad, or nobody knows the answer.
## [3][sampler - visualization for any shell command output](https://www.reddit.com/r/golang/comments/ej1qy6/sampler_visualization_for_any_shell_command_output/)
- url: https://github.com/sqshq/sampler
---

## [4][Introducing Horcrux, a program that lets you split your confidential files into encrypted fragments!](https://www.reddit.com/r/golang/comments/eiw4ma/introducing_horcrux_a_program_that_lets_you_split/)
- url: https://github.com/jesseduffield/horcrux
---

## [5][Private secure notes](https://www.reddit.com/r/golang/comments/ejc2ci/private_secure_notes/)
- url: https://www.reddit.com/r/golang/comments/ejc2ci/private_secure_notes/
---
Hi guys, I made small service that encrypts/decrypts messages. It was my learning experience with cryptography and Google Cloud Run from the other side. Comments are appreciated. 

Code: [https://github.com/blunext/obliviate/](https://github.com/blunext/obliviate/)

Sevrice: [https://securenote.io/](https://securenote.io/)
## [6][Authorization/Filter library/projects](https://www.reddit.com/r/golang/comments/ejby79/authorizationfilter_libraryprojects/)
- url: https://www.reddit.com/r/golang/comments/ejby79/authorizationfilter_libraryprojects/
---
Hi!

I am looking for options for some kind of not just an auth library ( casbin/ OPA) , but something which also allows filtering of attributes based on users authorisation level.

Ex:

    Struct A {
        id     int
        atttr1 string
        attr2 string
        attr3 string
    }

Role 1

    if id=100 Read  =&gt; attr1,attr2
    if id=200 Read =&gt; attr2

Role 2

    All ids allow Read attr2,attr3

I have a fairly complex and large data-set and would want to find  performant and architecturally correct/sensible way to do this. 

Currently I am thinking of using either of the libraries to pick a template and apply something like mergo to set default values to nil. Is there a better way to achive this ? Or any open source project does this already?

&amp;#x200B;

Thanks!
## [7][Go: Concurrency &amp; Scheduler Affinity](https://www.reddit.com/r/golang/comments/eiyog2/go_concurrency_scheduler_affinity/)
- url: https://medium.com/a-journey-with-go/go-concurrency-scheduler-affinity-3b678f490488
---

## [8][Am I the first person to write Golang in DOS](https://www.reddit.com/r/golang/comments/eis8fe/am_i_the_first_person_to_write_golang_in_dos/)
- url: https://i.redd.it/tjweeu406a841.jpg
---

## [9][Config service](https://www.reddit.com/r/golang/comments/eja1sn/config_service/)
- url: https://www.reddit.com/r/golang/comments/eja1sn/config_service/
---
Hey guys not sure if this is really adequate here but was wondering what do you usually use as configuration service. I've seen this https://cloud.spring.io/spring-cloud-config/reference/html/ but I think it uses github as a backend and that troubles me a little bit. Is using github as a backend scalable to 50 or so microservices (nvm the Java part of this post).

Thanks in advance
## [10][chaimleib/errors - informative and convenient errors for Go](https://www.reddit.com/r/golang/comments/ej8g1r/chaimleiberrors_informative_and_convenient_errors/)
- url: https://www.reddit.com/r/golang/comments/ej8g1r/chaimleiberrors_informative_and_convenient_errors/
---
&amp;#x200B;

[Doctor Gopher helping with an error.](https://preview.redd.it/xyte0w0c0h841.png?width=861&amp;format=png&amp;auto=webp&amp;s=3ec8d7c5bcd9539b81818319b321b7f042522504)

[https://github.com/chaimleib/errors](https://github.com/chaimleib/errors)

This is something I've been working on lately to make debugging easier for myself. What do you guys think?

* Stack trace with line numbers
* See the argument values
* See the module and function names
* Abbreviated `go.mod` module name
* Compatible with Go 1.13 `Unwrap()` API

Improved stack trace error output:

    main.main() prog.go:14 program failed
    main.FileHasHello("greet.txt") prog.go:24 could not open file
    open greet.txt: No such file or directory
    No such file or directory

Source ([try me!](https://goplay.space/#HE4BuAJaZYA))

    func main() {
    	b := errors.NewBuilder("")
    	if err := FileHasHello("greet.txt"); err != nil {
    		err = b.Wrap(err, "program failed")
    		fmt.Println(errors.StackString(err))
    		return
    	}
    }
    
    func FileHasHello(fpath string) error {
    	b := errors.NewBuilder("%q", fpath)
    	buf, err := ioutil.ReadFile(fpath)
    	if err != nil {
    		return b.Wrap(err, "could not open file")
    	}
    	if !bytes.Contains(buf, []byte("hello")) {
    		return b.Errorf("could not find `hello` in file")
    	}
    	return nil
    }
