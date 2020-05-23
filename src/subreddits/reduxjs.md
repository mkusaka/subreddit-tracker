# reduxjs
## [1][I worked with Redux for two years and I wrote a blog post about tips and tricks.](https://www.reddit.com/r/reduxjs/comments/goqbi3/i_worked_with_redux_for_two_years_and_i_wrote_a/)
- url: https://www.cyprientaque.com/blog/redux-tips
---

## [2][need some help passing an array to a component after actions dispatch](https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/)
- url: https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/
---
I am trying to pass an array from a container to a component inside of the container. the issue im having is that i dispatch my  actions in componentDidMount of the container. then i try passing the array that is acquired to the comp but it passes the initial state, an empty array, instead of the third and final state
## [3][What to put on the Redux state](https://www.reddit.com/r/reduxjs/comments/gnir1f/what_to_put_on_the_redux_state/)
- url: https://x8lucas8x.com/what-to-put-on-the-redux-state/
---

## [4][Experimental state management lib for React from Facebook](https://www.reddit.com/r/reduxjs/comments/gkahrf/experimental_state_management_lib_for_react_from/)
- url: https://blog.graphqleditor.com/recoil-react-state-management-library/
---

## [5][Recomendations for Enterprise-scale Nextjs(React)/Redux/Typescript arquitecture](https://www.reddit.com/r/reduxjs/comments/gjrykm/recomendations_for_enterprisescale/)
- url: https://www.reddit.com/r/reduxjs/comments/gjrykm/recomendations_for_enterprisescale/
---
I want to know useful patterns/arquitecture for big enterprise projects used in nextjs(react) with redux. 
i've seen some recomendations like [this](https://laniewski.me/javascript/react/redux/2019/02/28/enterprise-scale-react-redux-project-architecture.html) or [this](https://www.pluralsight.com/guides/how-to-organize-your-react-+-redux-codebase)  where application is split in `modules` that encapsulates all related things with it, Redux(actions, reducers, sagas), utils, ts-types, and jsx|tsx components and have one folder for shared things. I like this concept because it is easy to identify the related elements between the store and the ui layer inside modules, but I don't know if it's really scalable for big enterprise projects. I would like to hear recommendations and/or suggested articles.
## [6][[HELP] Redux store/mapping not available on init](https://www.reddit.com/r/reduxjs/comments/gitfgq/help_redux_storemapping_not_available_on_init/)
- url: https://www.reddit.com/r/reduxjs/comments/gitfgq/help_redux_storemapping_not_available_on_init/
---
So I have this component in React.  The {props.size} and {props.algorithm} only show up after I have submitted them. Upon init, my component's props are `options`, and `dispatch`, as opposed to `size` and `algorithm` as specified in `mapStateToProps`. 

mapDispatchToProp is given by default. It appears to work normally as the size and algorithm of the value does render on the page. But if I want to access size and algorithms before hitting submit, I have to do `{props.options.size}` and `{props.options.algorithm}`. After submitting `props.options` become undefined

Here's my component:

    const Options_Component = (props) =&gt; {
        const {register, handleSubmit, errors, reset} = useForm({
            validateCriteriaMode: 'all',
            mode: 'onSubmit',
            reValidateMode: 'onSubmit',
            defaultValues: {
                size: INITIAL_SIZE,
                algorithm: INITIAL_ALGORITHM
            }
        });
        
        const onSubmit = values =&gt; {
            const inputs = changeOptions({size: 50})
            props.dispatch(inputs);
        }
    
        return (
            &lt;header&gt;
                &lt;form onSubmit={handleSubmit(onSubmit)}&gt;
                    &lt;div&gt;{props.size}&lt;/div&gt;
                    &lt;div&gt;{props.algorithm}&lt;/div&gt;
                    &lt;input type="submit"/&gt;
                    &lt;input type="button" value="Reset to default" onClick={()=&gt; reset()}/&gt;
                &lt;/form&gt;
            &lt;/header&gt;
        );
    }
    
    const mapStateToProps = (state, ownProps) =&gt; ({
        size: state.options.size,
        algorithm: state.options.algorithm
    })
    
    export default connect(mapStateToProps)(Options_Component)

