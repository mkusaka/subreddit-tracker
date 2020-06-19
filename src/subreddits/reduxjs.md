# reduxjs
## [1][How Redux Toolkit can reduce your setup of Redux in your next React app](https://www.reddit.com/r/reduxjs/comments/hbc6ca/how_redux_toolkit_can_reduce_your_setup_of_redux/)
- url: https://medium.com//how-redux-toolkit-can-reduce-your-setup-of-redux-in-your-react-app-d87baab59268?source=friends_link&amp;sk=626b48e7ab94dff289177c14be3b7383
---

## [2][Parent componentDidMount and child componentDidMount.](https://www.reddit.com/r/reduxjs/comments/hbbc6n/parent_componentdidmount_and_child/)
- url: https://www.reddit.com/r/reduxjs/comments/hbbc6n/parent_componentdidmount_and_child/
---
I am trying to teach myself Redux and I am having trouble with an assignment.  I'm trying to figure out how to let a child use componentDidMount without the whole component tree, it belongs to, re-rendering.
Here's are some pics to explain my problem:
When I click on the [todo item](https://imgur.com/5SJ3QFS) I expect it to open up and reveal the details of the todo item including the [steps](https://imgur.com/bhSL24w).  This does happen BUT I only see the steps flash for a second and I'm back to seeing only the [todo list item](https://imgur.com/5SJ3QFS).  The todo items(parents) are fetched just like the steps(children).

Please let me know if you need more information!!!

Here is my code...


Child component:


    class TodoViewDetail extends React.Component {
        constructor(props) {
            super(props);
            this.props = props;
        }
        componentDidMount() {
            this.props.fetchSteps();
        }
        render() {
            const { body, id } = this.props.todo;
            return (
                &lt;div className="todo-details"&gt;
                    &lt;p&gt;{body}&lt;/p&gt;
                    &lt;StepListContainer todoId={id} /&gt;
                    &lt;button onClick={this.props.deleteTodo}&gt;Delete&lt;/button&gt;
                &lt;/div&gt;
            );
        }
    }

fetchSteps action:


    export const receiveSteps = (steps) =&gt;  ({
        type: RECEIVE_STEPS,
        steps: steps
    });
    export function fetchSteps(todoId) {
        return (dispatch, state) =&gt; {
            return stepAPIUtil.fetchSteps(todoId).then(
                    successfulStepsResponse =&gt; dispatch(receiveSteps(successfulStepsResponse)
                )
            )
        }
    }

ajax request:


    export function fetchSteps(todoId) {
        return $.ajax({
            type: "GET",
            url: `/api/todos/${todoId}/steps`,
        });
    }
## [3][I need help refactoring a middleware to use dependency injection.](https://www.reddit.com/r/reduxjs/comments/hb8u24/i_need_help_refactoring_a_middleware_to_use/)
- url: https://www.reddit.com/r/reduxjs/comments/hb8u24/i_need_help_refactoring_a_middleware_to_use/
---
I'm working on an application that uses legacy code from its version 1, where Redux middlewares were used to handle some actions. This middlewear would do something like `import Logger from './logger';` and `Logger.emit('some log');`.

In the new version of this application, `Logger` (among other objects, functions, and values) are no longer static modules. They are dependency injected:

    &lt;Application logger={Logger} /&gt;

The middleware can no longer import these variables this way, because they don't exist as simple modules. They exist within the React lifecycle as props.

What is the best way to handle this?

My team has two ideas, but I wanted to make sure we aren't overlooking anything obvious or potential problems.

1) Pass these values as a part of the action. Instead of `{ type: 'do', value: true }` it would now need to be `{ type: 'do', value: true, logger: Logger }`.

2) Create the store inside a `useMemo` on `Application` mount.

What is the optimal solution for injecting dependencies into actions?
## [4][Modern React/React Router Auth Best Practices with Redux Saga Firebase + React Hooks?](https://www.reddit.com/r/reduxjs/comments/hauvve/modern_reactreact_router_auth_best_practices_with/)
- url: /r/reactjs/comments/hauut4/modern_reactreact_router_auth_best_practices_with/
---

## [5][Create simple POS with React, Node and MongoDB #6: Redux Integration](https://www.reddit.com/r/reduxjs/comments/ha0wdc/create_simple_pos_with_react_node_and_mongodb_6/)
- url: https://blog.soshace.com/create-simple-pos-with-react-node-and-mongodb-6-redux-integration/
---

