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

## [3][cssbuttons: HTML &amp; CSS Buttons Collection Built in React](https://www.reddit.com/r/reactjs/comments/jcxy2p/cssbuttons_html_css_buttons_collection_built_in/)
- url: https://v.redd.it/0i2s6afhjot51
---

## [4][I turned some browser event APIs into React hooks.](https://www.reddit.com/r/reactjs/comments/jd7dus/i_turned_some_browser_event_apis_into_react_hooks/)
- url: https://www.reddit.com/r/reactjs/comments/jd7dus/i_turned_some_browser_event_apis_into_react_hooks/
---
I had some spare time on a plane flight, so I took the opportunity to hack away at one of my TODO sticky notes.

* [use-ctrl-key](https://www.npmjs.com/package/use-ctrl-key): Listen to if the Ctrl key is pressed. `const ctrlKey: boolean = useCtrlKey();`
* [use-key-down](https://www.npmjs.com/package/use-key-down): Given a key, listen to if it is pressed. `const isEnterKeyDown: boolean = useKeyDown('Enter');`
* [use-mouse-move](https://www.npmjs.com/package/use-mouse-move): Listen to mouse movement. `const [x, y] = useMouseMove();` Supports absolute, client, movement (relative), offset, page, and screen coordinates.
* [use-offline](https://www.npmjs.com/package/use-offline): Listens for network connectivity change events. `const isOffline: boolean = useOffline();`
* [use-page-transition](https://www.npmjs.com/package/use-page-transition): Listen for page transition changes. `const [isVisible, isPersistent] = usePageTransition();`
* [use-shift-key](https://www.npmjs.com/package/use-shift-key): Listen to if the Shift key is pressed. `const shiftKey: boolean = useShiftKey();`
## [5][I built an open-source React (Gatsby) + Tailwind Personal Portfolio and Blog Template](https://www.reddit.com/r/reactjs/comments/jd77qy/i_built_an_opensource_react_gatsby_tailwind/)
- url: https://github.com/RyanFitzgerald/devfolio
---

## [6][Trouble with TypeScript and React Context](https://www.reddit.com/r/reactjs/comments/jdctw3/trouble_with_typescript_and_react_context/)
- url: https://www.reddit.com/r/reactjs/comments/jdctw3/trouble_with_typescript_and_react_context/
---
Hi guys,

I've been struggling with using React Context and TypeScript. I've managed to get the context to work and pass data correctly to the subscribed components, however I'm struggling to know how to type the createContext function. I know it expects a defaultValue, but I'm confused what it should be.

I'm getting the following TypeScript error in my context file under the createContext function:

    (alias) createContext&lt;unknown&gt;(defaultValue: unknown): React.Context&lt;unknown&gt;
    import createContext
    Expected 1 arguments, but got 0.ts(2554)
    index.d.ts(385, 9): An argument for 'defaultValue' was not provided.
    Peek Problem (Alt+F8)
    No quick fixes available

This is my Context:

    import React, {useState, createContext} from 'react';
    
    interface NewPartyProviderProps {
        children: React.ReactNode
    }
    
    type NewParty = {
        name: string
    }
    
    export const NewPartyContext = createContext();
    
    export const NewPartyProvider = ({children}: NewPartyProviderProps) =&gt; {
        const [newParty, setNewParty] = useState&lt;NewParty&gt;({name: ""});
    
        return (
            &lt;NewPartyContext.Provider value={[newParty, setNewParty]}&gt;
                {children}
            &lt;/NewPartyContext.Provider&gt;
        )
    }

And I'm using it in a component like this:

    const NewPartyName = ({ navigation }: NewPartyNameProps) =&gt; {
    
      const [newParty, setNewParty] = useContext(NewPartyContext);
    
      return (...)
    
    }

Any help would be much appreciated! :)
## [7][How to change progress percent color in antd?](https://www.reddit.com/r/reactjs/comments/jdcs3q/how_to_change_progress_percent_color_in_antd/)
- url: https://www.reddit.com/r/reactjs/comments/jdcs3q/how_to_change_progress_percent_color_in_antd/
---
[https://codesandbox.io/s/jgu3j?file=/index.js](https://codesandbox.io/s/jgu3j?file=/index.js)

How to change the progress percent(90%) color from black to red?
## [8][Using Stackbit](https://www.reddit.com/r/reactjs/comments/jdf3q6/using_stackbit/)
- url: https://www.reddit.com/r/reactjs/comments/jdf3q6/using_stackbit/
---
A client of mine suggested that we try Stackbit but it feels like it isn't too "customizable" or flexible. I'd like to know your thoughts about it and the challenges you've faced using it.
## [9][[Material-ui] Create multiple slider knobs in one slider in material-ui, reactjs](https://www.reddit.com/r/reactjs/comments/jdelc7/materialui_create_multiple_slider_knobs_in_one/)
- url: https://www.reddit.com/r/reactjs/comments/jdelc7/materialui_create_multiple_slider_knobs_in_one/
---
[https://i.ibb.co/LNRH21c/multiple-slider-design.png](https://i.ibb.co/LNRH21c/multiple-slider-design.png)   


I am trying to create a slider which looks roughly similar to the photo attached. 

&amp;#x200B;

I know you can make multiple knobs by providing `defaultValue={[35,50,77]}` or something similar to the `&lt;Slider/&gt;` component as props. But I want to style the tracks and the individual knobs aswell. Is it possible? If not, is there any other way to achieve this apart from building my own component from scratch?
## [10][Started to build a simple blog with the Wordpress API and Reactjs](https://www.reddit.com/r/reactjs/comments/jdel9q/started_to_build_a_simple_blog_with_the_wordpress/)
- url: https://www.reddit.com/r/reactjs/comments/jdel9q/started_to_build_a_simple_blog_with_the_wordpress/
---
Just want to make sure im doing good in terms of clean code and " "way" of doing things:
https://github.com/morhaham/react-with-wordpress
## [11][Different ways to create Refs in React](https://www.reddit.com/r/reactjs/comments/jdbpku/different_ways_to_create_refs_in_react/)
- url: https://sumitkharche.hashnode.dev/different-ways-to-create-refs-in-react-ckgeraoen0c50pas18hzqa4u8
---

## [12][I am learning react.js and made my first ever todolist using react.](https://www.reddit.com/r/reactjs/comments/jdbn7j/i_am_learning_reactjs_and_made_my_first_ever/)
- url: https://www.reddit.com/r/reactjs/comments/jdbn7j/i_am_learning_reactjs_and_made_my_first_ever/
---
Github Link : [https://github.com/FSojitra/Simple-react-ToDo-with-crud-action](https://github.com/FSojitra/Simple-react-ToDo-with-crud-action)   
codesand box link : [https://codesandbox.io/s/awesome-water-rgjtr](https://codesandbox.io/s/awesome-water-rgjtr)
