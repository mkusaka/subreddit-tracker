# javascript
## [1][WTF Wednesday (July 08, 2020)](https://www.reddit.com/r/javascript/comments/hnjqff/wtf_wednesday_july_08_2020/)
- url: https://www.reddit.com/r/javascript/comments/hnjqff/wtf_wednesday_july_08_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Showoff Saturday (July 11, 2020)](https://www.reddit.com/r/javascript/comments/hp8pa9/showoff_saturday_july_11_2020/)
- url: https://www.reddit.com/r/javascript/comments/hp8pa9/showoff_saturday_july_11_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [3][Developer Handbook 2020 - was created to cover the most common technical questions and requirements appearing prior to job interviews, during onboarding or personal goals / career planning at our company Apptension.](https://www.reddit.com/r/javascript/comments/hopwh9/developer_handbook_2020_was_created_to_cover_the/)
- url: https://github.com/apptension/developer-handbook
---

## [4][Guide To Javascript Array Functions: Why you should pick the least powerful tool for the job](https://www.reddit.com/r/javascript/comments/hoz6hr/guide_to_javascript_array_functions_why_you/)
- url: https://jesseduffield.com/array-functions-and-the-rule-of-least-power/
---

## [5][[AskJS] Trick for destructuring re-assignment without parenthesis](https://www.reddit.com/r/javascript/comments/hp91ta/askjs_trick_for_destructuring_reassignment/)
- url: https://www.reddit.com/r/javascript/comments/hp91ta/askjs_trick_for_destructuring_reassignment/
---
For context of what I'm talking about, see either [here](https://stackoverflow.com/questions/27386234/object-destructuring-without-var) or [here on stackoverflow](https://stackoverflow.com/questions/48714689/javascript-re-assign-let-variable-with-destructuring/48714713) (short) or the [notes here on MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment#Assignment_without_declaration) (detailed).

I hate this requirement of parenthesis around an assignment, for me it seems to communicate things that are not true ("this is an expression, we are going to use the return value"). Also it doesn't allow for a semicolon-free coding style (which may be a good thing for some people, but I don't like it), since otherwise the parenthesis might be interpreted as trying to call the previous line as a function. Also it's cumbersome to wrap assignments.

So I've came up with the following trick for reassignment instead. Staying with the example from the second stackoverflow link, you can simply write

     let {} = { latitude, longitude } = props.userLocation.coords

This works, needs no parenthesis, needs no semicolons, and doesn't pollute the namespace with any more variables. And while it still doesn't communicate the correct thing clearly ("a destructuring reassignment is happening here"), at least it doesn't seem to communicate anything else either (or worst case it communicates "what the heck is this").

That's it, just wanted to let y'all know about this, maybe someone else finds this useful too. And, of course, if someone has an even better solution, I'm all ears.

----

Offtopic: I don't feel like the [AskJS] tag rings very true here as there's no explicit question in my post, but the guide says it's also for "debating best practices", so I guess this post should be ok.
## [6][An Open source code playground with javascript](https://www.reddit.com/r/javascript/comments/hp8e6h/an_open_source_code_playground_with_javascript/)
- url: https://thecodezine.com/free-online-code-editor-for-website/
---

## [7][ToastMaker - Open Source JavaScript Library for Showing Toast Notifications on Web Page](https://www.reddit.com/r/javascript/comments/hp85j0/toastmaker_open_source_javascript_library_for/)
- url: https://toastmaker.wirehall.com/
---

## [8][I'm Making a Video Series about Building a Bit-level Binary Library from the Ground up in Javascript. This One Implements and Explains Addition in Terms of Logic Gates.](https://www.reddit.com/r/javascript/comments/hoo885/im_making_a_video_series_about_building_a/)
- url: https://www.youtube.com/watch?v=VMOyiYRFm8A
---

## [9][The traversal order of object properties in ES6](https://www.reddit.com/r/javascript/comments/hp9nby/the_traversal_order_of_object_properties_in_es6/)
- url: https://2ality.com/2015/10/property-traversal-order-es6.html
---

## [10][Yarn 2.1! Git Workspaces, Focused Installs, Loose mode, Live Playground, ...](https://www.reddit.com/r/javascript/comments/hosfmn/yarn_21_git_workspaces_focused_installs_loose/)
- url: https://dev.to/arcanis/yarn-2-1-git-workspaces-focused-installs-loose-mode-live-playground-4kfc
---

## [11][Gamedev Patterns and Algorithms in Action with TypeScript](https://www.reddit.com/r/javascript/comments/hop9lw/gamedev_patterns_and_algorithms_in_action_with/)
- url: https://medium.com/@gregsolo/gamedev-patterns-and-algorithms-in-action-with-typescript-d29b913858e
---

## [12][[AskJS] How do you feel about the emergence of frameworks such as htmx and Intercooler?](https://www.reddit.com/r/javascript/comments/hp279x/askjs_how_do_you_feel_about_the_emergence_of/)
- url: https://www.reddit.com/r/javascript/comments/hp279x/askjs_how_do_you_feel_about_the_emergence_of/
---
In case anyone is wondering what they are, these frameworks are kind of a new way of building websites. For instance, htmx allows you to access Ajax and server sent events directly in HTML. I have been doing front-end dev for a long time (started back when Angular 1 was the hot new thing), but I definitely see benefits of using these technologies for some use cases.

Anyway, I would love to hear how you feel about it and maybe get a discussion going about server-side rendering making a comeback.

Links to frameworks mentioned:

[https://intercoolerjs.org/](https://intercoolerjs.org/)

[https://htmx.org/](https://htmx.org/)
