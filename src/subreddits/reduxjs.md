# reduxjs
## [1][Simplify your usage of Redux with redux-path](https://www.reddit.com/r/reduxjs/comments/g8xqru/simplify_your_usage_of_redux_with_reduxpath/)
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
## [2][Condition rendering failing in React Native Redux App](https://www.reddit.com/r/reduxjs/comments/g70svq/condition_rendering_failing_in_react_native_redux/)
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
## [3][Redux in Worker: Off-main-thread Redux Reducers and Middleware](https://www.reddit.com/r/reduxjs/comments/g6idzp/redux_in_worker_offmainthread_redux_reducers_and/)
- url: https://medium.com/@dai_shi/redux-in-worker-off-main-thread-redux-reducers-and-middleware-508e0cad8ac6?source=friends_link&amp;sk=e54dee252862e02d6a8a22c527547542
---

## [4][Fully Offline Progressive Web App Journal](https://www.reddit.com/r/reduxjs/comments/g5pdlv/fully_offline_progressive_web_app_journal/)
- url: https://www.reddit.com/r/reduxjs/comments/g5pdlv/fully_offline_progressive_web_app_journal/
---
Built with React / Redux on the Frontend and Django Rest on the backend. I have been keeping a journal since I was a kid. Sheltering in place has afforded me some extra time to develop an app to support my hobby of journaling. You don't have to sign up to start using it. I use the localStorage API to cache things locally. If you want to use the Django backend to sync journal entries between devices then that will require you to sign up. I hope you guys / gals enjoy! Stay safe out there!

