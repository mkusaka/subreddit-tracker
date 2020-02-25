# golang
## [1][My favorite birthday present! My friends do know me well! 💕](https://www.reddit.com/r/golang/comments/f94oza/my_favorite_birthday_present_my_friends_do_know/)
- url: https://i.redd.it/rt7bsyim00j41.jpg
---

## [2][This one for me](https://www.reddit.com/r/golang/comments/f98kz4/this_one_for_me/)
- url: https://i.redd.it/k1wqoa0ow1j41.jpg
---

## [3][GoUp. Pure Go tool, mainly designed to be used in CI to verify if any of your go.mod dependencies are outdated. Zero call to Go or Git command line tools. Major, major with minor and patch version are supported.](https://www.reddit.com/r/golang/comments/f99lee/goup_pure_go_tool_mainly_designed_to_be_used_in/)
- url: https://github.com/rvflash/goup
---

## [4][Minimalist Go Logger](https://www.reddit.com/r/golang/comments/f8xb7h/minimalist_go_logger/)
- url: https://www.reddit.com/r/golang/comments/f8xb7h/minimalist_go_logger/
---
Hello,

Please check out my super simple logging library and its colorful log viewing tool. This is my first open source project. All reviews and contributions are welcome.

Links:

[https://github.com/ermanimer/go-logger](https://github.com/ermanimer/go-logger)

[https://github.com/ermanimer/go-log-viewer](https://github.com/ermanimer/go-log-viewer)

&amp;#x200B;

[sample output with github.com\/ermanimer\/go-log-viewer](https://preview.redd.it/pey8qzyfixi41.png?width=654&amp;format=png&amp;auto=webp&amp;s=c37008972e1995b6f7aef9425798337a53a528c9)
## [5][Replace X509/ASN1 parser in http.Client](https://www.reddit.com/r/golang/comments/f9954v/replace_x509asn1_parser_in_httpclient/)
- url: https://www.reddit.com/r/golang/comments/f9954v/replace_x509asn1_parser_in_httpclient/
---
Hi, is there are way to replace the ASN1 parser as it rejects the server certificate due to ‘PrintableString contains invalid character’

The certificate can’t be changed, and the service can be connected to ok from other languages - the go parser is just more strict on the acceptable characters (the cert has things like underscore characters in the common name).

Any option in the transport that could be used for this?
## [6][Golang is not good for Fuchsia](https://www.reddit.com/r/golang/comments/f92nkc/golang_is_not_good_for_fuchsia/)
- url: https://fuchsia.googlesource.com/fuchsia/+/refs/heads/master/docs/project/policy/programming_languages.md
---

## [7][giu (rapid GUI framework) v0.3 is released on github.](https://www.reddit.com/r/golang/comments/f9a43x/giu_rapid_gui_framework_v03_is_released_on_github/)
- url: https://www.reddit.com/r/golang/comments/f9a43x/giu_rapid_gui_framework_v03_is_released_on_github/
---
[github.com/AllenDang/giu](https://github.com/AllenDang/giu)

giu is a cross platform rapid GUI framework for golang based on Dear ImGui.

The major feature of this release is SplitLayout which could create live resizing layout in few lines of code.

    package main
    
    import (
      g "github.com/AllenDang/giu"
    )
    
    func loop() {
      g.SingleWindow("splitter", g.Layout{
        g.SplitLayout("Split", g.DirectionHorizontal, true, 200,
          g.Layout{
            g.Label("Left panel"),
            g.Line(g.Button("Button1", nil), g.Button("Button2", nil)),
          },
          g.SplitLayout("Right panel", g.DirectionVertical, true, 200,
            g.Layout{},
            g.SplitLayout("HSplit", g.DirectionHorizontal, true, 200,
              g.Layout{},
              g.SplitLayout("VSplit", g.DirectionVertical, true, 100,
                g.Layout{},
                g.Layout{},
              ),
            ),
          ),
        ),
      })
    }
    
    func main() {
      wnd := g.NewMasterWindow("Splitter", 800, 600, 0, nil)
      wnd.Main(loop)
    }

The result is:

https://i.redd.it/01v7vifwl2j41.gif

&amp;#x200B;
## [8][Need some guidance with creating test](https://www.reddit.com/r/golang/comments/f97n4f/need_some_guidance_with_creating_test/)
- url: https://www.reddit.com/r/golang/comments/f97n4f/need_some_guidance_with_creating_test/
---
I am looking to automate the checks on our DNS servers and I am struggling, being new to Go.

&amp;#x200B;

Here is what I have 

&amp;#x200B;

&gt;package main  
import (  
shell "github.com/gruntwork-io/terratest/modules/shell"  
"bufio"  
"errors"  
"fmt"  
"io"  
"os"  
"os/exec"  
"strings"  
"sync"  
"syscall"  
"testing"  
"github.com/stretchr/testify/require"  
"github.com/gruntwork-io/terratest/modules/logger"  
 "testing"  
)  
type Command struct {  
Command           string // The command to run  
Args              \[\]string // The args to pass to the command  
WorkingDir        string // The working directory  
Env               map\[string\]string // Additional environment variables to set  
OutputMaxLineSize int // The max line size of stdout and stderr (in bytes)  
}  
  
&gt;  
&gt;func TestRunCommand(t, \*testing.T) {  
 cmd := Command{  
Command: "nslookup",  
Args : "www.google.com",  
WorkingDir : "/tmp",  
Env : "",  
OutputMaxLineSize : "1024",  
}  
shell.RunCommand(t, \*testing.T, cmd Command)  
}  


I get an error on the line 'shell.RunCommand(t, \*testing.T, cmd Command)', and I cannot work out why

&amp;#x200B;

I intend to use the runcommandandgetoutput to determine that the service is responding correctly.
## [9][Golang and architecture tips for larger applications](https://www.reddit.com/r/golang/comments/f8odq4/golang_and_architecture_tips_for_larger/)
- url: https://www.reddit.com/r/golang/comments/f8odq4/golang_and_architecture_tips_for_larger/
---
I have lots of experience in Java and PHP (the good kind, Laravel, Symfony). I write fully object oriented code. OO allows me to do things like:

\- Enforce a specific architecture. For example, to create a "Processor" for some event, I extend AbstractProcessor. This AbstractProcessor class has some abstract functions,  your editor will force you to implement them, giving the Abstract Code vital information that needs to work correctly, like getProcessorName() { return "SomeObjectProcessor" }, function getRepo() { return $this-&gt;someInjectedObject }. Then the abstract code can do its work. What is good with this approach is that, I only have to extent the base class. I don't have to remember anything else, abstract code is implemented one time, then whenever I create a new Child Class, I MUST provide the info, otherwise the program won't run at all.

I have a bad memory, so I usually use lots of tricks like this and that makes my life a lot easier. All these tricks use OO and Reflection on one level or other, and its either in RunTime or "tests" that check for the validity of the data. All the problems of the type "When I do that, I must remember to do a,b,c" are usually enforced by the architecture, or the tests fail. I found that this way of programming, takes more time to setup, but maintenance is 100x easier.

Now, I fucking love Go. I have rewritten some long running small tools, that use lots of parallel threads and IO, and the performance was x20 at least, CPU and RAM. All of them are fairly simple 200-500 lines of code per tool. Go is ideal for that, easy and has excellent performance. I also want to practice more for web apps in GO.

I know this is because of lack of experience with the language, of procedural languages in general, but I find writing larger applications without OO very difficult. I know I have to refresh the basics, then try to write larger things and I will find my way. I also need to study composition more, but my time is limited, so I wanted to take advantage your wisdom:

\- Did anyone had the similar issues ?

\- Do you know any book/resource that goes beyond the basics and talks more about project architecture and good practices using go ?

\- What tricks do you use to make your life easier in the future of the project ?
## [10][First, I wrote my own parser, then rewrite in ANTLR, now back to a handwritten parser.](https://www.reddit.com/r/golang/comments/f96laz/first_i_wrote_my_own_parser_then_rewrite_in_antlr/)
- url: https://www.reddit.com/r/golang/comments/f96laz/first_i_wrote_my_own_parser_then_rewrite_in_antlr/
---
I wrote one popular expression evaluation library: [https://github.com/antonmedv/expr](https://github.com/antonmedv/expr)

Here is the first implementation of the parser: [https://github.com/antonmedv/expr/blob/v1.1.4/parser.go](https://github.com/antonmedv/expr/blob/v1.1.4/parser.go)

After I was thinking about using ANTLR: [https://github.com/antonmedv/expr/blob/v1.4.5/parser/parser.go](https://github.com/antonmedv/expr/blob/v1.4.5/parser/parser.go)

But, bugs in ANTLR go bindings, lack of control of error massages, made we return back to the handwritten parser: [https://github.com/antonmedv/expr/blob/master/parser/parser.go](https://github.com/antonmedv/expr/blob/master/parser/parser.go)

Also, notice what removing all `if err != nil` in the latest parser evolution makes code readable.
