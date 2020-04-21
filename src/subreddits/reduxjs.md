# reduxjs
## [1][Redux-toolkit is the quickest and easiest way to write reducers and keep store state I’ve found so far](https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/g4tdqe/reduxtoolkit_is_the_quickest_and_easiest_way_to/
---
It is now the standard for writing redux logic.

I've been using redux-toolkit for 3 months and I am enjoying it so much! it simply boosts the productivity 🔥

It makes your code shorter and easier to understand and enforces you to follow best practice (normalizing, selectors, typing, etc… )

👇🏽You'll find concrete examples and code in the article below 👇🏽[https://blog.theodo.com/2020/01/reduce-redux-boilerplate/](https://blog.theodo.com/2020/01/reduce-redux-boilerplate/)
## [2][Decoupled State Interface](https://www.reddit.com/r/reduxjs/comments/g4nxu4/decoupled_state_interface/)
- url: https://github.com/gactjs/store/blob/master/docs/decoupled-state-interface.md
---

## [3][Normalized Reducer: an easier way to read and write normalized relational reducer state](https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/)
- url: https://www.reddit.com/r/reduxjs/comments/g3jdjv/normalized_reducer_an_easier_way_to_read_and/
---
[**Normalized Reducer**](https://github.com/brietsparks/normalized-reducer) helps you manage normalized relational state without requiring any reducer/action boilerplate.

You simply feed in a schema of the relationships, and it gives you the reducers, actions, and selectors to read and write state according to that schema. The [**docs quick start**](https://github.com/brietsparks/normalized-reducer#quick-start) shows how easy it is to set up and use. Check out the [**demo app**](https://brietsparks.github.io/normalized-reducer-demo/) to see its features in action.

Additionally, it integrates with Normalizr. First denormalize your data with Normalizr's [`normalize`](https://github.com/paularmstrong/normalizr/blob/master/docs/api.md#normalizedata-schema), and then pass that output into Normalized Reducer's [`fromNormalizr`](https://github.com/brietsparks/normalized-reducer#normalizr-integration). The final result will be normalized state that the reducer can handle. See the demo's Normalizr example to see shapes of the transformation.

To help decide whether this library is right for you, there is a [comparison to Redux ORM and Redux Toolkit's entity adapter](https://github.com/brietsparks/normalized-reducer#comparison-to-alternatives). Key points: Relational Reducer is lighter, zero dep, framework agnostic, and covers some important use-cases the others are missing. However, it (currently) has fewer selector features than Redux ORM and is less mature than both.

Thanks for reading, and I hope this helps build stuff faster!
## [4][Gact Store White Paper](https://www.reddit.com/r/reduxjs/comments/g3q2lj/gact_store_white_paper/)
- url: https://github.com/gactjs/store/blob/master/docs/white-paper.md
---

## [5][Keeping the userActions DRY in redux](https://www.reddit.com/r/reduxjs/comments/g2wey8/keeping_the_useractions_dry_in_redux/)
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
## [6][How do we try to get data directly under streams? why does it has undefined and data?](https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/)
- url: https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/
---
The problem is when I bring the data from backend, There will be undefined under streams reducer, but if I same data and act as fake backend, I wont get undefined.

[When fetching from real backend](https://preview.redd.it/m4jjwmzw4dt41.png?width=1768&amp;format=png&amp;auto=webp&amp;s=bf950010545e4fa108e579c8125f4f7316766e5e)

&amp;#x200B;

[When bringing data from fake backend](https://preview.redd.it/h8ayi9705dt41.png?width=1192&amp;format=png&amp;auto=webp&amp;s=63e32e08a66f8e725e1b7b555b2b29a51eb31356)
## [7][What is the difference between time travel and undo/redo functionality ?](https://www.reddit.com/r/reduxjs/comments/g2qf27/what_is_the_difference_between_time_travel_and/)
- url: https://www.reddit.com/r/reduxjs/comments/g2qf27/what_is_the_difference_between_time_travel_and/
---
Is time travel just capturing a whole snapshot of the state just for debugging purposes during development stage ?

Is undo/redo functionality a more efficient (cpu and RAM wise) version of time traveling that is supposed to be used by users ?
## [8][Tools for creating a client side database diagram .](https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/)
- url: https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/
---
I am trying to make an app in which the state is a little bit complex so I have to make it be like a normalized client side database as suggested in the redux docs  \[[1](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/)\] .

It would be extremely helpful for me if there is any kind of tool/app that allows me to create a client side normalized database diagram , like for example [this](https://erdplus.com/standalone) tool .

Unfortunately this tool does not provide me with types that exists in JS (array for example). I want the tool to allow me to define types like I do in typescript .

Since my database is a single object it would be nice for that tool to make a d.ts file for that object .

Also it would be extremely helpful if that tool would allow me to save my work .

It really feel painful defining my normalized client side database in a d.ts file , I think a tool like the one I described will boost productivity .

I am really noob regarding all that database thing , so sorry what I am asking sounds stupid .
## [9][Modern React Redux Tutorials with Redux toolkit - 2020](https://www.reddit.com/r/reduxjs/comments/fzx905/modern_react_redux_tutorials_with_redux_toolkit/)
- url: https://cloudnweb.dev/2020/04/modern-react-redux-tutotials-redux-toolkit/
---

## [10][Designing a normalized state : Arrays of ids should be used to indicate ordering .](https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/)
- url: https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/
---
This is mentioned [here](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/#designing-a-normalized-state) . What I do not understand is :

1. why arrays of ids should indicate ordering for the ids ? why cant the ids themselves represent ordering ?
2. what is the profit/best practice that we get from creating one more nest in the state to just add the `allIds` property ?
3. how are the unique ids produced ?
