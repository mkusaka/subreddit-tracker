# reduxjs
## [1][When should I use redux? Pros cons?](https://www.reddit.com/r/reduxjs/comments/gbs1o2/when_should_i_use_redux_pros_cons/)
- url: https://www.reddit.com/r/reduxjs/comments/gbs1o2/when_should_i_use_redux_pros_cons/
---
Hello,

I am very new to coding in general. While i was coding in react, I noticed I can use redux to have a global state. However, I am not too sure when I should use redux instead of just passing down props.

I have a component which has 2 layered 2 components. Something like this:

&gt; A  
&gt;  
&gt;/\\  
&gt;  
&gt;BC  
&gt;  
&gt;||  
&gt;  
&gt;DE

Should I be using redux if I want to communicate between D and E?  Also, what are the pros and cons of using redux? (such as performance) How would I know when to use redux or just pass down props?

&amp;#x200B;

Thanks for all the comments in advance
## [2][Why use separate constants for action types?](https://www.reddit.com/r/reduxjs/comments/g9pb5n/why_use_separate_constants_for_action_types/)
- url: https://www.reddit.com/r/reduxjs/comments/g9pb5n/why_use_separate_constants_for_action_types/
---
Relatively new to Redux. I'm familiar with the points as [why to use action creators](https://blog.isquaredsoftware.com/2016/10/idiomatic-redux-why-use-action-creators/), but I often see action types set up as a separate set of constants ([here's a basic example from a tutorial](https://github.com/rhysdavies1994/koa-react-todo-tutorial/blob/master/frontend/src/actions/todos.js)) rather than just declared directly in the action creator. The logic of why this is needed is less clear to me - I've heard it is to minimize typos, but as far as I can tell, it's just a bunch more typing and importing and exporting things, without fundamentally affecting what an action creator is doing.

Can someone explain, ideally with an example, why it is considered a good practice to do this? Or, in what situations is it important / not-so-important?
## [3][Simplify your usage of Redux with redux-path](https://www.reddit.com/r/reduxjs/comments/g8xqru/simplify_your_usage_of_redux_with_reduxpath/)
- url: https://www.reddit.com/r/reduxjs/comments/g8xqru/simplify_your_usage_of_redux_with_reduxpath/
---
No action, connection or other boilerplates with [redux-path](https://github.com/beauvaisbruno/redux-path).
~~~~
dispatchTo('foo/sub', value);
getFrom('foo/sub');
~~~~

## Basic example
~~~~
const Component = () =&gt; {
  const sub = useSelector(state =&gt; state.foo.sub);
  return &lt;div&gt;
    &lt;div onClick={() =&gt; dispatchTo('foo/sub', sub + 1)}&gt;
      Click to increment
    &lt;/div&gt;
    &lt;div onClick={() =&gt; dispatchTo('foo', {sub: getFrom('foo/sub') + 1})}&gt;
      Click to increment
    &lt;/div&gt;
    foo.sub: &lt;&gt;{sub}&lt;/&gt;
  &lt;/div&gt;
}
~~~~
## [4][Condition rendering failing in React Native Redux App](https://www.reddit.com/r/reduxjs/comments/g70svq/condition_rendering_failing_in_react_native_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/g70svq/condition_rendering_failing_in_react_native_redux/
---
 I'm trying to conditionally render my **redux app** based on if the user is logged in. The relevant &amp; condensed version of my code is below:

    let isLoggedIn = false;
    
    export default function App() {
      console.log('App Executing...');
      console.log('isLoggedIn: ', isLoggedIn);
      return (
        &lt;Provider store={store}&gt;
          &lt;NavigationContainer&gt;
            {isLoggedIn ? ContactsTab() : Login()}
          &lt;/NavigationContainer&gt;
        &lt;/Provider&gt;
      );
    }
    
    store.subscribe(() =&gt; {
      // Set isLoggedIn to true if token is received and reinvoke App()
      if (store.getState().user.token) {
        isLoggedIn = true;
        App();
      }
    });

The app starts with console logging **isLoggedIn: false** and displaying **Login()**(as expected). When I login on my phone using the correct credentials, **App()** is re-invoked console logging **isLoggedIn: true**(as expected) but it's still displaying **Login()**. If I set **isLoggedIn = true** inside the app function, the app successfully starts displaying the **ContactsTab()**.

What is happening here? Why is my app not moving to **ContactsTab()** when the value of **isLoggedIn** successfully changes to **true**? How can I fix this?

Thank you for reading along. I have been trying to debug this for the past 2 days with no success so any help would be greatly appreciated!
## [5][Redux in Worker: Off-main-thread Redux Reducers and Middleware](https://www.reddit.com/r/reduxjs/comments/g6idzp/redux_in_worker_offmainthread_redux_reducers_and/)
- url: https://medium.com/@dai_shi/redux-in-worker-off-main-thread-redux-reducers-and-middleware-508e0cad8ac6?source=friends_link&amp;sk=e54dee252862e02d6a8a22c527547542
---

## [6][Fully Offline Progressive Web App Journal](https://www.reddit.com/r/reduxjs/comments/g5pdlv/fully_offline_progressive_web_app_journal/)
- url: https://www.reddit.com/r/reduxjs/comments/g5pdlv/fully_offline_progressive_web_app_journal/
---
Built with React / Redux on the Frontend and Django Rest on the backend. I have been keeping a journal since I was a kid. Sheltering in place has afforded me some extra time to develop an app to support my hobby of journaling. You don't have to sign up to start using it. I use the localStorage API to cache things locally. If you want to use the Django backend to sync journal entries between devices then that will require you to sign up. I hope you guys / gals enjoy! Stay safe out there!

Frontend Source Code: [https://github.com/strap8/llexicon](https://github.com/strap8/llexicon)

Backend Source Code: [https://github.com/strap8/llexicon-db](https://github.com/strap8/llexicon-db)

Production link: [https://www.astraltree.com](https://www.astraltree.com/)
## [7][Redux-toolkit is the quickest and easiest way to write reducers and keep store state I’ve found so far](https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/
---
It is now the standard for writing redux logic.

I've been using redux-toolkit for 3 months and I am enjoying it so much! it simply boosts the productivity 🔥

It makes your code shorter and easier to understand and enforces you to follow best practice (normalizing, selectors, typing, etc… )

👇🏽You'll find concrete examples and code in the article below 👇🏽[https://blog.theodo.com/2020/01/reduce-redux-boilerplate/](https://blog.theodo.com/2020/01/reduce-redux-boilerplate/)
## [8][Decoupled State Interface](https://www.reddit.com/r/reduxjs/comments/g4nxu4/decoupled_state_interface/)
- url: https://github.com/gactjs/store/blob/master/docs/decoupled-state-interface.md
---

## [9][Normalized Reducer: an easier way to read and write normalized relational reducer state](https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/)
- url: https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/
---
[**Normalized Reducer**](https://github.com/brietsparks/normalized-reducer) helps you manage normalized relational state without requiring any reducer/action boilerplate.

You simply feed in a schema of the relationships, and it gives you the reducers, actions, and selectors to read and write state according to that schema. The [**docs quick start**](https://github.com/brietsparks/normalized-reducer#quick-start) shows how easy it is to set up and use. Check out the [**demo app**](https://brietsparks.github.io/normalized-reducer-demo/) to see its features in action.

Additionally, it integrates with Normalizr. First denormalize your data with Normalizr's [`normalize`](https://github.com/paularmstrong/normalizr/blob/master/docs/api.md#normalizedata-schema), and then pass that output into Normalized Reducer's [`fromNormalizr`](https://github.com/brietsparks/normalized-reducer#normalizr-integration). The final result will be normalized state that the reducer can handle. See the demo's Normalizr example to see shapes of the transformation.

To help decide whether this library is right for you, there is a [comparison to Redux ORM and Redux Toolkit's entity adapter](https://github.com/brietsparks/normalized-reducer#comparison-to-alternatives). Key points: Relational Reducer is lighter, zero dep, framework agnostic, and covers some important use-cases the others are missing. However, it (currently) has fewer selector features than Redux ORM and is less mature than both.

Thanks for reading, and I hope this helps build stuff faster!
## [10][Gact Store White Paper](https://www.reddit.com/r/reduxjs/comments/g3q2lj/gact_store_white_paper/)
- url: https://github.com/gactjs/store/blob/master/docs/white-paper.md
---

