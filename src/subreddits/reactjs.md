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
## [3][Hi Everyone, I need a recommendation for open-source flow builder boilerplates or toolkits that helps you build flows in React/Vanilla JS as fast as possible, like the one shown in the link. Thanks so much!](https://www.reddit.com/r/reactjs/comments/jtd4l9/hi_everyone_i_need_a_recommendation_for/)
- url: https://kapost-files-prod.s3.amazonaws.com/kapost/52f934574e14d3eb0c000856/studio/content/5b5caacdd3c8610027000015/html_bodies/1536711714-5173cfd8-dcc0-4313-8899-41eb366b046f/DF18%20Money%20Slide.png
---

## [4][The Best Icon Library I have Used So Far For React, Vue &amp; Plain HTML/CSS](https://www.reddit.com/r/reactjs/comments/jsw9d9/the_best_icon_library_i_have_used_so_far_for/)
- url: https://medium.com/@fletcherrippon/the-best-icon-library-i-have-used-so-far-for-react-vue-plain-html-css-450ac0ae508c
---

## [5][(Noob Question) I have a functional component which I import, how do i then use a function that I have written within the imported component?](https://www.reddit.com/r/reactjs/comments/jtgiex/noob_question_i_have_a_functional_component_which/)
- url: https://www.reddit.com/r/reactjs/comments/jtgiex/noob_question_i_have_a_functional_component_which/
---
Hi, I have 3 files:

    import React from 'react';
    import './App.css';
    import MovieList from './MovieList';
    
    function App() {
      return (
        &lt;div className="App"&gt;
          &lt;MovieList /&gt;
        &lt;/div&gt;
      );
    }
    
    export default App;
    

\--

    import React from 'react';
    
    const Movie = ({name, price}) =&gt; {
        return (
            &lt;div&gt;
                &lt;h3&gt;{name}&lt;/h3&gt;
                &lt;li&gt;{price}&lt;/li&gt;
            &lt;/div&gt;
        );
    };
    
    export default Movie;

\--

    import React, { useState } from 'react';
    import Movie from './Movie';
    
    const MovieList = () =&gt; {
    
        const [movies, setMovies] = useState([
            {
                name: 'Harry Potter',
                price: '£10',
                id: 23125
            },
            {
                name: 'Little Man',
                price: '£10.99',
                id: 16252
            },
            {
                name: 'Batman: Arkham Knight',
                price: '£5',
                id: 15632
            }
        ]);
    
        function addMovie() {
            setMovies(movies.concat({
                name: 'New Film', price: 'Free', id: 99999
            }));
        };
    
    
        return (
            &lt;div&gt;
                {movies.map(movie =&gt; (
                    &lt;Movie name={movie.name}
                        price={movie.price}
                        key={movie.id} /&gt;
                ))}
            &lt;/div&gt;
        );
    };
    
    export default MovieList;

How can I call the addMovie() function from inside App file please?

&amp;#x200B;

1. I tried adding form with button below the &lt;div&gt; in MovieList, this worked but would re-render the component so it would work for 0.2seconds before it re-rendered
2. I tried using dot notation from within App() - nothing.
3. I tried doing 'import MovieList, {addMovie} from './MovieList' - but it still wouldn't recognise it.

Any help please?
## [6][Looking for a mentor](https://www.reddit.com/r/reactjs/comments/jteoqx/looking_for_a_mentor/)
- url: https://www.reddit.com/r/reactjs/comments/jteoqx/looking_for_a_mentor/
---
TL;DR - I'm an entry-level programmer looking for a React mentor. Have some experience with Angular and now looking for someone to help me understand how React works.

F/26 from Europe here - as the title says, I am looking for someone to help me understanding the basics of React. It would be great if the person knows Angular too because I am already familiar with this framework, but it's not a must.

I've been stuck in the tutorials loop for some time, currently watching an Udemy course from Maximilian Sw*something* and I would like to start building real things now.

