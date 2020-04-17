# reduxjs
## [1][Keeping the userActions DRY in redux](https://www.reddit.com/r/reduxjs/comments/g2wey8/keeping_the_useractions_dry_in_redux/)
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
## [2][How do we try to get data directly under streams? why does it has undefined and data?](https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/)
- url: https://www.reddit.com/r/reduxjs/comments/g304ob/how_do_we_try_to_get_data_directly_under_streams/
---
The problem is when I bring the data from backend, There will be undefined under streams reducer, but if I same data and act as fake backend, I wont get undefined.

[When fetching from real backend](https://preview.redd.it/m4jjwmzw4dt41.png?width=1768&amp;format=png&amp;auto=webp&amp;s=bf950010545e4fa108e579c8125f4f7316766e5e)

&amp;#x200B;

[When bringing data from fake backend](https://preview.redd.it/h8ayi9705dt41.png?width=1192&amp;format=png&amp;auto=webp&amp;s=63e32e08a66f8e725e1b7b555b2b29a51eb31356)
## [3][What is the difference between time travel and undo/redo functionality ?](https://www.reddit.com/r/reduxjs/comments/g2qf27/what_is_the_difference_between_time_travel_and/)
- url: https://www.reddit.com/r/reduxjs/comments/g2qf27/what_is_the_difference_between_time_travel_and/
---
Is time travel just capturing a whole snapshot of the state just for debugging purposes during development stage ?

Is undo/redo functionality a more efficient (cpu and RAM wise) version of time traveling that is supposed to be used by users ?
## [4][Tools for creating a client side database diagram .](https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/)
- url: https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/
---
I am trying to make an app in which the state is a little bit complex so I have to make it be like a normalized client side database as suggested in the redux docs  \[[1](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/)\] .

It would be extremely helpful for me if there is any kind of tool/app that allows me to create a client side normalized database diagram , like for example [this](https://erdplus.com/standalone) tool .

Unfortunately this tool does not provide me with types that exists in JS (array for example). I want the tool to allow me to define types like I do in typescript .

Since my database is a single object it would be nice for that tool to make a d.ts file for that object .

Also it would be extremely helpful if that tool would allow me to save my work .

It really feel painful defining my normalized client side database in a d.ts file , I think a tool like the one I described will boost productivity .

I am really noob regarding all that database thing , so sorry what I am asking sounds stupid .
## [5][Modern React Redux Tutorials with Redux toolkit - 2020](https://www.reddit.com/r/reduxjs/comments/fzx905/modern_react_redux_tutorials_with_redux_toolkit/)
- url: https://cloudnweb.dev/2020/04/modern-react-redux-tutotials-redux-toolkit/
---

## [6][Designing a normalized state : Arrays of ids should be used to indicate ordering .](https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/)
- url: https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/
---
This is mentioned [here](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/#designing-a-normalized-state) . What I do not understand is :

1. why arrays of ids should indicate ordering for the ids ? why cant the ids themselves represent ordering ?
2. what is the profit/best practice that we get from creating one more nest in the state to just add the `allIds` property ?
3. how are the unique ids produced ?
## [7][Constantly updating in background](https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/)
- url: https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/
---
Totally new to react and redux, so please excuse my dumb questions.
I already managed to write a component, that opens a webcam and displays it, I also managed to add a button in this component to grab an image and send it to a handler function outside of this component.

Now I would like to do the following: I would like to write a page, that grabs a pic from the webcam, sends it via fetch to a server - waits for the reply and after the reply came does the same again. I’m currently still lost how to do that within the React - redux(-toolkit) framework.

Would love to get hints on it.
## [8][Difference between storing data in localStorage vs the state](https://www.reddit.com/r/reduxjs/comments/fy14l4/difference_between_storing_data_in_localstorage/)
- url: https://www.reddit.com/r/reduxjs/comments/fy14l4/difference_between_storing_data_in_localstorage/
---
I currently have an application using redux and I have it set that every time the app loads it checks the JWT token is valid and adds the user data to the state.

I was wondering what the differences are between calling the api and then storing data in the state every reload or storing the data once in `localStorage`? 

How the code is setup with calling the api and storing with redux. 

CHECK TOKEN
```
const token = localStorage.UserIdToken;
if (token) {
  const decodedToken = jwtDecode(token);

  if (decodedToken.exp * 1000 &lt; Date.now()) {
    store.dispatch(logoutUser());
  } else {
    store.dispatch({ type: SET_AUTHENTICATED });
    axios.defaults.headers.common['Authorization'] = token;
    store.dispatch(getUserData());
  }
}
```
`getUserData()`
```
export const getUserData = () =&gt; async dispatch =&gt; {
  try {
    const res = await axios.get('/user');
    dispatch({
      type: SET_USER,
      payload: res.data,
    });
  } 
  ...
};
```
## [9][Is there any difference b/w these two redux saga effects takeEvery() and while(true) { take() } ?](https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/)
- url: https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/
---
What are the pros and cons of these saga effects?
## [10][Dynadux, an alternative to Redux, simpler and powerful](https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/)
- url: https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/
---
Hello fellows

I have created a new alternative to the Redux library. It is much simpler and powerful. [https://github.com/aneldev/dynadux](https://github.com/aneldev/dynadux)

I would appreciate having your feedback!
