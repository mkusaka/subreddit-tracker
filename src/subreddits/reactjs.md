# reactjs
## [1][Beginner's Thread / Easy Questions (March 2020)](https://www.reddit.com/r/reactjs/comments/fbn1p2/beginners_thread_easy_questions_march_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fbn1p2/beginners_thread_easy_questions_march_2020/
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
- **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]
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
## [2][Who's Hiring? [March 2020]](https://www.reddit.com/r/reactjs/comments/fbn65q/whos_hiring_march_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fbn65q/whos_hiring_march_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/ex778e/whos_hiring_feb_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/f44wd7/whos_available_feb_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [3][🎉 Announcing React Table v7! 🎉](https://www.reddit.com/r/reactjs/comments/fgjw7u/announcing_react_table_v7/)
- url: https://github.com/tannerlinsley/react-table/releases/tag/v7.0.0?twitter=20200310
---

## [4][The official Redux+TypeScript template for Create-React-App is now available!](https://www.reddit.com/r/reactjs/comments/fgmydz/the_official_reduxtypescript_template_for/)
- url: https://github.com/reduxjs/cra-template-redux-typescript/releases/tag/v1.0.0
---

## [5][A React component to compare two images](https://www.reddit.com/r/reactjs/comments/fgg0p7/a_react_component_to_compare_two_images/)
- url: https://streamable.com/11iyc
---

## [6][What is the difference between const { foo} = props and using props.foo directly?](https://www.reddit.com/r/reactjs/comments/fgvwvj/what_is_the_difference_between_const_foo_props/)
- url: https://www.reddit.com/r/reactjs/comments/fgvwvj/what_is_the_difference_between_const_foo_props/
---
As the title says.

I noticed that there are codelines in project that I got that looks like this:

    export const FooComponent = (props: FooProps) =&gt; {
      const { id } = props;
    
      return (&lt;div&gt;{props.id}&lt;/div&gt;);
    }

What is the difference between:

    export const FooComponent = (props: FooProps) =&gt; {
      return (&lt;div&gt;{props.id}&lt;/div&gt;);
    }

This is a redux project so I guess its related to mutability of the vars.
## [7][Building Secure React Applications](https://www.reddit.com/r/reactjs/comments/fgx0cl/building_secure_react_applications/)
- url: https://youtu.be/O91hJJ5KMLs?list=PLEx5khR4g7PKMVeAqZdIHRdOwTM1yktD8
---

## [8][Next.js Tutorial - API Routes using SQL Database](https://www.reddit.com/r/reactjs/comments/fguhdy/nextjs_tutorial_api_routes_using_sql_database/)
- url: https://www.youtube.com/watch?v=PxiQDo0CmDE&amp;feature=share
---

## [9][Learning Jest and TDD Suggestions](https://www.reddit.com/r/reactjs/comments/fgweod/learning_jest_and_tdd_suggestions/)
- url: https://www.reddit.com/r/reactjs/comments/fgweod/learning_jest_and_tdd_suggestions/
---
Hey, guys, I'm planning to learn JEST and TDD in React, can you please give any tips and recommendations? I've never written any test and am very new to testing.
## [10][swyx Writing | React Single File Components Are Here](https://www.reddit.com/r/reactjs/comments/fgqzuw/swyx_writing_react_single_file_components_are_here/)
- url: https://www.swyx.io/writing/react-sfcs-here
---

## [11][Building a Minesweeper game using React Hooks](https://www.reddit.com/r/reactjs/comments/fgwy71/building_a_minesweeper_game_using_react_hooks/)
- url: https://www.ivaylopavlov.com/building-a-minesweeper-game-using-react-hooks/#.Xmjg0SWnzDs
---

## [12][Any people here who work with React daily but aren't thrilled by it?](https://www.reddit.com/r/reactjs/comments/fgrflz/any_people_here_who_work_with_react_daily_but/)
- url: https://www.reddit.com/r/reactjs/comments/fgrflz/any_people_here_who_work_with_react_daily_but/
---
The title.

So I've picked up React on the job. There are a lot of things which I like in particular. First, very wide ecosystem. Things like material-ui for example really cuts down the development costs and allows to rapidly build decent looking UI's while also providing design guidelines unlike common component libraries.

I particularly like Redux and how it enforces explicit structure. While it can be a bit unwieldy for smaller  tasks, it's godsent when it comes to debugging a complex application.

And I really appreciate JSX. Liked it from day one actually. Although perhaps angular 1.x left a bit of a bitter aftertaste in regards to html directives with fluff syntax on top which pretends to be html, but isn't.

I really like how it doesn't concern itself with classes all that much as a structural unit of code which is very prevalent in programming. IMO, often out of mare habit or syntactic familiarity rather than some pragmatic benefit.

However, having written a few non trivial applications and having one more down the pipelines to come, I'm not exactly thrilled about some of library features. Namely, how it stupidly rerenders everything down the component tree by default having you always explicitly marking not to rerender. Gotcha with PureComponent back in classes, memo everything by default when it comes to functional. Should component update / memo second argument when it comes to functional components on more complex cases. When I feel that the opposite should be true or some kind of reactivity would be more sensible approach.

Hooks feels like an improvement, but ultimately end up like surrogate for actual FRP. Now it's yet too difficult to draw conclusions, the jury is still out of how it will eventually work out. 

What are your gripes with the library you wish it did a bit differently if it could?
