# golang
## [1][Golang's repo displayed using Onefetch](https://www.reddit.com/r/golang/comments/jdek8w/golangs_repo_displayed_using_onefetch/)
- url: https://i.redd.it/ltz7slxd6ut51.png
---

## [2][Which Udemy course to purchase?](https://www.reddit.com/r/golang/comments/jdd7n2/which_udemy_course_to_purchase/)
- url: https://www.reddit.com/r/golang/comments/jdd7n2/which_udemy_course_to_purchase/
---
As the title says. I'm kinda bad at reading programming books, get distracted and bored very easily, but videos do very well for me.  Heard people say bad stuff about Go tutorials on udemy, especially Todd McLeod. Is it really that bad? I wanted to pick a good tutorial and practice on codewars.
## [3][go routine problem](https://www.reddit.com/r/golang/comments/jd96zc/go_routine_problem/)
- url: https://www.reddit.com/r/golang/comments/jd96zc/go_routine_problem/
---
Can someone tell me what I'm doing wrong?  Currently this script looks like its just hanging and never returning.  I've tested the snmp, and ip iteration independently and they worked. Once I introduced the go routine it looks like I've broken something. Can someone offer assistance as to what I've done wrong. Also I just read there is a better way of limiting go routines but in a pinch I just wanted to keep it under a thousand go routines.

    package main
    
    import (
    	"encoding/binary"
    	"fmt"
    	"log"
    	"net"
    	"sync"
    
    	g "github.com/gosnmp/gosnmp"
    )
    
    const snmpComm string = "public"
    
    func poller(ip string, wg *sync.WaitGroup) {
    	defer wg.Done()
    	// Default is a pointer to a GoSNMP struct that contains sensible defaults
    	// eg port 161, community public, etc
    	g.Default.Target = ip
    	g.Default.Community = snmpComm
    
    	err := g.Default.Connect()
    	if err != nil {
    		log.Fatalf("Connect() err: %v", err)
    	}
    	defer g.Default.Conn.Close()
    
    	oids := []string{"1.3.6.1.2.1.1.1.0", "1.3.6.1.2.1.1.5.0"}
    	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
    	if err2 != nil {
    		log.Fatalf("Get() err: %v", err2)
    	}
    
    	for i, variable := range result.Variables {
    		fmt.Printf("%d: oid: %s ", i, variable.Name)
    	}
    }
    
    func main() {
    	var wg sync.WaitGroup
    
    	_, ipv4Net, err := net.ParseCIDR("10.201.0.0/16")
    	if err != nil {
    		fmt.Println(err)
    	}
    
    	// convert IPNet struct mask and address to uint32
    	// network is BigEndian
    	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
    	start := binary.BigEndian.Uint32(ipv4Net.IP)
    
    	// find the final address
    	finish := (start &amp; mask) | (mask ^ 0xffffffff)
    
    	// loop through addresses as uint32
    	threadProtector := 0
    	for i := start; i &lt;= finish; i++ {
    		if threadProtector &lt;= 999 {
    			// convert back to net.IP
    			ip := make(net.IP, 4)
    			binary.BigEndian.PutUint32(ip, i)
    			ipstr := ip.String()
    			wg.Add(1)
    			go poller(ipstr, &amp;wg)
    			threadProtector++
    		} else if threadProtector == 1000 {
    			wg.Wait()
    			threadProtector = 0
    		}
    	}
    
    }
## [4][How do you deal with microservice versioning when using a message bus?](https://www.reddit.com/r/golang/comments/jdbo1i/how_do_you_deal_with_microservice_versioning_when/)
- url: https://www.reddit.com/r/golang/comments/jdbo1i/how_do_you_deal_with_microservice_versioning_when/
---
Hey all,

Got a front end API that sends messages to back end services. Right now, the message bus uses just string values as the topic name to send something on, or subscribe to. Also, some services can pub/sub with other services as well, not just the API services. With that in mind, do you put version values within the topic name string?

