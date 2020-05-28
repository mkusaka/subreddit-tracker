# reactjs
## [1][Beginner's Thread / Easy Questions (May 2020)](https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/
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
## [2][Who's Available? [May 2020]](https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/
---
We alternate between hirers (on the 1st of the month) and agencies/freelancers/jobseekers (on the 15th).  
If you are looking to post or reply to React job postings, please check [this month's Who's Hiring post here.][hiring:this month]

---

If your post or comment is removed wrongly, please [send a message][message:mods] to mods  
because Automods bot is not perfect :)

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
[available:last month]: https://www.reddit.com/r/reactjs/comments/fiv53t/whos_available_mar_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][Gatsby, Website-Building Startup Backed By Index Ventures, Raises $28 Million](https://www.reddit.com/r/reactjs/comments/grm6wt/gatsby_websitebuilding_startup_backed_by_index/)
- url: https://www.forbes.com/sites/davidjeans/2020/05/27/gatsby-website-building-startup-backed-by-index-ventures-raises-28-million/
---

## [4][React does not create the element, return an object, prints [Object object]?](https://www.reddit.com/r/reactjs/comments/gs68qk/react_does_not_create_the_element_return_an/)
- url: https://www.reddit.com/r/reactjs/comments/gs68qk/react_does_not_create_the_element_return_an/
---
`import React from 'react'`  
`class SearchCard extends React.Component{`  
 `ifProps=()=&gt;{`  
 `let user = this.props.user`  
 `return(&lt;li&gt;&lt;p&gt;{user.name}- @{user.username}&lt;/p&gt;&lt;/li&gt;) }`  
 `render(){`  
`return(&lt;div&gt;`  
 `{this.ifProps()}`  
`&lt;/div&gt;)}`  
`}`  
`export default SearchCard`

&amp;#x200B;

&amp;#x200B;

`import React from 'react'`  
`import SearchCard from './SearchCard'`  


`export default function SearchCardIndex(props){`

`function renderIndex(arr){`  
 `return arr.map((user,i)=&gt;{`  
 `debugger`  
 `if(user.attributes){`  
 `user = user.attributes`  
`}`  
 `return &lt;SearchCard user={user}/&gt;`  
   
`})`  
   
`}`  
 `const handleOnChange = (e)=&gt;{`  
 `let users=  props.users`  
 `let newList`  
 `let oldList = users.map(user=&gt;{`  
 `user = user.attributes`  
 `return {name: user.name.toLowerCase(),username: user.username.toLowerCase(),id: user.id}})`  
 `let word = e.target.value`  
 `if(word !==""){`  
 `newList = oldList.filter(user=&gt;{`  
 `return (user.username.includes(word)||(user.name.includes(word))`  
`)})`  
`}else{`  
 `newList=oldList`  
`}`  
 `let ul= document.querySelector(".searchIndex")`  
 `let list = renderIndex(newList)`  
 `ul.innerHTML = list`  
`}`  
 `return(&lt;div &gt;`  
 `&lt;input type="text" value="" onChange={(e)=&gt;handleOnChange(e)}/&gt;`  
 `&lt;ul className="searchIndex "&gt;`  
 `&lt;/ul&gt;`  
 `&lt;/div&gt;)`  
`}`
## [5][Working on NPM Package locally](https://www.reddit.com/r/reactjs/comments/gs5zi3/working_on_npm_package_locally/)
- url: https://www.reddit.com/r/reactjs/comments/gs5zi3/working_on_npm_package_locally/
---
I want to submit a PR for a feature I want within a package that I use in my React project. Is there a way that I can pull down the package locally, make changes on it, and test it within my project to submit my PR?
## [6][Render Component on click](https://www.reddit.com/r/reactjs/comments/gs5qoq/render_component_on_click/)
- url: https://www.reddit.com/r/reactjs/comments/gs5qoq/render_component_on_click/
---
Hi all, 

I am new to programming and I've been using Python and AWS, but now I got a project where I must use React as well. 

