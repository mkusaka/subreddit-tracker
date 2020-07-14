# reactjs
## [1][Beginner's Thread / Easy Questions (July 2020)](https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/
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
## [2][Who's Hiring? [July 2020]](https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/ha504b/whos_available_june_2020/
## [3][I made a Snapchat clone in the browser!](https://www.reddit.com/r/reactjs/comments/hqfm86/i_made_a_snapchat_clone_in_the_browser/)
- url: https://v.redd.it/t6bp0tgaqma51
---

## [4][Relay vs Apollo vs Amplify DataStore as client for a GraphQL backend](https://www.reddit.com/r/reactjs/comments/hqzm4l/relay_vs_apollo_vs_amplify_datastore_as_client/)
- url: https://www.reddit.com/r/reactjs/comments/hqzm4l/relay_vs_apollo_vs_amplify_datastore_as_client/
---
I’m building a React Native app and have realised I need a more scalable way of talking to my GraphQL backend. I work more as a quant than a dev so please forgive any misconceptions.

I’ve set up my backend using DynamoDB and have built a GraphQL layer over it using AppSync. The default way AWS suggests talking to AppSync isn’t easily scalable, plus I’ve ended up writing what is essentially a small library of hooks and other helpers to build local persistence using Recoil + AsyncStorage with atoms modelled after my graph schema.

So I’ve got 3 options:

* Apollo: big community, popular, nice API, AppSync uses (used? Can only find sources from 2017) it under the hood, built-in persistence

* Relay: very similar to Apollo, smaller community, API seems more complicated and setup seems a bit longer, but assume the flexibility could be beneficial further down the line

* AWS Amplify DataStore: works immediately, TypeScript type defs generated from my GraphQL schema are designed for this library, have to write hook wrappers, just very wary of committing to an Amplify library as I’ve been burned in the past by bugs and poor documentation (no fault of the devs, I think the Amplify guys are stretched too thin)

Is there a clear winner? Is Apollo still GraphQL king? I’ve heard Relay has come a long way recently.

Thanks
## [5][How to connect two APIs requests?](https://www.reddit.com/r/reactjs/comments/hr0cfk/how_to_connect_two_apis_requests/)
- url: https://www.reddit.com/r/reactjs/comments/hr0cfk/how_to_connect_two_apis_requests/
---
Hey guys, so i'm working on a movie app using [TheMovieDb](https://www.themoviedb.org/) API and the problem is that when I request Trending movies the genre of each movie is just an id and to get the genre text I have to make another request to get all genres and then get the name of the genres from the second request,so my question is how to do it because I tried but I'm stuck right now

    // I called each request in useEffect separately 
    
    
    &lt;div className={classes.SliderContainer}&gt;
     &lt;Slider {...settings}&gt; 
       {movies.map((movie) =&gt; ( 
         &lt;HomeSlider 
            genre={genres[movie.id]} // here's my lastest try but it didn't work 
            src={ImgUrl(movie.backdrop_path)} 
            title={movie.title}     
            key={movie.id} 
            rating={movie.vote_average} /&gt;   
       ))}  
     &lt;/Slider&gt; 
    &lt;/div&gt;

&amp;#x200B;
## [6][Whiteboard: React Hooks 🎣](https://www.reddit.com/r/reactjs/comments/hqzzb3/whiteboard_react_hooks/)
- url: https://blog.ranaemad.com/whiteboard-react-hooks-ckclrvccg0005fls16f1h80mc
---

## [7][Creating A React App With Webpack &amp; Babel From Scratch - Part 1](https://www.reddit.com/r/reactjs/comments/hqrh5e/creating_a_react_app_with_webpack_babel_from/)
- url: https://youtu.be/AC1bCJh_3O8
---

## [8][RxDeep: Reactive (Deep) State Management using RxJS](https://www.reddit.com/r/reactjs/comments/hr1ilx/rxdeep_reactive_deep_state_management_using_rxjs/)
- url: https://loreanvictor.github.io/rxdeep/
---

## [9][Function into class component](https://www.reddit.com/r/reactjs/comments/hr1aqs/function_into_class_component/)
- url: https://www.reddit.com/r/reactjs/comments/hr1aqs/function_into_class_component/
---
Hi,

