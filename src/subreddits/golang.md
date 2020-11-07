# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Static Dependency Analysis Tool for Go Files](https://www.reddit.com/r/golang/comments/jpm0aj/static_dependency_analysis_tool_for_go_files/)
- url: https://i.redd.it/h71qm0zrhrx51.gif
---

## [3][A Tour of GO moretypes 15: Why slice append extra 1 in capacity ?](https://www.reddit.com/r/golang/comments/jppmsp/a_tour_of_go_moretypes_15_why_slice_append_extra/)
- url: https://www.reddit.com/r/golang/comments/jppmsp/a_tour_of_go_moretypes_15_why_slice_append_extra/
---
In "A Tour of Go" on 15th of moretypes why "append(s, 2, 3, 4) " show "len=5 cap=6 \[0 1 2 3 4\]" instead of "len=5 cap=5 \[0 1 2 3 4\]" ? [question](https://tour.golang.org/moretypes/15):  
Code:  


    package main
    
    import "fmt"
    
    func main() {
    	var s []int
    	printSlice(s)
    
    	// append works on nil slices.
    	s = append(s, 0)
    	printSlice(s)
    
    	// The slice grows as needed.
    	s = append(s, 1)
    	printSlice(s)
    
    	// We can add more than one element at a time.
    	s = append(s, 2, 3, 4)
    	printSlice(s)
    }
    
    func printSlice(s []int) {
    	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    }
    
    Output:
    len=0 cap=0 []
    len=1 cap=1 [0]
    len=2 cap=2 [0 1]
    len=5 cap=6 [0 1 2 3 4]

I am new to learning go.
## [4][GORM Customise Join table](https://www.reddit.com/r/golang/comments/jpp2dz/gorm_customise_join_table/)
- url: https://www.reddit.com/r/golang/comments/jpp2dz/gorm_customise_join_table/
---
Hello everyone

I am trying to customize many2many table join. I have two tables from which I want to have taken the ids and want another field, which will tell me when the entry in the join table was made. The ids are coming fine, but the creation\_at is not updating and shows Null instead of time.

    // this is the table join struct which I want to make
    
    type UserChallenges struct {
    
    gorm.JoinTableHandler
    
    CreatedAt   time.Time
    
    UserID      int
    
    ChallengeID int
    
    }
    
    
    
    //hook before create
    
    func (UserChallenges) BeforeCreate(Db \*gorm.DB) error {
    
    Db.SetJoinTableHandler(&amp;User{}, "ChallengeId", &amp;UserChallenges{})
    
    return nil
    
    }

This is not giving any error on the build. Please tell me what I am missing so that I can get the creation time field in this.PS - The documentation of GORM on [gorm.io](https://gorm.io) is still showing SetupJoinTable method but it is deprecated in the newer version.
## [5][Rest API consume and host](https://www.reddit.com/r/golang/comments/jpp0ar/rest_api_consume_and_host/)
- url: https://www.reddit.com/r/golang/comments/jpp0ar/rest_api_consume_and_host/
---
&amp;#x200B;

I am new to go and also new to building rest APIs so I am trying to crunch a lot at once.

I am trying to request something, like a book by its id, and then host it locally so that if I write my local URL, like http://localhost:8080​/books?books=&lt;book-id&gt;  it would show me the specific result.

to try to be concrete, I need to connect the two. Get the information from that URL, so "consume" and also host it locally, specifically by ID. I am not sure how to do both at once.

To create the paths, I've been using gorilla mux

So separately, I've used this, which would give me all the books at once (URL is not real)

At the end I want to be able to put in my computer [http://localhost:8080/books/b5a92d0e-5fb4-43d4-ba60-c012135958e4](http://localhost:8080/species/b5a92d0e-5fb4-43d4-ba60-c012135958e4)

which would give me the book in json in my browser

&amp;#x200B;

pls be nice!!

`func main(){response, err := http.Get("https://bookibook.herokuapp.com/books/")if err != nil{fmt.Printf("there is no book with this ID %s\n", err)} else{data, _ := ioutil.ReadAll(response.Body)fmt.Println(string(data))`

and then this, which would create a local path for http://localhost:8080/books/&lt;id&gt;

`import ("fmt""log""net/http""github.com/gorilla/mux")`

`func getID(w http.ResponseWriter, r *http.Request) {vars := mux.Vars(r)fmt.Fprintf(w, "Get id %s\n", vars["id"])}`

`func main() {// Configure routes.router := mux.NewRouter()router.HandleFunc("/books/{id}/", getID).Methods(http.MethodGet)`

`// Start HTTP server.if err := http.ListenAndServe(":8080", router); err != nil {log.Fatal(err)}}`
## [6][Rust vs Go](https://www.reddit.com/r/golang/comments/jp4kvp/rust_vs_go/)
- url: https://bitfieldconsulting.com/golang/rust-vs-go
---

## [7][Unmarshal a json to a struct and the json object can scale ??](https://www.reddit.com/r/golang/comments/jpoe3a/unmarshal_a_json_to_a_struct_and_the_json_object/)
- url: https://www.reddit.com/r/golang/comments/jpoe3a/unmarshal_a_json_to_a_struct_and_the_json_object/
---
i am using a public api that has a json object that can scale 

for example 

`{`  
`"places" : {`

`"london" : 65,`

`"stockholm" : 14,`

`"somecity" : 3,`  
`....`  
`}`  
`}`  


and then you can have another object like this   
`{`  
`"places":{`

 `"morroco" : 30`

`}`

`}`

&amp;#x200B;

unfortunately the public api didn't put this stuff as an array because (i don't know  why )
## [8][Display multiple progress bars for all the goroutines](https://www.reddit.com/r/golang/comments/jpqm3f/display_multiple_progress_bars_for_all_the/)
- url: https://www.reddit.com/r/golang/comments/jpqm3f/display_multiple_progress_bars_for_all_the/
---
Hi guys.

I am working on a scenario where I start multiple goroutines each designed to do certain task and I want to display the progress of each goroutine using progress bar.

I have tried using uiprogress , mpb but all these progress bars fail when the number of bars created are high ex.100
https://github.com/gosuri/uiprogress/blob/master/example/full/full.go

When the number of bars provided are high the output starts overlapping leading to more than one progress bar getting display for each goroutine.


Can anybody help me with some approach for displaying progress bars only once per goroutine even when the number of goroutines are high

Or is there any other way apart from progress bars which I can use to display dynamic data from goroutines
## [9][Idiomatic Go Resources](https://www.reddit.com/r/golang/comments/jp3om9/idiomatic_go_resources/)
- url: https://medium.com/@dgryski/idiomatic-go-resources-966535376dba
---

## [10][Gron: A command line tool that makes JSON greppable](https://www.reddit.com/r/golang/comments/jp7vzq/gron_a_command_line_tool_that_makes_json_greppable/)
- url: https://github.com/tomnomnom/gron
---

## [11][Golang 1.15.4 released](https://www.reddit.com/r/golang/comments/joznmi/golang_1154_released/)
- url: https://www.reddit.com/r/golang/comments/joznmi/golang_1154_released/
---
[https://golang.org/doc/devel/release.html#go1.15.minor](https://golang.org/doc/devel/release.html#go1.15.minor)

&amp;#x200B;

[https://github.com/golang/go/issues?q=milestone%3AGo1.15.4+label%3ACherryPickApproved](https://github.com/golang/go/issues?q=milestone%3AGo1.15.4+label%3ACherryPickApproved)
