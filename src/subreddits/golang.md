# golang
## [1][PayPal Taps Go to Modernize and Scale: In our tightly managed environments where we run Go code, we have seen a CPU reduction of approximately ten percent with cleaner and maintainable code.](https://www.reddit.com/r/golang/comments/ilr81s/paypal_taps_go_to_modernize_and_scale_in_our/)
- url: https://go.dev/solutions/paypal/
---

## [2][All about PKCE and generate code_verifier and code_challenge through multiple languages including golang.](https://www.reddit.com/r/golang/comments/ilt9zb/all_about_pkce_and_generate_code_verifier_and/)
- url: https://www.loginradius.com/engineering/blog/pkce/
---

## [3][treedrawer: a Go module for drawing trees on the terminal](https://www.reddit.com/r/golang/comments/ilsocu/treedrawer_a_go_module_for_drawing_trees_on_the/)
- url: https://www.reddit.com/r/golang/comments/ilsocu/treedrawer_a_go_module_for_drawing_trees_on_the/
---
Hi Gophers

I'm a Go and reddit noob (this is my first post), so please be gentle.

I'd like to share with you my latest project: [treedrawer](https://github.com/m1gwings/treedrawer)

It is a simple Go module that can draw trees with several layers and several children for each node on the terminal. It can handle nodes with any type that satisfies the NodeValue interface, so you can have trees with integers, strings or custom types.

I started this project just for the sake of learning. I am currently studying "The Go Programming Language" and, in order to solve an exercise that required to implement the String() method on a tree, I ended up creating this.

I think that my code is far from "Good Golang code", so I'd like your feedback to discover my mistakes and improve. I also believe that maybe someone could find this module useful, but I'm not that sure.

Thanks in advance.
## [4][Introducing Clean Architecture by refactoring a Go project](https://www.reddit.com/r/golang/comments/il70h4/introducing_clean_architecture_by_refactoring_a/)
- url: https://www.reddit.com/r/golang/comments/il70h4/introducing_clean_architecture_by_refactoring_a/
---
Hi all!

It seems questions about clean architecture, DDD, and related patterns seem to come up on this subreddit pretty often. I'd like to share some ideas on this topic together with refactoring of a [real go project](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example). I hope someone finds it useful. :)

