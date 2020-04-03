# reactjs
## [1][Beginner's Thread / Easy Questions (April 2020)](https://www.reddit.com/r/reactjs/comments/fsqgep/beginners_thread_easy_questions_april_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fsqgep/beginners_thread_easy_questions_april_2020/
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
## [2][Who's Hiring? [April 2020]](https://www.reddit.com/r/reactjs/comments/fsqgf9/whos_hiring_april_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fsqgf9/whos_hiring_april_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/fbn65q/whos_hiring_march_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/fiv53t/whos_available_mar_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [3][Flowify - A new tab extension that I made. It uses React, Framer Motion, and Styled Components. https://addons.mozilla.org/en-US/firefox/addon/flowify/ https://github.com/Etesam913/flowify](https://www.reddit.com/r/reactjs/comments/ftrlpf/flowify_a_new_tab_extension_that_i_made_it_uses/)
- url: https://v.redd.it/7izazpne6gq41
---

## [4][Help Designing an Optimized Approach to Multi-Filtering Data](https://www.reddit.com/r/reactjs/comments/fu8et9/help_designing_an_optimized_approach_to/)
- url: https://www.reddit.com/r/reactjs/comments/fu8et9/help_designing_an_optimized_approach_to/
---
We have a table displaying data and a filter component. In this filter you can filter the different columns in the table.

Ex.

column A) "Event Types" You can filter by checkboxes of the different possible event types

column B) "Number" You can filter by slider 

column C) "Non enumerated text" you can filter by text field you enter value to match

My first approach to filtering this would be,

1) Loop through data
2) Loop through user selected filters

for each data (row in table) see if it matches the filter, if true push to a "filteredData" array.

