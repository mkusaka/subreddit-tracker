# reduxjs
## [1][Testing if Redux-Toolkit AsyncThunks are called in components](https://www.reddit.com/r/reduxjs/comments/jep6q0/testing_if_reduxtoolkit_asyncthunks_are_called_in/)
- url: https://www.reddit.com/r/reduxjs/comments/jep6q0/testing_if_reduxtoolkit_asyncthunks_are_called_in/
---
Hi,

I recently started writing tests for my app and am stuck in trying to test my components that call Redux-Toolkit's AsyncThunks.

I have a component that calls a few AsyncThunks on mount. I have a test setup for it:

    it('should dispatch actions', () =&gt; {
    	expect.hasAssertions();
    	const useDispatchSpy = jest.spyOn(redux, 'useDispatch');
    	const mockDispatchFn = jest.fn();
    	useDispatchSpy.mockReturnValue(mockDispatchFn);
    
    	mount(
    		&lt;Provider store={store}&gt;
    			&lt;Dashboard /&gt;
    		&lt;/Provider&gt;,
    	);
    
    	expect(mockDispatchFn).toHaveBeenCalledTimes(4);
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchCurriculumPrefixes());
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchHeatmap(''));
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchCurriculumFromPrefix(''));
    	expect(mockDispatchFn).toHaveBeenCalledWith(syncTimetable({
    		timetable: [], timestamp: '',
    	}));
    	useDispatchSpy.mockClear();
    });

The problem with this is, the tests give a very uninformative error on failing:

    expect(jest.fn()).toHaveBeenCalledWith(...expected)
    
    Expected: [Function anonymous]
    Received
           1: [Function anonymous]
           2: [Function anonymous]
           3: [Function anonymous]
    

How would you go about testing these calls? I just need to make sure that certain calls are made in the useEffect hooks.

Would love any help, Thanks!
## [2][Toolkit in Production?](https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/)
- url: https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/
---
I see that [https://redux.js.org/tutorials/essentials/part-2-app-structure](https://redux.js.org/tutorials/essentials/part-2-app-structure) Redux is pushing toolkit now. How many of you are using Redux Toolkit in production?
## [3][Where to add hooks after dynamically adding reducers](https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/)
- url: https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/
---
Hey all,

I've seen a few similar questions to this but thought I might go ahead and ask mine.

I'm dynamically adding reducers and middleware to my App but now I'm to the point where I want to tie those reducers I added to the Apps store for reference (Essentially, I'm trying to create self-contained widgets that adds what they need to the Redux store so that I may eventually add/remove components)

I'm not sure how to do it though. I'm using Redux Observables... I'm trying to useStore... or useSelector but I can't use them inside useEffect but I also can't try and access it before the reducers have been dynamically added in the use effect. I've done a ton of googling on this and I feel like I'm missing something key here....
## [4][pain free form validator that does not stand in your way](https://www.reddit.com/r/reduxjs/comments/j4wszk/pain_free_form_validator_that_does_not_stand_in/)
- url: https://github.com/euvoor/form
---

## [5][Why does async action creator in redux-thunk dispatches an action with payload rather than just returning returning an action object with payload ?](https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/)
- url: https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/
---
    export const fetchPost = () =&gt;async (dispatch) =&gt;{
    const resp = await jsonPlaceholder.get('/posts');
        dispatch({type:'FETCH_POST',payload:resp.data})
    }
## [6][Dynamically injecting reducers and sagas](https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/)
- url: https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/
---
I am new to redux and really need some resources on learning dynamic injection of reducers and sagas. I saw a few articles but I didn't understand anything. Please help! Thank you.
## [7][React Native With Redux In Expo](https://www.reddit.com/r/reduxjs/comments/iuhd55/react_native_with_redux_in_expo/)
- url: https://www.youtube.com/watch?v=MiJayg1eZvk&amp;feature=share
---

## [8][Visualize The Power Of Redux and Memoization In React](https://www.reddit.com/r/reduxjs/comments/iu3z0v/visualize_the_power_of_redux_and_memoization_in/)
- url: https://www.youtube.com/watch?v=KypVn6vGFWg&amp;feature=share
---

## [9][Advanced Express JS REST API [#1] Introduction | Building REST API Node JS | Full Course - Please subscribe](https://www.reddit.com/r/reduxjs/comments/itgx6u/advanced_express_js_rest_api_1_introduction/)
- url: https://m.youtube.com/watch?v=CLdkGgv9Miw&amp;list=PLs1waz0ZKTGO7agN0cntpe6ro6TIka0ow&amp;index=2&amp;t=0s
---

## [10][What's your approach to the size of redux state slices?](https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/)
- url: https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/
---
I like to have my redux state slices small and focused on specific parts of the interface.

But at the same time, I am feeling a growing resistance from the library to splitting the state, because some slice reducers end up wanting information from other slices of the state in order to decide how to update their portion of the state.

I've examined the [docs](https://redux.js.org/recipes/structuring-reducers/beyond-combinereducers#sharing-data-between-slice-reducers) to see what they advise. The first suggestion (passing parent state as the third argument) [does not seem to be supported by the typescript definitions](https://github.com/reduxjs/redux/pull/3465). The second suggestion — just use thunks — is something I am currently doing, but this both feels like a hack (thunks were intended to handle async code; that's why they are called thunks, or long-running computations), and ruins most of the elegance of redux-toolkit.

What has your experience been? How large are your state slices? Do you often need to look up data from other slices? How do you manage this?
