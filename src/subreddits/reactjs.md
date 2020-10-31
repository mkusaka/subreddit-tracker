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

## [3][Content Security Policy - protect your website from XSS attacks](https://www.reddit.com/r/reactjs/comments/jleg5w/content_security_policy_protect_your_website_from/)
- url: https://itsopensource.com/content-security-policy/
---

## [4][Run your React Native app on the web with React Native for Web](https://www.reddit.com/r/reactjs/comments/jlgo5q/run_your_react_native_app_on_the_web_with_react/)
- url: https://mmazzarolo.com/blog/2020-10-24-adding-react-native-web/
---

## [5][A responsive, composable react charting library with a hand-drawn style.](https://www.reddit.com/r/reactjs/comments/jkvvrp/a_responsive_composable_react_charting_library/)
- url: https://beizhedenglong.github.io/rough-charts/?path=/docs/roughcharts--page
---

## [6][Theming in React with variables (with default dark mode) 🌗](https://www.reddit.com/r/reactjs/comments/jl1ajl/theming_in_react_with_variables_with_default_dark/)
- url: https://www.reddit.com/r/reactjs/comments/jl1ajl/theming_in_react_with_variables_with_default_dark/
---
If you are sick of adding hundreds of dependencies in order to do a simple task, then this article is for you.  👻

This is how to implement a simple themes switcher with default dark mode step by step:

Firstly we are going to create a themes.js file and add the color variables 😓 :

    // the actual object holding the colors for each theme
    // you can always define colors with same values for all themes in :root
    const colors = {
      night: {
        '--lightest': '#e4e6eb',
        '--light': '#b0b3b8',
        '--medium': '#3a3b3c',
        '--dark': '#242526',
        '--darker': '#18191a',
      },
      day: {
        '--lightest': '#ffffff',
        '--light': '#f7f7f7',
        '--medium': '#dfe3ee',
        '--dark': '#8b9dc3',
        '--darker': '#3b5998',
      },
      coffee: {
        '--lightest': '#ece0d1',
        '--light': '#dbc1ac',
        '--medium': '#967259',
        '--dark': '#634832',
        '--darker': '#38220f',
      },
    };

Save the themes' names in array and the first theme as fallback theme 🤔 :

    const themes = Object.keys(colors);
    const fallback = themes[0];

Create functions for retrieving and saving the theme name 🤨 :

    // returns the last saved theme or the first available one
    // if the page is loaded for first time
    const getTheme = () =&gt; {
      const theme = localStorage.getItem('theme');
      return colors[theme] ? theme : fallback;
    };
    
    // simply saves the theme in the localStorage
    const saveTheme = (theme = fallback) =&gt; localStorage.setItem('theme', theme);

Then a function for actually changing the color variables from the colors object 😩 :

    // actually changing the colors depending on the theme
    const setColorScheme = (theme = fallback) =&gt; {
      Object.entries(colors[theme]).map(([color, value]) =&gt; {
        document.documentElement.style.setProperty(color, value);
      });
    };

And finally a function for export to set the theme 🤓 :

    // executes and returns a function for setting theme
    const setTheme = (() =&gt; {
      // it will load the last saved theme from the local storage
      // or fallback to the first available one
      // and then set it
      // happens on page load once
      setColorScheme(getTheme());
    
      return (theme) =&gt; {
        theme = colors[theme] ? theme : fallback;
        if (getTheme() !== theme) {
          setColorScheme(theme);
          saveTheme(theme);
        }
    
        return theme;
      };
    })();
    
    export { themes, getTheme, setTheme };

And this is a simple example on how to use it in ReactJS app (you can easily use it in any other):

