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
## [3][Is Module Federation for me?](https://www.reddit.com/r/reactjs/comments/jm1bl3/is_module_federation_for_me/)
- url: https://www.reddit.com/r/reactjs/comments/jm1bl3/is_module_federation_for_me/
---
We are working on a problem that is in the sweet spot for micro frontends:  App A, written by Team A, is about to hit production. App B, is yet-to-be written and is still just a twinkle in my eye. App B will appear as a new feature of App A down the road a bit. App B will appear in App A's nav bar, and will get all of somedomain-dot-com/appb. Each team will independently deploy their apps. App A "owns" the auth UI, and must be able to share the auth status with App B (and all future app-lets). It is anticipated that Team A might need to parcel out it's nav system into a separate shared container for App B to consume.

App A is, at this moment in time, a create-react-app app.  I've seen [examples](https://martinfowler.com/articles/micro-frontends.html) of micro frontends done with CRA.  App B aspires to be a NextJS app, and I anticipate that Next + Webpack 5 will play nicely by the time we need them to.  I am curious to find out what the level of effort will be to take App A, which "owns" user authentication UI and auth token creation, to be able to "hoist" the new feature from team B?  I am a total webpack nooob, and a devops feather-weight. I know how to solve the shared auth issue between to backends using NginX, http-only cookies, and a shared decrypt secret. 

What I am unable to ascertain, just by following that initial Webpack MF blog article example, is will MF get in the way of my planned solution using the above combo for shared auth UI?  As presented, it seems it would hit a brick wall when it comes to cross-domain requests. Can the federation's module resolution work through reverse proxy servers?

That's a lot of questions, and not really React-y, but CRA is the starting point, so I thought I'd through them all out here for discussion.
## [4][Run your React Native app on the web with React Native for Web](https://www.reddit.com/r/reactjs/comments/jlgo5q/run_your_react_native_app_on_the_web_with_react/)
- url: https://mmazzarolo.com/blog/2020-10-24-adding-react-native-web/
---

## [5][Is there a Carousel library that allows to add delay to slide changes in React?](https://www.reddit.com/r/reactjs/comments/jlx9pl/is_there_a_carousel_library_that_allows_to_add/)
- url: https://www.reddit.com/r/reactjs/comments/jlx9pl/is_there_a_carousel_library_that_allows_to_add/
---
Hi, I am relatively new to React and I am trying to include a Carousel that has a delay between slides to wait for an animation to complete, however I have not found a single library that offers this option. 

I know I could implement this from scratch myself, but as I am not very experienced with React I think it would end up having a few bugs sneaked in.

I ask this here because maybe someone who has been using React long enough has faced a similar situation and may have found a library or something like that that already does this, and well in case there is no one I guess I'll have to implement it myself.