Frontend Source Code: [https://github.com/strap8/llexicon](https://github.com/strap8/llexicon)

Backend Source Code: [https://github.com/strap8/llexicon-db](https://github.com/strap8/llexicon-db)

Production link: [https://www.astraltree.com](https://www.astraltree.com/)
## [5][Redux-toolkit is the quickest and easiest way to write reducers and keep store state I’ve found so far](https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/
---
It is now the standard for writing redux logic.

I've been using redux-toolkit for 3 months and I am enjoying it so much! it simply boosts the productivity 🔥

It makes your code shorter and easier to understand and enforces you to follow best practice (normalizing, selectors, typing, etc… )

👇🏽You'll find concrete examples and code in the article below 👇🏽[https://blog.theodo.com/2020/01/reduce-redux-boilerplate/](https://blog.theodo.com/2020/01/reduce-redux-boilerplate/)
## [6][Decoupled State Interface](https://www.reddit.com/r/reduxjs/comments/g4nxu4/decoupled_state_interface/)
- url: https://github.com/gactjs/store/blob/master/docs/decoupled-state-interface.md
---

## [7][Normalized Reducer: an easier way to read and write normalized relational reducer state](https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/)
- url: https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/
---
[**Normalized Reducer**](https://github.com/brietsparks/normalized-reducer) helps you manage normalized relational state without requiring any reducer/action boilerplate.

You simply feed in a schema of the relationships, and it gives you the reducers, actions, and selectors to read and write state according to that schema. The [**docs quick start**](https://github.com/brietsparks/normalized-reducer#quick-start) shows how easy it is to set up and use. Check out the [**demo app**](https://brietsparks.github.io/normalized-reducer-demo/) to see its features in action.

Additionally, it integrates with Normalizr. First denormalize your data with Normalizr's [`normalize`](https://github.com/paularmstrong/normalizr/blob/master/docs/api.md#normalizedata-schema), and then pass that output into Normalized Reducer's [`fromNormalizr`](https://github.com/brietsparks/normalized-reducer#normalizr-integration). The final result will be normalized state that the reducer can handle. See the demo's Normalizr example to see shapes of the transformation.

To help decide whether this library is right for you, there is a [comparison to Redux ORM and Redux Toolkit's entity adapter](https://github.com/brietsparks/normalized-reducer#comparison-to-alternatives). Key points: Relational Reducer is lighter, zero dep, framework agnostic, and covers some important use-cases the others are missing. However, it (currently) has fewer selector features than Redux ORM and is less mature than both.

Thanks for reading, and I hope this helps build stuff faster!
## [8][Gact Store White Paper](https://www.reddit.com/r/reduxjs/comments/g3q2lj/gact_store_white_paper/)
- url: https://github.com/gactjs/store/blob/master/docs/white-paper.md
---

## [9][Keeping the userActions DRY in redux](https://www.reddit.com/r/reduxjs/comments/g2wey8/keeping_the_useractions_dry_in_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/g2wey8/keeping_the_useractions_dry_in_redux/
---
So, I've defined userActions and I see the same pattern over and over again. So, just wanted to know if that's the correct way of writing action creators and thunks. Also, I've defined all the actions in a single file and wondering if I should separate them or not. The code works fine but clean code is always better. 

**userActions.js**

```js
import axios from "axios"

export const registerUser = (registrationData) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "REGISTRATION_STARTS" })
    try {
      const res = await axios.post(
        "http://localhost:3000/api/v1/users/register",
        registrationData
      )
      dispatch({
        type: "REGISTRATION_SUCCESS",
        data: { user: res.data.user },
      })
    } catch (err) {
      dispatch({
        type: "REGISTRATION_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const loginUser = (loginData, redirect) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "AUTH_STARTS" })
    try {
      const res = await axios.post(
        "http://localhost:3000/api/v1/users/login",
        loginData
      )
      dispatch({
        type: "AUTH_SUCCESS",
        data: { user: res.data.user },
      })
      localStorage.setItem("authToken", res.data.token)
      redirect()
    } catch (err) {
      dispatch({
        type: "AUTH_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const getCurrentUser = (token) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "AUTH_STARTS" })
    try {
      const res = await axios.get("http://localhost:3000/api/v1/users/me", {
        headers: {
          Authorization: token,
        },
      })
      dispatch({
        type: "AUTH_SUCCESS",
        data: { user: res.data.user },
      })
    } catch (err) {
      dispatch({
        type: "AUTH_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const logoutUser = () =&gt; {
  return (dispatch) =&gt; {
    dispatch({ type: "LOGOUT_USER" })
  }
}

export const addPost = (postData, redirect) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({
      type: "ADD_POST_STARTS",
    })
    try {
      const res = await axios.post(
        "http://localhost:3000/api/v1/posts/new",
        postData,
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `${localStorage.authToken}`,
          },
        }
      )
      dispatch({
        type: "ADD_POST_SUCCESS",
        data: { post: res.data.post },
      })
      redirect()
    } catch (err) {
      dispatch({
        type: "ADD_POST_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const getPost = (id) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "FETCHING_POST_START" })
    try {
      const res = await axios.get(`http://localhost:3000/api/v1/posts/${id}`)
      dispatch({
        type: "FETCHING_POST_SUCCESS",
        data: res.data.post,
      })
    } catch (error) {
      dispatch({
        type: "FETCHING_POST_FAILURE",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const updatePost = (id, postData, redirect) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "UPDATING_POST_START" })
    try {
      const res = await axios.put(
        `http://localhost:3000/api/v1/posts/${id}/edit`,
        postData
      )
      dispatch({
        type: "UPDATING_POST_SUCCESS",
        data: res.data,
      })
      redirect()
    } catch (error) {
      console.log(error)
      dispatch({
        type: "UPDATING_POST_FAILURE",
        data: { error: res.data.error },
      })
    }
  }
}

export const deletePost = (id) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({ type: "DELETING_POST_START" })
    try {
      const res = await axios.delete(
        `http://localhost:3000/api/v1/posts/${id}/delete`
      )
      dispatch({
        type: "DELETING_POST_SUCCESS",
        data: res.data.post,
      })
    } catch (error) {
      dispatch({
        type: "DELETING_POST_ERROR",
        data: { error: error },
      })
    }
  }
}

export const getListOfPosts = () =&gt; {
  return async (dispatch) =&gt; {
    dispatch({
      type: "FETCHING_POSTS_START",
    })
    try {
      const res = await axios.get("http://localhost:3000/api/v1/posts/list", {
        headers: {
          "Content-Type": "application/json",
        },
      })
      dispatch({
        type: "FETCHING_POSTS_SUCCESS",
        data: { posts: res.data.posts },
      })
    } catch (err) {
      dispatch({
        type: "FETCHING_POSTS_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}

export const getUserPosts = (id) =&gt; {
  return async (dispatch) =&gt; {
    dispatch({
      type: "FETCHING_USER_POSTS_START",
    })
    try {
      const res = await axios.get(
        `http://localhost:3000/api/v1/users/posts/${id}`,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      dispatch({
        type: "FETCHING_USER_POSTS_SUCCESS",
        data: res.data.userPosts,
      })
    } catch (err) {
      dispatch({
        type: "FETCHING_USER_POSTS_ERROR",
        data: { error: "Something went wrong" },
      })
    }
  }
}
```
## [10][How do we try to get data directly under streams? why does it has undefined and data?](https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/)
- url: https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/
---
The problem is when I bring the data from backend, There will be undefined under streams reducer, but if I same data and act as fake backend, I wont get undefined.

[When fetching from real backend](https://preview.redd.it/m4jjwmzw4dt41.png?width=1768&amp;format=png&amp;auto=webp&amp;s=bf950010545e4fa108e579c8125f4f7316766e5e)

&amp;#x200B;

[When bringing data from fake backend](https://preview.redd.it/h8ayi9705dt41.png?width=1192&amp;format=png&amp;auto=webp&amp;s=63e32e08a66f8e725e1b7b555b2b29a51eb31356)
