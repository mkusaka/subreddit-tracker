# golang
## [1][Best Server Code Generator for OpenAPI specifications in GoLang?](https://www.reddit.com/r/golang/comments/ieyjum/best_server_code_generator_for_openapi/)
- url: https://www.reddit.com/r/golang/comments/ieyjum/best_server_code_generator_for_openapi/
---
What is the best code generator out there for Golang REST APIs

I used the oapi-codegen utility but it often has trouble parsing schemas that go-swagger usually accepts. I am looking for a more stable solution. What does the community prefer?
## [2][Hacktoberfest is coming](https://www.reddit.com/r/golang/comments/iekqx8/hacktoberfest_is_coming/)
- url: https://www.reddit.com/r/golang/comments/iekqx8/hacktoberfest_is_coming/
---
I like hacktoberfest. I want to focus on Go projects this year, though - kind of my new year's resolution: "more Go. less Node/Python/Bash". 

I have _some_ projects in mind, but I'm siloed, of course (I only know the ones I know). Does anyone want to chime in with young, interesting Go projects that could use some help?
## [3][Use single struct for both API payload and DB?](https://www.reddit.com/r/golang/comments/ietxlo/use_single_struct_for_both_api_payload_and_db/)
- url: https://www.reddit.com/r/golang/comments/ietxlo/use_single_struct_for_both_api_payload_and_db/
---
Hey all,

So I am building my API payload struct (to match my json schema which I use for OpenAPI definition). So think of a typical User:

    type User struct {
    	Id             *uuid.UUID `json:"id"`
    	Name           string     `json:"name"`
    	Username       string     `json:"username"`
    	Password       string     `json:"password"`
    	Email          string     `json:"email"`
    	Role           string     `json:"role"`
    	LastSignedInAt *time.Time `json:"lastLogin"`
    	CreatedAt      *time.Time `json:"createdAt"`
    	UpdatedAt      *time.Time `json:"lastUpdated"`
    	Active         bool       `json:"active"`
    }

This basically matches the columns in my DB. However, upon login, I only return the name, email and role as part of the login response payload. I include the Id in the JWT token claims, but otherwise some of the data is really for internal only use.

Now, I know I can use the omit option or - option for not returning keys without values, etc.. that's fine. What I am hoping though, is that I can use the same structure to store (or retrieve) data from the DB as well. I am using Postgresql, pgx, and for retrieving data, the pgxscan library to scan in to the structure.

What I want to avoid is the old school Java DAO and DTO process.. where by we had to essentially copy the same data from one object to another... I want to "pass thru" the struct instance that is un-marshaled from say a request that includes JSON to the DB code so that it can write the struct to the DB.. or vice versa.. on a GET or GET ALL, it scans into a struct like above, and can then just "null out" some of the properties that shouldn't go back as part of the response.

I think this is doable.. but more so.. is it "good Go code" to do this.. or is it frowned upon. Still a noob with Go, so I want to avoid building this, then have some expert Go dev come in one day and be like WTF is this crap.. bad bad code. 

Thanks Go pros.
## [4][Float prints 0.0000](https://www.reddit.com/r/golang/comments/ieywov/float_prints_00000/)
- url: https://www.reddit.com/r/golang/comments/ieywov/float_prints_00000/
---
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	var framerate float32 = 1/24
    
    	temp := fmt.Sprintf("%.6f", framerate)
    
    	fmt.Println(temp)
    }

Why does this code print `0.000000`?
## [5][how many clients are connected to the websocket server ?](https://www.reddit.com/r/golang/comments/if2ko5/how_many_clients_are_connected_to_the_websocket/)
- url: https://www.reddit.com/r/golang/comments/if2ko5/how_many_clients_are_connected_to_the_websocket/
---
so i have a web socket server

every one in the front end should see a dot in his mouse position &amp; other "players " dots

the problem is how to know which is which

and how to send data to all players to update position ..

i am using the [golang.org/x/net/websocket](https://golang.org/x/net/websocket)  
edit : also how do i know when someone new connects to the websocket   

## [6][Minimalistic package to scan CSV rows into struct types](https://www.reddit.com/r/golang/comments/if1ghy/minimalistic_package_to_scan_csv_rows_into_struct/)
- url: https://pkg.go.dev/github.com/artyom/csvstruct
---

## [7][Client with Logger or without?](https://www.reddit.com/r/golang/comments/if0alh/client_with_logger_or_without/)
- url: https://www.reddit.com/r/golang/comments/if0alh/client_with_logger_or_without/
---
Hey Gophers,

I have a question which is a bit more theoretical:

At work we have a fine piece of Software which does a lot with data (no surprise so far). Lately I got a Client Library for another piece of Software to which we want to send data via the client. And here comes my question:

&amp;#x200B;

The client would like to have a Logger in its Constructor (sure as an Interface with two methods). Should a client library use a logger and if so, how should it do?

A:  Should it receive the logger from the using service (this means the already existing logger should support the interface which is not in our case) ?

B: Should we totally skip logging in the client library and just return errors ?

C: Should the client define its own logger and only be configured with LogLevel and Output ?

&amp;#x200B;

Thanks in advance for all your answers

cheers Jan
## [8][CGO problem: C call go function, and passing c function pointer as the parameter](https://www.reddit.com/r/golang/comments/if0a10/cgo_problem_c_call_go_function_and_passing_c/)
- url: https://www.reddit.com/r/golang/comments/if0a10/cgo_problem_c_call_go_function_and_passing_c/
---
I am trying to build golang as a static library for my C project using CGO. Everthing is fine so far, But I am having a problem passing C function to golang function as a parameter. I am not sure whether am I doing it correctly, or it is prohibited to do so?

in golang file

    //export go_example_callback
    func go_example_callback(pfoo unsafe.Pointer, p1 C.int) {
        foo := (*func(C.int))(pfoo)
    
        // call the c function pointer
        (*foo)(5)
    
        // would also love to store pfoo as global variables to use it in anywhere
    }

and then build as a static library using this command below:

    go build -o awesome.dylib -buildmode=c-shared awesome.go

link the static library in my C project, and then trying to use go\_example\_callback

in my C project 

    void pfoo(int x) {
        printf(x);
    }
    int main(void) {
        go_example_callback(&amp;pfoo, 5);
        // expect printing 5 in the console, but I am getting error instead
    }

It will be nice if I am able to call C function in golang through this way.  Is it possible to do so? If yes, how shld I do it correctly? Many thanks for the help!
## [9][How do you organize the instrumentation code?](https://www.reddit.com/r/golang/comments/iek9wx/how_do_you_organize_the_instrumentation_code/)
- url: https://www.reddit.com/r/golang/comments/iek9wx/how_do_you_organize_the_instrumentation_code/
---
Hi, I've recently started working with opentelemetry to instrument go application, which makes the functions really big as there is a lot of logic related to instrumentation(adding events/attributes/starting/finishing spans/traces). I wonder if there is a good design pattern/template for instrumenting code outside of the method's logic or somehow the size of the functions small?
## [10][Sharef](https://www.reddit.com/r/golang/comments/iep395/sharef/)
- url: https://www.reddit.com/r/golang/comments/iep395/sharef/
---
Hi All,

Just want to present this tool called [sharef](https://github.com/emiraganov/sharef), maybe someone find it useful also\*\*.\*\* It is command line tool that allows sharing p2p files over [WebRTC.](https://WebRTC.It) It is totally written in GO. Although there are some similar approaches, I just wanted to make tool more feature rich and easier to use. For me, it helps to move things around easier specially in some complex cloud setups. It is in some early stage. I have some more ideas around it, but we will see.

Any feedback is welcome.
