# golang
## [1][New release of DevDash, to create customizable dashboards in your terminal](https://www.reddit.com/r/golang/comments/gzk5z3/new_release_of_devdash_to_create_customizable/)
- url: https://www.reddit.com/r/golang/comments/gzk5z3/new_release_of_devdash_to_create_customizable/
---
Hello everybody!

I'm very proud to release the new version of [DevDash](https://github.com/Phantas0s/devdash), which allows you to create very flexible terminal dashboards.

What's the difference with the existing solutions, you might rightfully ask? You can personalize everything: what data you put in the widgets, from where, the color, the placement, and so on.

You can as well use very generic widgets to query yourself the APIs available!

With this new version, you can as well feed your nice dashboard widgets with data from command lines or scripts output, from your local host or from a remote host (via SSH).

The way to handle dashboard configurations has been simplified, too. I still think there is a lot of improvement to be made on that matter, so I will focus on that next.

The performance have been significantly improved, too.

If you like it, don't forget to give a sweet little Github star :)

* [Github repo](https://github.com/Phantas0s/devdash) 
* [Documentation](https://thedevdash.com)

&amp;#x200B;

https://preview.redd.it/lluiebevyu351.png?width=1920&amp;format=png&amp;auto=webp&amp;s=678a6e016c5ff92c7c731ed651af5780a62bb011

https://preview.redd.it/5zile4evyu351.png?width=1920&amp;format=png&amp;auto=webp&amp;s=88ad4abe600ca787fe0ab0b91c46ac1784b1c29a
## [2][Question about making GUIs in go](https://www.reddit.com/r/golang/comments/gzga4u/question_about_making_guis_in_go/)
- url: https://www.reddit.com/r/golang/comments/gzga4u/question_about_making_guis_in_go/
---
Hello I am wondering if anyone is currently making any desktop applications using go. If so what frameworks? I really like go and I was looking to build a desktop Gui with it. Any advice or information as to why this may or may not be a good idea is appreciated.
## [3][Phil Wadler: Featherweight Go - YouTube](https://www.reddit.com/r/golang/comments/gz19le/phil_wadler_featherweight_go_youtube/)
- url: https://www.youtube.com/watch?v=Dq0WFigax_c
---

## [4][Using gRPC to communicate between Go and Python programs - Ardan Labs](https://www.reddit.com/r/golang/comments/gz9kcr/using_grpc_to_communicate_between_go_and_python/)
- url: https://www.ardanlabs.com/blog/2020/06/python-go-grpc.html
---

## [5][Mapping complex SQL data in Go](https://www.reddit.com/r/golang/comments/gz7ar6/mapping_complex_sql_data_in_go/)
- url: https://www.reddit.com/r/golang/comments/gz7ar6/mapping_complex_sql_data_in_go/
---
I finally released a package which lets you map your SQL data onto Go structs while keeping track of has-one and has-many relationships. Feature requests, comments, stars greatly appreciated.  
[https://github.com/jackskj/carta](https://github.com/jackskj/carta)
## [6][Creating an API: How can I refactor this code? [Beginner][Code Review]](https://www.reddit.com/r/golang/comments/gzip32/creating_an_api_how_can_i_refactor_this_code/)
- url: https://www.reddit.com/r/golang/comments/gzip32/creating_an_api_how_can_i_refactor_this_code/
---
I am fairly new to Golang and I was creating an API to fetch commits from any Github repository using their API. I was then creating a frequency count, based on the date and then sending it to the frontend.

Currently, I only work in one main.go and I am not very aware of the programming pracices in Golang. I am using the Mux router to work with http.

 I have 2 routes.  
The '/' route is handled by getHome and '/updateRepository' is handled by updateRepository  
 handler.

I maintain a currentRepository variable globally.  
In both handler methods, I first get the timestamps of all the commits for that repository using Github's API using getCommitTimeStamps.

Then I parse the timestamps using parseTimeStamps, which first converts the RFC3339 date format to YYYY-MM-DD format and then basically create a map to store the frequency for each date.

The only extra method which updateRepository handler uses, is parseURL  
which basically extracts the Github Owner Name and Repository name from any Github URL. 

I plan to update this using an input text element on my Frontend.

init() initialises a new Client for enabling access to Github's API.

I have split everything into functions, what else can I do? What are the things I can clean?

    package main
    
    import (
        "context"
        "encoding/json"
        "fmt"
        "net/http"
        "strings"
        "time"
    
        "github.com/google/go-github/v31/github"
        "github.com/gorilla/mux"
        "golang.org/x/oauth2"
    )
    
    var client *github.Client
    var ctx context.Context
    
    type Repository struct {
        Owner string `json:"owner"`
        Name  string `json:"name"`
    }
    
    type Frequency struct {
        Date  string `json:"date"`
        Count int    `json:"count"`
    }
    
    var currentRepository = Repository{
        Owner: "randomusername",
        Name:  "foobar",
    }
    
    const GITHUB_AUTH_TOKEN string = "MY_AUTH_TOKEN"
    
    func init() {
        ctx = context.Background()
        ts := oauth2.StaticTokenSource(
            &amp;oauth2.Token{AccessToken: GITHUB_AUTH_TOKEN},
        )
        tc := oauth2.NewClient(ctx, ts)
        client = github.NewClient(tc)
    }
    
    func parseURL(repoURL *string) Repository {
        res := currentRepository
        if *repoURL != "" {
            index := strings.Index(*repoURL, "github.com")
    
            repo := *repoURL
            repo = repo[index+11:]
            var base []string
            base = strings.Split(repo, "/")
            var ans []string
            for _, e := range base {
                if e != "" {
                    ans = append(ans, e)
                }
            }
            res = Repository{
                Owner: ans[0],
                Name:  ans[1],
            }
        }
    
        return res
    }
    
    func getCommitTimeStamps(repo Repository) []time.Time {
        var commitTimeStamps []time.Time
    
        opt := &amp;github.CommitsListOptions{
            ListOptions: github.ListOptions{PerPage: 50},
        }
    
        var allRepositoryCommits []*github.RepositoryCommit
        for {
            repositoryCommits, resp, err := client.Repositories.ListCommits(ctx, repo.Owner, repo.Name, opt)
            if err != nil {
                fmt.Println(err)
            }
            allRepositoryCommits = append(allRepositoryCommits, repositoryCommits...)
            if resp.NextPage == 0 {
                break
            }
            opt.Page = resp.NextPage
        }
    
        for _, e := range allRepositoryCommits {
            commitTimeStamps = append(commitTimeStamps, *e.GetCommit().Author.Date)
        }
        return commitTimeStamps
    }
    
    func parseTimeStamps(commitTimeStamps []time.Time) []Frequency {
        const RFC3339FullDate = "2006-08-06"
        timeSeries := make([]string, len(commitTimeStamps))
        for i, t := range commitTimeStamps {
            timeSeries[i] = fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
        }
        m := make(map[string]int)
        for _, t := range timeSeries {
            m[t]++
        }
        var f []Frequency
    
        for k, v := range m {
            f = append(f, Frequency{
                Date:  k,
                Count: v,
            })
        }
        return f
    }
    
    func getHome(w http.ResponseWriter, r *http.Request) {
        commitTimeStamps := getCommitTimeStamps(currentRepository)
        commitTimeObjects := parseTimeStamps(commitTimeStamps)
    
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
        _ = json.NewEncoder(w).Encode(commitTimeObjects)
    }
    
    func updateRepository(w http.ResponseWriter, r *http.Request) {
        var url struct {
            URL string `json:"url"`
        }
    
        _ = json.NewDecoder(r.Body).Decode(&amp;url)
        currentRepository = parseURL(&amp;url.URL)
        fmt.Println(currentRepository)
        // currentRepository looks like {bradtraversy reactcharts}
        commitTimeStamps := getCommitTimeStamps(currentRepository)
        commitTimeObjects := parseTimeStamps(commitTimeStamps)
        // commitTimeObjects looks like [{2016-04-21 2} {2016-04-23 5} {2016-04-25 1}]
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "PUT")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        fmt.Println(currentRepository)
        _ = json.NewEncoder(w).Encode(commitTimeObjects)
    }
    
    func main() {
        router := mux.NewRouter()
        router.Use(mux.CORSMethodMiddleware(router))
        router.HandleFunc("/", getHome).Methods("GET", "OPTIONS")
        router.HandleFunc("/updateRepository", updateRepository).Methods("PUT", "OPTIONS")
    
        http.ListenAndServe(":8000", router)
    }
## [7][Which Graph database to use?](https://www.reddit.com/r/golang/comments/gzgjph/which_graph_database_to_use/)
- url: https://www.reddit.com/r/golang/comments/gzgjph/which_graph_database_to_use/
---
I need a graph database for my project, which will feature friend / connection and group, like Facebook, and connection degree (like LinkedIn's shows 1st, 2nd or 3rd degree connection) and I'm not sure which one to use.

Primary requirements:
1. Free and open sourced (must not be GPL).
2. Should support Go (as my back-end is fully in go).
3. Must be scalable.

Secondary requirements:
1. Should support backup, restore.
2. Should support encryption.

I tried neo4j, it's nice, but it's GPL.

As of now, considering Janusgraph, Dgraph (doesn't support backup, encryption for free version), Cayley.

Any suggestion, advice, experience will be highly appreciated.

Thanks in advance!
## [8][How to Write a CLI With “Just Enough” Architecture](https://www.reddit.com/r/golang/comments/gyxfvp/how_to_write_a_cli_with_just_enough_architecture/)
- url: https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/
---

## [9][Good concurrent/parallel programming books](https://www.reddit.com/r/golang/comments/gz9cg2/good_concurrentparallel_programming_books/)
- url: https://www.reddit.com/r/golang/comments/gz9cg2/good_concurrentparallel_programming_books/
---
Hello gophers!

I've been studying Golang for quite some time and have some understanding of basic concepts now. Even had some experience writing applications which are being used already. However, I still feel like I lack some fundamental patterns and concepts of concurrent and parallel programming besides the stuff presented and described in Go stdlib (mutexes, coroutines, shared memory, etc.)

So, are there any great books describing such general concurrent/parallel computation concepts that you could recommend? Especially if it's not strictly tied to any language except C/C++/Golang and for anyone with 3-4 years of programming work experience? Thanks in advance!
## [10][iron-go-project: github go repo template with batteries included!](https://www.reddit.com/r/golang/comments/gzj599/irongoproject_github_go_repo_template_with/)
- url: https://github.com/ironPeakServices/iron-go-project
---

