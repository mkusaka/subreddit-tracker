# reduxjs
## [1][Designing a normalized state : Arrays of ids should be used to indicate ordering .](https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/)
- url: https://www.reddit.com/r/reduxjs/comments/fz6fyl/designing_a_normalized_state_arrays_of_ids_should/
---
This is mentioned [here](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/#designing-a-normalized-state) . What I do not understand is :

1. why arrays of ids should indicate ordering for the ids ? why cant the ids themselves represent ordering ?
2. what is the profit/best practice that we get from creating one more nest in the state to just add the `allIds` property ?
3. how are the unique ids produced ?
## [2][Constantly updating in background](https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/)
- url: https://www.reddit.com/r/reduxjs/comments/fyb1pz/constantly_updating_in_background/
---
Totally new to react and redux, so please excuse my dumb questions.
I already managed to write a component, that opens a webcam and displays it, I also managed to add a button in this component to grab an image and send it to a handler function outside of this component.

Now I would like to do the following: I would like to write a page, that grabs a pic from the webcam, sends it via fetch to a server - waits for the reply and after the reply came does the same again. I’m currently still lost how to do that within the React - redux(-toolkit) framework.

Would love to get hints on it.
## [3][Difference between storing data in localStorage vs the state](https://www.reddit.com/r/reduxjs/comments/fy14l4/difference_between_storing_data_in_localstorage/)
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
## [4][Is there any difference b/w these two redux saga effects takeEvery() and while(true) { take() } ?](https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/)
- url: https://www.reddit.com/r/reduxjs/comments/fsd1cn/is_there_any_difference_bw_these_two_redux_saga/
---
What are the pros and cons of these saga effects?
## [5][Dynadux, an alternative to Redux, simpler and powerful](https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/)
- url: https://www.reddit.com/r/reduxjs/comments/frlhpu/dynadux_an_alternative_to_redux_simpler_and/
---
Hello fellows

I have created a new alternative to the Redux library. It is much simpler and powerful. [https://github.com/aneldev/dynadux](https://github.com/aneldev/dynadux)

I would appreciate having your feedback!
## [6][React, Redux and a little bit of math.](https://www.reddit.com/r/reduxjs/comments/fozpe4/react_redux_and_a_little_bit_of_math/)
- url: https://medium.com/@dmitriykharchenko/react-redux-and-a-little-bit-of-math-fe2c9a4a477c?source=friends_link&amp;sk=a826809258bc807c0919cdf63e49c766
---

## [7][Are there any good caching solutions or libraries for Redux?](https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/)
- url: https://www.reddit.com/r/reduxjs/comments/fojgse/are_there_any_good_caching_solutions_or_libraries/
---
Are there any best practices, guides, or libraries out there to simplify/standardize how caching is handled in Redux?

For now, I'm keeping a boolean in my reducers. If that value is false and I try to fetch data, it'll just no-op.
## [8][Tricks you never knew about Redux DevTool](https://www.reddit.com/r/reduxjs/comments/fkyd85/tricks_you_never_knew_about_redux_devtool/)
- url: https://blog.logrocket.com/redux-devtools-tips-tricks-for-faster-debugging/
---

## [9][Redux Fetch Action Array](https://www.reddit.com/r/reduxjs/comments/fkx5z1/redux_fetch_action_array/)
- url: https://www.reddit.com/r/reduxjs/comments/fkx5z1/redux_fetch_action_array/
---
Hi,  


Is there a way to Dispatch this "products" array somehow to a state ?   


https://preview.redd.it/j25z24mnwhn41.png?width=593&amp;format=png&amp;auto=webp&amp;s=57c7a45584ed17440aac22762c33559056a11e0e

Fetching it with:  


https://preview.redd.it/klg59sjzwhn41.png?width=329&amp;format=png&amp;auto=webp&amp;s=cd9eb470889b4387c3bead056ecfe53fec369cdf

Reducer;

&amp;#x200B;

https://preview.redd.it/rdw4nk05xhn41.png?width=429&amp;format=png&amp;auto=webp&amp;s=bfb82ea4736ba1567bc4a3cbeba97012fda24465
## [10][Filtering, Sorting and Pagination - Advanced Filtering with React and Redux](https://www.reddit.com/r/reduxjs/comments/fkmm0a/filtering_sorting_and_pagination_advanced/)
- url: https://blog.soshace.com/filtering-sorting-and-pagination-advanced-filtering-with-react-and-redux/
---