As an example, say I have a simple Email service, version 1. It only handles a TO and a message body. Json like {to: '...@...com', msg: 'blah blah'}.  Now, I get a more robust email service I can use, that can now support CC, html body, subject, etc. I also decide msg is not a good key.. so I modify the JSON payload to something like {to:..., subject:..., body:... cc:...}. Clearly the first version of my email service wont decode the new format correctly. But if I change the current service to the new format, then services still sending the old format wont work either. So.. in the case of event bus with topic strings... is the way to go to spin up a 2nd (new email service) while the 1st one still runs.. so that services now ready to use the new 2nd version, can work (using a new topic string.. either with v2 in it, or some new name like NewEmailService) and the current services using the old email service continue to work?
## [5][Type analysis for annotated empty interface{}s](https://www.reddit.com/r/golang/comments/jdfis4/type_analysis_for_annotated_empty_interfaces/)
- url: https://github.com/siadat/intertype#demo-1
---

## [6][My project for manage and run acceptance tests](https://www.reddit.com/r/golang/comments/jdf9k6/my_project_for_manage_and_run_acceptance_tests/)
- url: https://www.reddit.com/r/golang/comments/jdf9k6/my_project_for_manage_and_run_acceptance_tests/
---
Today i release 1.0.0-beta version of [at2k](https://github.com/ilya-mezentsev/acceptance-testing-2k) - project for manage and run acceptance tests.

It allows you write something like this
```
BEGIN
    createUserResponse = CREATE USER {"name": "Jim"}

    ASSERT createUserResponse.status EQUALS ok

    userResponse = GET USER ${createUserResponse.hash}

    ASSERT userResponse.status EQUALS ok
    ASSERT userResponse.data.name EQUALS Jim
    ASSERT userResponse.data.hash EQUALS ${createUserResponse.hash}
END

BEGIN
    getNotExistsUserResponse = GET USER not-exists-hash

    ASSERT getNotExistsUserResponse.status EQUALS error
    ASSERT getNotExistsUserResponse.data EQUALS user-not-found
END
```

To test code like this
```
users.get('/:hash', (req, res) =&gt; {
    const user = usersRepository.find(u =&gt; u.hash === req.params.hash);

    if (user) {
        res.status(200).send({
            status: 'ok',
            data: user
        });
    } else {
        res.status(200).send({
            status: 'error',
            data: 'user-not-found'
        });
    }
});

users.post('/', (req, res) =&gt; {
    const { name } = req.body;
    const hash = uuid4();

    usersRepository.push({
        hash, name
    });
    res.status(200).send({status: 'ok', hash});
});
```

I'm quite new in releasing projects (and posting on reddit ;-)). Feel free to review code or write comments here - I'll be glad to get feedback)

P.S. Web version is available [here](https://at2k-beta.site/). Please, don't overuse it, because server hasn't have enough of resources to handle "loads" of requests
## [7][Glow: Markdown reader for the terminal with a TUI and encrypted cloud stash](https://www.reddit.com/r/golang/comments/jcxanh/glow_markdown_reader_for_the_terminal_with_a_tui/)
- url: https://github.com/charmbracelet/glow
---

## [8][validating v2 (a Go library for validating structs and fields) is released](https://www.reddit.com/r/golang/comments/jcv6gt/validating_v2_a_go_library_for_validating_structs/)
- url: https://github.com/RussellLuo/validating
---

## [9][Gosaic - 4yrs but still a piece of art](https://www.reddit.com/r/golang/comments/jcrmou/gosaic_4yrs_but_still_a_piece_of_art/)
- url: https://github.com/atongen/gosaic
---

## [10][Wrap function definition over 2 lines](https://www.reddit.com/r/golang/comments/jde5mo/wrap_function_definition_over_2_lines/)
- url: https://www.reddit.com/r/golang/comments/jde5mo/wrap_function_definition_over_2_lines/
---
Hi all,

Hopefully a simple one. Is it possible to wrap a function definition over 2 lines. I am working on a project that likes lines to be hard wrapped at around 80 characters. The hard limit is 100 characters but 80 is preferable. Is it possible to wrap a go function definition over 2 lines? Something like:

```
func getEntityByShortId(el *openpgp.EntityList , id uint32)
    (entity *openpgp.Entity, err error) {
......
}
```

This doesn't work. Am I missing something or is it simply not possible?

Thanks in advance
