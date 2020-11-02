# reactjs
## [1][Beginner's Thread / Easy Questions (November 2020)](https://www.reddit.com/r/reactjs/comments/jlwguv/beginners_thread_easy_questions_november_2020/)
- url: https://www.reddit.com/r/reactjs/comments/jlwguv/beginners_thread_easy_questions_november_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Ask about React or anything else in its ecosystem :)

Stuck making progress on your app, need a feedback?  
Still Ask away! We’re a friendly bunch 🙂

---

## Help us to help you better

1. **Improve your chances** of reply by
   1. adding minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz] links
   1. describing what you want it to do (ask yourself if it's an [XY problem](https://meta.stackexchange.com/questions/66377/what-is-the-xy-problem))
   1. things you've tried. (Don't just post big blocks of code!)
1. **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
1. **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

Check out the sub's **sidebar**! 👉  
For rules and free resources~

**Comment here for any ideas/suggestions to improve this thread**

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[usehooks.com]: https://usehooks.com/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[being wrong on the internet]: https://xkcd.com/386/
[tweet organization]: https://twitter.com/dan_abramov/status/1027245759232651270?lang=en
[get started with redux]: https://www.reddit.com/r/reactjs/wiki/index#wiki_getting_started_with_redux
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Who's Hiring? [November 2020]](https://www.reddit.com/r/reactjs/comments/jlwj1x/whos_hiring_november_2020/)
- url: https://www.reddit.com/r/reactjs/comments/jlwj1x/whos_hiring_november_2020/
---
We alternate between **Who's Hiring** (on the 1st of the month, [most recent one here][hiring:most recent]) and **Who's Available** (on the 15th, [most recent one here][available:most recent])

Welcome to **the biggest React job board in the world!** This is like Hacker News' **Who's Hiring** but just for React. Top Level comments must be **Job Opportunities.**

