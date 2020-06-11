# reactjs
## [1][Who's Hiring? [June 2020]](https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [2][Beginner's Thread / Easy Questions (June 2020)](https://www.reddit.com/r/reactjs/comments/gukkex/beginners_thread_easy_questions_june_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gukkex/beginners_thread_easy_questions_june_2020/
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
## [3][Immer v7.0 - new `current` API](https://www.reddit.com/r/reactjs/comments/h0npbn/immer_v70_new_current_api/)
- url: https://github.com/immerjs/immer/commit/467ea5d3b5d062c084ce6d875f8d77e21f26965c
---

## [4][How would you style your components using styled-{components,system}/Material UI?](https://www.reddit.com/r/reactjs/comments/h0zckp/how_would_you_style_your_components_using/)
- url: https://www.reddit.com/r/reactjs/comments/h0zckp/how_would_you_style_your_components_using/
---
I am trying to determine which approach to styling is the best, so I was wondering what the general preference is for those using styled-components and styled-system/Material UI system. There are a few approaches, but my main question is about one-off components, such as when you are making a login screen.

1 - Wrap a component with styled-components and use css to style (inline with the css prop or wrapping with styled)


    const LoginContainer = styled(Paper)`
        width: 680px;
        padding: 75px 120px;
    `;
    const Logo = styled.img`
        margin-top: 0px;
        display: block;
        margin: 0 auto;
    `;
    const Login = () =&gt; (
        &lt;LoginContainer elevation={3}&gt;
            &lt;Logo src={logo} /&gt;
            &lt;LoginForm /&gt;
        &lt;/LoginContainer&gt;
    );


This approach avoids setting a lot of props in the component usage itself. However, considering the 'elevation' prop and how it sets box-shadow, it kind of means CSS props are set at 2 points: in the styled component and via the component props in Login. There's also more general CSS, which might make it easier to read for someone more accustomed to css instead of a styled-system approach. One big disadvantage in my eyes is that it takes a lot more css code for responsive design because of media queries (didn't include that in the examples).


2 - Wrap a component with styled-components, add styled-system props and set the values of those css props when you use that component

    const LoginContainer = styled(Paper)`
      ${spacing}
      ${sizing}
    `;
    const Logo = styled.img`
      ${display}
      ${spacing}
    `;
    const Login = () =&gt; (
      &lt;LoginContainer width={680} px="75px" py="120px" elevation={3}&gt;
        &lt;Logo mx="auto" src={logo} /&gt;
        &lt;LoginForm /&gt;
      &lt;/LoginContainer&gt;
    )


This approach keeps the setting of the css props in one place, making a responsive design easier too. However, defining what those props can be is quite some boilerplate for one-off components. Another point is that I'm not sure how legible the code will be if each prop has multiple values for the responsive design as opposed to solution 1 where you can have clearly defined sections in media queries. The documentation of styled-system also says to avoid this where possible and reuse components.


3 - Avoid creating wrapping styled components by having a Box component that implements all styled-system props and possibly create more nested levels of elements, but stops you from having to define one-off styled components


    const Login = () =&gt; (
      &lt;Paper elevation={3}&gt;
        &lt;Box width={680} px="75px" py="120px"&gt;
           &lt;Box mx="auto"&gt;
             &lt;img src={logo} /&gt;
             &lt;LoginForm /&gt;
           &lt;/Box&gt;
        &lt;/Box&gt;
      &lt;/Paper&gt;   
    );


This approach currently has my preference. I don't have to define one-off components, the props are all in one place and there's a clear distinction between what each level does (the paper determines the background and elevation, the box the internal spacing, etc.). However, there's clearly more elements needed to achieve the same result (I know that the innermost Box can be replaced by making the parent Box display="flex" justifyContent="center" but for the sake of this example and question I kept it this way to show you the accumulation of levels).
Indentation levels/readability can be increased by separating sections into components (more one-off components, hah), but I have no idea how this will affect the performance of the app in the long run.


