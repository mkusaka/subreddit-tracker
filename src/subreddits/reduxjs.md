# reduxjs
## [1][[HELP] Redux store/mapping not available on init](https://www.reddit.com/r/reduxjs/comments/gitfgq/help_redux_storemapping_not_available_on_init/)
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
## [2][At this point should my state (store) be viewable in the Redux dev tools? I am seeing Reddit's store there which is strange](https://www.reddit.com/r/reduxjs/comments/gi1cxi/at_this_point_should_my_state_store_be_viewable/)
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
## [3][Is it ok to make actions creators to depend of other reducers?](https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/)
- url: https://www.reddit.com/r/reduxjs/comments/ghu3wx/is_it_ok_to_make_actions_creators_to_depend_of/
---
I have an application with two reducers, both with tons of action creators, and 99% of the actions creators have to use data from another reducer to do some logic. I feel comfortable doing it, but I don't know if it is a bad or good pattern.

(Sorry guys, my English is bad)
## [4][tsrux: Typesafe and painless action creators and reducers for redux.](https://www.reddit.com/r/reduxjs/comments/ggi41k/tsrux_typesafe_and_painless_action_creators_and/)
- url: https://www.reddit.com/r/reduxjs/comments/ggi41k/tsrux_typesafe_and_painless_action_creators_and/
---
I created a new library for use with redux: [https://lusito.github.io/tsrux/](https://lusito.github.io/tsrux/)

It only weighs 300 Byte, but gives you type-safety for your actions and reducers and helps to reduce boilerplate code.

Previously I used [deox](https://deox.js.org/) (which inspired this library), but at work we have issues with the bundle size, since deox includes RXJS and without treeshaking (which we can't use for several reasons), this adds a lot of weight to the application.

Full documentation, unit-tests with 100% code-coverage and type-tests already included.

Let me know what you think.
## [5][How do you organize your file structure with Redux, Redux-Toolkit, react-router, and Redux-Saga?](https://www.reddit.com/r/reduxjs/comments/gf9eo7/how_do_you_organize_your_file_structure_with/)
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
## [6][Question about initializing state in the store](https://www.reddit.com/r/reduxjs/comments/gem4gj/question_about_initializing_state_in_the_store/)
- url: https://www.reddit.com/r/reduxjs/comments/gem4gj/question_about_initializing_state_in_the_store/
---
Hi,

I generated an application using the redux template and store my component state in slices. I have some sample data I want to feed into my component, but it's not something I want to put in my slice file as preloaded state. How do I populate it during initialization of the application? In the App component? Do I dispatch an action from App or outside of it?

Any suggestion will be appreciated!

Thanks,  
Guoliang Cao
## [7][Im trying to figure out what '...' means in return{... x, y, x}](https://www.reddit.com/r/reduxjs/comments/geajse/im_trying_to_figure_out_what_means_in_return_x_y_x/)
- url: https://www.reddit.com/r/reduxjs/comments/geajse/im_trying_to_figure_out_what_means_in_return_x_y_x/
---
Ive been struggling to find a good answer/explanation to this.  


what does the '...' mean in the redux initial state reducer method.

EXAMPLE:

&amp;#x200B;

 `export default function (state = {} , action) {`

`switch(action.type) {` 

`case GET_ORDER_STATUS:` 

`return { ...state, orderStatusData: action.payload.orderStatus };`   


`default: return state; }` 

`}`
## [8][Dispatching to non-Axios Actions?](https://www.reddit.com/r/reduxjs/comments/gdzzhn/dispatching_to_nonaxios_actions/)
- url: https://www.reddit.com/r/reduxjs/comments/gdzzhn/dispatching_to_nonaxios_actions/
---
Up until now I've only used Axios (fetch) requests from an API in my actions. What if I just want to change the Redux state from a React component like \`logout\`? If my function to \`logout()\` looks like this:

&amp;#x200B;

    logout = () =&gt; {localStorage.clear();this.props.clearCurrentUser()}

&amp;#x200B;

&amp;#x200B;

and my Reducer looks like this:

&amp;#x200B;

    const initialState = {currentUser: [],};
    
    export const currentUserReducer = (state = initialState, action) =&gt; {
    switch (action.type) {
    case "SET_CURRENT_USER":
    return { ...state, currentUser: action.payload }
    case "GET_CURRENT_USER":
    return { ...state, currentUser: action.payload }
    case "CLEAR_CURRENT_USER":
    return { ...state, currentUser: null }
    default:
    return state;}};

&amp;#x200B;

How do I write this in an \`action\`?
## [9][Do anyone have open source project that use Typescript and Redux-toolkit?](https://www.reddit.com/r/reduxjs/comments/gcktxf/do_anyone_have_open_source_project_that_use/)
- url: https://www.reddit.com/r/reduxjs/comments/gcktxf/do_anyone_have_open_source_project_that_use/
---
Please share the github link for learning
## [10][When should I use redux? Pros cons?](https://www.reddit.com/r/reduxjs/comments/gbs1o2/when_should_i_use_redux_pros_cons/)
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
