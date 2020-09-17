# golang
## [1][Datatable : opensource go in-memory table](https://www.reddit.com/r/golang/comments/iufaui/datatable_opensource_go_inmemory_table/)
- url: https://www.reddit.com/r/golang/comments/iufaui/datatable_opensource_go_inmemory_table/
---
Hi all,

Datatable is a Go package to manipulate tabular data, like an excel spreadsheet.

Select, sort, join, agregate, import from csv...See features and howto on github.

[https://github.com/datasweet/datatable](https://github.com/datasweet/datatable)

Please Try it and give us feedback.
## [2][Live debug go in prod using eBPF](https://www.reddit.com/r/golang/comments/itzsrc/live_debug_go_in_prod_using_ebpf/)
- url: https://www.reddit.com/r/golang/comments/itzsrc/live_debug_go_in_prod_using_ebpf/
---
Link: [https://blog.pixielabs.ai/blog/ebpf-function-tracing/post/](https://blog.pixielabs.ai/blog/ebpf-function-tracing/post/)

https://i.redd.it/l0rwtzsvkjn51.gif
## [3][Good restful api tutorial](https://www.reddit.com/r/golang/comments/iuib51/good_restful_api_tutorial/)
- url: https://www.reddit.com/r/golang/comments/iuib51/good_restful_api_tutorial/
---
Hi I am learning Go and love it.

I have worked through the Alex Edwards book, Lets Go and also the tour. I have done other dev, JS (inc node) and PHP as the primary over the recent years. 

I am now wanting to create a restful api for a server so it can be hooked up for front end service. (React not that it matters i guess). 

Just wondering if there is any good resources that show a best practice for it? I know there are plenty of results that come up on google but want to learn the correct way. PHP gave me a bad experience of going with tutorials as not all are great.

Would love to hear your thoughts!

Matthew
## [4][Missing Padding](https://www.reddit.com/r/golang/comments/iuf962/missing_padding/)
- url: https://www.reddit.com/r/golang/comments/iuf962/missing_padding/
---
I ran into and issue that gave me sleepless night, 

```
sign_test.go:31: got a29kkB5HdwZj+IQgtzIIuMh1Pl+8KooSIrrEC7WcrNwA5+8tufTHGNkzjwmocKfV5pl+BBpdqJnTG0w66ymIB2hlbGxv, want a29kkB5HdwZj+IQgtzIIuMh1Pl+8KooSIrrEC7WcrNwA5+8tufTHGNkzjwmocKfV5pl+BBpdqJnTG0w66ymIBw==
```

After hours wasted, i realized the problem may be from the `base64`


```go
const data = "AdRNDeBxM96UcaTQ+FijHLMCkm6pnkVUJPWkAOkTh3p106HzVCgeFkWfLCw1+KRzzQJTd39ZfoFkkSmuefPXAg=="

func main() {

	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal(err)
	}

	out := base64.RawStdEncoding.EncodeToString(b)

	fmt.Println("expected: ", data)
	fmt.Println("output  : ", out)

	if out != data {
		log.Fatal("somthing went wrong")
	}
}
```

#### Output
```
expected:  AdRNDeBxM96UcaTQ+FijHLMCkm6pnkVUJPWkAOkTh3p106HzVCgeFkWfLCw1+KRzzQJTd39ZfoFkkSmuefPXAg==
output  :  AdRNDeBxM96UcaTQ+FijHLMCkm6pnkVUJPWkAOkTh3p106HzVCgeFkWfLCw1+KRzzQJTd39ZfoFkkSmuefPXAg
2009/11/10 23:00:00 somthing went wrong
```

test: https://play.golang.org/p/BLzzAje2Q_c


Can anyone please clarify why the padding `==` is missing ?
## [5][gopls v0.5.0](https://www.reddit.com/r/golang/comments/itq7ih/gopls_v050/)
- url: https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.0
---

## [6][Github action golang test annotator](https://www.reddit.com/r/golang/comments/iuevjd/github_action_golang_test_annotator/)
- url: https://www.reddit.com/r/golang/comments/iuevjd/github_action_golang_test_annotator/
---
Hey,

I didn't find any GitHub action which annotates about failed tests, so I've implemented one. Available in the marketplace [https://github.com/marketplace/actions/golang-test-annotations](https://github.com/marketplace/actions/golang-test-annotations).
## [7][hasql: Library for accessing multi-host SQL databases from Go](https://www.reddit.com/r/golang/comments/itzdm7/hasql_library_for_accessing_multihost_sql/)
- url: https://github.com/yandex/go-hasql
---

## [8][docx: Simple pure Go (golang) library for creating DOCX file](https://www.reddit.com/r/golang/comments/itv2qw/docx_simple_pure_go_golang_library_for_creating/)
- url: https://github.com/gingfrederik/docx
---

## [9][Sending form data to a different request URL](https://www.reddit.com/r/golang/comments/iucg51/sending_form_data_to_a_different_request_url/)
- url: https://www.reddit.com/r/golang/comments/iucg51/sending_form_data_to_a_different_request_url/
---
Hello, I am trying to take form data submitted to a route on my website and forward it to another route.

So from my understanding there are two ways to submit form data of r.Values

http.PostForm()

but I cannot set the headers so I am not using this, and

[http.Post](https://http.Post)()

My question is, do I have to consider anything when sending r.Form (r.Values) from one route to another?

I tried:

resp, err := http.PostForm("[http://localhost:3000/](http://localhost:3000/sci/outbox)post", r.Form)

and

req, err := http.NewRequest("POST", "[http://localhost:3000/](http://localhost:3000/sci/outbox)post", strings.NewReader(r.Form.Encode))

and

req, err := http.NewRequest("POST", "[http://localhost:3000/](http://localhost:3000/sci/outbox)post", r.Body)

also

req, err := http.Post("[http://localhost:3000/](http://localhost:3000/sci/outbox)post", "multipart/form-data", r.Body)

non of these once it gets the [http://localhost:3000/](http://localhost:3000/sci/outbox)post retains any of the form or body data

Would anyone know how to correctly send formdata/values from one route to a next?

Any help would be appreciated.

&amp;#x200B;

Here is the client routing 

[https://paste.debian.net/1164143/](https://paste.debian.net/1164143/)

here is the server routing

[https://paste.debian.net/1164145/](https://paste.debian.net/1164145/)

Nothing is accessible on server side

&amp;#x200B;

I have tried to reference this [https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data](https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data) but for my use cases I am wondering how to send form data from one location to the next I am currently submiting a form to localhost:2000 and I want to send the same form data to localhost:3000
## [10][Why is golang extensively used in Network Security?](https://www.reddit.com/r/golang/comments/ityww8/why_is_golang_extensively_used_in_network_security/)
- url: https://www.reddit.com/r/golang/comments/ityww8/why_is_golang_extensively_used_in_network_security/
---
Hello! First post in this sub. Just a disclaimer, I am not here to ignite a war between programming languages, I am genuinely curious and unknowledgeable. 

I am an avid Rust user and played with the pros and cons Go vs Rust for quite a while before settling on Rust. This year has been crazy and I’ve changed industries and landed a role going into Penetration testing and network security. I’ve participated in loads of CTF events and all majority of pre-baked tools I have come across are either written in Golang or Python. Python I can understand why, but I don’t get the favoritism of Golang. Is this just coincidence of netsec over the years or does Golang provide tools that Rust doesn’t in terms of exploitation and netsec programs?

Thanks for the help.

Edit:

The purpose of my question is to see if I should allocate some of my time to learn Golang, but in my opinion programming languages are all just abstract and you could do mostly the same thing in any language. But I want to see if there is enough value for me to add golang to my “bucket”
