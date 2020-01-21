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
## [2][Who's Available? [Jan 2020]](https://www.reddit.com/r/reactjs/comments/eouupz/whos_available_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eouupz/whos_available_jan_2020/
---
We alternate between hirers (on the 1st of the month) and agencies/freelancers/jobseekers (on the 15th).  
If you are looking to post or reply to React job postings, please check [this month's Who's Hiring post here.][hiring:this month]

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
[available:last month]: https://www.reddit.com/r/reactjs/comments/dxxqdn/whos_available_nov_2019/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/
## [3][Babel 7.8.0 Released: we can now use ECMAScript 2020 new features like.](https://www.reddit.com/r/reactjs/comments/erry3c/babel_780_released_we_can_now_use_ecmascript_2020/)
- url: https://www.reddit.com/r/reactjs/comments/erry3c/babel_780_released_we_can_now_use_ecmascript_2020/
---
New features include:

* The nullish coalescing operator.

`const name = person.fullName ?? "Anonymous";`  which reads "if person.fullName is `null` or `undefined`, then `name = "Anonymous"`.

Same as `||`, but for nullish values, instead of falsy ones.

&amp;#x200B;

* Optional chaining

`const city = person.address?.city; // person.address could be not defined`

&amp;#x200B;

* Dynamic `import()`

`import getTargets from "@babel/helper-compilation-targets";`

`getTargets({`

`browsers: ["last 2 chrome versions"],node: 10`

`}) == { chrome: "77.0.0", node: "10.0.0" }`

&amp;#x200B;

I was able to use them in my immediately Create React App application by upgrading `@babel/core` and `react-scripts`.
## [4][To those who recommend beginners learn to build projects in Vanilla js...](https://www.reddit.com/r/reactjs/comments/erm07o/to_those_who_recommend_beginners_learn_to_build/)
- url: https://www.reddit.com/r/reactjs/comments/erm07o/to_those_who_recommend_beginners_learn_to_build/
---
THANK YOU. I’ve been trying to learn React on and off for a while and I recently did a 2-3,000 line project for a client where I had to manage state manually and work with a less than optimized file structure (my fault). 

Going from a project like that to React I really began to understand the problems that React solved, and that made my comprehension so much better. They weren’t just “concepts” anymore, but they were part of solutions to problems I had. 

Thank you!
## [5][The Opinionated Guide to React - new book by Sara Viera - Folder Structure, CRA/Next/Gatsby, TypeScript, Routing, State Mgmt, Animation, Styling, Forms, Dates, GraphQL and UI Toolkits](https://www.reddit.com/r/reactjs/comments/erd5ox/the_opinionated_guide_to_react_new_book_by_sara/)
- url: https://leanpub.com/opinionated-guide-to-react
---

## [6][The impact of GatsbyJS on performance, accessibility and scalability of modern webapps](https://www.reddit.com/r/reactjs/comments/ereon7/the_impact_of_gatsbyjs_on_performance/)
- url: https://youtu.be/p14g-Sep7HY?list=PLEx5khR4g7PKMVeAqZdIHRdOwTM1yktD8
---

## [7][Progressive Server Rendering + Progressive Client Hydration with React concurrent mode](https://www.reddit.com/r/reactjs/comments/ero1a1/progressive_server_rendering_progressive_client/)
- url: https://www.reddit.com/r/reactjs/comments/ero1a1/progressive_server_rendering_progressive_client/
---
I've been experimenting with the idea of rendering content progressively from the server and streaming it to the client in chunks with React and came up with a workable demo.

https://twitter.com/flexdinesh/status/1217454346780069888?s=20

Dan Abramov responded saying that it's possible to progressively hydrate the client side code now with the experimental concurrent mode and linked to his codesandbox example. I took concurrent mode for a spin and the results are promising. It still has a few gotchas here and there but now the idea of progressively sending HTML chunks from the server and progressively hydrating them in the client is realistic.

