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
## [3][Is redux really a good idea?](https://www.reddit.com/r/reactjs/comments/epxavs/is_redux_really_a_good_idea/)
- url: https://www.reddit.com/r/reactjs/comments/epxavs/is_redux_really_a_good_idea/
---
I was a redux fanboy and would defend everything about redux(when I was creating redux project). But now I'm maintaining several redux projects, I start to rethink and now posting this topic to ask for a discussion.

And when talking about redux, I'm referring to essentially all the state management tools on the market. Because other than maybe mobx-state-tree, everything is the same. Yeah, some might not ask for immutability and some may have a different way to write action object, but other than that everything propose the exact same structure of state management.

It sure provides benefit. As managing state become automatic as you don't have to make any decision as where to put the state and how to fetch them. And with dev tools you can pretty much click around the app and know which user input result in what state changes.

That said, still

1. I don't think everything in a global store the best practice. It might be a pleasure to write at first, but will be a nightmare for others to maintain. Adding a little pop window will require you the knowledge of so many implantation detail as in file structure and also writing code snippets everywhere. And the main defect of this approach is most states are not really global, simply stripe it away from where we expect it to be (the component with related UI) and cluster them in a big store is a huge mental setback for the maintainer of the code. Think that we all agree to gather code by logic approximation rather than format(css/js/html) nowadays, but redux keep those code that should stay together far away.
2. Having only some of the states in global store is also a headache. The problem is what could be the criteria for putting a state into the global store? How global should a state to be consider global? It might not be a problem for small project from a single developer but sure it can be hard to follow some vague rules collectively.
3. The boilerplate of redux is one of the worst to write. Switch....case....payload... action.type.... **Basically action.type is function name and action.payload is function parameters and reducer is the function body.** Can't we have a better implementation so we don't have to write a function in 2 files and 3 places? To me this structure is just not ideal at all.
4. Before anyone say anything about useReducer and context, they are essentially the same thing, except now we can have multiple stores, all the problem above still persist.

Ok enough about my rant. I'm just curious how do everyone mange their projects' state now in the 2020, and what are you general opinion about redux and other state management tools?
## [4][How To Not Have A Mess with React Hooks &amp; Redux](https://www.reddit.com/r/reactjs/comments/epv832/how_to_not_have_a_mess_with_react_hooks_redux/)
- url: https://www.reddit.com/r/reactjs/comments/epv832/how_to_not_have_a_mess_with_react_hooks_redux/
---
some of my lessons learned and what worked out **for me** using **redux** and **hooks**:

[https://orizens.com/blog/how-to-not-have-a-mess-with-react-hooks-and-redux/](https://orizens.com/blog/how-to-not-have-a-mess-with-react-hooks-and-redux/)
## [5][D3.js dashboard tutorial with react and cube.js](https://www.reddit.com/r/reactjs/comments/epm3na/d3js_dashboard_tutorial_with_react_and_cubejs/)
- url: https://d3-dashboard.cube.dev/
---

## [6][Suggestions for ideas for someone giving their first talk?](https://www.reddit.com/r/reactjs/comments/epy8aq/suggestions_for_ideas_for_someone_giving_their/)
- url: https://www.reddit.com/r/reactjs/comments/epy8aq/suggestions_for_ideas_for_someone_giving_their/
---
I've been a developer for 7 years but cannot for the life of me think of what I want to present as my first talk, I was thinking what are some of the interesting topics to cover in the react world currently? I guess hooks etc is overdone at this stage but is an example of an idea.

I know this is a bit vague but I have a solid understanding of most react things.

Any suggestions would be hugely appreciated, thanks!
## [7][Pros and cons of nesting Routes with react-router](https://www.reddit.com/r/reactjs/comments/epwrmq/pros_and_cons_of_nesting_routes_with_reactrouter/)
- url: https://www.reddit.com/r/reactjs/comments/epwrmq/pros_and_cons_of_nesting_routes_with_reactrouter/
---
For anyone that happens to use react-router or react-router-dom for routing with their projects, do you nest your routes? Why or why not?
## [8][How to avoid prop drilling?](https://www.reddit.com/r/reactjs/comments/epz6f8/how_to_avoid_prop_drilling/)
- url: https://www.reddit.com/r/reactjs/comments/epz6f8/how_to_avoid_prop_drilling/
---
Anyone who has been using react for a while knows what a nightmare prop drilling can be.  My question is what is the best way to avoid it.  I know context api and redux are two possible solutions.  Does anyone know of any others?  Thanks for your help.
## [9][Ryan Florence's preferred stack: Hasura + GraphQL + Cloudflare Workers + Webpack + React + Relay + React Router + Reach UI + Tailwind CSS](https://www.reddit.com/r/reactjs/comments/eposic/ryan_florences_preferred_stack_hasura_graphql/)
- url: https://twitter.com/ryanflorence/status/1217875244674797569
---

## [10][Conditional Rendering Methods in React](https://www.reddit.com/r/reactjs/comments/epvu5t/conditional_rendering_methods_in_react/)
- url: https://blog.soshace.com/conditional-rendering-methods-in-react/
---

## [11][How do you add an attribute that has no value to a DOM element from a component you're not allowed to modify?](https://www.reddit.com/r/reactjs/comments/epy6eb/how_do_you_add_an_attribute_that_has_no_value_to/)
- url: https://www.reddit.com/r/reactjs/comments/epy6eb/how_do_you_add_an_attribute_that_has_no_value_to/
---
&amp;#x200B;

[https://codesandbox.io/s/react-16-t3jrr](https://codesandbox.io/s/react-16-t3jrr)

How  do you add an attribute that has no value to a DOM element from a  component you're not allowed to differently modify? I am trying to do  this, but it seems because I am using functional components and they  don't get initialized the value of ref.current is undefined. Do I need  to modify the Wrapper? Is there a way to do this without modifying the  Wrapper? I rather not do that, but if there's no other way I don't mind  doing it. I need to add an attribute to the button inside FeatureRef.
## [12][React JS to traverse a folder and subfolders for files with specific extension](https://www.reddit.com/r/reactjs/comments/epy5ff/react_js_to_traverse_a_folder_and_subfolders_for/)
- url: https://www.reddit.com/r/reactjs/comments/epy5ff/react_js_to_traverse_a_folder_and_subfolders_for/
---
I have a button that I can upload a file, which is handled by a function that the file path is stored in etc.. 

upload button:

            &lt;Button
              className={classes.third}
              variant="contained"
              component="label"
            &gt;
              Upload File
              &lt;input
                type="file"
                style={{ display: "none" }}
                onChange={this.handleUpload}
              /&gt;
            &lt;/Button&gt;

function:

    handleUpload = e =&gt; {
     this.setState({
      [e.target.files]: e.target.files,
    });
    };

Is there anyway to just choose a folder and a function to traverse the folder tree for specific extensions? and it is stored in an array. I have done this in python but I am trying to move towards a browser based app. I have been looking everywhere online for a library and examples but I havent found much with React. Could someone possibly point me in the right direction? Thanks!