App.js

    import './styles/app.css';
    import { useState, useEffect } from 'react';
    import { themes, getTheme, setTheme } from './styles/theme.js';
    
    export default () =&gt; {
      const [currentTheme, nextTheme] = useState(getTheme());
    
      useEffect(() =&gt; {
        setTheme(currentTheme);
      }, [currentTheme]);
    
      return (
        &lt;div className="app-wrapper"&gt;
          &lt;select onChange={(event) =&gt; nextTheme(event.target.value)}&gt;
            {themes.map((theme, i) =&gt; (
              &lt;option selected={theme === currentTheme} key={i} value={theme}&gt;
                {theme}
              &lt;/option&gt;
            ))}
          &lt;/select&gt;
        &lt;/div&gt;
      );
    };

style.css

    .app-wrapper {
      width: 100vw;
      height: 100vh;
      background-color: var(--darker);
    }
    
    .app-wrapper select {
      position: absolute;
      top: 50%;
      left: 50%;
      -webkit-transform: translate(-50%, -50%);
              transform: translate(-50%, -50%);
      padding: 0.4rem 0.8rem;
      font-weight: bold;
      background: var(--dark);
      cursor: pointer;
      color: var(--light);
      border: 0.2rem solid var(--medium);
      outline: none;
    }
    
    .app-wrapper select:hover {
      color: var(--lightest);
    }

The actual result: [https://i.ibb.co/rKj4p83/switcher.gif](https://i.ibb.co/rKj4p83/switcher.gif)  😳
## [7][I made a package](https://www.reddit.com/r/reactjs/comments/jlcbt1/i_made_a_package/)
- url: https://www.reddit.com/r/reactjs/comments/jlcbt1/i_made_a_package/
---
If you're on this sub, you're probably aware of context - well I decided to try and reimplement this using some of React's basic API. 
I call it react-ancestor

[package](https://www.npmjs.com/package/react-ancestor)
[github](github.com/enahsor/ancestor)

I'd love to hear what you all of think of my attempt.
## [8][React not re-rendering new Context, and nested component is pulling stale context state leading to undefined issues. However, moving the code to another component solves this issue and forces a hard re-render.](https://www.reddit.com/r/reactjs/comments/jlcu07/react_not_rerendering_new_context_and_nested/)
- url: https://www.reddit.com/r/reactjs/comments/jlcu07/react_not_rerendering_new_context_and_nested/
---
I am building an Invoice view (This code is far from production as some design choices are still up in the air). This view has two tabs: Payments and Charges (aka Transactions). Clicking on either tab should force a clean un-mount and re-render of all nested components. The InvoiceDetail page conditionally renders a ListContext (with either a "txapi" or a "pmtapi" prop that is used to fetch the data and setState inside the constructor and componentDidMount), and either a PmtTable or a TxTable. Both are children of a ListContext that fetches data.

This does not work and leads to an internal validation error upon tab change when TxTable reads payment data, which is obviously of a different schema than transactions:

    if (view == "payments") {
    return (
    &lt;Paper variant="outlined"&gt;
          {invCard}
          &lt;Divider /&gt;
          &lt;ListContextProvider endpoint={"pmtapi/"} params={{ invoice: invID }}&gt;
                 &lt;PmtTable /&gt;
                 &lt;PageFooter /&gt;
          &lt;/ListContextProvider&gt;
    &lt;/Paper&gt; ); 
    }
    
    if (view == "transactions") {
    return (
    &lt;Paper variant="outlined"&gt;
          {invCard}
          &lt;Divider /&gt;
          &lt;ListContextProvider endpoint={"txapi/"} params={{ invoice: invID }}&gt;
                  &lt;TxTable /&gt;
                  &lt;PageFooter /&gt;
          &lt;/ListContextProvider&gt;
    &lt;/Paper&gt; ); 
    }

Now the following works fine when I move the ListContextProvider and its children into a separate component and import it :

    if (view == "payments") {
    return ( 
    &lt;Paper variant="outlined"&gt;
         {invCard} 
         &lt;Divider /&gt;
         &lt;ListContextProvider endpoint={"pmtapi/"} params={{ invoice: invID }}&gt;         
          &lt;PmtTable /&gt; 
              &lt;PageFooter /&gt; 
         &lt;/ListContextProvider&gt; 
    &lt;/Paper&gt; ); 
    } 
    
    if (view == "transactions") { 
    return ( 
    &lt;Paper variant="outlined"&gt; 
         {invCard} 
         &lt;Divider /&gt; 
         &lt;TxView /&gt; 
    &lt;/Paper&gt; ); 
    } 

I've tried everything else, from modifying the life cycle methods inside the Context to sticking the logic inside of immediately-invoked functions. I am assuming that React-Fiber's diffing algorithm recognizes that ListContextProvider's position has not changed, but the props have changed and this should return a re-render of its children. I added a comparison of prevProps to props inside of shouldComponentUpdate in ListContextProvider but to no avail.

And if the diffing is to blame here, why is wrapping the offending code inside nothing but essentially a named fragment or div not causing this issue? If I add an extra div wrapper to the original code to make one sub-tree different than the other, will this fix the issue? I really do not want to use a separate &lt;TxView&gt; component.

Thanks in advance for all of the help! Not an expert at React.

**EDIT**: Indeed the wrapper makes all the difference, and my suspicion of diffing being the issue is strengthened. I just stuck a div tag around one of the branches but not the other and it worked fine. I really am looking for a way to force a complete re-render.

**EDIT2**: Sticking the offending code inside a function component using the render props pattern forced a hard re-render and solved the issue. I am still interested in understanding what's going on. From my understanding:  


1) React Fiber realizes the DOM tree down to the `&lt;table&gt;`'s has not fundamentally changed, only the props and state are different. Even `&lt;TxTable&gt;` and `&lt;PmtTable&gt;` are actually implementations of the same Material-Table component.

1.5) Props and state updates are supposed to force a re-render, but the order may not be guaranteed? The Context updates are asynchronous and the `useContext()` hook may fire first before a props update. Because most of the DOM tree has not really changed, the nested component is updating before the Context is re-rendered.

