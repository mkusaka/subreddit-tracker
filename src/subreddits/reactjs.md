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
## [2][Who's Hiring? [August 2020]](https://www.reddit.com/r/reactjs/comments/i1u8a4/whos_hiring_august_2020/)
- url: https://www.reddit.com/r/reactjs/comments/i1u8a4/whos_hiring_august_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/hseduu/whos_available_july_2020/
## [3][React.FC vs JSX.Element](https://www.reddit.com/r/reactjs/comments/i4jx85/reactfc_vs_jsxelement/)
- url: https://www.reddit.com/r/reactjs/comments/i4jx85/reactfc_vs_jsxelement/
---
So another dev and I have both been assigned to a team that has recently lost all of it's devs and we're to clean up the leftover react code.

We've both agreed to get rid of unnecessary class components and move more towards a functional approach, but i've always used 

React.fc and he's always used JSX.element. Example below:

    const HelloWorld: React.FC&lt;Iprops&gt; = ({name}) =&gt; &lt;div&gt;`hi ${name}`&lt;/div&gt;

vs

    function HelloWorld({name}: Iprops): JSX.Element {
         return (&lt;div&gt;`hi ${name}`&lt;/div&gt;)
    }


I get it that React.FC is returning a ReactElement, but isn't JSX.Element also just a ReactElement without types? There would be no functional or performance difference between the two approaches correct? are there any pitfalls of using one over the other?
## [4][Building an offline-first app with React + RxDB + Hasura GraphQL](https://www.reddit.com/r/reactjs/comments/i4q4tw/building_an_offlinefirst_app_with_react_rxdb/)
- url: https://hasura.io/learn/graphql/react-rxdb-offline-first/introduction/
---

## [5][React + GraphQL boilerplate that scales easily](https://www.reddit.com/r/reactjs/comments/i44zx8/react_graphql_boilerplate_that_scales_easily/)
- url: https://blog.graphqleditor.com/graphql-first-fullstack-boilerplate/
---

## [6][Create simple POS with React.js, Node.js, and MongoDB #11: CRUD with Relation](https://www.reddit.com/r/reactjs/comments/i4qto1/create_simple_pos_with_reactjs_nodejs_and_mongodb/)
- url: https://blog.soshace.com/create-simple-pos-with-react-js-node-js-and-mongodb-11-crud-with-relation/
---

## [7][A thread of "advanced" React interview questions](https://www.reddit.com/r/reactjs/comments/i4a45q/a_thread_of_advanced_react_interview_questions/)
- url: https://twitter.com/_paulshen/status/1291065955594862593
---

## [8][Can I use React.forwardRef for mulitple refs in one component](https://www.reddit.com/r/reactjs/comments/i4rtbq/can_i_use_reactforwardref_for_mulitple_refs_in/)
- url: https://www.reddit.com/r/reactjs/comments/i4rtbq/can_i_use_reactforwardref_for_mulitple_refs_in/
---
I'll explain what I need to do because I may not need to do this by creating multiple refs.

My component has multiple things which I need to animate, think of it as a menu hamburger with 3 lines and I need to animate each one separately.

I need to get access to them from my parent component which is just my app.js component because I need to animate them within a timeline which involves other components.

Do I need to create a forwardRef to each element (line) in my menu hamburger that I am wanting to animate?

Or can I just create a forwardRef to the outter HTML element of the hamburger and then access the elements within this (somehow) when I do my animation.

