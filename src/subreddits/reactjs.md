# reactjs
## [1][Beginner's Thread / Easy Questions (October 2020)](https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/)
- url: https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Ask about React or anything else in its ecosystem :)

Stuck making progress on your app?  
Still Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** of reply by
   1. adding minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz] links
   1. describing what you want it to do (ask yourself if it's an [XY problem](https://meta.stackexchange.com/questions/66377/what-is-the-xy-problem))
   1. things you've tried. (Don't just post big blocks of code!)
1. **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
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
- and these React Hook recipes on [useHooks.com][usehooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[usehooks.com]: https://usehooks.com/
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
## [2][Help the React team write new docs—take the 2020 React Community Survey](https://www.reddit.com/r/reactjs/comments/j5iqj3/help_the_react_team_write_new_docstake_the_2020/)
- url: https://www.surveymonkey.co.uk/r/T58DPNS
---

## [3][Keeping up with React Libraries](https://www.reddit.com/r/reactjs/comments/jicfxg/keeping_up_with_react_libraries/)
- url: https://maxrozen.com/keeping-up-with-react-libraries/
---

## [4][I was able get copying cells from my react-virtualized grid to excel working! Here's how...](https://www.reddit.com/r/reactjs/comments/jieg9k/i_was_able_get_copying_cells_from_my/)
- url: https://v.redd.it/npexh5w9xfv51
---

## [5][Useful Hooks for Tired React Developers](https://www.reddit.com/r/reactjs/comments/jhsqeb/useful_hooks_for_tired_react_developers/)
- url: https://medium.com/javascript-in-plain-english/useful-custom-hooks-for-tired-react-devs-f2f602dc754f
---

## [6][Unit Testing React App with Jest/Enzyme](https://www.reddit.com/r/reactjs/comments/jidb1e/unit_testing_react_app_with_jestenzyme/)
- url: https://www.reddit.com/r/reactjs/comments/jidb1e/unit_testing_react_app_with_jestenzyme/
---
Hey folks,

Hope you are doing great. 

I have recently started exploring unit testing. I am a bit confused that what kind of tests should I write.

Example: Counter App has three buttons Increment/Decrement and a Reset button.

Should I write a test that checks whether a display field is present or not, because its obvious that it should be there.

I have written some tests for Counter App. Please have a look at it and let me know if I am going in the right direction  


Note: I am newbie

    describe('Counter', () =&gt; {
      it('renders Counter component without crashing', () =&gt; {
        shallow(&lt;Counter /&gt;);
      });
    
      it('contains a helpful text which explains our App', () =&gt; {
        const title = 'This is a Counter App. Use - button to decrease value and + button to increase value.';
        const wrapper = shallow(&lt;Counter /&gt;);
        expect(wrapper.contains(title)).toBe(true);
      });
    
      it('checks if inc/dec buttons are there', () =&gt; {
        const wrapper = shallow(&lt;Counter /&gt;);
        const incButton = wrapper.find('#increment').dive().text();
        const decButton = wrapper.find('#decrement').dive().text();
        expect(incButton).toEqual('+');
        expect(decButton).toEqual('-');
      });
    
      it('checks if initial value of counter is zero', () =&gt; {
        const wrapper = shallow(&lt;Counter /&gt;);
        const counter = wrapper.find('#counter-input').props().value;
        expect(counter).toBe(0);
      });
    
      it('checks if Increment button increases value by 1', () =&gt; {
        const wrapper = shallow(&lt;Counter /&gt;);
    
        const incButton = wrapper.find('#increment');
        incButton.simulate('click');
    
        const counter = wrapper.find('#counter-input').props().value;
        expect(counter).toBe(1);
      });
    
      it('checks if Decrement button decreases value by 1', () =&gt; {
        const wrapper = shallow(&lt;Counter /&gt;);
    
        const incButton = wrapper.find('#decrement');
        incButton.simulate('click');
    
        const counter = wrapper.find('#counter-input').props().value;
        expect(counter).toBe(-1);
      });
    
      it('checks Reset button functionality', () =&gt; {
        const wrapper = shallow(&lt;Counter /&gt;);
    
        const incButton = wrapper.find('#increment');
        incButton.simulate('click');
    
        const counter = wrapper.find('#counter-input').prop('value');
        expect(counter).toBe(1);
    
        wrapper.find('.reset-btn').simulate('click');
        expect(wrapper.find('#counter-input').prop('value')).toBe(0);
      });
    });
## [7][Is it tough to get projects as a freelancer in React Js ?](https://www.reddit.com/r/reactjs/comments/jidlhv/is_it_tough_to_get_projects_as_a_freelancer_in/)
- url: https://www.reddit.com/r/reactjs/comments/jidlhv/is_it_tough_to_get_projects_as_a_freelancer_in/
---
I have studied react js for the past 3 months and got myself convenient in every topic of React as well as its libraries and redux.
I have designed and developed a covid tracker, food delivery app and a shopping app using react. 
Will it be difficult for me to start freelancing in React although if I have a good looking portfolio with my projects hosted on the server for the client to refer ?

Any advice will be appreciated.
## [8][Estimating an UI application](https://www.reddit.com/r/reactjs/comments/ji7ain/estimating_an_ui_application/)
- url: https://www.reddit.com/r/reactjs/comments/ji7ain/estimating_an_ui_application/
---
How do you usually arrive at estimates? I have always faced backlashes on my gut feel estimates. Do we have any easy/ scientific way of estimating so that gut feel estimates can be more accurate. 


Consider this for a UI application based on react.
## [9][Using React with Parcel.js](https://www.reddit.com/r/reactjs/comments/ji9zph/using_react_with_parceljs/)
- url: https://www.reddit.com/r/reactjs/comments/ji9zph/using_react_with_parceljs/
---
So I've been interested in trying out Parcel and finally had some motivation. I've set my project us and I'm having problems with Parcel not reloading and seemingly not watching my files correctly.

I'm not doing anything crazy, I think.. I have a `src` folder with an `index.html` and `index.tsx` and my root `style.scss`. Whenever I save these they are immediately shown in my browser.

Inside src, I have a `components` and `scenes` file. When I import my `Home` component (from the `scenes` directory), this is immediately shown. However, if I make any changes to `Home` or any components in the `components` directory, nothing is ever updated which is really annoying and I can't work out what I'm doing wrong. 

The styles are changes in the components are updated if I stop and start the watcher (`yarn dev`) but otherwise nothing really updates.

Am I doing something wrong with my config?
## [10][Hosting images](https://www.reddit.com/r/reactjs/comments/jib3dr/hosting_images/)
- url: https://www.reddit.com/r/reactjs/comments/jib3dr/hosting_images/
---
Hi im currently making a full stack website using the MERN stack. I have a profile component which lets you upload an image and crop it like in twitter but im having trouble on where and how to store it. 

I read something suggesting using base64 but people said that is was slow.  Is there any other alternatives?
## [11][Write your React App with Fast Refresh in Deno.](https://www.reddit.com/r/reactjs/comments/jid8zt/write_your_react_app_with_fast_refresh_in_deno/)
- url: https://www.reddit.com/r/reactjs/comments/jid8zt/write_your_react_app_with_fast_refresh_in_deno/
---
Hi, let me introduce you to [Aleph.js](https://alephjs.org/), a React Framework in [Deno](https://deno.land), inspired by [Next.js](https://nextjs.org).

https://reddit.com/link/jid8zt/video/3hs8jzqjifv51/player

[https://alephjs.org](https://alephjs.org/)
## [12][I'm looking for a fullscreen carousel with pinch/pan/zoom features](https://www.reddit.com/r/reactjs/comments/jid5wv/im_looking_for_a_fullscreen_carousel_with/)
- url: https://www.reddit.com/r/reactjs/comments/jid5wv/im_looking_for_a_fullscreen_carousel_with/
---
I used to implement the carousel mysel and use react-zoom-pan-pinch on the pictures displayed.

Are there any more convenients solutions?

Thx