2) The above looks true according to unposted print statements. Is there a reliable way to influence the order, or suspend for a certain amount of time, of re-renders?

3) The above is re-inforced by a print statement inside ListContext's constructor that never fires. The `context.items` indeed change, but the constructor never fires even though the props are obviously different. How is this possible? And how can I change it?
## [9][[HELP]getting Unexpected token &lt; in JSON at position 0 even if im returning json from my flask backend](https://www.reddit.com/r/reactjs/comments/jleerx/helpgetting_unexpected_token_in_json_at_position/)
- url: https://www.reddit.com/r/reactjs/comments/jleerx/helpgetting_unexpected_token_in_json_at_position/
---
hey, I'm learning how to use flask and react together so that i can host my python projects online. I'm following [this article](https://blog.miguelgrinberg.com/post/how-to-create-a-react--flask-project) to do so. But as soon as i reload the page i get this error in the dev console -

`Uncaught (in promise) SyntaxError: Unexpected token &lt; in JSON at position 0`

i did some research on it and turns out its because the backend is not returning json, but in my case im returning json(as flask automatically jsonifies dictionaries).

&amp;#x200B;

Here is my code - [python code](https://pastebin.com/7SVKFqEg) , [react code](https://pastebin.com/tKkq5LgN)

&amp;#x200B;

Any help is appreciated as im literally pulling my hair apart

&amp;#x200B;

Thanks
## [10][How to make simple messaging using socket.io and react useEffect()](https://www.reddit.com/r/reactjs/comments/jl81pt/how_to_make_simple_messaging_using_socketio_and/)
- url: https://www.reddit.com/r/reactjs/comments/jl81pt/how_to_make_simple_messaging_using_socketio_and/
---
    function App() {
        const [messages, setMessages] = useState([{ type: "system", text: "Please stay nice!" }]);
        useEffect(() =&gt; {
          socket.current.on("messageSent", (data) =&gt; {
            setMessages([...messages, { type: "you", text: data.message }]);
          });
    
          socket.current.on("receiveMessage", (data) =&gt; {
            setMessages([...messages, { type: "partner", text: data.message }]);
          });
        }, []);
    
        return (
          &lt;&gt;
            &lt;span className="container"&gt;
              &lt;Chat messages={messages} /&gt;
            &lt;/span&gt;
          &lt;/&gt;
        );
      }

So this useEffect() is run only once to set-up socket.

Right now when I try to access messages from inside of useEffect I got only the starting value which is \[{ type: "system", text: "Please stay nice!" }\] also setMessages doesn't refresh props of my parent component and child component has only access to this starting value \[{ type: "system", text: "Please stay nice!" }\]. How to do this correctly?

.
## [11][Using localStorage with Redux to save favorites.](https://www.reddit.com/r/reactjs/comments/jktk0q/using_localstorage_with_redux_to_save_favorites/)
- url: https://www.reddit.com/r/reactjs/comments/jktk0q/using_localstorage_with_redux_to_save_favorites/
---
 

Live demo here: [https://moviesearcher-rust.vercel.app/](https://moviesearcher-rust.vercel.app/)  
Repository here: [https://github.com/wcelery/moviesearcher](https://github.com/wcelery/moviesearcher)  
What I'm trying to do is save my favorite movies in localStorage by clicking heart icon. First click adds item, second click removes it. The problem is ADD\_TO\_FAVORITES works fine, but REMOVE\_FROM\_FAVORITE crashes my app with Cannot read property 'toString' of undefined  
 error here:

    23 | case REMOVE_FROM_FAVORITES:
      24 |   return {
      25 |     ...state,
    &gt; 26 |     favorites: Object.keys(state.favorites)
         | ^  27 |       .filter((item) =&gt; item !== action.payload.id.toString())
      28 |       .reduce(
      29 |         (object, item) =&gt; ({ ...object, [item]: state.favorites[item] }),

Code, that handles clicks is in components/Movie.js:

    export default function Movie({ poster, title, id }) {
      const dispatch = useDispatch();
      const favorites = useSelector((state) =&gt; state.favoriteStore.favorites);
    
      const parse = (poster, title, id) =&gt; {
        return {
          poster: poster,
          title: title,
          id: id,
        };
      };
      let movie = parse(poster, title, id);
    
      let isFavourite = movie.id in favorites;
    
      const isStored = (value) =&gt; {
        if (value) {
          dispatch(removeFromFavorites(movie));
        } else {
          dispatch(addToFavorites(movie));
        }
      };

Here's my redux/reducers/favoritesReducer.js:

    import { ADD_TO_FAVORITES, REMOVE_FROM_FAVORITES } from "../config";
    
    const checkLocalStorage = () =&gt; {
      return Object.assign(
        {},
        ...Object.entries(localStorage).map(([id, value]) =&gt; ({
          [id]: JSON.parse(value),
        }))
      );
    };
    
    const initialState = {
      favorites: checkLocalStorage(),
    };
    
    export const favoritesReducer = (state = initialState, action) =&gt; {
      switch (action.type) {
        case ADD_TO_FAVORITES:
          //blah blah
        case REMOVE_FROM_FAVORITES:
          return {
            ...state,
            favorites: Object.keys(state.favorites)
              .filter((item) =&gt; item !== action.payload.id.toString())
              .reduce(
                (object, item) =&gt; ({ ...object, [item]: state.favorites[item] }),
                {}
              ),
          };
        default:
          return state;
      }
    };
## [12][React hook practical guide](https://www.reddit.com/r/reactjs/comments/jlaxss/react_hook_practical_guide/)
- url: https://www.reddit.com/r/reactjs/comments/jlaxss/react_hook_practical_guide/
---
Im trying to learn react hooks and i have the idea behind it but i want a practical tutorial. Is there any good video for this? Lets say they create a full stack project with it and explains every detail behind their decision (eg folder structure). 

A udemy course link is also appreciated!
