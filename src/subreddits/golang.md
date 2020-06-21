# golang
## [1][Playing with Go schedulers on a dual-core RISC-V](https://www.reddit.com/r/golang/comments/hcycss/playing_with_go_schedulers_on_a_dualcore_riscv/)
- url: https://embeddedgo.github.io/2020/06/21/playing_with_go_schedulers_on_a_dual-core_risc-v.html
---

## [2][Simple microservice boilerplate](https://www.reddit.com/r/golang/comments/hcvm6j/simple_microservice_boilerplate/)
- url: https://www.reddit.com/r/golang/comments/hcvm6j/simple_microservice_boilerplate/
---
Here's the [repo](https://github.com/diffuse/gloss)

I find myself copying/rewriting a lot of the same boilerplate whenever I create a new microservice, so I figured I would throw some of it in a standard setup and share it with you guys.  If anyone is interested, this is a bare-bones microservice deployment with Docker, chi, and pq SQL, to save me (and hopefully you) a bit of setup when you write new services.  The example code in the handler and database query function bodies exposes a simple integer counter and serves as a placeholder for actual application code.

Fork it, clone it, stick it in a stew.  Will happily take any advice/suggestions for improvements :)

Thanks!
## [3][gRPC over WebRTC](https://www.reddit.com/r/golang/comments/hd6hsb/grpc_over_webrtc/)
- url: https://github.com/jsmouret/grpc-over-webrtc
---

## [4][Meet VERMIN The smart virtual machines manager](https://www.reddit.com/r/golang/comments/hcvg4y/meet_vermin_the_smart_virtual_machines_manager/)
- url: https://github.com/mhewedy/vermin
---

## [5][Using TinyGo to create a Thermal Detonator from Star Wars](https://www.reddit.com/r/golang/comments/hd2lk2/using_tinygo_to_create_a_thermal_detonator_from/)
- url: https://www.reddit.com/r/golang/comments/hd2lk2/using_tinygo_to_create_a_thermal_detonator_from/
---
Created a project that I enjoyed working on and wanted to share [go-thermal](https://github.com/BattleBas/go-thermal).

A [thermal detonator](https://www.youtube.com/watch?v=bStsof9p7Mc) that uses an Arduino Nano to flicker LED's and play sound.

I used a chip called dfminiplayer to play sound, there wasn't any example how to use it using TinyGo, so I had to use the official documentation to partially recreate a library for it using Go which was an interesting challenge:  [link to source](https://github.com/BattleBas/go-thermal/blob/master/dfminiplayer.go)
## [6][How to Implement scheduler task the right way for many data](https://www.reddit.com/r/golang/comments/hd22gr/how_to_implement_scheduler_task_the_right_way_for/)
- url: https://www.reddit.com/r/golang/comments/hd22gr/how_to_implement_scheduler_task_the_right_way_for/
---
Hello, I am new in scheduler task and never use this in the back end, here is my case:

I have a user that have 2 reminders to email them when the meeting begin; let say the meeting is next week, and the user want to a reminder 2 days and 1 day before the meeting to send email to them
And sometimes user want to change the reminder as well, such as user want to reminder 3 days before the meeting or a day at 10 am before meeting begin.

Should I use only Cron job here to check the meeting reminder between on the date meeting ? Should I create the default reminder meeting for the user
Let say when the user add meeting today and decrement of that meeting time to 2 days before that meeting and then I should do Cron job to check Every day that reminder of meeting need to send email by today, is that correct??
What library to make this easier ? 

Thank you for you attention :)
Hope I get some advice for you
## [7][Share a distributed job management system : Asgard](https://www.reddit.com/r/golang/comments/hcrdko/share_a_distributed_job_management_system_asgard/)
- url: https://www.reddit.com/r/golang/comments/hcrdko/share_a_distributed_job_management_system_asgard/
---
##Source code

https://github.com/dalonghahaha/Asgard

## Introduction

Asgard is a distributed job management system designed to comprehensively solve resident progress applications, plan tasks, and schedule tasks.

## Architecture design

![Architecture design](https://raw.githubusercontent.com/dalonghahaha/Asgard/master/doc/Asgard.png)

- The Asgard system consists of web nodes, master nodes, and agent nodes.

- The main functions of the web node include instance management, group management, job configuration, job running status control, job running status viewing, log query

- The master node is responsible for the status monitoring of the agent node, and at the same time receives and dumps the runtime data reported by the agent node

- The agent node receives instructions from the web node and operates related operations in the corresponding server

- Master node and agent node exchange data through grpc protocol

## Web UI preview

![Home console](https://raw.githubusercontent.com/dalonghahaha/Asgard/master/doc/page1.png)

![Instance management](https://raw.githubusercontent.com/dalonghahaha/Asgard/master/doc/page2.png)

![Task management](https://raw.githubusercontent.com/dalonghahaha/Asgard/master/doc/page3.png)

![Monitoring information](https://raw.githubusercontent.com/dalonghahaha/Asgard/master/doc/page4.png)
## [8][Golang for Application installation automation?](https://www.reddit.com/r/golang/comments/hd57nf/golang_for_application_installation_automation/)
- url: https://www.reddit.com/r/golang/comments/hd57nf/golang_for_application_installation_automation/
---
Hey Guys,

i'm looking forward to build a tool to install an enterprise application (cross platform ), setup database (oracle/microsoft sql server), setup load balancer etc . Currently my plan is to go with python . 

As i have enough time in hand so i'm exploring all the possibilities.  Is it possible to do the same with golang i.e, building an automation tool to install enterprise application (cross platform), db setup, application/webserver setup etc using golang?

&amp;#x200B;

Thanks in Advance ;)
## [9][Is there a way to implement generics on a map?](https://www.reddit.com/r/golang/comments/hd0v05/is_there_a_way_to_implement_generics_on_a_map/)
- url: https://www.reddit.com/r/golang/comments/hd0v05/is_there_a_way_to_implement_generics_on_a_map/
---
I am attempting to do a filter function on a map but I am having trouble even doing the parameters for a generic map.
If I use just filter(type T)(s T) it lets the map through but then I cannot for loop through a type T.  Yet, if I use reflection with a fmt.Println(reflect.TypeOf(s)) it gives me type map[string]string.
I have tried params (type T)(s map[T]T) but it gives the error prog.go2:8:27: invalid map key type T.  Anyone figure this out because I cannot seem to find any resources online for generic maps.


````


func filter(type T)(s T) {

	fmt.Println(reflect.TypeOf(s))

        //wont let me do this with (s T)
	//for k, v := range s {
		//fmt.Printf("%s -&gt; %s\n", k, v)
	//}

}

func main() {

	kvs := map[string]string{"a": "apple", "b": "banana"}
	filter(kvs)
}


````

Playground example https://go2goplay.golang.org/p/WQDpNvA7ndq
## [10][Worldwide - a GameBoyColor emulator written in Go (Ebiten)](https://www.reddit.com/r/golang/comments/hchnmh/worldwide_a_gameboycolor_emulator_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/hchnmh/worldwide_a_gameboycolor_emulator_written_in_go/
---
[https://github.com/Akatsuki-py/Worldwide](https://github.com/Akatsuki-py/Worldwide)

https://i.redd.it/9znnudgpm0651.gif
