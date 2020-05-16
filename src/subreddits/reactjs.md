# reactjs
## [1][Beginner's Thread / Easy Questions (May 2020)](https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/
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
## [2][Who's Available? [May 2020]](https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/
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
[available:last month]: https://www.reddit.com/r/reactjs/comments/fiv53t/whos_available_mar_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][Modern React From The Beginning - Excursus: Strapi As Back-End](https://www.reddit.com/r/reactjs/comments/gksalu/modern_react_from_the_beginning_excursus_strapi/)
- url: https://youtu.be/CyaCQ3Qcbvw
---

## [4][Don't know what to test on your React App? Learn how to make a test list.](https://www.reddit.com/r/reactjs/comments/gkd9l3/dont_know_what_to_test_on_your_react_app_learn/)
- url: https://joaoforja.com/blog/learn-how-to-make-a-test-list/
---

## [5][Any reason not to cache all my functional components that don't take any props?](https://www.reddit.com/r/reactjs/comments/gklyyf/any_reason_not_to_cache_all_my_functional/)
- url: https://www.reddit.com/r/reactjs/comments/gklyyf/any_reason_not_to_cache_all_my_functional/
---
I've been noticing that ~60% of my components don't even take any props, and some of them get re-rendered all the time just because their parent re-render. Doing some tests with memoization with a handful of components I can totally see a noticeable difference when profiling the app if I start memoizing those components.

So the question is: is there any reason whatsoever why I shouldn't be wrapping all my functional components that don't take any props with the following wrapper?

```
const memoAlways = component =&gt; React.memo ( component, () =&gt; true );
```

---

Edit: if you think this question is about whether one should memoize all the things please read the question again more carefully. I'm talking about components with no props and a ~0 cost comparator function: `() =&gt; true`.
## [6][Creating module for API calls.](https://www.reddit.com/r/reactjs/comments/gkt84x/creating_module_for_api_calls/)
- url: https://www.reddit.com/r/reactjs/comments/gkt84x/creating_module_for_api_calls/
---
I was wondering if you should create module / class for handling the API. The class would contain all the api calls etc.  Would this be good approach?
## [7][Hey Guys, Check out this new react components library that i have created.](https://www.reddit.com/r/reactjs/comments/gkpqy5/hey_guys_check_out_this_new_react_components/)
- url: https://github.com/sha-el/sha-el-design
---

## [8][How to redirect react app to another localhost port to authenticate on load page if user not logged in?](https://www.reddit.com/r/reactjs/comments/gkr5id/how_to_redirect_react_app_to_another_localhost/)
- url: https://www.reddit.com/r/reactjs/comments/gkr5id/how_to_redirect_react_app_to_another_localhost/
---
React Route with windos.location.href = external url causes infinity reload and screen flickering on load page.

What I'm doing is that on load of the react page, I'm either returning a the following router under if "user object" is not defined, or the component if the user is logged in.

The code under works fine if I put it on a button and load the whole component, but on load of the page, when it returns this component automatically, it starts reload infinity.

React Route with windos.location.href = external url / port to another app causes infinity reload and screen flickering. in this case config.serverPort is 9000, and my react app is running on port 3000.

    return ( &lt;Router&gt; &lt;Route component={() =&gt; {             window.location.href = `http://localhost:${config.serverPort}/login`; return null; }} /&gt; &lt;/Router&gt;

meanwhile this works:

    return ( &lt;Router&gt; &lt;Route component={() =&gt; {             window.location.href = `http://www.vg.no`; return null; }} /&gt;

Is seems like as long as I have localhost in the url, it causes this infinity reload, even if I set port to the same as the app 3000. This is happening during rendering as I try to redirect the user to oath external login page if iser object is empty instead of rendering the app. Does anyone know how I can fix this?

It is worth to mention that the login service is on another external localhost port, so it is not a part of the react app.
## [9][background perspective effect](https://www.reddit.com/r/reactjs/comments/gkunes/background_perspective_effect/)
- url: https://www.reddit.com/r/reactjs/comments/gkunes/background_perspective_effect/
---
Hi, 

I am trying to implement an effect, that translates the background based on the mouse move. 

I got pretty close, but for some reason, the animation is a bit _laggy_. I thought it has something to do with the `requestAnimationFrame()`, but it did not solve the problem, it's still not smooth, and it seems I am failing to figure out why. 

I have wrapped this example into a [codesandbox](https://codesandbox.io/s/competent-mclaren-8vrg1?file=/src/components/Background/index.js), if someone could tell me how to fix this, it would be highly appreciated. :) 

Thank you!
## [10][Fakebooker - Universal Facebook Clone](https://www.reddit.com/r/reactjs/comments/gkiiic/fakebooker_universal_facebook_clone/)
- url: https://www.reddit.com/r/reactjs/comments/gkiiic/fakebooker_universal_facebook_clone/
---
Hey react devs! I've been developing this portfolio project for quite a few months and I would like you to check it out and tell me your opinion or any constructive critisism about it. Sadly I do not have the opportunity to create a video showcasing it so more people can see it but i guess it's alright. I truly believe that it's worth looking it up it would mean a lot to me. I've tried my best to make it pop out from the other usual clones so enjoy the app!

&amp;#x200B;

Github repo:  [https://github.com/KristianWEB/fakebooker-frontend](https://github.com/KristianWEB/fakebooker-frontend)

live website:  [https://fakebooker.com/](https://fakebooker.com/)

&amp;#x200B;

If you have any personal questions or whatever here's my twitter:  [https://twitter.com/KristianWEB7](https://twitter.com/KristianWEB7)
## [11][Recoil - state management lib for React](https://www.reddit.com/r/reactjs/comments/gka9x0/recoil_state_management_lib_for_react/)
- url: https://blog.graphqleditor.com/recoil-react-state-management-library/
---

## [12][In defense of the modern web - Rich Harris](https://www.reddit.com/r/reactjs/comments/gkjsfx/in_defense_of_the_modern_web_rich_harris/)
- url: https://dev.to/richharris/in-defense-of-the-modern-web-2nia
---

