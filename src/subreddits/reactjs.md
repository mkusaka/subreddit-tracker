# reactjs
## [1][Beginner's Thread / Easy Questions (Feb 2020)](https://www.reddit.com/r/reactjs/comments/exf7nq/beginners_thread_easy_questions_feb_2020/)
- url: https://www.reddit.com/r/reactjs/comments/exf7nq/beginners_thread_easy_questions_feb_2020/
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

---

## New to React?

### Check out the sub's sidebar!

🆓 Here are great, **free** resources! 🆓

- [Create React App][create react app]
- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Get started with Redux][get started with redux] by [/u/acemarke][/u/acemarke] (Redux Maintainer).
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- [Tyler McGinnis' 2018 Guide][tyler mcginnis' 2018 guide]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[create react app]: https://facebook.github.io/create-react-app/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[kent dodd's egghead.io course]: http://kcd.im/beginner-react
[tyler mcginnis' 2018 guide]: https://tylermcginnis.com/reactjs-tutorial-a-comprehensive-guide-to-building-apps-with-react/
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[robin wieruch's road to react]: https://roadtoreact.com/
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
## [2][Who's Available? [Feb 2020]](https://www.reddit.com/r/reactjs/comments/f44wd7/whos_available_feb_2020/)
- url: https://www.reddit.com/r/reactjs/comments/f44wd7/whos_available_feb_2020/
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
[available:last month]: https://www.reddit.com/r/reactjs/comments/eouupz/whos_available_jan_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/ex778e/whos_hiring_feb_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][I made my first real React App: a site to help people learn songs on the flute. I would love any and all feedback and constructive criticism!](https://www.reddit.com/r/reactjs/comments/f3y5cy/i_made_my_first_real_react_app_a_site_to_help/)
- url: https://www.reddit.com/r/reactjs/comments/f3y5cy/i_made_my_first_real_react_app_a_site_to_help/
---
Site: https://dbolesta.github.io/flute-cards/

Gif demo: https://i.imgur.com/1ZSE6w5.mp4

This is my first real/useful React App. I play the flute casually, and I got the idea for this app when I was frustrated hunting down the notes I wanted to play on a [flute fingering chart](https://i.imgur.com/meGU6Jx.jpg), especially when I wanted to play a whole melody. I thought it would be so convenient if I could just have those charts displaying the exact line of music I was learning, so I figured I would build it myself.

On my app, you can arrange notes in whatever order you'd like on a series of Rows. Each note is represented as a [Card](https://i.imgur.com/tAqcvyx.png) that displays the notes letter, position on a music staff, and the flute fingering diagram that shows you how to play that note on the flute.

Some handy features I was able to fit in:

 - Select a note from a Keyboard, a Music Staff, or a flute Register (low, mid, or high).
 - Hovering over any note will highlight that same note on every Selector.
 - Drag and drop Cards within a Row, between Rows, or the Rows themselves.
 - Play MIDI audio of any Row of notes to make sure the notes you constructed are the ones you want.
 - Save and Load as many "Decks" of Cards as you want, so you can come back to a song later.
 - A couple of sample "Decks" so you can see how songs can be constructed.

I still have some todos to get to, but I think it's at the point where it's functional and can be shown and used as intended.

I would love any feedback you have on this, but some specific things I would love to know are:

 - Does it perform well?
 - Do you understand how it works? Are you able to use the app easily?
 - Are there any glaring issues or bugs?
 - Are there any obvious features you think should be included?
 - Do you notice any specific React / technical improvements that could be made?

[GitHub link](https://github.com/dbolesta/flute-cards)
## [4][When to use useEffect or useLayoutEffect](https://www.reddit.com/r/reactjs/comments/f46gu9/when_to_use_useeffect_or_uselayouteffect/)
- url: https://aganglada.com/blog/useeffect-and-uselayouteffect/
---

## [5][Interview code challenge](https://www.reddit.com/r/reactjs/comments/f40x71/interview_code_challenge/)
- url: https://www.reddit.com/r/reactjs/comments/f40x71/interview_code_challenge/
---
Recently, I was invited in for a code challenge prerequisite to an interview. Upon arrival I was given a laptop with what seemed to be a basic scaffold of a React app, keep in mind this is for a React dev position. I opened the laptop and read the Readme with all of the instructions. There was also a basic node app with a get all todos, get todo by id, and a post for a todo and a few other endpoints. I was told I had 1 hour to do as much as I could. There was a mockup that seemed to have some pretty complex CSS involved. Including redux for todos was also a requirement.

The requirements for functionality included, display all todos, create a todo, delete a todo, and edit a todo amd to style these components appropriately. 

My question to you all is, is 1 hour enough to complete an entire project of this stature? Really just wanna gauge how I compare to the react community! Thanks guys!

EDIT: I felt this was a good test of my skills if not the best test I've taken in a while, just wish I'd had more time to complete/make it pretty.
## [6][20+ React Native Admin templates for 2019](https://www.reddit.com/r/reactjs/comments/f45ix1/20_react_native_admin_templates_for_2019/)
- url: https://bootstraplib.com/react-admin-templates/
---

## [7][IT IS WORKING! | React/Node/GraphQL PWA Manga Reader | Part 10 | Pair Programming](https://www.reddit.com/r/reactjs/comments/f47g9x/it_is_working_reactnodegraphql_pwa_manga_reader/)
- url: https://www.youtube.com/watch?v=YISWT2Q6NHo
---

## [8][My first MERN stack app: Day Planner](https://www.reddit.com/r/reactjs/comments/f49for/my_first_mern_stack_app_day_planner/)
- url: https://github.com/ahmetbcakici/DayPlanner
---

## [9][How to clone div content in react js](https://www.reddit.com/r/reactjs/comments/f461db/how_to_clone_div_content_in_react_js/)
- url: https://www.reddit.com/r/reactjs/comments/f461db/how_to_clone_div_content_in_react_js/
---
Helloooo..! At deep dive in reactjs. I've some trouble make some events most of the i got solution but in some cases i didn't get.   
Problem : How to clone div content   
I want to clone some div content :like duplicate. is that possible in reactjs.

Quick response is highly appreciated.  
Thanks
## [10][The future of React Native for Web · Issue #1538 · necolas/react-native-web · GitHub](https://www.reddit.com/r/reactjs/comments/f3y73w/the_future_of_react_native_for_web_issue_1538/)
- url: https://github.com/necolas/react-native-web/issues/1538
---

## [11][In the past 30 days, I made an open source design system: Looking for feedback!](https://www.reddit.com/r/reactjs/comments/f3pb73/in_the_past_30_days_i_made_an_open_source_design/)
- url: https://www.reddit.com/r/reactjs/comments/f3pb73/in_the_past_30_days_i_made_an_open_source_design/
---
TL;DR: 👉 [hacker-ui.com](https://hacker-ui.com/)

About a month ago from today, I started working on a design system for my own personal projects. I wanted to create my own design language because the rest of them just felt too complex, too outdated, or too branded (for example Bootstrap is too old, Material UI is too Google-y).

I wanted a design system that wasn’t too complicated or special but just something that’s unbranded and aesthetically pleasing.

Anyway 30 days from realizing that, I have a design system I’m ready to share!

It’s called “Hacker UI” and I’m currently pitching it as “the hacker’s go-to design system” (similar to how Bootstrap was a go-to).

Features:

- 25+ components
- 25+ codesandbox examples
- New CSS-in-JS solution
- Composable styles by default 
- Full Typescript support

**Be warned, the project is still young!**

There’s still a lot to be done with this library (mobile-friendly docs, optimizations, SSR support, a11y audits, CI and testing, animations etc) but I’m still super excited to share. I'm more posting to gauge interest for possible contributors at this time.

Please let me know what you think!

- [docs at hacker-ui.com](https://hacker-ui.com/)
- [github repo](https://github.com/ricokahler/hacker-ui)

---

Edit: Thanks for all the feedback and support. I've created [an issue on the GitHub for a rough idea of a roadmap/timeline for a stable 1.0 release](https://github.com/ricokahler/hacker-ui/issues/6). Check it out if you're curious!
## [12][Which frameworks and setup?!?!](https://www.reddit.com/r/reactjs/comments/f48pod/which_frameworks_and_setup/)
- url: https://www.reddit.com/r/reactjs/comments/f48pod/which_frameworks_and_setup/
---
I know you guys must get this a lot, but here we go. I am just starting to learn react after using pure HTML CSS and JS, i prefer a minimal architecture where i know exactly what everything does with minimal boilerplate. I have tried NextJS and Create React App, but there is just a lot going on that seems a bit extra. i read this article  [https://hackernoon.com/move-over-next-js-and-webpack-ba367f07545](https://hackernoon.com/move-over-next-js-and-webpack-ba367f07545) about creating your own toolchain with parcel instead of webpack, is this worth it? i would want to use JSX with my site too, so the script tags on an HTML page may not work.

&amp;#x200B;

Many thanks, if you need more info i will reply, but am feeling very indecisive as there are lots of options!