So my question is to you: Which approach do you/would you take?
If you have different examples, please let us know!
## [5][Next.JS Authentication without Third Party Libraries](https://www.reddit.com/r/reactjs/comments/h0n0tr/nextjs_authentication_without_third_party/)
- url: https://www.mikealche.com/software-development/how-to-implement-authentication-in-next-js-without-third-party-libraries
---

## [6][COVID-19 Globe - My bro and I built a web app that visualizes coronavirus cases as spikes on a globe](https://www.reddit.com/r/reactjs/comments/h0iosx/covid19_globe_my_bro_and_i_built_a_web_app_that/)
- url: https://covid19-globe.web.app/
---

## [7][What React dashboards have you had success with?](https://www.reddit.com/r/reactjs/comments/h0cgt0/what_react_dashboards_have_you_had_success_with/)
- url: https://www.reddit.com/r/reactjs/comments/h0cgt0/what_react_dashboards_have_you_had_success_with/
---
Looking to start a small project without a designer and am interested in using a dashboard with pre-fabricated components to speed things up. 

What have you used for this purpose that you’d recommend to other developers, and what was your experience like?

So far I’ve checked out work from creative tim (seems to be quite popular) as well as appwork, and a few others. 

Project should be fairly simple at first. Is there a case for avoiding a template and just trying to create some basic views/scenes from scratch?
## [8][Looking for a new good react project written by star react programmers](https://www.reddit.com/r/reactjs/comments/h0rvjw/looking_for_a_new_good_react_project_written_by/)
- url: https://www.reddit.com/r/reactjs/comments/h0rvjw/looking_for_a_new_good_react_project_written_by/
---
I am looking for a project written by "react pros" that mainly uses react hooks. And is available on github or something similar.

is there anything you can recommend?
## [9][How to eject axios interceptor if my reactjs app is setup like this?](https://www.reddit.com/r/reactjs/comments/h0xxoi/how_to_eject_axios_interceptor_if_my_reactjs_app/)
- url: https://www.reddit.com/r/reactjs/comments/h0xxoi/how_to_eject_axios_interceptor_if_my_reactjs_app/
---
How to eject axios interceptor if my reactjs app is setup like this?

&amp;#x200B;

App.tsx

    function App() {
    
        React.useEffect(() =&gt; {
            addInterceptor ();
            addResponseInterceptors();
    
            // Cleanup Interceptors
            return () =&gt; {
                // interceptors.removeInterceptor ();
            };
        }, []);
    
      return (
        Component that does a lot of stuff
      );
    }
    
    export default App;
    

Here is the interceptor.tsx

    
    const addInterceptor = () =&gt; {
        axios.interceptors.request.use(request =&gt; {     
              console.log(request);
              return request; });
    }
    
    function removeInterceptor (){
        What to put here in this code context?
    }
    
    export {addInterceptor, removeInterceptor };

How can I eject this interceptor from my App.tsx useeffect method?
## [10][Styled-components good practice?](https://www.reddit.com/r/reactjs/comments/h0wwlw/styledcomponents_good_practice/)
- url: https://www.reddit.com/r/reactjs/comments/h0wwlw/styledcomponents_good_practice/
---
So I started using styled-components to easily apply CSS to my components, but I was wondering, what are the good practices while using styled-components?

&amp;#x200B;

For instance, I'm currently doing a site using a styled component for each of the pages' top component from where I'm handling the style of the page and its child components. But I find that the file is starting to get quite long at some point depending on what you're doing, so I was wondering if I shouldn't have a styled component at the top of each of my page's child components, but then that would start to make a lots of files just for styling (because for clarity I create styled-components in their own file) and I don't really like that. So would it be a good practice to actually define the styled component within the file of the component for which it will be used?

&amp;#x200B;

How do you use styled components and organise your files with them?
## [11][Insight #3 - Use ImmerJS over lodash/set, ImmutableJS or plain JS](https://www.reddit.com/r/reactjs/comments/h0wif9/insight_3_use_immerjs_over_lodashset_immutablejs/)
- url: https://dev.to/sebastienlorber/insight-3-use-immerjs-over-lodash-set-immutablejs-or-plain-js-36bl
---

