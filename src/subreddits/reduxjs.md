# reduxjs
## [1][React, Redux and a little bit of math.](https://www.reddit.com/r/reduxjs/comments/fozpe4/react_redux_and_a_little_bit_of_math/)
- url: https://medium.com/@dmitriykharchenko/react-redux-and-a-little-bit-of-math-fe2c9a4a477c?source=friends_link&amp;sk=a826809258bc807c0919cdf63e49c766
---

## [2][Are there any good caching solutions or libraries for Redux?](https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/)
- url: https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/
---
Are there any best practices, guides, or libraries out there to simplify/standardize how caching is handled in Redux?

For now, I'm keeping a boolean in my reducers. If that value is false and I try to fetch data, it'll just no-op.
## [3][Tricks you never knew about Redux DevTool](https://www.reddit.com/r/reduxjs/comments/fkyd85/tricks_you_never_knew_about_redux_devtool/)
- url: https://blog.logrocket.com/redux-devtools-tips-tricks-for-faster-debugging/
---

## [4][Redux Fetch Action Array](https://www.reddit.com/r/reduxjs/comments/fkx5z1/redux_fetch_action_array/)
- url: https://www.reddit.com/r/reduxjs/comments/fkx5z1/redux_fetch_action_array/
---
Hi,  


Is there a way to Dispatch this "products" array somehow to a state ?   


https://preview.redd.it/j25z24mnwhn41.png?width=593&amp;format=png&amp;auto=webp&amp;s=57c7a45584ed17440aac22762c33559056a11e0e

Fetching it with:  


https://preview.redd.it/klg59sjzwhn41.png?width=329&amp;format=png&amp;auto=webp&amp;s=cd9eb470889b4387c3bead056ecfe53fec369cdf

Reducer;

&amp;#x200B;

https://preview.redd.it/rdw4nk05xhn41.png?width=429&amp;format=png&amp;auto=webp&amp;s=bfb82ea4736ba1567bc4a3cbeba97012fda24465
## [5][Filtering, Sorting and Pagination - Advanced Filtering with React and Redux](https://www.reddit.com/r/reduxjs/comments/fkmm0a/filtering_sorting_and_pagination_advanced/)
- url: https://blog.soshace.com/filtering-sorting-and-pagination-advanced-filtering-with-react-and-redux/
---

## [6][Can I dispatch an action every second?](https://www.reddit.com/r/reduxjs/comments/fknhxw/can_i_dispatch_an_action_every_second/)
- url: https://www.reddit.com/r/reduxjs/comments/fknhxw/can_i_dispatch_an_action_every_second/
---
I'm making a countdown timer with Redux, and the timer is stored in Redux. My doubt is about performance/design patterns: There is a problem with dispatching every second? 

&amp;#x200B;

(sorry, my English is bad XD)
## [7][Fetch data from json file =&gt; dispatch =&gt; Action (In production)](https://www.reddit.com/r/reduxjs/comments/fk6iuh/fetch_data_from_json_file_dispatch_action_in/)
- url: https://www.reddit.com/r/reduxjs/comments/fk6iuh/fetch_data_from_json_file_dispatch_action_in/
---
Hey,
Anyone out there knows if this is possible?

Currently fetch from localhost (with json server) for front end development, now i need to deploy this to production.

Any tips on approach with serverless, expressjs, mongodb, nodejs etc? 

Have been searching the internet but only finds articles for local development and saw alot of approaches with serverless express like proxying the localhost, serve static http server, etc.

*Got deployment setup against Netlify
## [8][How do you usually organize your redux stuff?](https://www.reddit.com/r/reduxjs/comments/fjuw3p/how_do_you_usually_organize_your_redux_stuff/)
- url: https://www.reddit.com/r/reduxjs/comments/fjuw3p/how_do_you_usually_organize_your_redux_stuff/
---
Hello everybody,I'm creating my first React SPA and I'm implementing both redux and react-redux for the first time.

Since I'm using react-router to render different page-components, I was thinking to split redux actions, reducers, etc. for each component, but to begin to write these components I've put everything in one folder, like this (I've omitted the rest of the project because there are a lot of files!)

    src
    ├── Selector
    │   └── ...
    ├── public
    │   ├── index.html
    │   ├── index.tsx
    │   └── styles.less
    ├── redux
    │   ├── actions.ts
    │   ├── model.ts
    │   ├── reducers.ts
    │   └── store.ts
    ├── app.tsx
    └── model.ts

So, how you are used to organizing all the redux stuff in your projects? Thank you! 🚀
## [9][React Context API with async hooks as an alternative to state management](https://www.reddit.com/r/reduxjs/comments/fjie4o/react_context_api_with_async_hooks_as_an/)
- url: https://medium.com//when-you-finally-decided-to-rid-yourself-of-redux-8fff0624d2fb?source=friends_link&amp;sk=812065aa580515002a12a051554404e2
---

## [10][reduxform question on "paths"](https://www.reddit.com/r/reduxjs/comments/fig13b/reduxform_question_on_paths/)
- url: https://www.reddit.com/r/reduxjs/comments/fig13b/reduxform_question_on_paths/
---
Hey all, wasn't sure where else to go so I decided to try my luck with Reddit. I am currently trying to make an application for my school assignment, which is where I have to meet the demands of a client in terms of writing a program. 

What I am currently trying to do is make a "glorified spreadsheet".

Anyway, I am trying to make a series of input fields and display said data. Where I am currently stuck is that I want to create a so-called path when 1 option is selected.

 &lt;select *name*="category" *defaultValue*={todo.category} *onChange*={handleChange}&gt;  
                    &lt;option *value*="option1"&gt;option 1&lt;/option&gt;  
                    &lt;option *value*="option2"&gt;option 2&lt;/option&gt;  
 &lt;/select&gt;

So as you can see above, in the input, after a few other things have been inputed by the client, I am trying to get a "path" (not sure what to call it), where selecting 1 of these options provides specific options.

i.e.: selecting option1 (fruit, for example), and i then can display another select of possible fruits. and if option2 is selected then i will do another drop down menu, but of veggies etc.

Hope I am making sense and hope to get some input on what to do. I tried if statements, but that didn't work. I am quite new to programming and the ultimate goal of this is to give my client and my school a fully functioning program (using react, redux, thunk) that fulfils both their criteria.

Sorry for being a dummy, I just got lost :((
