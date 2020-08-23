# javascript
## [1][The JetBrains WebStorm team is here to answer your questions](https://www.reddit.com/r/javascript/comments/idc09e/the_jetbrains_webstorm_team_is_here_to_answer/)
- url: https://www.reddit.com/r/javascript/comments/idc09e/the_jetbrains_webstorm_team_is_here_to_answer/
---
Hi [r/javascript](https://www.reddit.com/r/javascript/)! We, the WebStorm team, are excited to announce our first AMA. We’ve never done anything like this before, but we feel the time has come to try something new.

If you’ve never heard of WebStorm, it is a JavaScript IDE by JetBrains. It comes with out-of-the-box support for lots of popular technologies and lets you do most of your development tasks right inside it. More information is available on [our website](https://www.jetbrains.com/webstorm/).

We’ll start answering your questions at 12 pm UTC on the 24th of August and will be doing this until 5 pm UTC. You can ask us about anything related to WebStorm or the JavaScript support in any other JetBrains IDEs like IntelliJ IDEA Ultimate, PhpStorm, or PyCharm Professional.

**Feel free to submit your questions ahead of time.** This thread will be used for both questions and answers.

Your questions will be answered by:

* Andrey Starovoyt (WebStorm Team Lead), [u/anstarovoyt](https://www.reddit.com/user/anstarovoyt/),
* Ekaterina Prigara (WebStorm Product Manager), [u/prigara](https://www.reddit.com/user/prigara/),
* Ekaterina Ryabukha (WebStorm Product Marketing Manager), [u/ryababukha](https://www.reddit.com/user/ryababukha),
* Piotr Tomiak (WebStorm Software Developer), [u/piotrtomiak](https://www.reddit.com/user/piotrtomiak), and
* Konstantin Ulitin (WebStorm Software Developer), [u/konstantin\_ulitin](https://www.reddit.com/user/konstantin_ulitin)
## [2][Showoff Saturday (August 22, 2020)](https://www.reddit.com/r/javascript/comments/ieh338/showoff_saturday_august_22_2020/)
- url: https://www.reddit.com/r/javascript/comments/ieh338/showoff_saturday_august_22_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [3][Transduction in JavaScript](https://www.reddit.com/r/javascript/comments/if0qqk/transduction_in_javascript/)
- url: https://medium.com/weekly-webtips/transduction-in-javascript-fbe482cdac4d
---

## [4][A website that lets users upload and draw their own fourier epicycles. A type of drawing which is created by taking the fourier transform of an image. It also gives a brief explanation of the mathematics connecting fourier series and revolving epicycles.](https://www.reddit.com/r/javascript/comments/ien5a7/a_website_that_lets_users_upload_and_draw_their/)
- url: https://www.myfourierepicycles.com/
---

## [5][[AskJS] Getting Tail Call Optimisation When Transpiling To JavaScript!](https://www.reddit.com/r/javascript/comments/if1vcq/askjs_getting_tail_call_optimisation_when/)
- url: https://www.reddit.com/r/javascript/comments/if1vcq/askjs_getting_tail_call_optimisation_when/
---
A Functional Programming style is seeing some significant adoption in non-traditional Functional Programming languages. JavaScript being a clear example, and the Dart language being another. Dart supports Either and Option through the Dartz package and also completely immutable collections via built\_collections, and so does JavaScript with union-type.js, immutable.js and others. I believe Dart is a very exciting language, which is not only a modern statically typed and easy to write language, but also compilable to various platforms and also **completely transpilable to JavaScript**!

The one major drawback, from a Functional Programming style perspective, of JavaScript and Dart, is that they currently don't support Tail Call Optimisation. This can change however. There is currently a new open issue with the Dart language team at [GitHub](https://github.com/dart-lang/language/issues/1159) that is discussing the benefits of Tail Recursion, the importance of Tail Call Optimisations and even possible ways to implement this into the language, compilers and transpiler (thus language guaranteed TCO in JavaScript).

Some interesting reading about the feasibility for Tail Call Optimisation is coming out in this [thread](https://github.com/dart-lang/language/issues/1159#issuecomment-678710258).

I quote a short snippet of the update:

&gt;These types of issues have been grappled with for years and numerous solutions have been implemented to work around similar constraints. Looping/Chicken Scheme/ShadowChicken/Trampoline... in various types of recursions. Compiler annotations ... to demarcate expectation so that the compiler can error if the said expectations can't be met. We can also turn to C++ compilers using loops and other TCO strategies, JavaScriptCore/WebKit having implemented ECMAScript 6 Tail Position Calls; Scala, Clojure, Kotlin that have successfully implemented TCOs by working around the JVM having no such support, ....

This is our chance to influence the Dart Language/Compiler teams to accept this feature into the Dart language and give us a JavaScript transpilable language that supports TCO. So, if you want to be able to run Tail Recursive functions in Constant Space (that is, without the fear of **Stack Overflow errors**) and those that are experimenting with/loving the Functional Programming style of coding, head on over to [GitHub](https://github.com/dart-lang/language/issues/1159#issuecomment-678710258) and thumbs-up the latest comments on this open issue and let your voices be heard.
## [6][GridFile - Mongoose Schema for MongoDB GridFS](https://www.reddit.com/r/javascript/comments/iez5x3/gridfile_mongoose_schema_for_mongodb_gridfs/)
- url: https://github.com/abskmj/gridfile
---

## [7][[AskJS] How do you guys pick colors for your projects?](https://www.reddit.com/r/javascript/comments/ieq442/askjs_how_do_you_guys_pick_colors_for_your/)
- url: https://www.reddit.com/r/javascript/comments/ieq442/askjs_how_do_you_guys_pick_colors_for_your/
---
I've been using the default TailwindCSS colors, but they are a bit too bland/flat for me.

I was trying [https://coolors.co/](https://coolors.co/) which is great, but I don't really have any sense of what looks good. Like i'll see one color I like as a main color, but trying to find colors that go well with that color and match the kind of style I want feels impossible.
## [8][To understand it better, I've simulated JavaScript "for await" loop with "while" loop](https://www.reddit.com/r/javascript/comments/if353p/to_understand_it_better_ive_simulated_javascript/)
- url: https://gist.github.com/noseratio/721fea7443b74a929ea93c8f6a18cec4#file-async-generator-js-L30
---

## [9][[AskJS] Would you read this blog post?](https://www.reddit.com/r/javascript/comments/if1wft/askjs_would_you_read_this_blog_post/)
- url: https://www.reddit.com/r/javascript/comments/if1wft/askjs_would_you_read_this_blog_post/
---
I was working on a quick easy to follow blog introducing PyPi and python ecosystem to Javascript developers.

I started wondering if anyone would like to read it. You know we all have those moments of doubt. Let me know what you think.
## [10][I made a small 2d raycasting simulation. Any feedback would be great.](https://www.reddit.com/r/javascript/comments/iecc7k/i_made_a_small_2d_raycasting_simulation_any/)
- url: https://aydencook03.github.io/simulations/raycasting.html
---

## [11][A web framework that uses React Hooks for declarative data persistence and reactivity](https://www.reddit.com/r/javascript/comments/ieh0bf/a_web_framework_that_uses_react_hooks_for/)
- url: https://github.com/redia-server/redia
---

## [12][W3C Keyboard events with Types](https://www.reddit.com/r/javascript/comments/iexd64/w3c_keyboard_events_with_types/)
- url: https://github.com/ashubham/w3c-keys/blob/master/README.md
---

