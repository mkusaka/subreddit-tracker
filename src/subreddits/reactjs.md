# reactjs
## [1][Beginner's Thread / Easy Questions (Jan 2020)](https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/
---
[Previous threads][wiki previous threads] can be found in the Wiki.

Got questions about React or anything else in its ecosystem? Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## 🆘 Want Help with your Code? 🆘

- **Improve your chances** by putting a minimal example to either [JSFiddle][jsfiddle], [Code Sandbox][code sandbox] or [StackBlitz][stackblitz].
  - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
  - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
- **Pay it forward!** Answer questions even if there is already an answer - multiple perspectives can be very helpful to beginners. Also there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].
- **[Learn by teaching][learn by teaching]** &amp; **[Learn in public][learn in public]** - It not only helps the asker but also the answerer.

---

## New to React?

### Check out the sub's sidebar!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp](https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/)
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]
- [FreeCodeCamp's React course](https://www.freecodecamp.org/news/learn-react-course/)
- [Flavio Copes' React handbook](https://reacthandbook.com/)
- New to Hooks? Check Amelia Wattenberger's [Thinking in React Hooks](https://wattenberger.com/blog/react-hooks)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[official getting started page]: https://reactjs.org/docs/getting-started.html
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[robin wieruch's road to react]: https://roadtoreact.com/
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Who's Hiring? [Jan 2020]](https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/
---
We alternate between **Who's Hiring** (on the 1st of the month, [most recent one here][hiring:most recent]) and **Who's Available** (on the 15th, [most recent one here][available:most recent])

_I am sorry for the missing "Who's Available [Dec 2019]" 🙇‍♂️_

Welcome to **the biggest React job board in the world!** This is like Hacker News' **Who's Hiring** but just for React. Top Level comments must be **Job Opportunities.**

⚠️ NEW: WE ARE REQUESTING EVERYBODY FOLLOW [THE HN Who's Hiring FORMAT][format:hiring:hn]

**Company inc. | Job Title | City/State Location | Full-time/Part-Time | On-site/Remote | (Optional) Salary range | Website jobs page, other hard requirements etc.**

examples:

- **Thorn | San Francisco or Remote (US based) | Full-time Contract | $100k - $150k | Software Engineer | https://www.wearethorn.org/**
- **PolicyStat | Full-Stack Python+Django Software Engineer | Indianapolis, Vancouver, or REMOTE | Full Time | +\$80k**

Please include as much information as possible. **If you are remote-friendly, or open to sponsoring work visas to your country, say so! These are the top 2 questions!**

If you are looking for jobs, send a PM to the poster or post in our [Who's Available thread!][available:most recent]

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/e4pud0/whos_hiring_dec_2019/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/dxxqdn/whos_available_nov_2019/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [3][I build an open api lists repository](https://www.reddit.com/r/reactjs/comments/ekjfqc/i_build_an_open_api_lists_repository/)
- url: https://www.reddit.com/r/reactjs/comments/ekjfqc/i_build_an_open_api_lists_repository/
---
I was so tired to find an open API that is (actually) useful.

 I wish there was a list of all the open APIs I can use. 

So, I googled and found a useful repository at github. ([Link](https://github.com/public-apis/public-apis))

But it seems that repo isn't managed anymore. (Look at the issues and pr requests waiting there)

So I forked that and make new one.

&amp;#x200B;

Anyone whose interested can check or contribute. ([Link](https://github.com/public-api-lists/public-api-lists))

Thanks.
## [4][Redux pubsub?](https://www.reddit.com/r/reactjs/comments/ektu3d/redux_pubsub/)
- url: https://www.reddit.com/r/reactjs/comments/ektu3d/redux_pubsub/
---
Maybe I don't yet have the right mental model of Redux as I've now seen it used in different ways in a couple of, admittedly, messy projects. But can I use Redux with Thunk as a traditional pubsub type system?

If so, how do I connect events (i.e. actions) to other calls that have side effects, without having to explicitly design out a chain of dispatch calls? I'm assuming the right way to do it is to have the listeners watch for actions and know what to do.

The *listeners* get called in the reducers, but I can't take the payload and call another service to update the backend -  I can, of course, update the front end state, but it's now out of sync on the backend. Where do I attach the API call if the reducer can't make it.

For example:

* Click button
* Callback makes post to update service A, and returns payload of state.A, dispatching action.A
* Reducer B listens for Action.A and sees payload. It needs to take payload.A and update state.B, but should actually do it by calling service B and updating it's state when action.B is dispatched.

How do I trigger A and B without having to explicitly call dispatchA().then(payload.A, dispatchB(payload.A))?
## [5][My best VSCode Linting/Formatting configuration for Typescript Projects](https://www.reddit.com/r/reactjs/comments/eks63o/my_best_vscode_lintingformatting_configuration/)
- url: https://medium.com/javascript-in-plain-english/my-best-vscode-linting-formatting-configuration-for-typescript-projects-ef400ed9b78f?source=friends_link&amp;sk=d1a797a25de1e668bcd4e69d247f2be4
---

## [6][Sharing Code Between React and React-Native: What Not to Share - Ben Ellerby](https://www.reddit.com/r/reactjs/comments/eke9og/sharing_code_between_react_and_reactnative_what/)
- url: https://youtu.be/QO7SkFqRd7s
---

## [7][How often do you/people use Redux in React projects?](https://www.reddit.com/r/reactjs/comments/eklce7/how_often_do_youpeople_use_redux_in_react_projects/)
- url: https://www.reddit.com/r/reactjs/comments/eklce7/how_often_do_youpeople_use_redux_in_react_projects/
---
I am lerning React.js on Udemy and the lecturer says that every React project he has worked on used Redux and everyone who is working with React should know Redux.

At the same time, I find Redux quite complex. At lest it introduces a new level of complexity which seems unneeded (for me as a learner). 

**Questions:** Do you use Redux in your projects? Why? Does a developer really need to know Redux if he or she wants to work with React?
## [8][Please advise on why I can't get Remote jobs with my Resume with my experience and skills](https://www.reddit.com/r/reactjs/comments/eku4ko/please_advise_on_why_i_cant_get_remote_jobs_with/)
- url: https://www.reddit.com/r/reactjs/comments/eku4ko/please_advise_on_why_i_cant_get_remote_jobs_with/
---
I have worked in really challenging roles and did get good recommendations from previous employers. For some reason, I can't even get an interview chance.

What is wrong with my Resume??
## [9][Advanced React with TypeScript and Redux Course: Build a Sudoku App](https://www.reddit.com/r/reactjs/comments/ektr15/advanced_react_with_typescript_and_redux_course/)
- url: https://www.youtube.com/watch?v=MUs4oAudvO0
---

## [10][Hooks + Redux + Tests open source project](https://www.reddit.com/r/reactjs/comments/ektonl/hooks_redux_tests_open_source_project/)
- url: https://www.reddit.com/r/reactjs/comments/ektonl/hooks_redux_tests_open_source_project/
---
I’m looking for an open source project that uses hooks and redux, with plenty of tests, especially complex integration tests. 
I want to learn about the architecture and testing of a non-trivial app. 

I already checked video courses (I prefer real code), github and top posts in this /r. Not satisfied with my findings so far (I‘m fairly new to react, is not easy to find repos for three specific topics combined).

Do you have a favorite project in mind that cover all the bases?

Thanks.
## [11][JavsScript TV : Tech talks on React, CSS and everything JavaScript](https://www.reddit.com/r/reactjs/comments/ektinl/javsscript_tv_tech_talks_on_react_css_and/)
- url: http://jsgeeks.tv/
---

## [12][Code a Movie APP in React Native | React Native Tutorial for Beginners](https://www.reddit.com/r/reactjs/comments/eksmer/code_a_movie_app_in_react_native_react_native/)
- url: https://www.youtube.com/watch?v=aZYCEGyMIN0&amp;feature=share
---

