# reduxjs
## [1][Does mapStateToProps run first before the component pulls in the props passed down into it?](https://www.reddit.com/r/reduxjs/comments/gz7hpt/does_mapstatetoprops_run_first_before_the/)
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
## [2][I am getting Reddit's store in my app and not the one I created](https://www.reddit.com/r/reduxjs/comments/gvduz0/i_am_getting_reddits_store_in_my_app_and_not_the/)
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
## [3][[Q] Why so much work just for a global storage space?](https://www.reddit.com/r/reduxjs/comments/gv47ka/q_why_so_much_work_just_for_a_global_storage_space/)
- url: https://www.reddit.com/r/reduxjs/comments/gv47ka/q_why_so_much_work_just_for_a_global_storage_space/
---
 I was thinking that perhaps redux has too much indirection, maybe? I mean its a bit too much separation of concern with all the mapdispatch and mapstate and actions and reducers . If the point is to use a global store why not just import a singleton class and use its state with plain getters &amp; setters? or some other object with application level scope ?  Thanks in advance for reading and giving this some thought. 
## [4][Easy Peasy the React Redux wrapper](https://www.reddit.com/r/reduxjs/comments/gv3kc9/easy_peasy_the_react_redux_wrapper/)
- url: https://medium.com//easy-peasy-the-react-redux-wrapper-b31a5911c5e3?source=friends_link&amp;sk=b5d0c558e24e3e0c7f40cf58dff17d70
---

## [5][Dynamic dependency injection with Redux](https://www.reddit.com/r/reduxjs/comments/gutj9m/dynamic_dependency_injection_with_redux/)
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
## [6][Nullable state best practice (typescript) [redux toolkit, createSlice - set/allow state to be null ]](https://www.reddit.com/r/reduxjs/comments/guns93/nullable_state_best_practice_typescript_redux/)
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
## [7][Do you connect every component which fetches data to Redux?](https://www.reddit.com/r/reduxjs/comments/guhewr/do_you_connect_every_component_which_fetches_data/)
- url: https://www.reddit.com/r/reduxjs/comments/guhewr/do_you_connect_every_component_which_fetches_data/
---
 It just seems crazy to add so much boilerplate code for every data-fetching component. I've seen this pattern a lot. Any motives for it?It seems that simple ApiCall in componentDidMount / useEffect would work better...
## [8][What happens when Redux is compiled/minified/deployed?](https://www.reddit.com/r/reduxjs/comments/gr46kl/what_happens_when_redux_is/)
- url: https://www.reddit.com/r/reduxjs/comments/gr46kl/what_happens_when_redux_is/
---
This is a random conceptual question but  I was just thinking, at the end of the day when you launch something and you have static files, they're deployed on some website/app. What does Redux/reducer/actions become... just an object/methods?
## [9][[Puzzler] Action creator using function.name](https://www.reddit.com/r/reduxjs/comments/gqze1l/puzzler_action_creator_using_functionname/)
- url: https://www.reddit.com/r/reduxjs/comments/gqze1l/puzzler_action_creator_using_functionname/
---
Today I encountered interesting bug which was caused by action creators like this:

    export function setCartStatus(status) {
        return {
            type: setCartStatus.name,
            cartStatus: status
        }
    }

All creators had unique names, I checked this several times manually across all modules. All other components (store, reducers etc) were also ok — the problem is only in the code given above. Can you spot one?

Note: comments under the post mention solution.

Answer is under the spoiler: &gt;! webpack minimizes function names, and different functions from different modules may happen to have the same minimized name. !&lt;
## [10][Frontend Cache Question](https://www.reddit.com/r/reduxjs/comments/gqqzkc/frontend_cache_question/)
- url: https://www.reddit.com/r/reduxjs/comments/gqqzkc/frontend_cache_question/
---
Hey everyone! So for the last few weeks I’ve been designing and beginning to implement a GraphQL API using Prisma. I’m at a stage where I have enough of an API foundation to think about developing frontend. I’ll be using react-native and I have a few questions regarding fetching data from my API. I want to be able to provide users with an offline experience and therefore require cache. I also want to use subscriptions (real-time functionality in GraphQL). I’ve found Apollo Client and seen that it has a lot of good reviews, however, I’m not a huge fan of the built in React components that display the data. I haven’t used them so I can’t really be sure that I don’t like them, however I don’t think they’ll be great for query testing which isn’t a huge deal since I’ll be thoroughly testing my API and Prisma is tested. On the other hand I’ve used redux in the past and am wondering if it has the possibly of acting as a sort of cache when paired with simple https requests. Any thoughts are appreciated! 🙏
