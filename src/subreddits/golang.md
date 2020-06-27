# golang
## [1][Trubka, a Swiss army CLI knife for Apache Kafka, built in Go](https://www.reddit.com/r/golang/comments/hgjrpj/trubka_a_swiss_army_cli_knife_for_apache_kafka/)
- url: http://github.com/xitonix/trubka
---

## [2][gometry is a GO library which offers geometry processing algorithms.](https://www.reddit.com/r/golang/comments/hgnopw/gometry_is_a_go_library_which_offers_geometry/)
- url: https://github.com/roboticeyes/gometry
---

## [3][Why its defined like this? why not true = 1 &amp; false = 0. Can some one help me?](https://www.reddit.com/r/golang/comments/hgpf9o/why_its_defined_like_this_why_not_true_1_false_0/)
- url: https://i.redd.it/21inq86bre751.png
---

## [4][memplot: generate .PNG memory usage plots of any process from a single binary](https://www.reddit.com/r/golang/comments/hg7gb4/memplot_generate_png_memory_usage_plots_of_any/)
- url: https://github.com/0x0f0f0f/memplot
---

## [5][goarrange: Automatic arrangement of Go source code](https://www.reddit.com/r/golang/comments/hgpvck/goarrange_automatic_arrangement_of_go_source_code/)
- url: https://github.com/jdeflander/goarrange
---

## [6][Free Go programming ebook from Digital Ocean](https://www.reddit.com/r/golang/comments/hft9zb/free_go_programming_ebook_from_digital_ocean/)
- url: https://www.digitalocean.com/community/books/how-to-code-in-go-ebook
---

## [7][what is the time complexity of this golang solution to left-rotating a slice?](https://www.reddit.com/r/golang/comments/hgpeap/what_is_the_time_complexity_of_this_golang/)
- url: https://www.reddit.com/r/golang/comments/hgpeap/what_is_the_time_complexity_of_this_golang/
---
Hi Gophers,

New to go and currently practicing 'easy' level algorithms, using go.

Please help me understand the time complexity of the below solution.

Since we only iterate over the slice once, it seems to be O(N) time. we only iterate to the length of d, so even if the input slice has length of 1M, if d is say, 5, we will only iterate 5 times. but then, we create a new Slice based on the length of the full array, on every iteration.

Any explanation appreciated, and would you implement it differently?

Original problem from HackerRank,  **Arrays: Left Rotation.**

    package main
    
    import "fmt"
    
    func main() {
    	slc := []int32{1, 2, 3, 4, 5}
    	rotLeft(slc, 4)
    }
    
    func rotLeft(a []int32, d int32) []int32 {
    	fmt.Println("Starting left rotation on the slice")
    	var newSlice []int32
    	for i := int32(0); i &lt; d; i++ { 
    		newSlice = a[1:len(a)]
    		newSlice = append(newSlice, a[0])
    		a = newSlice
    	}
    	return a
    }
## [8][How to implement POA in go lang?](https://www.reddit.com/r/golang/comments/hgnvti/how_to_implement_poa_in_go_lang/)
- url: https://www.reddit.com/r/golang/comments/hgnvti/how_to_implement_poa_in_go_lang/
---
I've been studying up on proof of authority for quite some time now, wanted to implement it in go lang as I'm in the process of building my own supply chain blockchain using go lang. Not able to get any good resources for the same, any resources?
## [9][Writing a library for creating CRUD REST APIs](https://www.reddit.com/r/golang/comments/hgjqq6/writing_a_library_for_creating_crud_rest_apis/)
- url: https://www.reddit.com/r/golang/comments/hgjqq6/writing_a_library_for_creating_crud_rest_apis/
---
Trying to gauge people here on the idea, and maybe get some interesting design implementations if people have an idea. Currently no code has been released to let people play with it, since it's something I'm working on alongside a project that is using it (so I improve on it as I see use-cases in an actual project).

Essentially, it allows me to rapidly create a whole API. Just added a new table that needs an API endpoint? I can duplicate an existing endpoint (or just have a "template") and then modify it with the new tables type and the properties for that table the endpoint will process.

Example post config:

    func PostUser(w http.ResponseWriter, r *http.Request) {
    	var user db.User
    	var api = rest.NewRequest(r, &amp;w)
    	conf := config.Config{
    		Properties: []config.Property{
    			{
    				Name: "UserID",
    				Type: "int",
    			},
    			{
    				Name:          "Username",
    				Type:          "str",
    				Required:      true,
    				CheckConflict: true,
    			},
    			{
    				Name:     "Password",
    				Type:     "str",
    				Required: true,
    			},
    			{
    				Name:     "Firstname",
    				Type:     "str",
    				Required: true,
    			},
    			{
    				Name:     "Lastname",
    				Type:     "str",
    				Required: true,
    			},
    			{
    				Name:          "Email",
    				Type:          "str",
    				Required:      true,
    				CheckConflict: true,
    			},
    			{
    				Name: "Department",
    				Type: "str",
    			},
    		},
    	}
    	problem, err := api.Post(&amp;user, conf)
    	if err != nil {
    		log.Error(err, "system", "api/users")
    		return
    	}
    
    	if problem.StatusCode != 0 {
    		return
    	}
    
    	id := strconv.Itoa(user.UserID)
    	r.RequestURI = r.RequestURI + id + "/"
    	getUser(w, r, id)
    }

So the steps are pretty simple, register a new rest request, giving it the response writer and request variables. It uses these to automatically send the "correct" response back, weather it's the success of creating the resource, it telling you that a required property was missing, or that a server-side error occurred.

Notes: It uses the [HAL specification](https://tools.ietf.org/html/draft-kelly-json-hal-08) as a base, which might be a turn-off for some people, but I argue it will really ruin any development experience. And for returning problems, it uses a [problem details specification](https://tools.ietf.org/html/rfc7807) for returning an error in a standardized way.

&amp;#x200B;

Anything people would really like to see from such a project?
## [10][Cato: Automated configuration documentation library for Go Projects.](https://www.reddit.com/r/golang/comments/hgadun/cato_automated_configuration_documentation/)
- url: https://github.com/cs3org/cato
---