The problem is this approach only checks if true for ANY of the filters and pushes to "filteredData" I need to ensure the each data row matches ALL filters before pushing. But any approach I take feels slow for the size of dataset.
## [5][The callback in S3.putObject().on('httpUploadProgress', callback) doesn't update a React state array correctly](https://www.reddit.com/r/reactjs/comments/fu87ap/the_callback_in_s3putobjectonhttpuploadprogress/)
- url: https://www.reddit.com/r/reactjs/comments/fu87ap/the_callback_in_s3putobjectonhttpuploadprogress/
---
[I also posted this on Stack Overflow](https://stackoverflow.com/posts/61004673/edit)

[Github repo](https://github.com/BitterDone/ReactAwsS3UploadProgress)

Apologies for the formatting. It looks better on the SO post.

My desired goal is a state array that's updated in real-time with the upload progress of multiple file.

Each callback is able to update the % progress for its file at the correct index in the state array \`uploadProg\`, but it's like they can't see the updated \`uploadProg\` that result from other callbacks.

In the image below, the first \`&gt; (3) \[0,0,0\]\` is logging \`uploadProg\` before updating, the second is logging \`newUploadProg\` after updating.

What am I missing here? Is this a misunderstanding of closures in Javascript?

Thank you for any insights!

[Imgur screencap of uploadProg state array console.log'd](https://i.stack.imgur.com/2C3zY.png)

 

`import React, { useState } from 'react';`  
`import logo from './logo.svg';`  
`import './App.css';`  
`import AWS from 'aws-sdk';`  
`function App() {`  
 `const [selectedFiles, setSelectedFiles] = useState({})`  
 `const [uploadProg, setUploadProg] = useState([])`  
 `var s3 = new AWS.S3({`  
 `apiVersion: '2012-10-17',`  
 `accessKeyId: 'accessKeyId',`  
 `secretAccessKey: 'secretAccessKey',`  
`});`  
 `const handleFile = event =&gt; {`  
 `setSelectedFiles(event?.target?.files)`  
 `setUploadProg(new Array(Object.keys(event?.target?.files).length).fill(0))`  
`}`  
 `const upload = event =&gt; {`  
 `Object.keys(selectedFiles).forEach((key, index) =&gt; {`  
 `console.log(key)`  
 `console.log(selectedFiles[key])`  
 `var params = {`  
 `Body: selectedFiles[key],`  
 `Bucket: "uploadprogress",`  
 `Key: \`exampleobject${index}\`,`  
`};`  
 `s3.putObject(params)`  
`.on('httpUploadProgress', (progressEvent, response) =&gt; {`  
 `const newUploadProg = [...uploadProg]`  
 `const percent = parseInt(100 * progressEvent.loaded / progressEvent.total)`  
 `newUploadProg[index] = percent`  
 `console.log(uploadProg)`  
 `console.log(newUploadProg)`  
 `console.log('______________')`  
 `setUploadProg(newUploadProg)`  
`})`  
`.send((err, data) =&gt; {`  
 `if (err) console.log(err, err.stack);`  
 `else console.log(data);`  
`})`  
`})`  
`}`  
 `return (`  
 `&lt;div className="App"&gt;`  
 `&lt;header className="App-header"&gt;`  
 `&lt;img src={logo} className="App-logo" alt="logo" /&gt;`  
 `&lt;input type="file" name="file" onChange={handleFile} multiple /&gt;`  
 `&lt;button type="button" class="btn btn-success btn-block" onClick={upload}&gt;Upload&lt;/button&gt;`  
 `{uploadProg.map(percent =&gt; (&lt;div style={{`  
 `margin: 5,`  
 `width: 100,`  
 `height: 10,`  
 `backgroundColor: 'green'`  
`}}&gt;`  
 `&lt;div style={{`  
 `height: 8,`  
 `width: percent,`  
 `backgroundColor: 'yellow',`  
`}}&gt;`  
 `&lt;/div&gt;`  
 `&lt;/div&gt;))}`  
 `&lt;/header&gt;`  
 `&lt;/div&gt;`  
`);`  
`}`  
`export default App;`
## [6][Any tips for a code review?](https://www.reddit.com/r/reactjs/comments/fu841r/any_tips_for_a_code_review/)
- url: https://www.reddit.com/r/reactjs/comments/fu841r/any_tips_for_a_code_review/
---
Hi guys, 
I’m interviewing for a company, they code mostly frontend code. My work experience is 100% backend. 
I was requested to code a rather simple single page site. 

I finished it and they want me to perform a code review with them sometime next week. 

I am pretty confident with what I did but I’m not so sure what kind of questions to expect from them. 
If anyone has any experience I would love to get any tips on what should I expect or on what things should I think about before the code review.

Thanks alot for your help.
## [7][Trouble understanding rendering when state is updated](https://www.reddit.com/r/reactjs/comments/fu83u9/trouble_understanding_rendering_when_state_is/)
- url: https://www.reddit.com/r/reactjs/comments/fu83u9/trouble_understanding_rendering_when_state_is/
---
Hey, I am fairly new to react, and have been playing around with the `useState` hook. I have the following code:

    import React, { useState } from 'react';
    import Label from './Label';
    import Button from './Button';
    import './App.css';
    
    
    function App() {
      const [number, setNumber] = useState(0);
    
      function clickHandler() {
        setNumber(number + 1);
      }
      return(
        &lt;div&gt;
          &lt;Label&gt;{number}&lt;/Label&gt;
          &lt;Button clickHandler={clickHandler}&gt;Tap&lt;/Button&gt;
        &lt;/div&gt;
      );
    }
    
    export default App;
    
    import React, { useEffect } from 'react';
    import './App.css';
    
    function Label(props) {
    
      useEffect(() =&gt; {
        // Why is this called?
        console.log("label updated");
      });
      return &lt;div className="Label"&gt;{props.children}&lt;/div&gt;
    }
    
    export default Label;
    
    import React, { useEffect } from 'react';
    import './App.css';
    
    function Button(props) {
    
      useEffect(() =&gt; {
        // Why is this called?
        console.log("button updated");
      });
      return (
        &lt;button className="Button" onClick={props.clickHandler}&gt;
          {props.children}
        &lt;/button&gt;
      );
    }
    
    export default Button;

When I hit the button the log in the `Button` class prints `"button updated"`, I am using the `useEffect` hook to detect this. Why is the button rendering again? From what I have read, only the components related to the state are re-rendered if the state is updated but the `Label` class is not related to the state. Am I missing something here?
## [8][Any good tutorial that shows you how to use the profiling tools for detecting memory leaks?](https://www.reddit.com/r/reactjs/comments/fu82bo/any_good_tutorial_that_shows_you_how_to_use_the/)
- url: https://www.reddit.com/r/reactjs/comments/fu82bo/any_good_tutorial_that_shows_you_how_to_use_the/
---
Any good tutorial that shows you how to use the profiling tools for detecting memory leaks? It's one of the most advanced skills for React developers, but I don't know a single tutorial that covers it. Any suggestion?
## [9][Error handling to be done in the front-end or the back-end?](https://www.reddit.com/r/reactjs/comments/fu54uk/error_handling_to_be_done_in_the_frontend_or_the/)
- url: https://www.reddit.com/r/reactjs/comments/fu54uk/error_handling_to_be_done_in_the_frontend_or_the/
---
I'm creating an application with React as front-end and Flask as back-end. I had a doubt regarding error handling, there are two ways of doing it, one by handling all the errors in the front-end (React) or directly creating error templates in the back-end (Flask) and sending it to the front-end.

Which method is more efficient ?
## [10][Is it best to use prop drilling, context api or redux in this scenario?](https://www.reddit.com/r/reactjs/comments/fu5wrp/is_it_best_to_use_prop_drilling_context_api_or/)
- url: https://www.reddit.com/r/reactjs/comments/fu5wrp/is_it_best_to_use_prop_drilling_context_api_or/
---
My app has a sidebar which can open and close. A `isSidebarOpen` prop is passed down to some child components within the `Header` and `Sidebar` components as they also need to know about it as well. I was wondering if this is ok? Or should I put that `isSidebarOpen` property into a store like redux or else use the Context API?

&amp;nbsp;

    function App(): JSX.Element {
      const [isSidebarOpen, setSidebarOpen] = useState(false);
    
      const toggleSidebar = (): void =&gt; {
        setSidebarOpen(!isSidebarOpen);
      };
    
      return (
        &lt;Layout&gt;
          &lt;Header isSidebarOpen={isSidebarOpen} onMenuButtonClicked={toggleSidebar} /&gt;
          &lt;Sidebar isSidebarOpen={isSidebarOpen} onSidebarLinkClicked={toggleSidebar} /&gt;
          &lt;Footer /&gt;
        &lt;/Layout&gt;
      );
    }

&amp;nbsp;

Both `Header` and `Sidebar` have child components which use this prop as well so there is about 3 levels of prop drilling.
## [11][I'm a React Dev and I Make $10/hr](https://www.reddit.com/r/reactjs/comments/ftxbsa/im_a_react_dev_and_i_make_10hr/)
- url: https://www.reddit.com/r/reactjs/comments/ftxbsa/im_a_react_dev_and_i_make_10hr/
---
This is my first job, so I'm trying to have a good attitude and get as much experience as I can, but I can't help but feel like something's wrong here.

Okay, full disclosure, I make $10/hr PLUS $25 per 'Story Point'. You devs know what I'm talking about. In my shop, we estimate based on 5 points/day, so I average around $22/hr, but it varies.

And when I'm doing un-pointed work, such as meetings, planning, code-review, coordinating with other team members, and even bug-fixing, it's hard for me not to think about the fact that at this moment, for this particular portion of my job, I make $10/hr.

Has anyone ever heard of a pay-by-point arrangement like this? Am I crazy, or is this low-key exploitative? Add to that the fact that our senior Front-End guy just left, and I'm now taking on a huge amount of responsibility with nobody to answer my questions or guide big decisions. For $10/hr.

After six months, it's time to get out of here. So I'm posting in hopes that someone can help me find a new opportunity with competitive compensation, and more opportunity for mentorship and growth.

Would love to hear from anyone who can offer some perspective, or has a line on some work!

(Will send resume/personal info separately. Trying not to let my employer know that I'm looking for work on reddit :)
## [12][Request for help on an open source AWS Amplify project that allows you to create CMS back-ends for your APIs in breakneck speed using react-admin](https://www.reddit.com/r/reactjs/comments/fu4j8t/request_for_help_on_an_open_source_aws_amplify/)
- url: https://www.reddit.com/r/reactjs/comments/fu4j8t/request_for_help_on_an_open_source_aws_amplify/
---
[AWS Amplify](https://aws-amplify.github.io) is great for generating infrastructure from the CLI - love that. Unfortunately managing your data and files easily is pretty crappy, short of trawling DynamoDB/S3 console (yuck) or building a complete CMS (yikes). 

It's all good though, I've got that problem solved (with your help).

## TLDR

I'm looking for people who are familiar with Amplify and React, and want to contribute to an open source project called [`ra-aws-amplify`](https://github.com/mayteio/ra-aws-amplify) that lets you build CMS front-ends for Amplify projects really quickly - with some help from [`react-admin`](https://github.com/marmelab/react-admin). [Here's some good first issues](https://github.com/mayteio/ra-aws-amplify/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) &lt;3

## So what's this breakneck thing you're talking about?

As I mentioned - actually managing the data in DynamoDB is a pain in the ass via the console - and I set out to make this easier with [`ra-aws-amplify`](https://github.com/mayteio/ra-aws-amplify). It's a toolchain that lets you build a CMS front-end for your amplify projects using [`react-admin`](https://github.com/marmelab/react-admin) at breakneck speed.

I was using Amplify for an app platform that manages locations, apps and device tracking. I was able to build the back-end for it within a day with the current features in `ra-aws-amplify` - letting me focus on integrating it into the client app (where the real features are built) which is great.

I want to make it better though. And I'm asking for some help in creating some great open source software. 

## Quick example

Let me first give you an example of how easy it is though. To create a *really* rough back-end that supports `API` and `Auth`, given the following Amplify GraphQL Schema: 

```graphql
type Post @model @auth({rules: [allow: owner]}) {
  id: ID!
  title: String!
  content: String
}
```

Here's your entire Admin component:
```js
// App.js
import React, { useEffect, useState } from 'react';
import { Admin, Resource, ListGuesser, EditGuesser } from 'react-admin';
import { useAmplifyDataProvider, useAuthProvider } from 'ra-aws-amplify';

// grab your amplify generated code
import config from './aws-exports';
import schema from './graphql/schema.json';
import * as queries from './graphql/queries';
import * as mutations from './graphql/mutations';

function App() {
  // dataProvider that connects Amplify API with react-admin
  const dataProvider = useAmplifyDataProvider({ config, schema, queries, mutations });
  // authProvider that connects Amplify Auth with react-admin
  const authProvider = useAuthProvider();

  return
    &lt;Admin dataProvider={dataProvider} authProvider={authProvider}&gt;
      &lt;Resource 
        name="Post" 
        list={ListGuesser} 
        create={EditGuesser} 
        edit={EditGuesser} 
      /&gt;
    &lt;/Admin&gt;
  )
}
```

This will create an admin app that allows you to create, edit, list and delete your Posts. Complete with authentication. Breakneck speed.

## Details

For those who are unacquainted, `react-admin` lets you rapidly build admin screens and forms without worrying too much about building individual components, form validation, and all the lovely jazz that comes with building a CMS from scratch. Here's their demo video:

[![react-admin-demo](https://marmelab.com/react-admin/img/react-admin-demo-still.png)](https://vimeo.com/268958716)

The way it works, is basically this: `react-admin` calls your API in a generic way and expects the response in a certain format. `ra-aws-amplify` (specifically, the `useAmplifyDataProvider` hook) is the missing connector here, that translates react-admin queries into Amplify graphql queries (using the output from `amplify codegen`) and parses the response into a react-admin friendly format. It also [handles authentication for you](https://github.com/mayteio/ra-aws-amplify#authentication-and-sign-in-with-auth).

On top of that, I've also built out tools and components for uploading images and files via `Storage`, including support for `private` and `protected` files - it's all in the [README](https://github.com/mayteio/ra-aws-amplify#image-upload-with-storage), check it out.

## What do you need help with?
Glad you asked. Most pressingly, [I need help with handling many-to-many connections](https://github.com/mayteio/ra-aws-amplify/issues/8).
Other awesome features would be
- [Filtering and sorting models](https://github.com/mayteio/ra-aws-amplify/issues/13)
- [List Pagination with nextToken](https://github.com/mayteio/ra-aws-amplify/issues/11) (essential) 
- [Making models/lists searchable using `@searchable`](https://github.com/mayteio/ra-aws-amplify/projects/1).

If there are typescript buffs out there, typing the package would be great (I have, shamefully, been littering `any`s everywhere) or people who get a kick out of creating rock-solid unit tests, the [dataProvider needs testing](https://github.com/mayteio/ra-aws-amplify/issues/3).

Anyone who wants to contribute - I am more than happy to zoom or chat with you, or even write better comments in the code where you don't understand something. Or talk you through the ins and outs of react-admin if you're not familiar. 

So that's it - let me know if you want to contribute and let's get cracking on this!
