# reduxjs
## [1][Asynchronous actions, Redux store and race conditions.](https://www.reddit.com/r/reduxjs/comments/h7j059/asynchronous_actions_redux_store_and_race/)
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
## [2][Best practice for actions?](https://www.reddit.com/r/reduxjs/comments/h13l2y/best_practice_for_actions/)
- url: https://www.reddit.com/r/reduxjs/comments/h13l2y/best_practice_for_actions/
---
At my work we use something along the lines of every real "action" having a pending, success, and fail action. Out of curiosity I checked some online resources and I'll see more of a SET vs GET sort of thing for actions. Just wondering if there is a best practice for this sort of thing for my own projects?

Thanks
## [3][Why should I write unit test for actionCreators?](https://www.reddit.com/r/reduxjs/comments/h0ue75/why_should_i_write_unit_test_for_actioncreators/)
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
## [4][Can you use reducers across sibling components?](https://www.reddit.com/r/reduxjs/comments/h0dbjw/can_you_use_reducers_across_sibling_components/)
- url: https://www.reddit.com/r/reduxjs/comments/h0dbjw/can_you_use_reducers_across_sibling_components/
---
This might seem like a dumb question but this current code architecture I'm working with, each "sibling component" think left/right panels, have their own reducers(obviously?). Then they're joined into a parent reducer eg. `allReducers`.

So for the sake of an example we have: left panel, right panel

If right-panel has some state it's maintaining, can left-panel use it(without using that primary combined parent reducer).

Anyway I know this is hard to imagine without code, also we're using saga which I don't know off hand what it's for. The saga files have function generators inside them. I don't think it's relevant.
## [5][Does mapStateToProps run first before the component pulls in the props passed down into it?](https://www.reddit.com/r/reduxjs/comments/gz7hpt/does_mapstatetoprops_run_first_before_the/)
- url: https://www.reddit.com/r/reduxjs/comments/gz7hpt/does_mapstatetoprops_run_first_before_the/
---
This is probably a weird question but as I trace through(console log execution) of events when loading a component that is using `mapStateToProps` the value I set in the reducer state is what I see on the immediate load of the component.

I have a case where there is a destructured "key" in the props that is the same name/key as an added prop from `mapStateToProps` so I'm wondering if this is expected or a clash/technically an error...

And testing if I take out the `key: val` entry in the reducer state, the destructured prop is still there but the value is undefined.

edit: here's a better idea of what I'm saying

`const componentName = (props) =&gt; {`

`const { isLoading } = props;`

`}`

`const mapStateToProps = state =&gt; ({`

`isLoading: state.reducerState.isLoading`

`})`

`// rest of dispatch/connect`

I don't know if this fully captures the issue, since we also have connected components going on, a main reducer/saga...
## [6][I am getting Reddit's store in my app and not the one I created](https://www.reddit.com/r/reduxjs/comments/gvduz0/i_am_getting_reddits_store_in_my_app_and_not_the/)
- url: https://www.reddit.com/r/reduxjs/comments/gvduz0/i_am_getting_reddits_store_in_my_app_and_not_the/
---
I made a small app that fetches restaurant data, displays it, and allows filtering and searching cities and all the usual stuff using useReducer and now I am trying to refactor it to use Redux. I am not using Redux ToolKit yet but I plan on learning it for my next project.

Anyway, I made my rootReducer, brought in Provider and createStore and hooked all that up but in Redux devtools I see Reddit's store for my user and not my app store. Here is my index.js code. And I did install redux and react-redux and I see it in package.json. I have not connected any components yet but should I not see my store in redux dev tools?

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
    
    ...

Thanks in advance for nay help

