# reactjs
## [1][Beginner's Thread / Easy Questions (August 2020)](https://www.reddit.com/r/reactjs/comments/i1u5g1/beginners_thread_easy_questions_august_2020/)
- url: https://www.reddit.com/r/reactjs/comments/i1u5g1/beginners_thread_easy_questions_august_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Got questions about React or anything else in its ecosystem?  
Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** by adding a minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz].
    - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
    - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
1. **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**! 👉

🆓 Here are great, **free** resources!

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- New to Hooks? Check out [Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- and these React Hook recipes on [useHooks.com][useHooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[useHooks.com]: https://usehooks.com/
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
## [2][Who's Available? [August 2020]](https://www.reddit.com/r/reactjs/comments/iaggwf/whos_available_august_2020/)
- url: https://www.reddit.com/r/reactjs/comments/iaggwf/whos_available_august_2020/
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
[available:last month]: https://www.reddit.com/r/reactjs/comments/hseduu/whos_available_july_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/i1u8a4/whos_hiring_august_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][My very first Portfolio site after learning react.js.](https://www.reddit.com/r/reactjs/comments/ig9829/my_very_first_portfolio_site_after_learning/)
- url: https://mirozdevkota.com.np/
---

## [4][Code review please? 💩](https://www.reddit.com/r/reactjs/comments/ig6m5s/code_review_please/)
- url: https://www.reddit.com/r/reactjs/comments/ig6m5s/code_review_please/
---
Hi everyone! I'm working on a React project that uses Context for state management instead of Redux. Can anyone review my project with their thoughts?

Project: [github.com/junnac/pomoplan](https://github.com/junnac/pomoplan)

Feel free to use [@vinylemulator](https://www.reddit.com/user/vinylemulator/)'s [codecomments.dev](https://codecomments.dev/) to provide comments for the files!

I've been trying to research best practices for implementing Context with hooks to refactor my project. The React docs demonstrate a basic implementation to explain the API, but there are no notes about best/common practices. I've seen the following approaches:

* Creating custom hooks for accessing a context value
* Directly interfacing with the Context from within a consuming component by accessing the context value with `useContext`
* Managing the Context value with `useState` in the context provider component
* Managing the Context value with `useReducer`
* Creating container components (similar to Redux container components created via the [connect()](https://react-redux.js.org/api/connect) function)

Any and all feedback is super appreciated! Thank you!!
## [5][JS/TS Debugger with Time-Traveling, Hot-swapping, API and Persistent State, zero-config for create-react-app projects (VSCode plugin)](https://www.reddit.com/r/reactjs/comments/ifnj3m/jsts_debugger_with_timetraveling_hotswapping_api/)
- url: https://v.redd.it/ovm2zvvvvxi51
---

## [6][Remember (notetaking application)](https://www.reddit.com/r/reactjs/comments/igaytj/remember_notetaking_application/)
- url: https://v.redd.it/9fvgqz7u35j51
---

## [7][Tippy JS Tutorial | How to build tooltips in ReactJs](https://www.reddit.com/r/reactjs/comments/ig5zon/tippy_js_tutorial_how_to_build_tooltips_in_reactjs/)
- url: https://www.youtube.com/watch?v=36WZjFR7UMc&amp;feature=share
---

## [8][useEffect: 4 Tips Every Developer Should Know](https://www.reddit.com/r/reactjs/comments/ig9ajr/useeffect_4_tips_every_developer_should_know/)
- url: https://medium.com/swlh/useeffect-4-tips-every-developer-should-know-54b188b14d9c
---

## [9][Navbar background different for landing and other pages.](https://www.reddit.com/r/reactjs/comments/ig8wqs/navbar_background_different_for_landing_and_other/)
- url: https://www.reddit.com/r/reactjs/comments/ig8wqs/navbar_background_different_for_landing_and_other/
---
I want to create a navbar which should have the background of header on landing page and different color on other pages  
for example  
landing page :  [http://glavet.de/](http://glavet.de/) , [https://demos.creative-tim.com/argon-design-system-react/#/landing-page](https://demos.creative-tim.com/argon-design-system-react/#/landing-page)   


other pages:  [https://demos.creative-tim.com/argon-design-system-react/#/documentation/app-navigation](https://demos.creative-tim.com/argon-design-system-react/#/documentation/app-navigation) ,  [http://glavet.de/faq/](http://glavet.de/faq/)   
what approach should I use to achieve.  

Thanks.
## [10][Uber Homepage Clone](https://www.reddit.com/r/reactjs/comments/igbwiv/uber_homepage_clone/)
- url: https://www.reddit.com/r/reactjs/comments/igbwiv/uber_homepage_clone/
---
&amp;#x200B;

 [https://uberclone.netlify.app/](https://uberclone.netlify.app/)
## [11][Webpack - The Easiest Way](https://www.reddit.com/r/reactjs/comments/igbq9c/webpack_the_easiest_way/)
- url: https://medium.com/@bhavesh12/webpack-the-easiest-way-a14dcb270cb9?source=friends_link&amp;sk=4a46bcdddf07a1a340427118c0c99b63
---

## [12][Personal/Portfolio website after 3 months of bootcamp.](https://www.reddit.com/r/reactjs/comments/igbmsv/personalportfolio_website_after_3_months_of/)
- url: https://www.reddit.com/r/reactjs/comments/igbmsv/personalportfolio_website_after_3_months_of/
---
Hey all, 

I completed a coding bootcamp last week and made this site as small portfolio for potential employers.

[Pawel Jaskolski ](http://www.paweljaskolski.com)

Its built with React, TypeScript, JavaScript, Node, Express and MySql...also some Bootstrap and styled-components for the good looks.

Please take a look, gimme some love if you feel like it or some constructive criticism.

Thank you all!