[https://threedots.tech/post/introducing-clean-architecture/](https://threedots.tech/post/introducing-clean-architecture/)
## [5][Turn your Go REST API to GraphQL using Hasura Actions](https://www.reddit.com/r/golang/comments/ilsouk/turn_your_go_rest_api_to_graphql_using_hasura/)
- url: https://hasura.io/blog/turn-your-go-rest-api-to-graphql-using-hasura-actions/
---

## [6][Twitter authentication with Go Language and Goth](https://www.reddit.com/r/golang/comments/ilsarf/twitter_authentication_with_go_language_and_goth/)
- url: https://www.loginradius.com/engineering/blog/twitter-authentication-with-golang-and-goth/
---

## [7][mechanics of :=](https://www.reddit.com/r/golang/comments/ilmz80/mechanics_of/)
- url: https://www.reddit.com/r/golang/comments/ilmz80/mechanics_of/
---
so if i do something like this;

`for i := 0 ... {
  bar
}`

does i get reallocated with every iteration of the loop? or does it just assign to the existing i?
## [8][Anonymous Struct or Interface in flexible json result?](https://www.reddit.com/r/golang/comments/ilt4s8/anonymous_struct_or_interface_in_flexible_json/)
- url: https://www.reddit.com/r/golang/comments/ilt4s8/anonymous_struct_or_interface_in_flexible_json/
---
I want to pull some data using JSON response and fetch it through the struct in golang. Thing is the contents change and it is really hard to come up with the right method as everything I tried does not seem to work, that's why I would really appreciate some help.

&amp;#x200B;

Original working code, you can copy and paste the following JSON response to check for consistency:

`{"table_contents":[{"id":100,"description":"text100"},{"id":101,"description":"text101"},{"id":1,"description":"text1"}]}`

    package main
    
    import (
    	"fmt"
    	"encoding/json"
    )
    
    type MyStruct1 struct {
    	TableContents []struct {
    		ID          int
    		Description string
    	} `json:"table_contents"`
    }
    
    func main() {
    	result:= []byte(`{"table_contents":[{"id":100,"description":"text100"},{"id":101,"description":"text101"},{"id":1,"description":"text1"}]}`)
    	var container MyStruct1
    	err := json.Unmarshal(result, &amp;container)
    	if err != nil {
    		fmt.Println(" [0] Error message: " + err.Error())
    		return
    	}
    }
    

&amp;#x200B;

New code looks like this, sometimes we get nested array as well.

`{"table_contents":[[{"id":100,"description":"text100"},{"id":101,"description":"text101"}],{"id":1,"description":"text1"}]}`

I wanted to make it work with the code above and this is what I have tried so far:

&amp;#x200B;

1st Try:

    package main
    
    import (
    	"fmt"
    	"encoding/json"
    )
    
    type MyStruct1 struct {
    	TableContents [][]struct {
    		ID          int
    		Description string
    	} `json:"table_contents"`
    }
    
    func main() {
    	result:= []byte(`{"table_contents":[[{"id":100,"description":"text100"},{"id":101,"description":"text101"}],{"id":1,"description":"text1"}]}`)
    	var container MyStruct1
    	err := json.Unmarshal(result, &amp;container)
    	if err != nil {
    		fmt.Println(" [0] Error message: " + err.Error())
    		return
    	}
    }
    

&gt;\[0\] Error message: json: cannot unmarshal object into Go struct field MyStruct1.table\_contents of type \[\]struct { ID int; Description string } 

You might have noticed I have added another set of square brackets in TableContents inside the struct. That did not seem to work.

&amp;#x200B;

2nd Try:

    package main
    
    import (
    	"fmt"
    	"encoding/json"
    )
    
    type MyStruct1 struct {
    	TableContents []interface{} `json:"table_contents"`
    }
    
    func main() {
    	result:= []byte(`{"table_contents":[[{"id":100,"description":"text100"},{"id":101,"description":"text101"}],{"id":1,"description":"text1"}]}`)
    	var container MyStruct1
    	err := json.Unmarshal(result, &amp;container)
    	if err != nil {
    		fmt.Println(" [0] Error message: " + err.Error())
    		return
    	}
    	
    	for i := range container.TableContents {
    		fmt.Println(container.TableContents[i].Description)
    	}
    }
    

&gt; ./prog.go:22:41: container.TableContents\[i\].Description undefined (type interface {} is interface with no methods) 

&amp;#x200B;

This time I tried utilising interface, but it is not possible to retrieve its values once they have been processed or I am doing it wrong?

&amp;#x200B;

\---

In short, we can receive either responses and we have to process it through the struct:

`{"table_contents":[{"id":100,"description":"text100"},{"id":101,"description":"text101"},{"id":1,"description":"text1"}]}`

`{"table_contents":[[{"id":100,"description":"text100"},{"id":101,"description":"text101"}],{"id":1,"description":"text1"}]}`

&amp;#x200B;

I would really like for someone to show me a quick example if anyone is aware how to do this
## [9][A tool for ETL and many other jobs developed in Go](https://www.reddit.com/r/golang/comments/iliu1s/a_tool_for_etl_and_many_other_jobs_developed_in_go/)
- url: /r/ETL/comments/il76z7/dixer_v110_is_out/
---

## [10][Getting TLS certificates with echo](https://www.reddit.com/r/golang/comments/ilqigb/getting_tls_certificates_with_echo/)
- url: https://www.reddit.com/r/golang/comments/ilqigb/getting_tls_certificates_with_echo/
---
Hello I'm having a hard time getting my TLS certificates on a freshly installed debian 10 VPS.

Domain name is "deiz.fr".

Using echo atm, following the most basic example from their websites :

```

e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")  
e.Logger.Fatal(e.StartAutoTLS(":443"))

```


The server starts, but whenever I logged to the domain I get this error message :
&gt; http: TLS handshake error from 31.33.213.24:59575: acme/autocert: unable to satisfy  https://acme-v02.api.letsencrypt.org/acme/authz-v3/6956493770" for domain "deiz.fr": no viable challenge type found

I built a minimum example equivalent with gorilla and I manage to get the certificate :

```
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested with TLS certificates")
	})
	dataDir := "/var/www/.cache/certificates"
	certManager := &amp;autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("deiz.fr", "www.deiz.fr"),
		Cache:      autocert.DirCache(dataDir),
	}
	server := &amp;http.Server{
		Handler:   router,
		Addr:      ":https",
		TLSConfig: &amp;tls.Config{GetCertificate: certManager.GetCertificate},
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	server.ListenAndServeTLS("", "")
}

```

Is there something I am missing with echo ?