What I am looking for is someone to create a board with tickets maybe and then check my PR-s/projects from time to time. The reason why I am looking for this specific thing is that I want to learn the whole process of building an app in React, instead of doing the same things from tutorials over and over again.
I want to understand what I am writing.
I don't need babysitting, I love challenges and searching for solutions, I also respect everyone's time, so I won't be bothering potential mentor every single day.

I know that time is money, I am apologizing in advance if anyone gets offended by this - I honestly feel stupid for asking here for someone to help me without any wage, but I am broke AF and desperate for a job. This is a career change for me, not because of money and fancy title, but because I enjoy programming and the whole process of creating something from nothing makes me happy. Moving tickets from TO-DO to DONE is also fun :)

If there is anyone interested, please contact me in PM or chat.
Thank you.
## [7][Best way to host super-low-traffic nextjs site with a persistent backend storage?](https://www.reddit.com/r/reactjs/comments/jte4xi/best_way_to_host_superlowtraffic_nextjs_site_with/)
- url: https://www.reddit.com/r/reactjs/comments/jte4xi/best_way_to_host_superlowtraffic_nextjs_site_with/
---
Hey! I'm a teacher and I'm looking for the best way to share my materials with my students. I've got a static nextjs site in place, but I'd like to add some personalisation into it.

## Specifications and requirements

- the site is using react+nextjs
- under 50 unique monthly users, all of them in the same region (EU)
- most of them visit the site at most thrice a day
- most of the content is static
- I'd like to add the ability to log in and manually mark some articles as read (which should be persisted on the backend somehow), as well se view your profile and all your read articles
- I don't care much for the backend language, but Haskell, Node.js or Python would be preferable
- I want to pay as little as possible

## My ideas

I think for a site with such a low traffic, the nextjs' incremental rendering + some serverless functions could be all I need; if only I knew how to cheaply connect the serverless functions to a cheap SQLite instance or some similar persistent storage. 

But maybe I'm thinking wrong; I just discovered this whole "serveless" bussiness yesterday. That's why I'm asking for help.
## [8][How I learned reactive programming by re-building RxJs from scratch](https://www.reddit.com/r/reactjs/comments/jt407o/how_i_learned_reactive_programming_by_rebuilding/)
- url: https://blog.bitsrc.io/how-i-learned-reactive-programming-by-re-building-rxjs-from-scratch-975d12e4dde4
---

## [9][Will react-native work with OpenJDK 11 as of november 2020?](https://www.reddit.com/r/reactjs/comments/jtb6xx/will_reactnative_work_with_openjdk_11_as_of/)
- url: https://www.reddit.com/r/reactjs/comments/jtb6xx/will_reactnative_work_with_openjdk_11_as_of/
---
This is my first post here on Reddit. I want to ask if 
OpenJDK 11 would work with react-native or do I need to use 
Version 8?
## [10][Extreme Noob in React and programming in general](https://www.reddit.com/r/reactjs/comments/jswi1p/extreme_noob_in_react_and_programming_in_general/)
- url: https://www.reddit.com/r/reactjs/comments/jswi1p/extreme_noob_in_react_and_programming_in_general/
---
I made my first app, and I am super proud of it. A search image app that spits out images consuming Unsplash API.

link: [https://gauravm-bit.github.io/image-search/](https://gauravm-bit.github.io/image-search/)

Github : [https://github.com/gauravm-bit/image-search](https://github.com/gauravm-bit/image-search)
## [11][The What, Why, and How of Using a Skeleton Loading Screen](https://www.reddit.com/r/reactjs/comments/jt05tc/the_what_why_and_how_of_using_a_skeleton_loading/)
- url: https://medium.com/better-programming/the-what-why-and-how-of-using-a-skeleton-loading-screen-e68809d7f702
---

## [12][What open-source library/tool/component do you wish existed for React?](https://www.reddit.com/r/reactjs/comments/jtfept/what_opensource_librarytoolcomponent_do_you_wish/)
- url: https://www.reddit.com/r/reactjs/comments/jtfept/what_opensource_librarytoolcomponent_do_you_wish/
---
Have you had a very special need that a library could handle but it did not exist? or just simply something that would be nice that we could have it on React? Or something that is already available for Vue or Angular, and you wish it was also available for React?
