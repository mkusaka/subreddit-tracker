# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][How I Structure Go Packages](https://www.reddit.com/r/golang/comments/hm34kq/how_i_structure_go_packages/)
- url: https://bencane.com/stories/2020/07/06/how-i-structure-go-packages/
---

## [3][[How to?] Design a Chat application](https://www.reddit.com/r/golang/comments/hm2ck3/how_to_design_a_chat_application/)
- url: https://www.reddit.com/r/golang/comments/hm2ck3/how_to_design_a_chat_application/
---
I am developing a synced and persists (like Messenger, Telegram where messages will be synced with server, and be available on all the devices and be available from new devices when logging in, unlike WhstsApp) chat application, with basic functionalities like one-to-one chat, group chat, ability to delete a message, ability to delete a message for everyone and ability to edit a message. 
I've developed the basic websocket for one-to-one and group chat, however, I'm not sure how to develop the following features:

****##NOTE: If you don't wanna read the whole thing, jump to point 3, as 3 is the most important question.****

1. "Delete for me": In case of group or one to one chat, a person can delete a certain message and it will be deleted for all his synced devices. Even if he logs in using a new devices, all his old chat will be fetched and displayed there, apart from the deleted message. The other parties of the chat, however, will still have the message, and it will be displayed from their devices.
The naive approach will be to maintain a separate chat table for each person, as by doing that deletion will be really easy.
But the issue will be that for a group of 500 people, every message will be duplicated 500 times in the server!
To avoid that, another approach might be to have a centralized message table and add a "visibleTo" field to message, which will by default have ids of all the recipients, but whenever an user deletes a message, the visibleTo field will remove his id. I think second approach will be better. What do you think?
2. Delete for everyone / edit will simply reach out to the server, and edit or delete the message.

3. Now, the most important question is syncing.
Suppose the client is offline. Another user send few new messages or edit or deletes a message. After hours the client comes online. How to reflect these changes there? Obviously it'll not be optimal to always fetch all the messages and update the client. How to do it efficiently? Also, a client sends few messages or edits / deletes few messages while being online. How to reflect these changes in the server, and also to other recipients of the messages?

Any suggestion, idea is welcomed!

Thanks in advance!
## [4][Is there a package that multiplex based on string pattern but not for http router?](https://www.reddit.com/r/golang/comments/hlxxtm/is_there_a_package_that_multiplex_based_on_string/)
- url: https://www.reddit.com/r/golang/comments/hlxxtm/is_there_a_package_that_multiplex_based_on_string/
---
I want a switch statement based on a certain named params that execute an arbitrary function. Preferably using {} pattern so that it is easier for users, instead of regex patterns.

Basically what an http router does but not for handling http requests.

Edit: I ended up making [my own library]( https://github.com/didip/switcheroo ) as I wanted something simple and lightweight.
## [5][I’m writing a lossless compression algorithm. Is there anyway I can write individual bits to a file in go?](https://www.reddit.com/r/golang/comments/hm6yd4/im_writing_a_lossless_compression_algorithm_is/)
- url: https://www.reddit.com/r/golang/comments/hm6yd4/im_writing_a_lossless_compression_algorithm_is/
---

## [6][Please help me with eazye](https://www.reddit.com/r/golang/comments/hm6ulp/please_help_me_with_eazye/)
- url: https://www.reddit.com/r/golang/comments/hm6ulp/please_help_me_with_eazye/
---
I find package eazye ([github](https://github.com/jprobinson/eazye)) to simple read email from outlook server

Found it in [this subreddit](https://www.reddit.com/r/golang/comments/3eqy9y/receiving_emails_in_go/)

Here is my code:

    func main() {
    	info:=eazye.MailboxInfo{}
    	info.Host= "serveraddr:26"
    	info.TLS=false
    	info.User= "username@domain.com"
    	info.Pwd= "pass"
    	info.Folder= "Inbox"
    	ch,err:=eazye.GenerateAll(info,false,false)
    	if err!=nil{
    		log.Println(err)
    	}
    	log.Println(&lt;-ch)
    }

I get in out:

    2020/07/06 15:29:34 uid search failed: imap: bad response tag ("220 serveraddr ESMTP Company")
    2020/07/06 15:29:34 {{&lt;nil&gt; &lt;nil&gt; [] 0001-01-01 00:00:00 +0000 UTC   [] [] false 0} &lt;nil&gt;}

Am I stupid or I use lib wrong?  
Thanks.
## [7][Usecase-Survey](https://www.reddit.com/r/golang/comments/hm6qnv/usecasesurvey/)
- url: https://www.reddit.com/r/golang/comments/hm6qnv/usecasesurvey/
---
Hi all!

&amp;#x200B;

My current work-environment is heavily based on Java and Spring and I'd like to convince my fellow coworkers that Go might be a great option for some of our use cases.

Therefore I've created a short survey with the goal to collect some common use-cases.

[https://forms.gle/CY5BaE23hpdYwFrx5](https://forms.gle/CY5BaE23hpdYwFrx5)

&amp;#x200B;

It would be very nice of you to take 5 minutes to participate, the results will be published here:

[https://github.com/ledex/go-survey-results](https://github.com/ledex/go-survey-results)

&amp;#x200B;

Thanks in advance!
## [8][Building Go Services with DDD Approach / Eddy Kiselman](https://www.reddit.com/r/golang/comments/hlj8dc/building_go_services_with_ddd_approach_eddy/)
- url: https://youtu.be/YfLPZOpJQjY
---

## [9][Case Study - How a startup in Fintech increases throughput by 62% doing Load Testing](https://www.reddit.com/r/golang/comments/hm5yup/case_study_how_a_startup_in_fintech_increases/)
- url: https://www.reddit.com/r/golang/comments/hm5yup/case_study_how_a_startup_in_fintech_increases/
---
I wrote a blog post on how I helped a Fintech SaaS startup increase throughput on their platform while lowering their costs at the same time using Rungutan.

https://rungutan.com/blog/startup-fintech-increase-throughput-load-testing/

FULL DISCLOSURE: I'm the CTO of Rungutan.
## [10][Broadcast: Send on one channel and receive on many. An example of a generic library using reflect.](https://www.reddit.com/r/golang/comments/hm5iuw/broadcast_send_on_one_channel_and_receive_on_many/)
- url: https://github.com/malcsm/broadcast
---

## [11][Mocking your SQL database in Go tests has never been easier.](https://www.reddit.com/r/golang/comments/hlhj8v/mocking_your_sql_database_in_go_tests_has_never/)
- url: https://github.com/cockroachdb/copyist
---

