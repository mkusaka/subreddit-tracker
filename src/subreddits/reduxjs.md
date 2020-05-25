# reduxjs
## [1][Red - Type-safe, composable, boilerplateless reducers ; https://github.com/betafcc/red](https://www.reddit.com/r/reduxjs/comments/gq1rj9/red_typesafe_composable_boilerplateless_reducers/)
- url: https://v.redd.it/t0wnti4hdt051
---

## [2][Is the constructor even needed when using Redux?](https://www.reddit.com/r/reduxjs/comments/gp5xfa/is_the_constructor_even_needed_when_using_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/gp5xfa/is_the_constructor_even_needed_when_using_redux/
---
I haven’t yet had any reason to use it since I place all the props I need in mapStateToProps which goes into the connect function. I was pretty good at passing props around in React but now that I have added Redux I’m not too sure how props are passed around. 

I’m making an application with Redux because I want to expand my web development “toolkit”. So even though my application doesn’t really need Redux I’m just curious about how to use it. I feel like I’ve been taught Redux badly or maybe I’m just not getting it but if you know of any good resources for me to learn that would be great too.
## [3][I worked with Redux for two years and I wrote a blog post about tips and tricks.](https://www.reddit.com/r/reduxjs/comments/goqbi3/i_worked_with_redux_for_two_years_and_i_wrote_a/)
- url: https://www.cyprientaque.com/blog/redux-tips
---

## [4][need some help passing an array to a component after actions dispatch](https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/)
- url: https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/
---
I am trying to pass an array from a container to a component inside of the container. the issue im having is that i dispatch my  actions in componentDidMount of the container. then i try passing the array that is acquired to the comp but it passes the initial state, an empty array, instead of the third and final state
## [5][What to put on the Redux state](https://www.reddit.com/r/reduxjs/comments/gnir1f/what_to_put_on_the_redux_state/)
- url: https://x8lucas8x.com/what-to-put-on-the-redux-state/
---

## [6][Experimental state management lib for React from Facebook](https://www.reddit.com/r/reduxjs/comments/gkahrf/experimental_state_management_lib_for_react_from/)
- url: https://blog.graphqleditor.com/recoil-react-state-management-library/
---

## [7][Recomendations for Enterprise-scale Nextjs(React)/Redux/Typescript arquitecture](https://www.reddit.com/r/reduxjs/comments/gjrykm/recomendations_for_enterprisescale/)
- url: https://www.reddit.com/r/reduxjs/comments/gjrykm/recomendations_for_enterprisescale/
---
I want to know useful patterns/arquitecture for big enterprise projects used in nextjs(react) with redux. 
i've seen some recomendations like [this](https://laniewski.me/javascript/react/redux/2019/02/28/enterprise-scale-react-redux-project-architecture.html) or [this](https://www.pluralsight.com/guides/how-to-organize-your-react-+-redux-codebase)  where application is split in `modules` that encapsulates all related things with it, Redux(actions, reducers, sagas), utils, ts-types, and jsx|tsx components and have one folder for shared things. I like this concept because it is easy to identify the related elements between the store and the ui layer inside modules, but I don't know if it's really scalable for big enterprise projects. I would like to hear recommendations and/or suggested articles.
## [8][[HELP] Redux store/mapping not available on init](https://www.reddit.com/r/reduxjs/comments/gitfgq/help_redux_storemapping_not_available_on_init/)
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
## [9][At this point should my state (store) be viewable in the Redux dev tools? I am seeing Reddit's store there which is strange](https://www.reddit.com/r/reduxjs/comments/gi1cxi/at_this_point_should_my_state_store_be_viewable/)
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
## [10][Is it ok to make actions creators to depend of other reducers?](https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/)
- url: https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/
---
I have an application with two reducers, both with tons of action creators, and 99% of the actions creators have to use data from another reducer to do some logic. I feel comfortable doing it, but I don't know if it is a bad or good pattern.

(Sorry guys, my English is bad)