## [6][Advice for a frontend dev — firing multiple api calls at once](https://www.reddit.com/r/reduxjs/comments/h9u0mp/advice_for_a_frontend_dev_firing_multiple_api/)
- url: https://www.reddit.com/r/reduxjs/comments/h9u0mp/advice_for_a_frontend_dev_firing_multiple_api/
---
I'm struggling to find the right words to convey what I'm looking for in my research, so I thought I would ask the reddit community. I am looking for best practices to create records in different tables at once.

An example would be, user registration. Say you need to create 6 records for that user when they sign up, which need to connect. By connect I mean in the sense that if a user and team were created on user registration, the `userID` would need to be included in the team's members array. So the records need to fire in order so the relationship is properly recorded. So user would need to be created first, then the team record so I can add the `userID` to the team's members. Also note that the user record would need to be updated later on (after the team's record is created) with the `teamID` under the user's teams.

So as you can see it feels a bit all over the place. Currently I have multiple API calls being fired on user submit. While I have this working using `redux`, `firebase` and `react` — I foresee a lot of potential errors happening and feel as if I am not doing this in the most efficient way. I want to do this correctly and happy to do the research, I'm just not exactly sure what I am looking for. I was hoping for some guides, information, search terms, etc — basically anything to help me understand this concept more throughly if that makes sense.
## [7][Dispatching for one type of global state affects the other. Why?](https://www.reddit.com/r/reduxjs/comments/h7t6sq/dispatching_for_one_type_of_global_state_affects/)
- url: https://www.reddit.com/r/reduxjs/comments/h7t6sq/dispatching_for_one_type_of_global_state_affects/
---
I'm using Redux with my React Hooks simple counter project. It worked without any bugs or problems when the only global state was a simple integer with +/- buttons to. Then I added a second global state for light/dark themes and found that the add/subtract buttons affect the light/dark variable! I think I'm misusing the useDispatch() hook or combining reducers incorrectly. I've tried moving things into different containers and fiddled a lot with the syntax. In the code below I have omitted `import` and `export` statements for brevity:  
  
App.js:  
  
    const App = () =&gt; {
        return (
          &lt;div className="App"&gt;
            &lt;ThemeBar /&gt;
            &lt;Counter /&gt;
          &lt;/div&gt;
        );
    }
  
ThemeBar.js:  
  
    const ThemeBar = () =&gt; {
        const theme = useSelector(state =&gt; state.theme.themeIsLight)
        const dispatch = useDispatch();
    
        return (
            &lt;div&gt;
                &lt;ThemeOutput value={theme} /&gt;
                &lt;ThemeButton label="Toggle Theme"
                             clicked={()=&gt;dispatch({type: 'TOGGLE_THEME'})}/&gt;
            &lt;/div&gt;
        );
    };
  
Counter.js:  
  
    const Counter = () =&gt; {
        const count = useSelector(state =&gt; state.number.counter);
        const dispatch = useDispatch();
    
            return (
                &lt;div&gt;
                    &lt;CounterOutput value={count} /&gt;
                    &lt;CounterControl label="Increment"
                                    clicked={()=&gt;dispatch({ type: 'INCREMENT'})} /&gt;
                    &lt;CounterControl label="Decrement"
                                    clicked={()=&gt;dispatch({ type: 'DECREMENT'})} /&gt;
                    &lt;CounterControl label="Add 5"
                                    clicked={()=&gt;dispatch({ type: 'ADD', value: 5})} /&gt;
                    &lt;CounterControl label="Subtract 5"
                                    clicked={()=&gt;dispatch({ type: 'SUBTRACT', value: 5})} /&gt;
                &lt;/div&gt;
            );
    };  
  
themeReducer.js:  
  
    const themeReducer = (state = initialState, action) =&gt; {
        console.log('Theme: ' + state.themeIsLight);
        if (action.type === 'TOGGLE_THEME')
            return { themeIsLight: !state.themeIsLight};
        return state;
    };
  
globalNumberReducer.js:  
  
    const globalNumberReducer = (state = initialState, action) =&gt; {
        switch (action.type) {
            case 'INCREMENT': return {counter: state.counter + 1};
            case 'DECREMENT': return {counter: state.counter - 1};
            case 'ADD': return {counter: state.counter + action.value};
            case 'SUBTRACT': return {counter: state.counter - action.value};
            default: return state;
        }
    };
  
index.js:  
  
    const rootReducer = combineReducers({
        number: globalNumberReducer,
        theme:themeReducer
    });
    
    const store = createStore(rootReducer);
    console.log(store.getState());
    
    
    ReactDOM.render(
        &lt;Provider store={store}&gt;&lt;App /&gt;&lt;/Provider&gt;, document.getElementById('root')
    );
    registerServiceWorker();
## [8][Asynchronous actions, Redux store and race conditions.](https://www.reddit.com/r/reduxjs/comments/h7j059/asynchronous_actions_redux_store_and_race/)
- url: https://www.reddit.com/r/reduxjs/comments/h7j059/asynchronous_actions_redux_store_and_race/
---
**The problem:**

I have a list of buttons that when each is clicked I dispatch some actions to fetch various bits of data from a server and populate the Redux store. Lets say the store looks like the following:

`{ customerDetails: { name: "John", age: 25 }, tasks: [] }`

and the actions are split into 2, the first fetches customer details and the other fetches tasks.

If the user clicks the first button to see Customer1 and then clicks to see Customer2 before all the async actions for Customer1 have come back then the Redux store can get into a state where it'll be populated with Customer1 data if the async actions from the first button come back \_after\_ the async actions from the second button click (Customer2).

&amp;#x200B;

I can think of a few solutions to this, but all seem slightly complex..

1. Store a `latestButtonClickedID` in the store and when all the async actions come back they can then check whether they were dispatched with the same ID as the one that is in the store now. If the IDs match then we update the store, if the IDs do not match then we do nothing with the API response (i.e. do not update the store). This means I would have to pass in `latestButtonClickedID` into a lot of my actions that fetch data, not really ideal.
2. Change the store to look like  `{ buttonID1: { customerDetails: { ... }, tasks [] }, buttonID2: { customerDetails: { ... }, tasks [] } }` I'd then need to pass the correct data into the presentational component by using a selector that selects the data for the correct `buttonID`, where the `buttonID` is set synchronously on button click. This would also involve passing `buttonID`'s around everywhere.
3. Some sort of promise cancellation? Can anyone point me in the right direction for this if possible?
4. Disable any other button from being clicked until all promises of fetching the data have resolved. I can't really do this as the buttons are from a 3rd party library.
5. redux-saga `takeLatest`... I'd rather not as the team have found this to be fairly complex for our needs.

&amp;#x200B;

Any other suggestions? Thanks.
## [9][Best practice for actions?](https://www.reddit.com/r/reduxjs/comments/h13l2y/best_practice_for_actions/)
- url: https://www.reddit.com/r/reduxjs/comments/h13l2y/best_practice_for_actions/
---
At my work we use something along the lines of every real "action" having a pending, success, and fail action. Out of curiosity I checked some online resources and I'll see more of a SET vs GET sort of thing for actions. Just wondering if there is a best practice for this sort of thing for my own projects?

Thanks
## [10][Why should I write unit test for actionCreators?](https://www.reddit.com/r/reduxjs/comments/h0ue75/why_should_i_write_unit_test_for_actioncreators/)
- url: https://www.reddit.com/r/reduxjs/comments/h0ue75/why_should_i_write_unit_test_for_actioncreators/
---
Reference from [Official docs](https://redux.js.org/recipes/writing-tests#action-creators):

In Redux, action creators are functions which return plain objects. When testing action creators, we want to test whether the correct action creator was called and also whether the right action was returned.

But my question is WHY?

ActionCreator is a function that returns an object. It's a pure function. All it does is returning an object with whatever data is passed while calling. Is there any risk of not testing this function?

    export function addTodo(text) {
      return {
        type: types.ADD_TODO,
        text
      }
    }

Its more like creating two objects and doing deep comparison.

&gt;expect(actions.addTodo(text)).toEqual(expectedAction)

The only scenario I can think of where unit tests can be useful is when someone accidentally changes this function to return an object with type: `types.EDIT_TODO`

But  is this the only case? Are there any other benefits  of writing tests cases for functions that does the obvious?  


EDIT: I do use `redux-saga` for managing async actions (fetching data through API calls etc) and I do write unit tests for sagas. I'm only concerned  about writing unit tests for action creators!
