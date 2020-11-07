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
## [3][How to manage state in React apps with useReducer and useContext hooks](https://www.reddit.com/r/reactjs/comments/jplgzy/how_to_manage_state_in_react_apps_with_usereducer/)
- url: https://dev.to/amanhimself/how-to-manage-state-in-react-apps-with-usereducer-and-usecontext-hooks-4eoc
---

## [4][Best practices for styled components](https://www.reddit.com/r/reactjs/comments/jp5lbe/best_practices_for_styled_components/)
- url: https://www.reddit.com/r/reactjs/comments/jp5lbe/best_practices_for_styled_components/
---
Are there any best practices for styled components? In terms of naming conventions, storing in same or different files? Any good open source examples of medium/large scale applications using styled components?

I've made a decent sized app with styled components but i found it quite annoying to have to navigate the same file to modify the css, and it felt cluttered. Especially in files where the react components are already large.
## [5][Use Redux middleware to reduce redux thunk api action boilerplate](https://www.reddit.com/r/reactjs/comments/jpqma4/use_redux_middleware_to_reduce_redux_thunk_api/)
- url: https://www.reddit.com/r/reactjs/comments/jpqma4/use_redux_middleware_to_reduce_redux_thunk_api/
---
If you have written code like below to call api using redux thunk again and again and got frustrated with it. There is a way to avoid this boilerplate using redux middleware.

    const FETCH_POSTS_LOADING = 'FETCH_POSTS_LOADING'
    const FETCH_POSTS_SUCCESS = 'FETCH_POSTS_SUCCESS'
    const FETCH_POSTS_FAILURE = 'FETCH_POSTS_FAILURE'
    
    export const postLoading = () =&gt; ({ type: FETCH_POSTS_LOADING })
    export const postLoadingSuccess = (response) =&gt; ({
      type: FETCH_POSTS_SUCCESS,
      payload: response
    })
    export const postLoadingFailure = (errors) =&gt; ({
      type: FETCH_POSTS_FAILURE,
      payload: errors
    })
    
    export function fetchPosts() {
      return function (dispatch) {
        dispatch(postLoading())
    
        return axios.request({
          url: 'https://www.example.come/posts.json',
          method: 'get'
        })
          .then(({ data }) =&gt; {
            dispatch(postLoadingSuccess(data))
          })
          .catch(errors =&gt; {
            next(postLoadingFailure(errors))
          })
      }
    }