## [12][Almost finished an online React course from Udemy of around 40 hours - I have some questions and I would love some quick tips of what you think is important for someone new](https://www.reddit.com/r/reactjs/comments/h0wb2w/almost_finished_an_online_react_course_from_udemy/)
- url: https://www.reddit.com/r/reactjs/comments/h0wb2w/almost_finished_an_online_react_course_from_udemy/
---
So, first of all, I honestly enjoyed it. It is my entry into the web development scene, I am coming from languages like Java and Swift. I think he managed to touch many topics and tried to put a lot of weight into doing stuff "the correct way" - which of course is subjective, but in general there is some sort of consensus. I know I could probably Google some of this stuff, find some niche Medium articles etc., but I would also just like to get a small discussion going on. I know could just jump straight into a project, but I would like to have a clean project. So, I have a list of questions basically and I would love your input/tip of something very specific to React, that a beginner should remember. There were some parts where I thought *"there must be a better way to do this"*.

&amp;#x200B;

1. Keep all CSS files in one place, or have something like? `component/post/post.js` and `component/post/post.module.css`
2. There’s a bonus module of around 3 hours of React hooks – Should I do it? My understand is that it’s another way to do local state handling, so instead of classes you have functions with hooks. However, I am honestly quite impressed and enjoy Redux a lot, so the state management I will have in the local components are stuff like loading, error etc. I am bit unsure where the actual gain is using hooks? 
3. Shared functions across the whole application. Is it “okay” to just have a folder in the root of src with functions such as “form validation” which is used across multiple components?
4. Exporting a class only for testing purposes. I have a class x and I want to write jest tests for it – the only place I want to import is in the tests, so in theory I am shipping source code which is modified only for testing purposes, by “export class x” (I guess my concern here is coming from stuff like Java, where I wouldn't want to make a function/variable public just for testing purposes)
5. Nested states and immutability. So if I have a nested object, I need to spread the children as well, so I don't make a shallow copy but a deep copy. Handling state seems to be sensitive topic and I can introduce bugs without potentially having any idea about it, because, even if I don't change state immutable, the chance that I see an error upon testing etc. is pretty low? Basically, he created a function which he imported everywhere - but even this wouldn't help with deeply nested components  
`export constant updateObject = (oldObject, updatedProperties) =&gt; {`  
 `return {`  
 `...oldObject,`  
 `...updatedProperties`  
 `}`  
`}`
6. PropTypes - okay to use? Coming from languages with strict typing, one of the things I really have a tough time with, is the typing - I really miss it. I just feel like it is trying to force something which isn't meant to be there and if no one uses it, I won't start to make it a habit to use it.
7. Networking(!!!) - He uses axios and imports it in several modules and just call the http requests directly in the modules. With some experience from Swift, I am used to have a central networking component which imports the networking library, makes the http calls and returns the objects needed. So I would assume I would have a file like networking.js, which imports axios and all the http requests (can be split up into several files if it is a large project) and returns it back, so this networking.js is imported to the modules instead, so there arent ad hoc networking requests.

I think this is more or less it. I think it is pretty cool and man, Redux is awesome. I am pretty impressed with it and it is super effective. I know that it shouldn't be abused either, but it's cool.

About tips, I honestly dread the CSS part. Are there some cool libraries/websites, where there are some default styling for projects that you like? I don't want a template for now, because I would like to keep control. I will still make my own toolbars, sidedrawers., but some styling wouldn't hurt.

Otherwise, some specific React/JS libraries you just love and you think more people should know about? I don't want to depend too much on 3rd party libraries, but if there's something really cool/good I don't mind. An example of what I mean is something like [https://projects.lukehaas.me/css-loaders/](https://projects.lukehaas.me/css-loaders/). Very trivial and not functionality related really, but just cool and so easy to integrate and just makes it feel nicer. (I know that that is not React related tho). 

Thanks in advance :-)