My App - Option component is a child of SortVisualizer

    function App() {
      return (&lt;Provider store={OptionStore}&gt;
        &lt;SortVisualizer/&gt;
      &lt;/Provider&gt;)
    }
    
    export default App;
    

Reducer and createStore:

    export default function (state = initialState, {type, payload}) {
        console.log({type, payload})
        switch (type) {
            case CHANGE_OPTIONS:
                console.log(type, payload)
                const {size = initialValue.INITIAL_SIZE, algorithm = initialValue.INITIAL_ALGORITHM} = payload
                return {size: size, algorithm: algorithm}
            default: return {...state}
        }
    }
    
    export const changeOptions = options =&gt; ({
        type: CHANGE_OPTIONS,
        payload: options
    })
    
    export default combineReducers({options}); &lt;--- Root Reducer
    
    export default createStore(rootReducer);
## [7][At this point should my state (store) be viewable in the Redux dev tools? I am seeing Reddit's store there which is strange](https://www.reddit.com/r/reduxjs/comments/gi1cxi/at_this_point_should_my_state_store_be_viewable/)
- url: https://www.reddit.com/r/reduxjs/comments/gi1cxi/at_this_point_should_my_state_store_be_viewable/
---
Just refactoring a small app to use Redux and hooks and noticing something strange. In index.js I have:   


    ...
    
    import { createStore } from 'redux';
    import { Provider } from 'react-redux';
    import { rootReducer } from './reducers/rootReducer.js';
    const store = createStore(rootReducer);
    
    ReactDOM.render(
      &lt;Provider store={store}&gt;
        &lt;React.StrictMode&gt;
          &lt;App /&gt;
        &lt;/React.StrictMode&gt;
      &lt;/Provider&gt;,
      document.getElementById('root')
    );

My root reducer is defined in another file I am importing and I have confirmed that is firing in the console. Yet for some reason when I inspect the redux dev tools I see a bunch of data related to my reddit user acccount oddly enough and not the initial state I defined in rootReducer which is passed to createStore  


At this point should I not be able to see my store in redux dev tools?
## [8][Is it ok to make actions creators to depend of other reducers?](https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/)
- url: https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/
---
I have an application with two reducers, both with tons of action creators, and 99% of the actions creators have to use data from another reducer to do some logic. I feel comfortable doing it, but I don't know if it is a bad or good pattern.

(Sorry guys, my English is bad)
## [9][tsrux: Typesafe and painless action creators and reducers for redux.](https://www.reddit.com/r/reduxjs/comments/ggi41k/tsrux_typesafe_and_painless_action_creators_and/)
- url: https://www.reddit.com/r/reduxjs/comments/ggi41k/tsrux_typesafe_and_painless_action_creators_and/
---
I created a new library for use with redux: [https://lusito.github.io/tsrux/](https://lusito.github.io/tsrux/)

It only weighs 300 Byte, but gives you type-safety for your actions and reducers and helps to reduce boilerplate code.

Previously I used [deox](https://deox.js.org/) (which inspired this library), but at work we have issues with the bundle size, since deox includes RXJS and without treeshaking (which we can't use for several reasons), this adds a lot of weight to the application.

Full documentation, unit-tests with 100% code-coverage and type-tests already included.

Let me know what you think.
## [10][How do you organize your file structure with Redux, Redux-Toolkit, react-router, and Redux-Saga?](https://www.reddit.com/r/reduxjs/comments/gf9eo7/how_do_you_organize_your_file_structure_with/)
- url: https://www.reddit.com/r/reduxjs/comments/gf9eo7/how_do_you_organize_your_file_structure_with/
---
Hello,

I've been using the Rails style of folder structure for a while and I think I've come to hate it. Splitting actions, reducers, and containers has lead to a lot of bloat in file structure and I've found it confusing to explain redux to jr devs given how spread out the logic is from each other.

I've recently spiked Redux-Saga and Redux-toolkit (late to the party, been a while since I've had to create a store from scratch), and I like both.

I've typically broken down stuff like this:  
router (contains my route logic)

components (globally shared components)

pages (individual pages that are imported into the router, contain components local to the page)

services, helpers, hooks and the like up there.

If I do a more feature style, should I encompass a page and a Navbar under a /features folder? Or how have projects you've worked on done it?