[Here](https://blog.sockosoft.com/power-of-redux-middleware/) is the blog post that shows the approach to reduce this boilerplate
## [6][Why does this event listener trigger](https://www.reddit.com/r/reactjs/comments/jpqdoh/why_does_this_event_listener_trigger/)
- url: https://www.reddit.com/r/reactjs/comments/jpqdoh/why_does_this_event_listener_trigger/
---
Hey React Community,

Im currently developing a React App using the new hooks (okay, they are not that new, but last time I developed react stuff, class components were still king). Nevertheless im trying to implement a dropdown that closes if you click outside. For that i register a document event handler in a useEffect hook, like this:

```
  useEffect(() =&gt; {
      document.addEventListener('click', handleClickOutside)
      document.addEventListener('keydown', handleEsc)
    return () =&gt; {
      document.removeEventListener('click', handleClickOutside)
      document.removeEventListener('keydown', handleEsc)
    }
  }, [isOpen])
```

But for some odd reason the same click event that triggers the isOpen to toggle also causes this event handler to trigger, one hacky workaround i found is this, but it seems like there needs to be a better solution.

```
  useEffect(() =&gt; {
    setTimeout(() =&gt; {
      document.addEventListener('click', handleClickOutside)
      document.addEventListener('keydown', handleEsc)
    }, 0);
    return () =&gt; {
      document.removeEventListener('click', handleClickOutside)
      document.removeEventListener('keydown', handleEsc)
    }
  }, [isOpen])
```

The handleClickOutside looks like this
```
  function handleClickOutside(e: MouseEvent) {
    if (dropdownRef.current &amp;&amp; !dropdownRef.current.contains(e.target as Node)) {
      onClose()
    }
  }
```

This should actually for the click before the render not have any reference so failing in the first condition and no further click should happen after the render, so not sure why its closing immediately. The timeout can circumvent this by delaying the register of the event listener, but i dont think it should happen in the first place.

Complete example in this code sandbox: [https://codesandbox.io/s/cool-poitras-efnoj?file=/src/App.js](https://codesandbox.io/s/cool-poitras-efnoj?file=/src/App.js)
## [7][Is Redux Toolkit being used in companies](https://www.reddit.com/r/reactjs/comments/jpmrxz/is_redux_toolkit_being_used_in_companies/)
- url: https://www.reddit.com/r/reactjs/comments/jpmrxz/is_redux_toolkit_being_used_in_companies/
---
I've been trying to learn redux a little better and it seems that many tutorials online have still not adopted the toolkit approach to writing redux.

Have companies adopted redux toolkit? I am also looking to improve my skills for applying to jobs so I was wondering if it's more valuable to learn the old way of writing redux than the new way?
## [8][Need help with bundling images with webpack for react component library.](https://www.reddit.com/r/reactjs/comments/jpmr6w/need_help_with_bundling_images_with_webpack_for/)
- url: https://www.reddit.com/r/reactjs/comments/jpmr6w/need_help_with_bundling_images_with_webpack_for/
---
Hey fellow developers!   


I am struggling with something that I have never dealt before. I am creating a component library using react . A few components are importing images and placing them as a background using `style={{backgroundImage: imgSrcHere}}` it loads images just fine while I am developing the library, when I build it with webpack (using a `file-loader` for all images ) I can see these images added to my dist folder. Now, when I publish the lib and import components inside create react app project, the components with image have an inline style as expected `background-image: url("6a4eaca085cb36d75d5e16181091010b.png"); background-position: left center;` But I don\`t see the images being rendered and the network tab explains it with a status code 304.  Obviously, [`http://localhost:3000/`](http://localhost:3000/)`6a4eaca085cb36d75d5e16181091010b.png` is not where the image is located....I don\`t know what to do or even how to google this problem.

For now I solved it with webpacks `url-loader` but I don\`t like this solution because it adds \~300KB to my bundled index.js
## [9][Expenses | Full Stack React and .NET Core Application](https://www.reddit.com/r/reactjs/comments/jpnwoo/expenses_full_stack_react_and_net_core_application/)
- url: https://www.youtube.com/watch?v=MSTYmdzlU-c
---

## [10][how do I start learning react?](https://www.reddit.com/r/reactjs/comments/jpq85v/how_do_i_start_learning_react/)
- url: https://www.reddit.com/r/reactjs/comments/jpq85v/how_do_i_start_learning_react/
---
 have basic JS and web-dev skills but I'm new to react and I find it difficult to keep myself concentrated on a 50 hour Udemy course. but I've always been able to be fully concentrated on doing my own stuff. like developing my own app. so I want to learn react by doing but I'm not sure where to start. obviously you need to learn basic concepts first before start coding. how did you guys do that? should I just overcome the boredom and finish Udemy course first any how?
## [11][what is a good react datasheet package?](https://www.reddit.com/r/reactjs/comments/jpo2xz/what_is_a_good_react_datasheet_package/)
- url: https://www.reddit.com/r/reactjs/comments/jpo2xz/what_is_a_good_react_datasheet_package/
---
I really need something like \[[Notion.so](https://Notion.so)\]([https://www.notion.so/](https://www.notion.so/)) table or excels spreadsheet where I can create formulas. Also, I need it to be very customizable on the level of frontend (the looking) and on the backend, where I can create custom formulas or implement javascript code in the cells.

I tried \[jexcel3\]([https://bossanova.uk/jexcel/v3/](https://bossanova.uk/jexcel/v3/)) but it is very primitive and I users can't edit footer nor drag and drop columns and rows, Also I tried \[ag-grid\]( [https://www.ag-grid.com/example.php](https://www.ag-grid.com/example.php)) but it doesn't support formulas at all, finally I tried \[react-datasheet\]([https://nadbm.github.io/react-datasheet/](https://nadbm.github.io/react-datasheet/)) but users can't edit footers as well and it doesn't have a pattern recognition where I can write 1 and other cell 2 and select the tow cells and drag the selection to have automatically 1,2,3,4,5...
## [12][Jexcel: How can I use updateTable with react?](https://www.reddit.com/r/reactjs/comments/jpnm5g/jexcel_how_can_i_use_updatetable_with_react/)
- url: https://www.reddit.com/r/reactjs/comments/jpnm5g/jexcel_how_can_i_use_updatetable_with_react/
---
I have this problem, I want to use Jexcel with react, we can pass the option updateTable  
 to manage how going to see the cell, how to render, we have this arguments: instance, cell, col, row, val, id  
 where instance and cell are native nodes, then, for change an entry to image, for example, we could use cell.appendChild(...)  
 or cell.innerHTML=...  
 I want to know if exist a right way to use react components, specifically I want put inside a &lt;NavLink /&gt;  
 which have the click event to navigate to another route (using react-router-dom of course). my code:

\--------------------------------------------------------------------------------------------------------------------------  


 class ....{  
.....  
   updateTable(instance, cell, col, row, val, id) {       
if (col === 1) {  
if (!val) return;         
/\*  
\* this NOT WORKING:  
\* let img = &lt;img src={val} style={{ height: 40, width: '100%', objectFit: 'cover' }} /&gt;;         
\* cell.appendChild(img);  
\*/  
cell.innerHTML=\`&lt;img src="${val}" /&gt;\`;  
}  
   }     
componentDidMount() {  
 let opts = {  
...  
freezeColumns: 1,  
tableOverflow: true,  
updateTable: this.updateTable,  
lazyLoading: true  
};  
this.el = jexcel(this.wrapper.current, opts);  
   }  
.....  
}

\--------------------------------------------------------------------------
