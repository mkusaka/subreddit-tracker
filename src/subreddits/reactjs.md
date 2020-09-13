# reactjs
## [1][Who's Hiring? [September 2020]](https://www.reddit.com/r/reactjs/comments/ikn3vo/whos_hiring_september_2020/)
- url: https://www.reddit.com/r/reactjs/comments/ikn3vo/whos_hiring_september_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/i1u8a4/whos_hiring_august_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/iaggwf/whos_available_august_2020/
## [2][Beginner's Thread / Easy Questions (September 2020)](https://www.reddit.com/r/reactjs/comments/illwv0/beginners_thread_easy_questions_september_2020/)
- url: https://www.reddit.com/r/reactjs/comments/illwv0/beginners_thread_easy_questions_september_2020/
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
## [3][I just published another vscode extension that allows you to search through 20+ free icon sets and paste them into your code all within the editor.](https://www.reddit.com/r/reactjs/comments/irs96o/i_just_published_another_vscode_extension_that/)
- url: https://v.redd.it/c7vkbsarlum51
---

## [4][Serverless Video Chat App using Firebase and WebRTC in React](https://www.reddit.com/r/reactjs/comments/iraa9e/serverless_video_chat_app_using_firebase_and/)
- url: https://www.youtube.com/watch?v=-d45WHNU9J4&amp;feature=share
---

## [5][How can I run an API call AFTER the results from my useSelector() hook?](https://www.reddit.com/r/reactjs/comments/iry5hs/how_can_i_run_an_api_call_after_the_results_from/)
- url: https://www.reddit.com/r/reactjs/comments/iry5hs/how_can_i_run_an_api_call_after_the_results_from/
---
I am trying to use some data that will come from my `useSelector()`  
 hook in a network API call. However, with the code below, I get the error `TypeError: cannot read property 'query' of undefined.`

I understand this is because the query has not come back from the `useSelector()`  
 yet. Is there a way I can wait for that data THEN call the API?

`const pageContent = useSelector(getPageContent);`

  
`useEffect(() =&gt; {`  
`axios`  
`.get('https://www.googleapis.com/youtube/v3/search', {\`  
`params: {`  
`key: process.env.API_KEY,`  
`part: 'snippet',`  
`type: 'video',`  
`q: pageContent.data.query,`  
`},`  
`})`  
`.then((res) =&gt; res)`  
`.then((data) =&gt; console.log(data));`  
`}, []);`
## [6][Would you use useMemo for memoizing components?](https://www.reddit.com/r/reactjs/comments/irxtsd/would_you_use_usememo_for_memoizing_components/)
- url: https://www.reddit.com/r/reactjs/comments/irxtsd/would_you_use_usememo_for_memoizing_components/
---
I saw in a project I am working on a use of useMemo for a small (a potential additional icon) component within a larger component. Is it something you would do? Is it a common?
## [7][Help with react-testing-library and Firebase](https://www.reddit.com/r/reactjs/comments/irxobp/help_with_reacttestinglibrary_and_firebase/)
- url: https://www.reddit.com/r/reactjs/comments/irxobp/help_with_reacttestinglibrary_and_firebase/
---
*tl;dr*: How do I mock firebase(initializeApp, auth, etc) using jest and react-testing-library

I'm creating a form component that uses firebase to authenticate a user, either using email-password or google. I'm trying to test the component but can't mock firebase(don't know how, tbh). Can someone please help me? How do I mock firebase(initializeApp, auth, etc) using jest and react-testing-library?

For example, how do I check if `firebase.initializeApp()` is called by my component on mount, using jest and react-testing-library?

Also, what are some good resources that i can use to learn jest and react-testing-library.
## [8][What do I do?](https://www.reddit.com/r/reactjs/comments/irnkiy/what_do_i_do/)
- url: https://www.reddit.com/r/reactjs/comments/irnkiy/what_do_i_do/
---
Hello everyone. hope anyone can give me some advice or guide me.

I'm an engineer student and decided to start programming on my own because yes, I got a lot of logic thinking on my career but I was missing a lot of new tech that jobs are looking for to apply.

I have some background with Visual Basic (my first programming language when I was a teenager) , C++, Java and C where I learned the fundamentals about it and now I'm focused with Javascript and React (MERN) / React Native because I love how it works. I have been doing FreeCodeCamp (1800hs done), Udemy, TreeHouse, Scrimba, SoloLearn and even more among free and paid courses but I feel like I need to get my hands on real projects. I have my github (It's growing and I feel very proud of the green dots showing) where I upload my practices and some apps that I made and I try to work everyday because despite all, I enjoy programming.

And here comes my question.. how to get a job where my skills can be used? I mean, I find all job offers looking for Senior developers and if they are for Jr developers, they ask for 2 years of experience and even more... that's crazy.

I'm not giving up because I love to work in the tech field but its kinda frustrating right now.
## [9][Creating a Custom React Hook: Use Window Size](https://www.reddit.com/r/reactjs/comments/irlpdr/creating_a_custom_react_hook_use_window_size/)
- url: https://www.youtube.com/watch?v=OHvJqOjToes
---

## [10][I built a SERVERLESS VIDEO CHAT app in REACTJS using WEBRTC and FIREBASE](https://www.reddit.com/r/reactjs/comments/ir7cxa/i_built_a_serverless_video_chat_app_in_reactjs/)
- url: https://v.redd.it/z8a5iznyrnm51
---

## [11][Using Context with Radio Button](https://www.reddit.com/r/reactjs/comments/irx0we/using_context_with_radio_button/)
- url: https://www.reddit.com/r/reactjs/comments/irx0we/using_context_with_radio_button/
---
I am trying to set context on a category react component using the context API, which occurs when a radio button is checked. Whenever I click, nothing is happening. I want the radio button to remember the previously checked item when the component is unmounted. Am i doing something wrong?

[https://imgur.com/a/Ea1dWAl](https://imgur.com/a/Ea1dWAl)

The other styles are Checkmark which is a span, and the Container is a label. Thanks in advance.
## [12][Covid-19 Map: Complete Tutorial using react leaflet, hooks and bootstrap - Choropleth map](https://www.reddit.com/r/reactjs/comments/irkvxl/covid19_map_complete_tutorial_using_react_leaflet/)
- url: https://youtu.be/4cliojOu3as
---