A working example of what I would like to achieve can be found [here](https://bridgelanding.qodeinteractive.com/elements/qode-slider/).
## [6][How to use redux in case of nested object?](https://www.reddit.com/r/reactjs/comments/jlzv5j/how_to_use_redux_in_case_of_nested_object/)
- url: https://www.reddit.com/r/reactjs/comments/jlzv5j/how_to_use_redux_in_case_of_nested_object/
---
Hello, I was wondering how to handle nested state.  Suppose, I have Project state and it has properties Id and Task. Also, Task is an object with properties Id and Name. How should I set my redux in order to edit Name properties in Task?
## [7][How to persist form data?](https://www.reddit.com/r/reactjs/comments/jm29mj/how_to_persist_form_data/)
- url: https://www.reddit.com/r/reactjs/comments/jm29mj/how_to_persist_form_data/
---
If there is a form usually we send the record and save it in db on form submission. 
But what would be the best practice to persist data as making call to backend on every update wont be a good idea and local storage would only work on the same machine/browser and  wont be available across every device?
## [8][Wanted to dive into MERN? Here's a Full-stack project to start with :)](https://www.reddit.com/r/reactjs/comments/jm0ufd/wanted_to_dive_into_mern_heres_a_fullstack/)
- url: https://youtu.be/ngc9gnGgUdA
---

## [9][Re-render not being applied appropriately](https://www.reddit.com/r/reactjs/comments/jm0avj/rerender_not_being_applied_appropriately/)
- url: https://www.reddit.com/r/reactjs/comments/jm0avj/rerender_not_being_applied_appropriately/
---
I'm a newbie in ReactJS and I'm finding some unexpected behaviour made by React (Or maybe my concepts aren't at best !)

Here's the thing:

I'm preparing a test page using NextJS and inside the page, I'm rendering a component '*Questions*'

Inside &lt;*Questions&gt;*, I have a block as follows:

&amp;#x200B;

`const Questions = ({ data }) =&gt; {const {current_question_no,current_question,queids,get_question,authToken,answers,} = data;`

&amp;#x200B;

`const [userPref,setUserPref] = useState(1)`

`const renderedChoices = answers.map((a) =&gt; {return (&lt;Answersanswer_id={a.id}option={a.option}user_pref={a.user_pref}onUserPrefClick={handleUserPref}&gt;&lt;/Answers&gt;);});`

`const handleUserPref = (spanRef) =&gt; {//manipulating the form ref directlyif(!spanRef.current.innerText){//assign the valuespanRef.current.innerText = userPrefsetUserPref(userPref+1) //state changes}};`

`return(`

`&lt;QuestionBlock&gt;`

`{renderedChoices}`

`)`

This is the code for my "*Answers*" component:

`import React from "react";const Answers = ({ answer_id, option, user_pref, onUserPrefClick }) =&gt; {const span_ref = React.createRef();return (&lt;label &gt;&lt;inputtype="checkbox"id={answer_id}class="check-hidden"value={user_pref}onClick={(e) =&gt; onUserPrefClick(span_ref)}/&gt;&lt;div class="question-body"&gt;&lt;div class="number"&gt;&lt;span ref={span_ref}&gt;{user_pref}&lt;/span&gt;&lt;/div&gt;&lt;p&gt;{option}&lt;/p&gt;&lt;/div&gt;&lt;/label&gt;);};export default Answers;`

&amp;#x200B;

As you can see, I'm allowing users to click on a span element inside my &lt;Answers&gt; component, and whenever the user clicks on it, I'm invoking `handleUserPref` (of Questions component) via a prop named `onUserPrefClick`

The `handleUserRef` of &lt;Question&gt;s Component assigns the value of `userRef` (piece of state) to span ref element. Thereafter, it increments the the `userPref` by calling `setUserPref(userPref+1)`

When a piece of state changes, it must rerender the entire component which must ALSO rerender my &lt;Answers&gt; Component ! If this happens, incidentally, my value of the span element must change to the original `user_pref` prop passed to &lt;Answers&gt; Component - But it is NOT CHANGING - In fact, it is remaining the same as the one assigned from the `userRef` state piece in &lt;Questions&gt; Component !

&amp;#x200B;

Can anyone explain the matter please !
## [10][Populate my chat app! 🔥Hello everyone, I deployed a messenger app to firebase and I request you guys to populate my project with more messages, so that it looks more lively! Feel free to suggest new features. Appreciate it!](https://www.reddit.com/r/reactjs/comments/jlzwqv/populate_my_chat_app_hello_everyone_i_deployed_a/)
- url: https://messenger-nowfel2k.web.app/
---

## [11][What's a good React library for avatar selections?](https://www.reddit.com/r/reactjs/comments/jlum7v/whats_a_good_react_library_for_avatar_selections/)
- url: https://www.reddit.com/r/reactjs/comments/jlum7v/whats_a_good_react_library_for_avatar_selections/
---
Something similar to Github for example. With the file upload, cropping etc.

Googling didn't yield many popular results. I wanted to know if there's some standard library that developers default to.

Thanks!

EDIT: Clarification added.
## [12][Spooky Singles: A Phoenix + React site I made that's a dating site for Ghosts. Happy Halloween!](https://www.reddit.com/r/reactjs/comments/jlmy8j/spooky_singles_a_phoenix_react_site_i_made_thats/)
- url: https://spooky.singles/
---

