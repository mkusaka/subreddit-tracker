# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][A Go unikernel running on x86 bare metal](https://www.reddit.com/r/golang/comments/jfuusy/a_go_unikernel_running_on_x86_bare_metal/)
- url: https://github.com/icexin/eggos
---

## [3][Anyone using go-migrate with a team?](https://www.reddit.com/r/golang/comments/jfvwtg/anyone_using_gomigrate_with_a_team/)
- url: https://www.reddit.com/r/golang/comments/jfvwtg/anyone_using_gomigrate_with_a_team/
---
So we are using go-migrate and we have an issue with versions. Say dev A creates a new set of migrations (via cli) like:


    20201022081817_foobar.up.sql
    20201022081817_foobar.down.sql


At dev B creates a set of migrations in another branch like:


    20201020081917_baz.up.sql
    20201020081917_baz.down.sql


Now the migration tool cant migrate depending on whats run first.

Now dev A cant run dev B migrations, and vice-versa. This is because go-migrate only knows about the latest migration, and does not track previous ones. Its turned out to be a real PITA so handle these every time database changes happen.

TLDR

If you merge a branch with "older" migrations these cant be run with pg-migrate because it does not "know" they where not already run.

Im just wondering if there is a good way of managing these cases, or maybe another library that can track previous migrations and know what to apply when running up.
## [4][When Too Much Concurrency Slows You Down (Golang)](https://www.reddit.com/r/golang/comments/jfi21j/when_too_much_concurrency_slows_you_down_golang/)
- url: https://medium.com/@_orcaman/when-too-much-concurrency-slows-you-down-golang-9c144ca305a
---

## [5][If I use ioutil.ReadAll, and then set a struct member to some small part of that file, does the entire file remain in memory?](https://www.reddit.com/r/golang/comments/jfy9vq/if_i_use_ioutilreadall_and_then_set_a_struct/)
- url: https://www.reddit.com/r/golang/comments/jfy9vq/if_i_use_ioutilreadall_and_then_set_a_struct/
---
Once upon a time, I thought I read a post explaining why loading an entire file in memory, and then referencing parts of that file, can be dangerous, because it keeps the entire file in memory. I can't seem to find that post, so I'm asking my question here. 

Use case: I have thousands of compressed CSV files (each hundreds of MB uncompressed) in cloud storage. I created a utility function that downloads the file, unzips it, and loads it in memory, returning the entire file contents as a string. 

Then my code usually takes that string and slaps a CSV reader on it, and parses it. In some of my parsing code, I'll do something like:

`myStruct.Member = rec[0]`

Where `rec[0]` is one field from one CSV record. 

The thing that has me concerned: does this mean my program will *always* retain the entire file in memory, because I have created a reference to it in this way?

And if so, how do I copy the value without retaining a reference to the file?
## [6][My Six Years of Experience of as a Go Programming Language Mentor in India](https://www.reddit.com/r/golang/comments/jfwrav/my_six_years_of_experience_of_as_a_go_programming/)
- url: https://shijuvar.medium.com/my-six-years-of-experience-of-as-a-go-programming-language-mentor-in-india-67854dcf1b95
---

## [7][PSA: Think of the Less/less functions in the sort package as actually being ComesBefore](https://www.reddit.com/r/golang/comments/jfei3w/psa_think_of_the_lessless_functions_in_the_sort/)
- url: https://www.reddit.com/r/golang/comments/jfei3w/psa_think_of_the_lessless_functions_in_the_sort/
---
As in `less(i,j)` means *item i comes-before item j*, which I think is more helpful than "less" which doesn't really tell you where it'll end up in the list, per se, and makes it more obvious how to sort a list in "reverse" order, or any arbitrary multi-dimensional order.
## [8][Multi-source file server: serve from 15TB files on baremetal + google cloud storage (public url) via single link](https://www.reddit.com/r/golang/comments/jfv67u/multisource_file_server_serve_from_15tb_files_on/)
- url: https://github.com/codenoid/file-server
---

## [9][A Cloud Native Distributed Streaming Network Telemetry](https://www.reddit.com/r/golang/comments/jff8t1/a_cloud_native_distributed_streaming_network/)
- url: https://github.com/yahoo/panoptes-stream
---

## [10][A simple helper to restart Docker containers with newer versions of images pulled from registry.](https://www.reddit.com/r/golang/comments/jfcus1/a_simple_helper_to_restart_docker_containers_with/)
- url: https://www.reddit.com/r/golang/comments/jfcus1/a_simple_helper_to_restart_docker_containers_with/
---
Helps k8s-less provisioning and updating to the newer version of a stateless container without pains of recalling the command-line options used to start the container.

[https://github.com/jdevelop/repull](https://github.com/jdevelop/repull)
## [11][program that adds the percentage (decided by the user) of a number (decided by the user) to that number, which added several times with itself gives me another number (decided by the user) added to its percentage (the one decided at the beginning)](https://www.reddit.com/r/golang/comments/jfnt3e/program_that_adds_the_percentage_decided_by_the/)
- url: https://www.reddit.com/r/golang/comments/jfnt3e/program_that_adds_the_percentage_decided_by_the/
---
Hi, I'm a new user and (I've also just started University) and I'm already having problems coding...

I'll give an example and paste my code:

&amp;#x200B;

example:

20 (number decided) \* 15/100 (percentage decided) + 20 = 23 (the number I want to reach by adding a certain number of my choice plus the percentage decided before \[15/100\])

&amp;#x200B;

For example, let's take the number 2 as the number to add (with the for cycle):

&amp;#x200B;

2+2\*15/100=2.30 (the first print)

2+2\*15/100=2.30+2.30=4.60 (second print)

...

&amp;#x200B;

until you reach the last print that would be 23.

&amp;#x200B;

&amp;#x200B;

My try (completely wrong as I don't know how to proceed/change it in a way that would work):

    package main
    
    import "fmt"
    
    func main() {
      var n, i, accumulatore float64
      var percentuale int
      fmt.Println("")
      fmt.Print("Inserire valore: ")
      fmt.Scan(&amp;n)
      fmt.Print("Inserire percentuale (int): ")
      fmt.Scan(&amp;percentuale)
    
      aumento := n*((percentuale/100)+1)
      accumulatore = aumento
    
      fmt.Println("")
      fmt.Println(aumento)
      fmt.Println("")
      
      if (n+(n*15/100))%aumento == 0 {
        for i=0; i&lt;n; i++{
        fmt.Println(accumulatore)
        accumulatore += aumento
        for accumulatore == n/aumento {
          break
        }
      }
      
    }
    
    }

could someone kindly help me?
