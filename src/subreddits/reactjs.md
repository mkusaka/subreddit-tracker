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
## [3][We have a new subreddit snoo thanks to /u/fjellet!](https://www.reddit.com/r/reactjs/comments/ems5l0/we_have_a_new_subreddit_snoo_thanks_to_ufjellet/)
- url: http://share.pickaxe.nl/react-snoo.png
---

## [4][Build a CRUD Application with React and Apollo GraphQL - CodeSource.io](https://www.reddit.com/r/reactjs/comments/en3xz7/build_a_crud_application_with_react_and_apollo/)
- url: https://codesource.io/build-a-crud-application-with-react-and-apollo-graphql/
---

## [5][React-three-fiber v4 released](https://www.reddit.com/r/reactjs/comments/en4z85/reactthreefiber_v4_released/)
- url: https://twitter.com/0xca0a/status/1215617845314445312
---

## [6][React Hooks Tutorial | Build Yelp in React #12 | starting to build Yelp's sub-navigation bar](https://www.reddit.com/r/reactjs/comments/en7g14/react_hooks_tutorial_build_yelp_in_react_12/)
- url: https://www.youtube.com/watch?v=URaGI2AvKoQ
---

## [7][Best Practice when using Context?](https://www.reddit.com/r/reactjs/comments/en76h6/best_practice_when_using_context/)
- url: https://www.reddit.com/r/reactjs/comments/en76h6/best_practice_when_using_context/
---
Following is my sample code when using context

&amp;#x200B;

I have created a class called createDataContext which has following code 

&amp;#x200B;

`import React, { useReducer } from "react";`  
`export default (reducer, actions, defaultValue) =&gt; {`  
 `const Context = React.createContext();`  
 `const Provider = ({ children }) =&gt; {`  
 `const [state, dispatch] = useReducer(reducer, defaultValue);`  
 `const boundActions = {};`  
 `for (let key in actions) {`  
 `boundActions[key] = actions[key](dispatch);`  
`}`  
 `return (`  
 `&lt;Context.Provider value={{ state, ...boundActions }}&gt;`  
 `{children}`  
 `&lt;/Context.Provider&gt;`  
`);`  
  `};`  
 `return { Context: Context, Provider: Provider };`  
`};`  


Following is the code for one of my context 

&amp;#x200B;

`const authReducer = (state, action) =&gt; {`  
 `switch (action.type) {`  
 `case "error":`  
 `return {`  
`...state,`  
`...{ loading: false },`  
`...{ errorMessage: action.payload }`  
`};`  
 `case "customerMobileNumberLogIn":`  
 `return {`  
`...state,`  
`...{ data: action.payload },`  
`...{ loading: false },`  
`...{ errorMessage: "" }`  
`};`  
   
 `case "inProgress":`  
 `return { ...state, ...{ loading: true }, ...{ errorMessage: "" } };`  
 `default:`  
 `return state;`  
  `}`  
`};`

&amp;#x200B;

`const customerMobileNumberLogin = dispatch =&gt; {`

`..do api call with axios` 

`}`

&amp;#x200B;

`export const { Provider, Context } = createDataContext(`  
 `authReducer,`  
  `{`  
 `customerMobileNumberLogin`  
  `},`  
  `[]`  
`);`  


&amp;#x200B;

I am using ReactRouter

`import { Provider as AuthProvider } from "./context/AuthContext";`

`ReactDOM.render(`  
 `&lt;AuthProvider&gt;`  
 `&lt;Router&gt;`  
 `&lt;RootWithAuth /&gt;`  
 `&lt;/Router&gt;`  
 `&lt;/AuthProvider&gt;,`  
 `document.getElementById("root")`  
`);`  
Now suppose I have many api calls and I decide to create separate context file for each api call, won't this approach slow down my app?The context api is causing the same problem as redux. I am not 100 percent sure if this is the right way of doing things but I followed Stephen Grider udemy React Native course for using this approach in Reactjs as well.
## [8][Workflow using Storybook](https://www.reddit.com/r/reactjs/comments/en6nn1/workflow_using_storybook/)
- url: https://www.reddit.com/r/reactjs/comments/en6nn1/workflow_using_storybook/
---
*I hope this is related enough to React.js, I don't know a sub where I can post text where it would fit better.*

I'm working on a React project based on Next.js, using TypeScript and styled-components. I've just implemented Storybook as a system to test my UI components. I am also trying out some testing, but I'm really new to it and I know just the basics. 

