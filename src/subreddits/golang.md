# golang
## [1][golang-dev &amp; the next few months](https://www.reddit.com/r/golang/comments/flzoc1/golangdev_the_next_few_months/)
- url: https://groups.google.com/forum/#!topic/golang-dev/UxvN1W2B-zg
---

## [2][What do employers look in GitHub repo of golang services?](https://www.reddit.com/r/golang/comments/fma0dz/what_do_employers_look_in_github_repo_of_golang/)
- url: https://www.reddit.com/r/golang/comments/fma0dz/what_do_employers_look_in_github_repo_of_golang/
---
Apologies if already asked, I have started learning golang and am writing a few backend services. What are the things that employers specifically seek in such personal projects that could help me stand out?
## [3][Golang MVC framework](https://www.reddit.com/r/golang/comments/fmdody/golang_mvc_framework/)
- url: https://github.com/RobyFerro/go-web
---

## [4][Design pattern to avoid hitting the rate limit on an external API resource](https://www.reddit.com/r/golang/comments/flyvzg/design_pattern_to_avoid_hitting_the_rate_limit_on/)
- url: https://www.reddit.com/r/golang/comments/flyvzg/design_pattern_to_avoid_hitting_the_rate_limit_on/
---
I have to get the PR reviews from hundreds of repos, lets say I have these in batches of 20 PRs for ever repo

like this 

    {RepoA: #1, #2, ..........#20}
    {RepoA: #21, #22, ..........#40}
    .
    .
    {RepoZ: #1, #2, ..........#20}

Something like this, and I want to run goroutines for all these different batches, but the Github API limit is 5000 requests per hour

Is there a design pattern that would apply to this problem to solve it in an elegant way,

also should I launch the goroutines for a specific number of batches and use the rest later or should I run all the goroutines and pause them when I hit the rate limit and resume them later on once a certain amount of time has passed?

I know for server-side APIs there are patterns like token bucket and leaky bucket and all but I am a client of these APIs is there something similar for throttling/regulating the frequency of requests from my application?
## [5][Why are 2 arrays of equivalent interface types not assignable to each other](https://www.reddit.com/r/golang/comments/fm682s/why_are_2_arrays_of_equivalent_interface_types/)
- url: https://www.reddit.com/r/golang/comments/fm682s/why_are_2_arrays_of_equivalent_interface_types/
---
I'm not sure why this doesn't compile
    
    package main
    
    import "fmt"
    
    type FromLibA interface {
    	Something()
    }
    
    type FromLibB interface {
    	FromLibA
    }
    
    type FromLibBType1 struct {
    }
    
    func (*FromLibBType1) Something() {}
    
    func main() {
    	var a FromLibB
    	var b *FromLibBType1
    	a = b // works
    
    	var arr []FromLibB
    	var brr []*FromLibBType1
    	arr = brr // doesn't work
    
    	fmt.Println("%v, %v, %v, %v", a, b, arr, brr)
    }

Is there any way to make it work other than for loop over the array to copy it?
## [6][Anyone know of a Go discord?](https://www.reddit.com/r/golang/comments/fm5wvs/anyone_know_of_a_go_discord/)
- url: https://www.reddit.com/r/golang/comments/fm5wvs/anyone_know_of_a_go_discord/
---
I want to be more immersed in the go community anyone know of a discord. Also looking for any development discord’s in general 
Thanks in advance
## [7][gopls 0.3.4 update release notes](https://www.reddit.com/r/golang/comments/fln72e/gopls_034_update_release_notes/)
- url: https://github.com/golang/go/issues/33030#issuecomment-601280048
---

## [8][buffalo database](https://www.reddit.com/r/golang/comments/fm4zv5/buffalo_database/)
- url: https://www.reddit.com/r/golang/comments/fm4zv5/buffalo_database/
---
Hi, guys. I try with Buffalo. But I have some problems at start. How I must create migrations?

I tried to generate and I have empty fizz files. And then in database I have just table schema\_migration without my tables. Can someone help me with this problem, please?

I will be very grateful. 

Thanks for attention
## [9][Testing a change in an imported package (newbie)](https://www.reddit.com/r/golang/comments/fm4umi/testing_a_change_in_an_imported_package_newbie/)
- url: https://www.reddit.com/r/golang/comments/fm4umi/testing_a_change_in_an_imported_package_newbie/
---
*Warning: I am very new to golang. I hope I get the terminology correct*

I am trying to update a package `github.com/terraform-providers/terraform-provider-gitlab` but it has an imported package from `github.com/xanzy/go-gitlab`. The change I want to make requires changes in both packages and I don't want to create a pull request until I test everything. I have no issues changing and building the files in `go-gitlab` but don't have a good way to test it independently. The problem I have is building `terraform-provider-gitlab` with the changes I put into `go-gitlab`. I don't know how to tell `terraform-provider-gitlab` to use my changes. I think it has something to do with the `go.mod` file calling the specific version of `go-gitlab`. My changes only exist on my local computer and my `go-gitlab` fork. I tried to modify the files in `terraform-provider-gitlab/vendor/github.com/xanzy/go-gitlab/` but I still think it's looking somewhere else because of some errors I see. 

Example Error: `gitlab/resource_gitlab_group_cluster.go:175:40: cluster.ManagementProject undefined (type *gitlab.GroupCluster has no field or method ManagementProject)`
## [10][I'm working on a Grafana alertManager for Lark - Glark](https://www.reddit.com/r/golang/comments/flzwi2/im_working_on_a_grafana_alertmanager_for_lark/)
- url: https://www.reddit.com/r/golang/comments/flzwi2/im_working_on_a_grafana_alertmanager_for_lark/
---
This is my first project in Go, chat bot integration works well so far, I'd love some criticisms and suggestions!

[https://github.com/nehsus/glark](https://github.com/nehsus/glark)
