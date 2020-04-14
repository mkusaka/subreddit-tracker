# reduxjs
## [1][Tools for creating a client side database diagram .](https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/)
- url: https://www.reddit.com/r/reduxjs/comments/g0n3zg/tools_for_creating_a_client_side_database_diagram/
---
I am trying to make an app in which the state is a little bit complex so I have to make it be like a normalized client side database as suggested in the redux docs  \[[1](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/)\] .

It would be extremely helpful for me if there is any kind of tool/app that allows me to create a client side normalized database diagram , like for example [this](https://erdplus.com/standalone) tool .

Unfortunately this tool does not provide me with types that exists in JS (array for example). I want the tool to allow me to define types like I do in typescript .

Since my database is a single object it would be nice for that tool to make a d.ts file for that object .

Also it would be extremely helpful if that tool would allow me to save my work .

It really feel painful defining my normalized client side database in a d.ts file , I think a tool like the one I described will boost productivity .

I am really noob regarding all that database thing , so sorry what I am asking sounds stupid .
## [2][Modern React Redux Tutorials with Redux toolkit - 2020](https://www.reddit.com/r/reduxjs/comments/fzx905/modern_react_redux_tutorials_with_redux_toolkit/)
- url: https://cloudnweb.dev/2020/04/modern-react-redux-tutotials-redux-toolkit/
---

## [3][Designing a normalized state : Arrays of ids should be used to indicate ordering .](https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/)
- url: https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/
---
This is mentioned [here](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/#designing-a-normalized-state) . What I do not understand is :

1. why arrays of ids should indicate ordering for the ids ? why cant the ids themselves represent ordering ?
2. what is the profit/best practice that we get from creating one more nest in the state to just add the `allIds` property ?
3. how are the unique ids produced ?
## [4][Constantly updating in background](https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/)
- url: https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/
---
Totally new to react and redux, so please excuse my dumb questions.
I already managed to write a component, that opens a webcam and displays it, I also managed to add a button in this component to grab an image and send it to a handler function outside of this component.

Now I would like to do the following: I would like to write a page, that grabs a pic from the webcam, sends it via fetch to a server - waits for the reply and after the reply came does the same again. I’m currently still lost how to do that within the React - redux(-toolkit) framework.

Would love to get hints on it.
## [5][Difference between storing data in localStorage vs the state](https://www.reddit.com/r/reduxjs/comments/fy14l4/difference_between_storing_data_in_localstorage/)
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
## [6][Is there any difference b/w these two redux saga effects takeEvery() and while(true) { take() } ?](https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/)
- url: https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/
---
What are the pros and cons of these saga effects?
## [7][Dynadux, an alternative to Redux, simpler and powerful](https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/)
- url: https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/
---
Hello fellows

I have created a new alternative to the Redux library. It is much simpler and powerful. [https://github.com/aneldev/dynadux](https://github.com/aneldev/dynadux)

I would appreciate having your feedback!
## [8][React, Redux and a little bit of math.](https://www.reddit.com/r/reduxjs/comments/fozpe4/react_redux_and_a_little_bit_of_math/)
- url: https://medium.com/@dmitriykharchenko/react-redux-and-a-little-bit-of-math-fe2c9a4a477c?source=friends_link&amp;sk=a826809258bc807c0919cdf63e49c766
---

## [9][Are there any good caching solutions or libraries for Redux?](https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/)
- url: https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/
---
Are there any best practices, guides, or libraries out there to simplify/standardize how caching is handled in Redux?

For now, I'm keeping a boolean in my reducers. If that value is false and I try to fetch data, it'll just no-op.
## [10][Tricks you never knew about Redux DevTool](https://www.reddit.com/r/reduxjs/comments/fkyd85/tricks_you_never_knew_about_redux_devtool/)
- url: https://blog.logrocket.com/redux-devtools-tips-tricks-for-faster-debugging/
---

