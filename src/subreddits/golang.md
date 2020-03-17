# golang
## [1][Go just reached 70k stars on GitHub!](https://www.reddit.com/r/golang/comments/fjvxxk/go_just_reached_70k_stars_on_github/)
- url: https://imgur.com/7OC82j7
---

## [2][Get current coronavirus case data using Go](https://www.reddit.com/r/golang/comments/fjzutb/get_current_coronavirus_case_data_using_go/)
- url: https://github.com/montanaflynn/covid-19
---

## [3][Kafka implemented in Golang with built-in coordination (No ZK dep, single binary install, Cloud Native)](https://www.reddit.com/r/golang/comments/fk1367/kafka_implemented_in_golang_with_builtin/)
- url: https://github.com/nash-io/jocko
---

## [4][Does anyone have an example/reference using MongoDB Change Streams w/ Websockets?](https://www.reddit.com/r/golang/comments/fk55ew/does_anyone_have_an_examplereference_using/)
- url: https://www.reddit.com/r/golang/comments/fk55ew/does_anyone_have_an_examplereference_using/
---
I'm working on a project where I want to implement realtime functionality for a task management app. I'm completely fine on the front-end side of things, however, I'm struggling to conceptualize how this will work server-side. The only real examples I've found online utilize a 3rd party service which is unnecessary. The overall concept is:

&amp;#x200B;

* All current clients connect to a WS endpoint to receive changes (Vue + Vuex for the state, but this really isn't important for my question)
* A client creates a new task POSTing to an API endpoint that writes to Mongo -&gt; triggers change stream
* All *other* clients receive the changes over WS and the app's state is mutated to reflect those changes.

Does anyone have an example of how to do this? Thanks in advance!
## [5][go-git has a new home! v5 and some explanations](https://www.reddit.com/r/golang/comments/fjiye7/gogit_has_a_new_home_v5_and_some_explanations/)
- url: https://github.com/go-git/go-git/wiki/go-git-has-a-new-home!-v5-and-some-explanations
---

## [6][Sending data from sever to client](https://www.reddit.com/r/golang/comments/fk1uw1/sending_data_from_sever_to_client/)
- url: https://www.reddit.com/r/golang/comments/fk1uw1/sending_data_from_sever_to_client/
---
Hello i am designing a matchmaking sever for my game. I click on find match client ip is sent to server. Match is found. (here is where i don't know how to implement) how to send client the result of matchmaking? (result of matchmaking is an ip)
## [7][Different behavior based on number of arguments consumed](https://www.reddit.com/r/golang/comments/fk1m44/different_behavior_based_on_number_of_arguments/)
- url: https://www.reddit.com/r/golang/comments/fk1m44/different_behavior_based_on_number_of_arguments/
---
I am looking at type assertions and it looks like the behavior is different based on whether you save the value to one variable, or two variables.

&amp;#x200B;

Is this something that's specific to type assertions, or are there other cases in which we can have this behavior ?

&amp;#x200B;

Sample code: [https://play.golang.org/p/upnwMWeqdTk](https://play.golang.org/p/upnwMWeqdTk)
## [8][How to log errors without exiting the program?](https://www.reddit.com/r/golang/comments/fk3l9n/how_to_log_errors_without_exiting_the_program/)
- url: https://www.reddit.com/r/golang/comments/fk3l9n/how_to_log_errors_without_exiting_the_program/
---
I am running a task every 5 minutes and I want to log the error (if it faced the error) without exiting the program but the problem is \`log.Fatal()\` is exiting the program and \`log.Panic()\` will call \`panic()\` which again exits the program.  
How can I log the error without exiting the program?
## [9][Processing in the console](https://www.reddit.com/r/golang/comments/fk0p5g/processing_in_the_console/)
- url: /r/processing/comments/fjr0u6/processing_in_the_console/
---

## [10][Trying to convert Python to Golang](https://www.reddit.com/r/golang/comments/fjybos/trying_to_convert_python_to_golang/)
- url: https://www.reddit.com/r/golang/comments/fjybos/trying_to_convert_python_to_golang/
---
I'm new to golang and trying to use it for networking and found some Python code which is doing what I'm trying to do but uses a hash table in a way I haven't seen before.

This is the initialization

    ip_dict = {}

This is the code I'm having trouble understanding

    src_ip_addr_str = socket.inet_ntoa(ip_hdr.src)
    dst_ip_addr_str = socket.inet_ntoa(ip_hdr.dst)
    
    if syn_flag and not ack_flag:
        if src_ip_addr_str in ip_dict:
            ip_dict[src_ip_addr_str]['SYN'] += 1
        else:
            ip_dict[src_ip_addr_str] = {'SYN':1, 'SYN-ACK':0}

I'm having trouble understanding the code in the 2nd if-statement and the else statement. 

How is it possible to interact with a hash table in 2 seemingly different formats? The first hash table indexing appears to need \['SYN'\] at the after the source's ip address string, \[src\_ip\_addr\_str\] whereas in the else-statement, no \['SYN'\] is needed and the input isn't just an integer.

I'm not sure how to initialize a hash table to do this. I initially tried this:

    var synMap = make(map[string]int)
    ...
    if tcp.SYN == true &amp;&amp; tcp.ACK == false {
        srcIPString := ip.SrcIP.String()
        value, ok := synMap[srcIPString]
        if ok {
    	value["SYN"]++
        } else {
    	synMap[srcIPString] = {"SYN": 1, "SYN-ACK": 0}
        }
    }

but am getting the error for the line "synMap\[srcIPString\] = {"SYN": 1, "SYN-ACK": 0}":  

syntax error: unexpected {, expecting expression

Can someone please explain what these 2 different lines are doing and how it's possible to interact with the same hash table in this way? I'm trying to convert these 2 lines to golang and having a lot of trouble.