⚠️ NEW: WE ARE REQUESTING EVERYBODY FOLLOW [THE HN Who's Hiring FORMAT][format:hiring:hn]

**Company inc. | Job Title | City/State Location | Full-time/Part-Time | On-site/Remote | (Optional) Salary range | Website jobs page, other hard requirements etc.**

examples:

- **Thorn | San Francisco or Remote (US based) | Full-time Contract | $100k - $150k | Software Engineer | https://www.wearethorn.org/**
- **PolicyStat | Full-Stack Python+Django Software Engineer | Indianapolis, Vancouver, or REMOTE | Full Time | +\$80k**

Please include as much information as possible. **If you are remote-friendly, or open to sponsoring work visas to your country, say so! These are the top 2 questions!**

If you are looking for jobs, send a PM to the poster or post in our [Who's Available thread!][available:most recent]  
(I am sorry folks. I didn't post 'Who's Available' for October...)

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/j32odm/whos_hiring_and_rreactjs_moderator_applications/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/itrbgt/whos_available_september_2020/
## [3][In desperate need of a MERN/React engineer for help](https://www.reddit.com/r/reactjs/comments/jm7hxu/in_desperate_need_of_a_mernreact_engineer_for_help/)
- url: https://www.reddit.com/r/reactjs/comments/jm7hxu/in_desperate_need_of_a_mernreact_engineer_for_help/
---
So my co-founder just left and gave me the entirety of his code developed in the MERN stack. Not being a software/web development engineer (although I am very experienced Python and C++ developper), I had to learn the fundamentals of Node, React, MongoDB and Express to be able to re-host what he built in a new domain, set up my own AWS etc...

Everything is running fine in development mode on my machine, but when he sent me the files, he did not include his .env.production environment variable, and I can't figure out how to set it up myself. The result of this is that the React code is working perfectly on Localhost, but not on my local network. Naturally, it is not working either when I try to set it up on the domain I owe through AWS. Due to a disagreement, he's not willing to provide any help with what he considers is troubleshooting and outside the scope of our separation agreement...

This is a much too broad question for stack overflow, and I guess that talking to someone directly would be the best way to fix the problem - thus me reaching out here... Anyone willing to help with this? You'd save my life...

&amp;#x200B;

UPDATE:

To everyone who replied and/or DMed me, thank you so much for being so helpful. u/RangerCoder was extremely kind and helped me fix the issue (well the one I knew of at least LOL). I had a chat with my co-founder at the end, and he sent me the env files although they were useless since he was doing everything from the bash environment anyways. He was overall not very helpful for sure, and I honestly am quite annoyed at him, but most important is that things are (mostly) fixed now. Thanks again to everyone.  
## [4][How do you deal with null values from an api?](https://www.reddit.com/r/reactjs/comments/jmj7hl/how_do_you_deal_with_null_values_from_an_api/)
- url: https://www.reddit.com/r/reactjs/comments/jmj7hl/how_do_you_deal_with_null_values_from_an_api/
---
Hey guys, suppose I get a users info from the api that returns something like

    { 
        name: "username", 
        subscription: { status: "active/pending verification/etc.", subscription_ends: date }
     } 

but if the user is not a subscriber it returns something like 

    { 
        name: "username", 
        subscription: null
    } 

and i want to display data according to the subscription status currently I do something like 

    {data.subscription ? 
        ( data.subscription.status === "active"
            ? ( &lt;button&gt;Cancel Subscription&lt;/button&gt; ) 
            : ( &lt;button&gt;Subscribe&lt;/button&gt; ) 
        : ( &lt;button&gt;Subscribe&lt;/button&gt; ) 
    }

basically if i don't do "data.subscription  ? () : () " first then if data.subscription is null I'll get errors like cannot read data of null etc.

Is there a better way of checking for null values from the api than the one I'm currently using?
## [5][Blog App Using MERN Stack](https://www.reddit.com/r/reactjs/comments/jmkan6/blog_app_using_mern_stack/)
- url: https://medium.com/mehulkothari05/blog-app-using-mern-stack-b0b4d69d7ea1?source=friends_link&amp;sk=48181cdd4f1988f88718cbdd0ce35d49
---

## [6][Redux vs Recoil / Differences and Simple API explanation](https://www.reddit.com/r/reactjs/comments/jmmq6l/redux_vs_recoil_differences_and_simple_api/)
- url: https://www.reddit.com/r/reactjs/comments/jmmq6l/redux_vs_recoil_differences_and_simple_api/
---
Video review of Recoil JS - *strengths and weaknesses* **\[spanish\]**  

&amp;#x200B;

&amp;#x200B;

[https://youtu.be/uHTqYBg\_zkE](https://youtu.be/uHTqYBg_zkE)
## [7][I just write a RSS reader using electron + react + typescript!](https://www.reddit.com/r/reactjs/comments/jmn2uv/i_just_write_a_rss_reader_using_electron_react/)
- url: https://github.com/Saul-Mirone/homura
---

## [8][Has anyone used this folder structure in their projects?](https://www.reddit.com/r/reactjs/comments/jmiqm3/has_anyone_used_this_folder_structure_in_their/)
- url: https://www.reddit.com/r/reactjs/comments/jmiqm3/has_anyone_used_this_folder_structure_in_their/
---
I'm wondering if anyone has anyone used this folder structure in their projects?  
If you have, do you recommend it? Or if not what is your preferred folder structure?  
[https://medium.com/@alexmngn/how-to-better-organize-your-react-applications-2fd3ea1920f1](https://medium.com/@alexmngn/how-to-better-organize-your-react-applications-2fd3ea1920f1)
## [9][Trouble finding a job - what now?](https://www.reddit.com/r/reactjs/comments/jmlj16/trouble_finding_a_job_what_now/)
- url: https://www.reddit.com/r/reactjs/comments/jmlj16/trouble_finding_a_job_what_now/
---
I'm having trouble finding a job as front-end developer with react, because i have 0 commercial experience with it. I have commercial experience with HTML, CSS, jquery, basic JS, but i haven't been doing that for a while.

Here is my github with some react apps i made for practice - [https://github.com/red4211](https://github.com/red4211)

I feel really demotivated by my inability to get a job, and now i have to look for jobs outside of web development.

What do i do next? Is it actually realistic for me to hope for a job with React when i have 0 commercial experience w it? Haven't worked out for me yet.

If not, what are the jobs in the field i actually have a chance to get?
## [10][Show r/reactjs: fast arrow, a vscode extension/snippet to enable you to write arrow function in less than 1 second.](https://www.reddit.com/r/reactjs/comments/jmlcgo/show_rreactjs_fast_arrow_a_vscode/)
- url: https://www.reddit.com/r/reactjs/comments/jmlcgo/show_rreactjs_fast_arrow_a_vscode/
---
One of the thing that bugs me when writing javascript is the need to type out `() =&gt; {}`. 

That's a lot of button pressed and that's kinda annoying. This is why I created [fast arrow.](https://github.com/vinliao/fast-arrow) Now I type an arrow function really fast.

Here's a short gif to illustrate: [gif](https://raw.githubusercontent.com/vinliao/fast-arrow/master/images/faster.gif)
## [11][Why we need to abstract dependencies, a real world example with react router](https://www.reddit.com/r/reactjs/comments/jm6uoo/why_we_need_to_abstract_dependencies_a_real_world/)
- url: https://davemanton.medium.com/why-we-need-to-abstract-dependencies-a-real-world-example-with-react-router-8b06d1540264
---

## [12][A simple react shopping cart that demonstrates fundamentals of React](https://www.reddit.com/r/reactjs/comments/jmdjtu/a_simple_react_shopping_cart_that_demonstrates/)
- url: https://github.com/dilab/react-simple-cart
---

