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
## [2][Who's Hiring? [Jan 2020]](https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/
---
We alternate between **Who's Hiring** (on the 1st of the month, [most recent one here][hiring:most recent]) and **Who's Available** (on the 15th, [most recent one here][available:most recent])

_I am sorry for the missing "Who's Available [Dec 2019]" 🙇‍♂️_

Welcome to **the biggest React job board in the world!** This is like Hacker News' **Who's Hiring** but just for React. Top Level comments must be **Job Opportunities.**

⚠️ NEW: WE ARE REQUESTING EVERYBODY FOLLOW [THE HN Who's Hiring FORMAT][format:hiring:hn]

**Company inc. | Job Title | City/State Location | Full-time/Part-Time | On-site/Remote | (Optional) Salary range | Website jobs page, other hard requirements etc.**

examples:

- **Thorn | San Francisco or Remote (US based) | Full-time Contract | $100k - $150k | Software Engineer | https://www.wearethorn.org/**
- **PolicyStat | Full-Stack Python+Django Software Engineer | Indianapolis, Vancouver, or REMOTE | Full Time | +\$80k**

Please include as much information as possible. **If you are remote-friendly, or open to sponsoring work visas to your country, say so! These are the top 2 questions!**

If you are looking for jobs, send a PM to the poster or post in our [Who's Available thread!][available:most recent]

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/e4pud0/whos_hiring_dec_2019/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/dxxqdn/whos_available_nov_2019/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [3][Styled components v5.0.0 released](https://www.reddit.com/r/reactjs/comments/eo839o/styled_components_v500_released/)
- url: https://github.com/styled-components/styled-components/releases/tag/v5.0.0
---

## [4][React-spring visualizer, tweak your spring configuration](https://www.reddit.com/r/reactjs/comments/eo5qse/reactspring_visualizer_tweak_your_spring/)
- url: https://react-spring-visualizer.com/
---

## [5][Railroad Solitaire: made with React, svg, material-ui](https://www.reddit.com/r/reactjs/comments/eohypz/railroad_solitaire_made_with_react_svg_materialui/)
- url: https://www.reddit.com/r/reactjs/comments/eohypz/railroad_solitaire_made_with_react_svg_materialui/
---
I made recently as a Christmas present for a family member who enjoys this game. It is a variant called Russian Solitaire, give it a try!

https://www.railroadsolitaire.com/

Made with React, material-ui component library, and svg cards.
## [6][Trouble with CSS Modules and modifying text color in Antd](https://www.reddit.com/r/reactjs/comments/eol3dg/trouble_with_css_modules_and_modifying_text_color/)
- url: https://www.reddit.com/r/reactjs/comments/eol3dg/trouble_with_css_modules_and_modifying_text_color/
---
    import React from 'react';
    import { Typography, Icon } from 'antd';
    import { Link } from "react-router-dom";
    import styles from './Header.module.css';
    
    const { Title } = Typography;
    
    const Header = () =&gt; {
        return (
            &lt;div className={ styles.container }&gt;
                &lt;div className={ styles.flexLinks }&gt;
                    &lt;Link to="/home"&gt;
                         &lt;Title className = {styles.link} level = {3}&gt;
                           home
                        &lt;/Title&gt;
                    &lt;/Link&gt;
                    &lt;Link to="/contact"&gt;
                       &lt;Title className = {styles.link} level = {3}&gt;
                         contact
                      &lt;/Title&gt;
                    &lt;/Link&gt;
                &lt;/div&gt;
                &lt;div style={ styles.mobileIcon }&gt;
                    &lt;Icon style={ styles.menu } type="menu" /&gt;
                &lt;/div&gt;
            &lt;/div&gt;
        );
    }
    export default Header;
    
    //Header.module.css file
    
    .container {
        background-image: url("../../images/dark-paths.png");
        padding: 30px;
        text-align: center;
        display: block;
    }
    
    .link {
      color: white;
    }
    
    .flexLinks {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
    }

The links inside Title are still black but should be white. The container styling works and if I use an inline style ie

    &lt;Title style = {{color: 'white'}} level = {3}&gt;
         home
    &lt;/Title&gt;

it works fine. I don't want to do this just to make it more manageable overall and consistent. Btw appending !important after it also works but I'm just trying to find the cleanest solution for future reference.

Thanks!
## [7][(Help) Can't import function in a .test.js file](https://www.reddit.com/r/reactjs/comments/eokulm/help_cant_import_function_in_a_testjs_file/)
- url: https://www.reddit.com/r/reactjs/comments/eokulm/help_cant_import_function_in_a_testjs_file/
---
Hi,

I have a basic create-react-app. I installed all Jest dependencies, and created two files:

\- hello.js, which contains a "hello world" function

\- hello.test.js which imports hello() and test it.

I then run yarn jest. The test only works if the function is directly written in the test file. It crashes if I try to import it. Why?

hello.test.js:

    import { hello } from "./hello";
    
    describe("HELLO WORLD", () =&gt; {
      it("returns hello world", () =&gt; {
        expect(hello()).toEqual("hello world!");
      });
    });

hello.js:

    export const hello = () =&gt; "hello world!";

Also, the Jest crashes because of App.js - which I have not even modified -:

     SyntaxError: /Users/mynamee/Desktop/helloapp/src/App.test.js: Unexpected token (6:31)
    
          4 |
          5 | test('renders learn react link', () =&gt; {
        &gt; 6 |   const { getByText } = render(&lt;App /&gt;);
            |                                ^
          7 |   const linkElement = getByText(/learn react/i);
          8 |   expect(linkElement).toBeInTheDocument();
          9 | });

Thanks for your help!
## [8][Implementing Loading Placeholder with React Content Loader](https://www.reddit.com/r/reactjs/comments/eod92c/implementing_loading_placeholder_with_react/)
- url: https://www.kohei.dev/posts/react-content-loader?hl=en-US
---

## [9][Environment Settings for Build Once Deploy Everywhere Packages in React](https://www.reddit.com/r/reactjs/comments/eojsr9/environment_settings_for_build_once_deploy/)
- url: https://www.daniellittle.dev/environment-settings-for-build-once-packages-in-react
---

## [10][styled media helper v2](https://www.reddit.com/r/reactjs/comments/eojnsl/styled_media_helper_v2/)
- url: https://www.reddit.com/r/reactjs/comments/eojnsl/styled_media_helper_v2/
---
Easy media query usage for styled-components

[https://github.com/dvpnt/styled-media-helper](https://github.com/dvpnt/styled-media-helper)
## [11][Just launched the first stage of our side project - reserve your user ID now! ⚡️](https://www.reddit.com/r/reactjs/comments/eokh9q/just_launched_the_first_stage_of_our_side_project/)
- url: https://www.reddit.com/r/reactjs/comments/eokh9q/just_launched_the_first_stage_of_our_side_project/
---
Hey, redditors!

I and my friend are working on a side project, that's similar to reddit in the core but will have different features built on top. It's focused around creating like-minded communities. [Everchat](https://everch.at/) will be a place where you can send encrypted text and voice messages, images, videos, files, and more, free from any regulation.

Our mission is to secure, protect, and empower the Internet freedom, and the freedom of speech in general, and our vision is to build the most secure group messaging out there, but the time (and hopefully our future community) will tell.

Right now, you can support us with subscribing to our [Product Hunt Ship page](https://www.producthunt.com/upcoming/everchat) if you have a Product Hunt account, or reserve your user ID directly at [Everchat](https://everch.at/).

All feedback is welcome, and I'm hoping to open up some Internet freedom- or content regulation-related discussion here. 🙂
## [12][Which is the better way to validate my title text?](https://www.reddit.com/r/reactjs/comments/eohqfj/which_is_the_better_way_to_validate_my_title_text/)
- url: https://www.reddit.com/r/reactjs/comments/eohqfj/which_is_the_better_way_to_validate_my_title_text/
---
    state = {
        title: "Text",
        error: null
      };
      changeTitle = e =&gt; {
        this.setState({ title: e.target.value });
        this.validateTitle();
      };
    
      validateTitle() {
        this.setState(state =&gt; {
          return {
            error: state.title.length === 0 ? "Title cant be blank" : null
          };
        });
      }

should I put `validateTitle` in the `componentDidUpdate?`
