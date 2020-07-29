# reactjs
## [1][Beginner's Thread / Easy Questions (July 2020)](https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/
---
You can find [previous threads][wiki previous threads] in the wiki.

Got questions about React or anything else in its ecosystem?  
Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## 🆘 Want Help with your Code? 🆘

- **Improve your chances** by adding a minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz].
  - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
  - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
- **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- [New to Hooks? Check Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- and these React Hook recipes on [useHooks.com](https://usehooks.com/) by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[thinking in react hooks]: https://wattenberger.com/blog/react-hooks
[freecodecamp's react course]: https://www.freecodecamp.org/news/learn-react-course/
[microsoft frontend bootcamp]: https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[kent dodd's egghead.io course]: http://kcd.im/beginner-react
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
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
## [2][Who's Available? [July 2020]](https://www.reddit.com/r/reactjs/comments/hseduu/whos_available_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hseduu/whos_available_july_2020/
---
We alternate between hirers (on the 1st of the month) and agencies/freelancers/jobseekers (on the 15th).  
If you are looking to post or reply to React job postings, please check [this month's Who's Hiring post here.][hiring:this month]

---

If your post or comment is removed wrongly, please [send a message][message:mods] to mods  
because Automods bot is not perfect :)

---

Top Level comments must be Agencies and React Devs available for contract/permanent work.

Please include Location or any other Requirements in your comment. You can choose to use this format if it helps:

## (Fulltime | Contract | USA | Remote)

or

## (Agency | Europe | Remote)

Then we recommend adding a 2-3 sentence bio as well.

Not required, but may help:

- Link to Github/Portfolio
- Notable [r/reactjs][r/reactjs] submissions
- Preferred stack
- Former companies or clients
- Design or backend dev experience
- anything else you consider relevant. Put on your best show!
- Listing years of experience NOT required, it's a poor metric

If you are looking to hire, you can send a PM, or reply so that others might see your job opening.  
**Note**: Due to the sensitive nature of availability while currently in a job, users may be using alternate accounts.

For more ideas on what to include, look at the [last Who's Available posts][available:last month].

If you just want some portfolio feedback, check the stickied post below.

Good luck! #WriteOnceApplyEverywhere

[r/reactjs]: https://www.reddit.com/r/reactjs/
[available:last month]: https://www.reddit.com/r/reactjs/comments/ha504b/whos_available_june_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][Rubber Mesh Swipe Transition ( three-js &amp; react-spring)](https://www.reddit.com/r/reactjs/comments/hzfig9/rubber_mesh_swipe_transition_threejs_reactspring/)
- url: https://v.redd.it/9br9rls4xld51
---

## [4][What is a common best practice for the UI when calling a single API to load the content for a page?](https://www.reddit.com/r/reactjs/comments/hzvraq/what_is_a_common_best_practice_for_the_ui_when/)
- url: https://www.reddit.com/r/reactjs/comments/hzvraq/what_is_a_common_best_practice_for_the_ui_when/
---
Hi there, fairly new to React. I'm curious to what design patterns people use when calling a single api to get some data for a page. For example I have a route `/pages/list`. I need to call the API `get_pages` on mount to populate the data of a list.

Is it really necessary to use a loading indicator for this? I would like to just transition to the page naturally like a traditional website does. How would I achieve this? I considered:

1. Calling the API before calling the route but if there is a delay the UI will just do nothing and the user will be confused. I noticed on a traditional website there is a loading bar in the browser. Is it possible to use this with react?
2. First route and render null, then call the api and render on success, this has the disadvantage that the user sees a blank page.

What are your thoughts?
## [5][Create simple POS with React.js, Node.js, and MongoDB #10: CRUD Supplier](https://www.reddit.com/r/reactjs/comments/hzwjxh/create_simple_pos_with_reactjs_nodejs_and_mongodb/)
- url: https://blog.soshace.com/create-simple-pos-with-react-js-node-js-and-mongodb-10-crud-supplier/
---

## [6][Anyone can explain the following syntax to me: function Layout({children}) {...}](https://www.reddit.com/r/reactjs/comments/hzxgg2/anyone_can_explain_the_following_syntax_to_me/)
- url: https://www.reddit.com/r/reactjs/comments/hzxgg2/anyone_can_explain_the_following_syntax_to_me/
---
Reading Next.js docs, I came across this part, which I can't understand. Can anyone explain why the parameter of the function is inside an object?

    function Layout({ children }) { 
        return &lt;div&gt;{children}&lt;/div&gt; 
    } 
    export default Layout
## [7][Guide to React Context](https://www.reddit.com/r/reactjs/comments/hzhdfn/guide_to_react_context/)
- url: https://ui.dev/react-context/
---

## [8][Animating the Progress Percent Change in React - Time to Hack](https://www.reddit.com/r/reactjs/comments/i00no9/animating_the_progress_percent_change_in_react/)
- url: https://time2hack.com/animating-the-progress-percent-change-in-react/
---

## [9][Reactivated.app is an open-source webapp that keeps an eye on your JS dependency versions (built with React / NestJs / TypeScript / Chakra UI)](https://www.reddit.com/r/reactjs/comments/hzzu7v/reactivatedapp_is_an_opensource_webapp_that_keeps/)
- url: https://github.com/premieroctet/reactivated-app
---

## [10][Request header field x-auth-token is not allowed by Access-Control-Allow-Headers in preflight response.](https://www.reddit.com/r/reactjs/comments/hzwv1t/request_header_field_xauthtoken_is_not_allowed_by/)
- url: https://www.reddit.com/r/reactjs/comments/hzwv1t/request_header_field_xauthtoken_is_not_allowed_by/
---
I'm having the error in the title, as [shown here](https://imgur.com/a/dZSTXNL). I have zero clues as to what it means. I do have the CORS extension enabled, after all, as can be seen in the same screenshot. [Here's my code](https://pastebin.com/sei11i7X), if it helps. I'm COMPLETELY LOST as to why this is happening, and my job is threatened by this bullshit.
## [11][How to capture a part of URL in localStorage to be used later on a different route?](https://www.reddit.com/r/reactjs/comments/hzy95a/how_to_capture_a_part_of_url_in_localstorage_to/)
- url: https://www.reddit.com/r/reactjs/comments/hzy95a/how_to_capture_a_part_of_url_in_localstorage_to/
---
Hey, I’m stuck with an issue which I’m unable to resolve.
I’m working on a webapp whose one of the routes would be of the format 
webappName.ai/?invite_code=value&amp;fireside=yes

From this I have to capture the invite_code in localStorage.

How do you do this?
## [12][I have a question from react developers as a beginner.](https://www.reddit.com/r/reactjs/comments/hzxn4n/i_have_a_question_from_react_developers_as_a/)
- url: https://www.reddit.com/r/reactjs/comments/hzxn4n/i_have_a_question_from_react_developers_as_a/
---
I'm taking a udemy course on react, after learning the basics(props, state,...) I want to build a small project. I was wondering what was your first project when you started?