Here's an elaborate tweet thread on the results of the demo —

https://twitter.com/flexdinesh/status/1218824439237767168?s=20

In short —

- CSR is super slow to show content in slower networks 'cos of bundle download wait time. 
- SSR can become a bottle neck if one of the api's takes a longer time to generate content. 
- Progressive SSR generates and chunks HTML to the client as and when it is rendered in the server.
- With React suspense and concurrent mode, each HTML chunk can be hydrated as and when it is received.
- In slower networks, chunks will be received before bundle is downloaded. In such scenarios, you show the content to the user even before the bundle is downloaded. 

I think this is pretty awesome and can open up a lot of perf cases.
## [8][React hook that helps developers use google spreadsheet as their data table (API endpoint)](https://www.reddit.com/r/reactjs/comments/era6xf/react_hook_that_helps_developers_use_google/)
- url: https://github.com/idw111/use-google-spreadsheet
---

## [9][Article Suggestions](https://www.reddit.com/r/reactjs/comments/erstpm/article_suggestions/)
- url: https://www.reddit.com/r/reactjs/comments/erstpm/article_suggestions/
---
I'm passionate about simplifying web topics (which I know or can learn) as possible as I can.

For example, you can check out how I simplified React Components Lifecyle in this article - https://dillionmegida.com/p/understanding-react-component-lifecycle/

So I'm throwing out this question.

Do you have any web topic (most especially React) which you find difficult to comprehend?

If yes, state them in the comments and I'll do my best on them.

I'm trying to use this means to learn and to also help out. Thanks.
## [10][Here is a new tutorials series on #GraphQL tooling, this time generating #typescript definitions and #reactjs #hooks using @apollographql CLI and @UriGoldshtein codegen cli. In Feb Im planning state machines using @DavidKPiano ’s xstate and react.](https://www.reddit.com/r/reactjs/comments/eriugf/here_is_a_new_tutorials_series_on_graphql_tooling/)
- url: https://m.youtube.com/playlist?list=PLI7l2IqBd0QuuUjnTxgt0Zb-vOXxv5g9S
---

## [11][swc-compiler with react and webpack](https://www.reddit.com/r/reactjs/comments/ersqwp/swccompiler_with_react_and_webpack/)
- url: https://www.reddit.com/r/reactjs/comments/ersqwp/swccompiler_with_react_and_webpack/
---
swc is a typescript / javascript compiler. It consumes a javascript or typescript file which uses recently added features like async-await and emits javascript code which can be executed on old browsers.

Example with React and Webpack(loader -&gt; swc-loader)

https://github.com/playground-code/react-swc-compiler
## [12][Code readability while using hooks](https://www.reddit.com/r/reactjs/comments/erse78/code_readability_while_using_hooks/)
- url: https://www.reddit.com/r/reactjs/comments/erse78/code_readability_while_using_hooks/
---
On the previous project I worked on, which was using React 15, I formed some practices for keeping my code clean and easy to read. My main objective was keeping the `render()` method as small as possible. To that effect, I would move pieces of `jsx` into their own methods. For instance:

    class SomeComponent extends React.Component {
        render() {
            return (
                &lt;&gt;
                    {this.renderSectionA()}
                    {this.renderSectionB()}
                    {this.renderSectionC()}
                &lt;/&gt;
            )
        }
    }

Now, on my current project, I have the possibility of leveraging hooks, but I still want to keep the practice of separating pieces of `jsx` into separate functions. But, since I dont want my functions to be recreated at every rerender, would something like this make sense ?

    function SomeComponent(props) {
        return (
            {renderSectionA(props)}
            {renderSectionB(props)}
        )
    }
    
    function renderSectionA(props) {}
    
    function renderSectionB(props) {}
        
    

My biggest concern are cases where I will need to send either state, event handlers or variables created in the component to the outer functions, thus having to pass a big number of paramaters.