EDIT: formatting
## [7][[Q] Why so much work just for a global storage space?](https://www.reddit.com/r/reduxjs/comments/gv47ka/q_why_so_much_work_just_for_a_global_storage_space/)
- url: https://www.reddit.com/r/reduxjs/comments/gv47ka/q_why_so_much_work_just_for_a_global_storage_space/
---
 I was thinking that perhaps redux has too much indirection, maybe? I mean its a bit too much separation of concern with all the mapdispatch and mapstate and actions and reducers . If the point is to use a global store why not just import a singleton class and use its state with plain getters &amp; setters? or some other object with application level scope ?  Thanks in advance for reading and giving this some thought. 
## [8][Easy Peasy the React Redux wrapper](https://www.reddit.com/r/reduxjs/comments/gv3kc9/easy_peasy_the_react_redux_wrapper/)
- url: https://medium.com//easy-peasy-the-react-redux-wrapper-b31a5911c5e3?source=friends_link&amp;sk=b5d0c558e24e3e0c7f40cf58dff17d70
---

## [9][Dynamic dependency injection with Redux](https://www.reddit.com/r/reduxjs/comments/gutj9m/dynamic_dependency_injection_with_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/gutj9m/dynamic_dependency_injection_with_redux/
---
I've been hustling with this for a few days now and I can't find a satisfactory answer anywhere else.

I'm currently working as an Ethereum Dapp Engineer and the crypto ecosystem is known to have some very poorly implemented libraries which almost everyone depends upon.

