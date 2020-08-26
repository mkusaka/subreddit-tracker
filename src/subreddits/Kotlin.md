# Kotlin
## [1][Shrinking a Kotlin binary by 99.2%](https://www.reddit.com/r/Kotlin/comments/igrt6a/shrinking_a_kotlin_binary_by_992/)
- url: https://jakewharton.com/shrinking-a-kotlin-binary/
---

## [2][Writing J2ME app in Kotlin](https://www.reddit.com/r/Kotlin/comments/igyt7k/writing_j2me_app_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/igyt7k/writing_j2me_app_in_kotlin/
---
I'm interested in developing a Java ME app using LWUIT. I've developed Java server-side and client-side for Android, but never used ME. Recently, I started working with Kotlin and found it saves a great deal of effort, is less verbose, and the resulting code is easier to read. 

Are there guides or examples of a J2ME app developed using Kotlin or Kotlin and Java? Alternatively, are there reasons why this might not be a great idea? I'm aware that Kotlin JAR files will be larger, and not all features (i.e. Reflection) are available.
## [3][Coroutines in Kotlin by Venkat Subramaniam](https://www.reddit.com/r/Kotlin/comments/igy18p/coroutines_in_kotlin_by_venkat_subramaniam/)
- url: https://www.youtube.com/watch?v=jT2gHPQ4Z1Q
---

## [4][[Article] Snippets | List to String with Examples | Join operations and advance use cases](https://www.reddit.com/r/Kotlin/comments/igxzfd/article_snippets_list_to_string_with_examples/)
- url: https://chetangupta.net/list-to-string
---

## [5]["Wrap" an abstract method](https://www.reddit.com/r/Kotlin/comments/igway6/wrap_an_abstract_method/)
- url: https://www.reddit.com/r/Kotlin/comments/igway6/wrap_an_abstract_method/
---
I'm sure there is probably a design pattern or standard way of doing this but I don't know what it is, so rather than invent my own, I thought I'd ask!

I have a 3rd party library which has an abstract class like this

`abstract class Job{ abstract fun run() }`

Now, I want to automatically log the start and the end of all jobs into a file / table / whatever, so what I want to do is have the person writing the job not need to worry about it - they just implement the run method and the logging gets added automatically.

I could do this by having a class "LoggedJob", which implements run and has a differently named abstract method

    abstract class LoggedJob{
      override fun run(){
        log.start()
        myRun()
        log.end()
      }
      abstract fun myRun()
    }

but is there no way for me to keep the method name "run" and add logging automatically - basically wrap the abstract method with some extra functionality (open a transaction, write to a log, whatever).

I realise AOP could probably do this but it seems overkill for such a simple requirement - is there an idiomatic way of doing this?
## [6][Anyone using Kotlin 1.4 with Springboot in Production?!](https://www.reddit.com/r/Kotlin/comments/igdg24/anyone_using_kotlin_14_with_springboot_in/)
- url: https://www.reddit.com/r/Kotlin/comments/igdg24/anyone_using_kotlin_14_with_springboot_in/
---
Just wanted to check before I start working on a project. 
Is anyone using Kotlin 1.4 with springboot  with micro-services in *production*?
How is it and what are your new findings? 
Also do you want to give me suggestions if any?

I am already using older version of Kotlin in production, thinking to use a latest one.
## [7][gRPC with Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/ig51fl/grpc_with_kotlin_coroutines/)
- url: https://codingwithmohit.com/grpc/grpc-kotlin-coroutines/
---

## [8][refactor method (apply DRY) using generic method or design pattern](https://www.reddit.com/r/Kotlin/comments/igg7a5/refactor_method_apply_dry_using_generic_method_or/)
- url: https://www.reddit.com/r/Kotlin/comments/igg7a5/refactor_method_apply_dry_using_generic_method_or/
---
 Can anyone help me? I need guidance on how to solve this problem.  
 just guide me to use a design pattern to get rid of the repeated code.  
[https://codereview.stackexchange.com/questions/248415/refactor-method-using-generic-method-or-design-pattern](https://slack-redir.net/link?url=https%3A%2F%2Fcodereview.stackexchange.com%2Fquestions%2F248415%2Frefactor-method-using-generic-method-or-design-pattern)
## [9][App developer - Co-founder/CTO](https://www.reddit.com/r/Kotlin/comments/igsrre/app_developer_cofoundercto/)
- url: https://www.reddit.com/r/Kotlin/comments/igsrre/app_developer_cofoundercto/
---
I am building a team to launch a tech Start-up. 

Product - App, category- Social Media

We are looking for a professional code/app developer, who can create and manage the app.

Remote, {until it gets big}

Role Responsibilities:- CTO

1. At the early stage - you need to create and manage the app on your own. Afterwards, you need to manage the team of developers and direct the tech department of the Start-up.

2. You can improvise the design, creativity and showcase your skills.

Perks:

1. You will be working closely with the CEO.
2. You will be a part of decision making board members.
3. You will get a co-founder position in the start-up.
4. You will get stakes &amp; shares in the start-up.
5. You will be your own Boss.

Role Requirements:

1. Commitment and time management. 
2. You will be working in an entrepreneurial environment.
3. You need to think big.

We are looking for a long term trustworthy relationship. 

All interested individuals can DM me or send Your Resume at - saurabh914921@gmail.com
## [10][I made a thing, I'm sure all Kotlin lovers will appreciate!](https://www.reddit.com/r/Kotlin/comments/iflmgn/i_made_a_thing_im_sure_all_kotlin_lovers_will/)
- url: https://i.redd.it/l5np0cey2xi51.jpg
---