I am fairly new to react and programming in general. I started programming last year with the basics(CSS, HTML, Java) and since a few weeks I am  working on a project that uses react, typescript, Material-ui etc. The team almost exclusively uses class components  and I have to migrate some of the functional components into class components. So here is the code that I have (Link: [demo](https://material-ui-pickers.dev/demo/datetime-picker) ):

&amp;#x200B;

    import React, { Fragment, useState } from "react";
    import { DateTimePicker } from "@material-ui/pickers";
    
    function BasicDateTimePicker() {
    const [selectedDate, handleDateChange] = useState(new Date());
    
    return (
        &lt;Fragment&gt;
            &lt;DateTimePicker
                value={selectedDate}
                disablePast
                onChange={handleDateChange}
                label="System Date and Time"
                showTodayButton
            /&gt;
        &lt;/Fragment&gt;

            );
    }
    export default BasicDateTimePicker;
    

&amp;#x200B;

And I want to make it into a class Component.

It should, in the end, look something like this:

    import React, { Fragment, useState } from "react";
    import { DateTimePicker } from "@material-ui/pickers";
    
    interface OptionalCompProps {}
    
    interface CompProps extends OptionalCompProps {
    }
    
    type CompAllProps = CompProps &amp; WithStyles&lt;ClassNames&gt; &amp; WithApplicationBaseProps &amp; WithSnackbarProps;
    
    class ControlPanel extends React.Component&lt;CompAllProps, CompState&gt; {
    	static defaultProps: OptionalCompProps = {};
    
    	constructor(props: CompAllProps) {
    		super(props);
    		this.state = {
    		    date: undefined,
    		};
    	}
    
    private selectedDate = () =&gt; useState(new Date());
    private onChange = () =&gt; useState(this.selectedDate);
    
    render(): React.ReactNode {
        const { date } = this.state;
            return(
                &lt;MuiPickersUtilsProvider utils={DateFnsUtil}   
                     &lt;Fragment&gt;	                             
                          &lt;DateTimePicker 
                                disablePast
                                size="small"
                                value={this.onChange}
                                onChange={this.selectedDate}
                                label="System Date and Time: "
                                showTodayButton
                          /&gt;
    	        &lt;/Fragment&gt;
                &lt;/MuiPickersUtilsProvider&gt;
                    );
    }
    
    
    export default withSnackbar(withApplicationBase(withStyles(componentStyles)(ControlPanel)));
    
    

I know this is wrong but its what I have... I am able to change the date/time with it, but the datetimepicker doesnt take the input/unable to accept. Can anyone help converting the funtional component to a class component?
## [10][Functional components, class, or mix?](https://www.reddit.com/r/reactjs/comments/hr0qe8/functional_components_class_or_mix/)
- url: https://www.reddit.com/r/reactjs/comments/hr0qe8/functional_components_class_or_mix/
---
If I'm building a new app should I go with functional based components, class based, or case dependent?

&amp;#x200B;

[View Poll](https://www.reddit.com/poll/hr0qe8)
## [11][Best way to manage/maintain/implement plugins's code with a core application](https://www.reddit.com/r/reactjs/comments/hr0hg1/best_way_to_managemaintainimplement_pluginss_code/)
- url: https://www.reddit.com/r/reactjs/comments/hr0hg1/best_way_to_managemaintainimplement_pluginss_code/
---
Let's first clarify the situation. I'm currently working on a product with a configurable part for each client and a core part where you have the application kernel. For the frontend part we have implemented something like nylas mail ([https://nylas.github.io/nylas-mail/guides/React.html](https://nylas.github.io/nylas-mail/guides/React.html)). The client is able to register some components, changing theme etc... It's possible to have access to some functions from the core trough a specific object defined inside global variable. The client part is builded on his side then every components is registered during runtime. It works very well an most of time this is so far enough for clients. 

&amp;#x200B;

But recently we are facing to a problem. Some clients want even more specific and bigger specific fonctionnalities. And this is not possible to just say "No it's now possible to do that". We must find a way to  implement and have something more flexible. So we are currently looking for the better solution to manage that. Of course with the current implementation we could able to do that and say "Let's just continue to add some fonctionalities inisde the global object from core". But somethimes this is very specific thing and another problem is that by always adding more possibilities it's will becomes untenable. 

Another problem with the current situation is that all dependencies are duplicated inisde the core and client parts and inside client part (just React is passed as a global variable). In this situation,  We have no access to the core DOM inside the client part (for example no access to all providers). And this is very problematic. For example you want do a drag and drop between a component inside the core and another one inside the client. It's not possible... 

The goal now is to find a better solution where we can easily add very specific code while having something maintainable for the next few years.

&amp;#x200B;

A first possible way to implement that could be to inject directly the client code during build time. 

\- A first possiblity could be to add the client as part of the core (as a dependcies) or the opposite and build everything together. 

\- Another solution could be to replace components. React work as components inside a tree so why not say ok let's define specific components inside the client part then add/replace them during build time with a specific babel plugin maybe. It's could be interesting but what's going on if we switch soon to TypeScript... ? 

&amp;#x200B;

I didn't find a lot's of documentation who speak arround the problematic to inject code during build time. 

Feel free to share you feeback, what you are currently using, your idea etc...
## [12][About the comparison algorithm behind useEffect](https://www.reddit.com/r/reactjs/comments/hr00ls/about_the_comparison_algorithm_behind_useeffect/)
- url: https://twitter.com/tinloof/status/1282976284696473601
---