I need multi-wallet support in my app and this is a lot of work for now. My team uses this library called [web3-react](https://github.com/NoahZinsmeister/web3-react). It is a React-centric lib that allows for easy integration with different wallet providers.

To integrate this with Redux, I need to tap into it at the higher levels of my app and dispatch some actions. This way I can store the current connector and library in the store:

    function Initializer({ children }) {
      const dispatch = useDispatch();
      const web3React = useWeb3React();
      const { account, library, connector } = web3React;
    
      React.useEffect(() =&gt; {
        dispatch(
          changeProvider({
            connector,
            library,
           })
        );
      }, [connector, library, dispatch]);
    
      React.useEffect(() =&gt; {
        dispatch(changeAccount(account));
      }, [account, dispatch]);
    
      return children;
    }

**Problem #1**: `connector` and `library` are non-serializable and this is a NO-NO [according to the official style-guide](https://redux.js.org/style-guide/style-guide#do-not-put-non-serializable-values-in-state-or-actions).

Okay, let's find some place else for them to be.

**Approach #1:**

Since I'm using [@reduxjs/toolkit](https://redux-toolkit.js.org/), by default it comes with `redux-thunk`.

I created the following thunk to deal with provider connection:

```
export const activateConnector = Object.assign(
  ({ activate, setError, connectorName }) =&gt; async (dispatch, getState) =&gt; {
    const currentConnectorName = selectConnectorName(getState());
    if (currentConnectorName === connectorName) {
      return;
    }

    try {
      dispatch(activateConnector.pending(connectorName));

      await activate(getConnector(connectorName), (err) =&gt; setError(err), true);

      dispatch(activateConnector.fulfilled(connectorName));
    } catch (err) {
      dispatch(activateConnector.rejected(err.message));
    }
  },
  {
    pending: createAction("web3/activateConnector/pending"),
    fulfilled: createAction("web3/activateConnector/fulfilled"),
    rejected: createAction("web3/activateConnector/rejected"),
  }
);
```

In this scenario, my other thunks depend on `library` (to be more precise, they depend on both on `library` and some smart contract instances that depend on `library`). Since I can't put it in the store, I thought about using `thunk.withExtraArgument` API.

**Problem #2:** `withExtraArgument` assumes the extra arg is resolved by the time the store is created. However, since the user can change the wallet provider at any time, I need a way to overwrite in runtime. That doesn't seem possible and [redux-thunk maintainers don't seem to eager to add such functionality](https://github.com/reduxjs/redux-thunk/issues/277).

I managed to workaround that by injecting a mutable object in the thunk and using a custom middleware to change the reference whenever the library changes:

    const createApi = (library) =&gt; {
      return {
        // This is a contrived example. Most methods are not just a passthrough.
        async getBalance(account) {
          return library.getBalance(account);
        },
      };
    };
    
    const services = {
      api: new Proxy(
        {},
        {
          get: (target, prop, receiver) =&gt; {
            return () =&gt; Promise.reject(new Error("Not initialized"));
          },
        }
      ),
    };
    
    const store = configureStore({
      reducer: rootReducer,
      middleware: [
        // this should probably be exported from the web3Slice.js file
        (store) =&gt; (next) =&gt; (action) =&gt; {
          if (changeLibrary.match(action)) {
            services.api = createApi(action.payload);
            // do not forward this action
            return;
          }
    
          return next(action);
        },
        thunk.withExtraArgument(services),
        ...getDefaultMiddleware({
          thunk: false,
        }),
      ],
    });

The `Initializer` component is changed a little bit now:

    //...
      React.useEffect(() =&gt; {
        dispatch(
          activateConnector({
            connectorName: "network",
            activate,
            setError,
          })
        );
      }, [activate, dispatch, setError]);
    
      React.useEffect(() =&gt; {
        dispatch(changeAccount(account));
      }, [account, dispatch]);
    
      React.useEffect(() =&gt; {
        dispatch(changeLibrary(library));
      }, [library, dispatch]);
    // ...

**Problem #3:** While this solves the problem, this looks like a JavaScript-ey version of the Service Locator pattern (which many consider to be an anti-pattern).

It also just basically mutable global state, which could cause inconsistent state.

Imagine that I have thunk A and thunk B, which must perform sequential operations always as A -&gt; B. 

    const thunkA = () =&gt; async (dispatch, getState, { api }) =&gt; {
      // ...
      await api.doLongProcess();
    
      // ... dispatch some actions
    
      dispatch(thunkB());
    }
    
    const thunkB = () =&gt; async (dispatch, getState, { api }) =&gt; {
      await api.doOtherThing();
    }

While `api.doLongProcess` is in course, a `changeLibrary` event arrives. That will cause the `api` dependency to be changed, so what happens next is that when `thunkB` is called with the newer `api` instance. This is a big problem.

What I believe should happpen is that upon `api` change, all in-course operations depending on it should be cancelled. That is not an easy thing to pull out with this setup.

Does anyone have a suggestion on how to approach this?
## [10][Nullable state best practice (typescript) [redux toolkit, createSlice - set/allow state to be null ]](https://www.reddit.com/r/reduxjs/comments/guns93/nullable_state_best_practice_typescript_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/guns93/nullable_state_best_practice_typescript_redux/
---
Hey!

I'm trying to create a slice in which its state can be null.

    import { createSlice, PayloadAction } from '@reduxjs/toolkit';
    import { RootState, AppThunk } from 'app/store';
    interface GlobalErrorState {
        name?: string;
        message?: string;
        code?: string;
    }
    
    const initialState: GlobalErrorState|null= null;
    
    const globalErrorSlice = createSlice({
        name: 'globalError',
        initialState: initialState as (GlobalErrorState|null),
        reducers: {
            errorOccurred: (state, action: PayloadAction&lt;CommonAPIError&gt;) =&gt; {
                return action.payload;
            },
            errorDismissed: (state) =&gt; {
                return null;
            },
        },
    }});
    

Then, taking advantage of useSelector in a component would look as below:

    const globalError = useSelector((state: RootState) =&gt; {
            return state.globalError;
    }, shallowEqual);
    if (globalError) {
            return &lt;CommonErrorPage message={globalError.message} /&gt;;
    }

Few questions regarding the abovementioned approach:

1. Is there any best practice for setting a state to null? Would it be better if the state had wrapped the error object? 

&amp;#8203;

    
interface GlobalErrorState {
        error?: {
            name?: string;
            message?: string;
            code?: string;
        };
    }
    
    const initialState: GlobalErrorState = {error:undefined};

2. For typescript experts - any other way to make the first example work without the 'as'?

    initialState: initialState as (GlobalErrorState|null)

Thank you!