I would like to be able to automatically run all my tests when I commit to a feature branch, and if it's possible some sort of canceling/failing the commit if the tests fail. I would also like to prevent merging of a feature branch into master if the tests have failed.  I heard about CI systems but I don't know where to start. Should I use CircleCI? Where do I get started? Are there any developers willing to advise me on where to get started with this? I work on all my projects alone and I use GitHub. Thanks in advance!
## [9][React Apollo Graphql Streaming Soundboard Application!](https://www.reddit.com/r/reactjs/comments/en4qoe/react_apollo_graphql_streaming_soundboard/)
- url: https://www.reddit.com/r/reactjs/comments/en4qoe/react_apollo_graphql_streaming_soundboard/
---
Live streaming an app for my future twitch streams. The app is a soundboard that can be managed on the computer and then the sounds can get triggered from the phone. I want to make it so it can be added as a browser source in obs and streamlabs. Usually I get a few friends on here to help answer my questions and others. Come join in. help out or sit back and learn a bit. [https://twitch.tv/jamie\_337\_nichols](https://twitch.tv/jamie_337_nichols)
## [10][A book-list app](https://www.reddit.com/r/reactjs/comments/en4ft9/a_booklist_app/)
- url: https://www.reddit.com/r/reactjs/comments/en4ft9/a_booklist_app/
---
I want to make a book list app in react. It contains three inputs - Title, Author, ISBN and two buttons - Submit, Delete but I want to break this app into separate components and make communication among these components for managing data flow in these components, so how can I design this app in react components please give me suggestions I am beginner in react. Thank you.
## [11][Can someone explain how this clock app works in written English?](https://www.reddit.com/r/reactjs/comments/en4bsk/can_someone_explain_how_this_clock_app_works_in/)
- url: https://www.reddit.com/r/reactjs/comments/en4bsk/can_someone_explain_how_this_clock_app_works_in/
---
Clock found here: [http://demo1.downloader.tech/](http://demo1.downloader.tech/) 

This clock app is really impressive to me, like, it's beautiful. I'd like to know how it works.

The code for the app can be found here:  [https://www.golangprograms.com/react-js-projects-for-beginners/react-js-compass-clock.html](https://www.golangprograms.com/react-js-projects-for-beginners/react-js-compass-clock.html)

It's only 125 lines of React code and 115 lines of CSS. I don't know how to read all of it though. Can someone explain? It looks pretty complicated so maybe explaining in written English is too big of a challenge...

I understand almost all of the individual lines of code but I don't get how it all comes together. The presence of the $ symbol in the code also confuses me -- I don't know what it does.

I also see this line as being particularly important yet hard to understand: 

    array = length =&gt; Array.from({length}).map((v, k) =&gt; k).map(x=&gt;x+1);
## [12][Uploading DICOM (dcm) files from React App to django server](https://www.reddit.com/r/reactjs/comments/en37bo/uploading_dicom_dcm_files_from_react_app_to/)
- url: https://www.reddit.com/r/reactjs/comments/en37bo/uploading_dicom_dcm_files_from_react_app_to/
---
Hi there,

I need help with uploading a dataset of dicom files from my React App to a Django server.

The hierarchy of the files are as follows:

&gt;main/  
&gt;  
&gt;title/setOne/view  
&gt;  
&gt;1.dcm  
2.dcm  
&gt;  
&gt;title/setTwo/view  
&gt;  
&gt;1.dcm  
2.dcm

&amp;#x200B;

So the use case is that the user is going to upload the "main" directory and a script is supposed to generate the final list of objects to be posted to the server.

Example of request needed to be sent to server:

&gt;Array of objects  
&gt;  
&gt;The objects contain name of set and an array of files.  
&gt;  
&gt;`[`  
&gt;  
&gt;`{`  
&gt;  
&gt;`"name": setOne,`  
&gt;  
&gt;`"scans": [`  
&gt;  
&gt;`{"scan_image": 1.dcm},`  
&gt;  
&gt;`{"scan_image": 2.dcm}`  
&gt;  
&gt;`]`  
&gt;  
&gt;`},`  
&gt;  
&gt;`{`  
&gt;  
&gt;`"name": setOne,`  
&gt;  
&gt;`"scans": [`  
&gt;  
&gt;`{"scan_image": 1.dcm},`  
&gt;  
&gt;`{"scan_image": 2.dcm}`  
&gt;  
&gt;`]`  
&gt;  
&gt;`}`  
&gt;  
&gt;`]`

&amp;#x200B;

The problem I am facing is that for some reason the file objects upon posting them to the server using an axios POST request are all null. I have been looking around so much for a solution to this but with no luck. Hope I find it here.

The GitHub Repo: [https://github.com/omargaber/Lesion-Lab-UI](https://github.com/omargaber/Lesion-Lab-UI)  
The file where all the action happens: [https://github.com/omargaber/Lesion-Lab-UI/blob/master/src/Sidebar.jsx](https://github.com/omargaber/Lesion-Lab-UI/blob/master/src/Sidebar.jsx)