I managed to solve most of the tasks, but I got stuck on the next one - In short: When the button on one component is clicked, the content of that component is replaced with different  content

Longer version:

I have a component called Home content, where some basic divs and a button are rendered. What I need to achieve is when that button is clicked, content that is fetched via API should show instead of the current content. I am not sure what is the best way to do it, but my idea was to create a separate component, called Generate URL, that will be in charge of fetching and rendering the data.

I have attached the code example here: [https://jsfiddle.net/me8gLzsp/](https://jsfiddle.net/me8gLzsp/)

When I just pass my Generate component to the Home content in 'normal way' &lt;Genereate /&gt;, it works as it should, but I don't know how to do it on click.

What would be the best approach here?
## [7][96: Pete Hunt on Seven Years of Reconsidering Best Practices | React Podcast](https://www.reddit.com/r/reactjs/comments/gs5qc9/96_pete_hunt_on_seven_years_of_reconsidering_best/)
- url: https://reactpodcast.com/96
---

## [8][Fireact Part 1: Firebase authentication in React](https://www.reddit.com/r/reactjs/comments/gs5h62/fireact_part_1_firebase_authentication_in_react/)
- url: https://medium.com/@chaoming_li/fireact-part-1-firebase-authentication-in-react-b835f6beacc5?sk=26ed71ccd439fe97103416d7f4039cd7
---

## [9][Nowadays - Class or Function Based Components?](https://www.reddit.com/r/reactjs/comments/gs5b9u/nowadays_class_or_function_based_components/)
- url: https://www.reddit.com/r/reactjs/comments/gs5b9u/nowadays_class_or_function_based_components/
---
Hello, I have been learning React from tutorials based on older versions of React where of course Class based components are used as the smart containers and function based components as dumb, stateless components. However at the time this courses was created the react hooks was just released and they weren't widely used. Nowadays since the Hooks are fully-supported what approach should I use:  
1. Only function based components, where some of them use the useState hooks and are containers?

2. Or this old approach dividing components into class or function based?
## [10][Introducing React Modal - an extremely minimal, completely unopinionated modal library](https://www.reddit.com/r/reactjs/comments/grpgp4/introducing_react_modal_an_extremely_minimal/)
- url: https://www.reddit.com/r/reactjs/comments/grpgp4/introducing_react_modal_an_extremely_minimal/
---
Hi all—I've been a member of this subreddit for years and am excited to share the first of many open-source projects that my digital design firm [TRBL](https://trbl.design) is releasing in the future. We're a team of designers and developers and one of our goals is to open-source as much of the code we write as possible. 

The code we release will be aimed squarely at solving problems in a minimal and unobtrusive way, where developers can execute on complex experiences as-designed, without making concessions based on limitations from the packages that they rely on.

React Modal is our first release. We often find effective solutions from modal-based navigation and experiences, and this package was built because we found ourselves fighting against existing modal implementations. We needed a completely unstyled, unopinionated modal implementation that could flex and bend to make our highly varied, deliberate projects come to life as designed.

[https://github.com/trouble/react-modal](https://github.com/trouble/react-modal)

We use this package ourselves and have it in production on many sites and applications. There's a lot more coming from us in the future, including a very large open-source project that will be revealed in a few months from now. 

Would love to know what you all think.
## [11][Two Steps Forward, One Step Back - a critique on Atomic CSS (basscss/tachyons/tailwinds) by one of the pioneers of Atomic CSS](https://www.reddit.com/r/reactjs/comments/grmyrl/two_steps_forward_one_step_back_a_critique_on/)
- url: https://jxnblk.com/blog/two-steps-forward/
---

## [12][Ideas for personal portfolio](https://www.reddit.com/r/reactjs/comments/gs0ohs/ideas_for_personal_portfolio/)
- url: https://www.reddit.com/r/reactjs/comments/gs0ohs/ideas_for_personal_portfolio/
---
Hey, I'm a full stack developer and work mainly in React. I wanted to create a personal portfolio but I'm not very good at UI and UX. So, can anyone provide some ideas or inspiration for the portfolio. Thank you.
