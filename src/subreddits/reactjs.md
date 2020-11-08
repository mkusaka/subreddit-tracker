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
## [3][I created TwitterFOMO—see the most liked tweets in webdev at one place. Used Next.js, ChakraUI, and react-spring. (Soon to be open-sourced!)](https://www.reddit.com/r/reactjs/comments/jpxw93/i_created_twitterfomosee_the_most_liked_tweets_in/)
- url: https://v.redd.it/ny9bwvo8qvx51
---

## [4][React js entry/junior level interview questions. HELP](https://www.reddit.com/r/reactjs/comments/jqan4l/react_js_entryjunior_level_interview_questions/)
- url: https://www.reddit.com/r/reactjs/comments/jqan4l/react_js_entryjunior_level_interview_questions/
---
Im having an Interview this week Im terrified and I want to prepare myself.

What are the commonly ask reactjs questions for entry/junior level developer? 
Another questions is do they ask “leet code” type of questions during an interview for a front end position?

Thank you!
## [5][My window-styled portfolio - Winport](https://www.reddit.com/r/reactjs/comments/jq9st0/my_windowstyled_portfolio_winport/)
- url: https://khang-nd.github.io
---

## [6][Beginner question about to-do list](https://www.reddit.com/r/reactjs/comments/jqav9g/beginner_question_about_todo_list/)
- url: https://www.reddit.com/r/reactjs/comments/jqav9g/beginner_question_about_todo_list/
---
In what part of the code are we setting that the newly added todos(through addTodo func)are being set to isCompleted:false so that they are not crossed out from the beginning(since we are only adding text in addTodo)?

    import React,{useState,useEffect} from 'react';
    import './App.css';
    
    const LOCAL_STORAGE_KEY='todo_app.todos'
    
    function Todo({todo,index,completeTodo,removeTodo}){
    return (
    
    &lt;div className='todo'&gt;
    &lt;div style={{textDecoration:todo.isCompleted? 'line-through':null}}&gt; 
      {todo.text}
    
    &lt;button onClick={()=&gt;completeTodo(index)}&gt;Complete&lt;/button&gt;
    &lt;button onClick={()=&gt;removeTodo(index)}&gt;X&lt;/button&gt;
    &lt;/div&gt;
    &lt;/div&gt;
    
    
    )
    
    }
    
    function TodoForm({addTodo}){
    const[value,setValue]=useState('')
     
    function submitHandler(e){
    e.preventDefault();
    if(!value){
      return
    }
    addTodo(value)
    setValue('')
    } 
    
    return(
      &lt;form onSubmit={submitHandler}&gt;
      &lt;input type="text" value={value} onChange={e=&gt;setValue(e.target.value)}/&gt;
        &lt;/form&gt;
    )
     
    
    
     
      
    }
    
    function App(){
    const[todos,setTodos]=useState([
      {text:'Go out with a girl',
    isCompleted:false
    },
    {text:'Learn React',
    isCompleted:false
    },
    {text:'Read a book',
    isCompleted:false
    },
    ])
    useEffect(()=&gt;{
    const storedTodos=JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY))
    if (storedTodos){
    setTodos(storedTodos)
    }
    },[])
    useEffect(()=&gt;{
      localStorage.setItem(LOCAL_STORAGE_KEY,JSON.stringify(todos))
    },[todos])
    function addTodo(text){
     const newTodos=[...todos,{text}]
     setTodos(newTodos)
    }
    function completeTodo(index){
    const newTodos=[...todos]
    newTodos[index].isCompleted=true
    setTodos(newTodos)
    
    }
    function removeTodo(index){
    const newTodos=[...todos]
    newTodos.splice(index,1)
    setTodos(newTodos)
    
    }
    return(
    &lt;div className="app"&gt;
     &lt;div className="todo-list"&gt;
      {todos.map((todo,index)=&gt;(
        &lt;Todo todo={todo} index={index} key={index} completeTodo={completeTodo} removeTodo={removeTodo}/&gt;
      ))}
      &lt;TodoForm addTodo={addTodo}/&gt;
     &lt;/div&gt; 
    &lt;/div&gt;
      )
    
    
    }
    export default App;