I'm using functional React for this.
## [9][React SSR using react-snap](https://www.reddit.com/r/reactjs/comments/i4rrjw/react_ssr_using_reactsnap/)
- url: https://www.reddit.com/r/reactjs/comments/i4rrjw/react_ssr_using_reactsnap/
---
So, I'm working on a project using ReactJS (CRA) as my front end framework and it has a blog wherein we're using [ghost.io](https://ghost.io), the idea is we're using its API for us to fetch the articles and post / show it in our website, it is working but since blog needs to have unique meta tags, SEO friendly, etc. I'm kinda having a problem with it. React snap only works when you're trying to build and deploy but if I add another one or a new blog it doesn't crawl the new one's SEO and meta tags and if you try to use FB Debugger, meta tags cannot be shown or can't be seen since the website is on production already. My option is re-build and re-deploy every time that there's a new blog for me to actually get the new meta tags or set the meta tags, this one is not really a good solution.

Anyone here can help me or suggest any good dynamic SSR? Wherein I don't have to worry every time that there's a new blog. Using another framework like Next or Razzle is not an option. Thank you!
## [10][need someone help review code on populate react-select](https://www.reddit.com/r/reactjs/comments/i4jx6x/need_someone_help_review_code_on_populate/)
- url: https://www.reddit.com/r/reactjs/comments/i4jx6x/need_someone_help_review_code_on_populate/
---
Hi Guy 

on my update data page use react-select for handle dropdown, and so  require to populate dropdown and make selected options. so I 've tried by using two request 

the first request, populate a full list of dropdown 

the second request, populate other form and make selected

problem here

1. redux state will overwrite when we make first request and continue with second request because I  using same state name try solve by creating new reducer case for receive only option data
```javascript
   switch (type) {
        case BRANCH_FETCHING:
            return { ...state, isFetching: true, isError: false, result: null };
        case BRANCH_FAILED:
            return { ...state, isFetching: false, isError: true, result: null };
        case BRANCH_SUCCESS:
            return { ...state, isFetching: false, isError: false, result: payload };
        case BRANCH_CLEAR:
            return { ...state, result: null, isFetching: false, isError: false };
        case FETCHOPTION_SUCCESS:
            return { ...state, isFetching: false, isError: false, options: payload };
        default:
            return state;
    }
```
2. react select doesn't provide method to set defaultValue after we populate dropdown solve by hold-on until data from reducer ready 
```javascript
 if (branchReducer.result) {
                return (
                    &lt;div class="form-group "&gt;
                        &lt;Select
                            name="pos_machines"
                            defaultValue={branchReducer.result
                                ? branchReducer.result.pos_machines.map(val =&gt; {
                                    return {
                                        'value': val._id,
                                        'label': val.alias
                                    }
                                }) : null}
                            onChange={setMultiselect}

                            isMulti
                            closeMenuOnSelect={false}
                            options={branchReducer.options
                                ? branchReducer.options : null}
                        /&gt;
                    &lt;/div&gt;

                )
            } else {
                return null; 
            }
````
anything that improves code quality 
thank you for your advice
## [11][Issue with hooks](https://www.reddit.com/r/reactjs/comments/i4qx7h/issue_with_hooks/)
- url: https://www.reddit.com/r/reactjs/comments/i4qx7h/issue_with_hooks/
---
I am working on a project right now and trying to learn and use hooks more. In my component I am making a call to my api using axios and trying to store the queried json data in a useState hook like so

Hook:

const \[games, setGames\] = useState(\[\]);

...

[axios.post](https://axios.post)(url, payload)

.then((res) =&gt; {

setGames(res.data);

}

&amp;#x200B;

I know that that is not how I am supposed to use the setGames part of useState but I have tried several other methods and nothing works. I know that the post request is returning the data that I want, but when I try and put it into my components state it comes back as an empty array

&amp;#x200B;

&amp;#x200B;

EDIT: Problem solved. Added a useEffect hook to console.log my games state once it had been updated. Forgot about useState's asynchronous properties. Thanks to everybody for offering guidance
## [12][Shared State with React Hooks and Context API 👾](https://www.reddit.com/r/reactjs/comments/i4mhxk/shared_state_with_react_hooks_and_context_api/)
- url: https://blog.sabinthedev.com/shared-state-with-react-hooks-and-context-api-ckdhvq3eq002rlts1b9m90twt
---

