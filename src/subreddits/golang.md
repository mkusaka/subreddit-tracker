# golang
## [1][Can we talk about how user-friendly Go's time formatting is?](https://www.reddit.com/r/golang/comments/ht8m3s/can_we_talk_about_how_userfriendly_gos_time/)
- url: https://www.reddit.com/r/golang/comments/ht8m3s/can_we_talk_about_how_userfriendly_gos_time/
---
I've been working with Go for a while now, but I'd never had much cause to get into time arithmetic in Go. Coming from JavaScript, SQL, and Python, I've always avoided date-math where possible--Python makes it bearable, but SQL and JavaScript tend to be annoying. The worst with many languages is learning the format strings required to output times in whatever way you want. For example, in the dialect of SQL I use most often, formatting a DATETIME in ISO format requires a format string of "%Y-%m-%d %H-%M-%S", or "%F %T", if you happen to know about that shortcut. And remembering the difference between a capital or lowercase M has messed me up several times.

Imagine my joy when I saw how the time.Format function worked.

Seriously, if you haven't seen it, go look it up. Bravo, designers.
## [2][reflink: very small package for btrfs/xfs copy-on-write file copy in native go](https://www.reddit.com/r/golang/comments/htfdhh/reflink_very_small_package_for_btrfsxfs/)
- url: https://github.com/KarpelesLab/reflink
---

## [3][GoScope - Gin Gonic Middleware](https://www.reddit.com/r/golang/comments/htfbep/goscope_gin_gonic_middleware/)
- url: https://www.reddit.com/r/golang/comments/htfbep/goscope_gin_gonic_middleware/
---
Watch incoming requests and outgoing responses as well as logs from your Go Gin  application. All is logged into a database for persistence and paginated  for performance.

The aim of this application is to be a plug and play addition to your  application, not a hurdle, thus to setup, you only require a one-liner  in your main function.

Any addition, contribution and criticism to the project is highly welcome.

[https://github.com/averageflow/goscope](https://github.com/averageflow/goscope)

&amp;#x200B;

https://preview.redd.it/am0x43yuslb51.png?width=800&amp;format=png&amp;auto=webp&amp;s=c5679a54b5d511b05673fbad15bbc9b616552f66
## [4][[2006.09973] Breaking Type-Safety in Go: An Empirical Study on the Usage of the unsafe Package](https://www.reddit.com/r/golang/comments/ht7hyl/200609973_breaking_typesafety_in_go_an_empirical/)
- url: https://arxiv.org/abs/2006.09973
---

## [5][Wordpress in Golang?](https://www.reddit.com/r/golang/comments/htehlt/wordpress_in_golang/)
- url: https://www.reddit.com/r/golang/comments/htehlt/wordpress_in_golang/
---
Hello all,

Would it be a good idea to convert/re-write wordpress into golang? Personally when I use a self hosted wordpress, pages takes 2 to 4 seconds to load with default theme and no additional plugins. There are many tweaks that we can make to php and apache using which performance can be improved however still wordpress will get run in php.

As I understood, golang creates a compiled binary and leads to improved performance, would it be good idea to port WP to golang? In my observations, I traced latest wordpress using Xdebug, it ran over 40K functions, can its performance improve by injecting everything into binaries that will make wordpress to run as a service? (I am not saying 40K function = slow performance)

I understood there would be another story of plugins and themes but is it good idea to spend time on designing this thing?
## [6][Clivern/Beetle - Kubernetes multi-cluster deployment automation service.](https://www.reddit.com/r/golang/comments/ht0ebt/clivernbeetle_kubernetes_multicluster_deployment/)
- url: https://github.com/Clivern/Beetle
---

## [7][Interfaces and the Outside World](https://www.reddit.com/r/golang/comments/htfs67/interfaces_and_the_outside_world/)
- url: https://github.com/Evertras/go-interface-examples/tree/master/outside-world
---

## [8][New to golang trying to convert simple pyhton script to golang.](https://www.reddit.com/r/golang/comments/htf9nl/new_to_golang_trying_to_convert_simple_pyhton/)
- url: https://www.reddit.com/r/golang/comments/htf9nl/new_to_golang_trying_to_convert_simple_pyhton/
---
Hi

Im new to golang . And i am trying to learn by converting my simple python scripts. But i have no idea how to do it.

here is my python script.

    import math
    
    barlength = 400
    cutlength = int(input("what is your cut length? In cm: "))
    
    print('{} / {} = '.format(barlength, cutlength))
    print(barlength / cutlength)
    
    lengtheverybar = (barlength / cutlength)
    cutbars = math.floor(lengtheverybar)
    
    print("You will get", cutbars, "every 4 meters bar.")

Could any one help me out?
## [9][Go 1.14.6 is out now](https://www.reddit.com/r/golang/comments/hsqdak/go_1146_is_out_now/)
- url: https://www.reddit.com/r/golang/comments/hsqdak/go_1146_is_out_now/
---
[https://golang.org/dl/](https://golang.org/dl/)  
[https://golang.org/doc/go1.14](https://golang.org/doc/go1.14)  
[https://golang.org/doc/devel/release.html#go1.14](https://golang.org/doc/devel/release.html#go1.14)
## [10][What are your thoughts on using Go for REST APIs?](https://www.reddit.com/r/golang/comments/htb1gh/what_are_your_thoughts_on_using_go_for_rest_apis/)
- url: https://www.reddit.com/r/golang/comments/htb1gh/what_are_your_thoughts_on_using_go_for_rest_apis/
---
I have had the pleasure of learning Go lately and I must say that I like many aspects of the language. Type safety, go routines, native performance. It has the feel of C/C++ without all of the nonsense and I look forward to continuing to use it.

However, coming from the C#, Java, Python, and Ruby world, creating REST APIs in Go seems to be quite arduous compared to these other languages. What could be accomplished with relatively few lines of code in those languages seems to take much, much more setup and plumbing to get even a basic API up and running.

What are your thoughts and experiences on using Go as a web service development language? Am I maybe missing some helpful packages that would streamline the process a bit?