&amp;#x200B;
## [7][Displaying long list of related articles](https://www.reddit.com/r/reactjs/comments/jqbora/displaying_long_list_of_related_articles/)
- url: https://www.reddit.com/r/reactjs/comments/jqbora/displaying_long_list_of_related_articles/
---
I made this site for my tai chi students near the beginning of covid, [https://breathetaichi.com](https://breathetaichi.com/)  
It  was my first project, and I'd like to go back to it and make it  portfolio worthy.  My only other piece I really like to show is [https://protopla.net](https://protopla.net/)

I'd  love to hear any critiques of either site, but the my main question  would be is there a standard way to display as much info as I have on my  tai chi site?  It feels kinda messy, and I'm too close to it  to be able to see what it should be. 

Thanks!
## [8][Nextjs supported CDN outside US](https://www.reddit.com/r/reactjs/comments/jqblx8/nextjs_supported_cdn_outside_us/)
- url: https://www.reddit.com/r/reactjs/comments/jqblx8/nextjs_supported_cdn_outside_us/
---
Any one knows a CDN that belongs to a non-US startup which can host NextJS apps?
## [9][implementing build react app inside another website](https://www.reddit.com/r/reactjs/comments/jq7hkl/implementing_build_react_app_inside_another/)
- url: https://www.reddit.com/r/reactjs/comments/jq7hkl/implementing_build_react_app_inside_another/
---
Hey guys, I made a chatbot with reactjs and ran npm run build.

Now I wat to use the built code inside other websites. So my chatbot will popup on other websites.

How can I implement it to other codebases?
## [10][Why my asynchronous function findOne request to wrong pat?](https://www.reddit.com/r/reactjs/comments/jqb1jk/why_my_asynchronous_function_findone_request_to/)
- url: https://www.reddit.com/r/reactjs/comments/jqb1jk/why_my_asynchronous_function_findone_request_to/
---
 I created a function to check if in my database a user exists according to the `sub` parameter.

**Front-end**

    async function exist (sub)  {  
    const response = await InfoDataService.get(sub);   
    console.log(response);       
    }    

**InfoDataService**

    import http from "../http-common.js";   
    const get = sub =&gt; {  
    return http.get(`/users/${sub}`);   
    };   
    export default { get, };  

**Controller**

    // Find a single User with sub  
    exports.findOne = (req, res) =&gt; {   
    const sub = req.params.sub;  
    User.findOne({sub})     
    .then(data =&gt; { res.send(data);     
    })     
    .catch(err =&gt; {   
    res.status(500).send({ message: "Error retrieving User with sub=" + sub});  });  
     };  

**Route**

    module.exports = app =&gt; {  
    const users = require("../controllers/user.controller.js");  
    const router = require("express").Router();   
    
    // Retrieve a single Tutorial with sub   
    router.get("/:sub", users.findOne);  
     
    app.use('/api/users', router);  
    };  

The path for post is `api/users`, while for get it is `api/users/:sub`. Checking the network I found that the asynchronous function makes the request to `api/users/undefined`. Where would the error be?

The problem is in the asynchronous function. Create, update, getAll etc ... also work well in this API. I also have an API for published posts, structured the same way and it works. I don't understand why only the asynchronous function follows a different and wrong path.
## [11][How to select options in a custom dropdown menu with arrow keys?](https://www.reddit.com/r/reactjs/comments/jqauun/how_to_select_options_in_a_custom_dropdown_menu/)
- url: https://www.reddit.com/r/reactjs/comments/jqauun/how_to_select_options_in_a_custom_dropdown_menu/
---
In a custom dropdown, There is an input type text and a different container floating below the input which contains the list of options. How do I select the options in the container floating below using the arrow keys when I focus on the input? (I've attached a screen-capture below)

&amp;#x200B;

Here's a link to the trimmed down version of my project if anyone wants to see the actual code: 

[https://github.com/paritosh-34/absolute\_dropdown](https://github.com/paritosh-34/absolute_dropdown)

The project is created with next.js and typescript.

&amp;#x200B;

[ScreenCapture](https://reddit.com/link/jqauun/video/y8vke2e1d0y51/player)
## [12][Disable page scrolling, but allow sidebar scrolling](https://www.reddit.com/r/reactjs/comments/jqarb2/disable_page_scrolling_but_allow_sidebar_scrolling/)
- url: https://www.reddit.com/r/reactjs/comments/jqarb2/disable_page_scrolling_but_allow_sidebar_scrolling/
---
Hi!  


I have a feeling this is a simple issue, but I can't seem to figure it out. I've built a simple website with React and Next.JS. On mobile, I have a navigation bar that opens from the side and has position: absolute. When the sidebar is opened, I do not want users to be able to scroll the page. However, since the content on the sidebar is more than the average screen can display, I do want users to scroll inside the sidebar. Is there a universal or commonly used way to approach this?  


Thanks in advance,  
Quinten
