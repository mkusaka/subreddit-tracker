# reduxjs
## [1][redux-saga style-guide [x-post]](https://www.reddit.com/r/reduxjs/comments/eik2ak/reduxsaga_styleguide_xpost/)
- url: https://www.reddit.com/r/reactjs/comments/eik1od/reduxsaga_styleguide
---

## [2][Redux in 30 seconds](https://www.reddit.com/r/reduxjs/comments/eiktfu/redux_in_30_seconds/)
- url: https://medium.com/@pulkitsingh01/redux-in-30-seconds-4527722aa068
---

## [3][Understand and implement your own Redux](https://www.reddit.com/r/reduxjs/comments/ehoea2/understand_and_implement_your_own_redux/)
- url: https://buttercms.com/blog/understand-and-implement-your-own-redux
---

## [4][Simplifying WebSockets in RxJS.](https://www.reddit.com/r/reduxjs/comments/ehbvdf/simplifying_websockets_in_rxjs/)
- url: https://medium.com//simplifying-websockets-in-rxjs-a177b887f3b8?source=friends_link&amp;sk=a932c9f76ea91c8a763cc1963dad337e
---

## [5][Working with Nested Objects](https://www.reddit.com/r/reduxjs/comments/eghyy3/working_with_nested_objects/)
- url: https://www.reddit.com/r/reduxjs/comments/eghyy3/working_with_nested_objects/
---
Relatively new to React-Redux here, and wondering if someone can direct me to examples of how to work with updating nested objects. I'm aware that \`combineReducers\` can be used, but I've not yet seem examples of how Redux can be coded to deal with levels-within-levels of data storage.

Thanks for any help on this!
## [6][Partial Implementation of Redux's hook API in Pure React Hooks](https://www.reddit.com/r/reduxjs/comments/eg3dtz/partial_implementation_of_reduxs_hook_api_in_pure/)
- url: https://www.reddit.com/r/reduxjs/comments/eg3dtz/partial_implementation_of_reduxs_hook_api_in_pure/
---
Here's the link: [https://codesandbox.io/s/upbeat-hooks-5hu6f](https://codesandbox.io/s/upbeat-hooks-5hu6f)   


The implementation is in the store.js file. It implements *useStore*, *useDispatch* and *useSelector* using only pure React Hooks.   


Why did I make it? Just because I had some free time during the holiday and noticed that Redux had a hook based API and was curious if I could basically make the same thing (I've been using a lot of hooks lately)
## [7][Common Redux Pitfalls](https://www.reddit.com/r/reduxjs/comments/eepifw/common_redux_pitfalls/)
- url: https://blog.alessandrom.dev/common-redux-pitfalls/
---

## [8][Why Redux advises against using actions as setters?](https://www.reddit.com/r/reduxjs/comments/ed3609/why_redux_advises_against_using_actions_as_setters/)
- url: https://www.reddit.com/r/reduxjs/comments/ed3609/why_redux_advises_against_using_actions_as_setters/
---
Hey folks, I've been looking into Redux best practices and I found that [Redux style guide strongly recommends against using actions as setters](https://redux.js.org/style-guide/style-guide#model-actions-as-events-not-setters), at first I agreed with it but after some introspection I fail to see the issue with this, like imagine you want to fetch the user data, you would have something like this:

    async function fetchUser (userId) {
        dispatch({ type: 'SET_USER_LOADING' })
        try {
            const payload = await service.fetchUser(userId)
            dispatch({ type: 'SET_USER', payload })
        } catch (error) {
            dispatch({ type: 'SET_USER', error })
        } finally {
            dispatch({ type: 'SET_USER_LOADED' })
        }
    }

Why is it considered a bad practice to do that? Or am I missing the point?

&amp;#x200B;

EDIT: Thanks a lot to all of your replies, I believe my doubts came from poor naming of my actions, going back to the style guide it seems much clearer on the issue. You guys are breathtaking!
## [9][Simplified "Redux + Redux-Rematch" for Tiny Bundles](https://www.reddit.com/r/reduxjs/comments/ecn5os/simplified_redux_reduxrematch_for_tiny_bundles/)
- url: https://www.reddit.com/r/reduxjs/comments/ecn5os/simplified_redux_reduxrematch_for_tiny_bundles/
---
I'd like to get some feedback on something I've been working on. It's a tiny alternative to Redux and Redux-Rematch that aims to make the redux approach to state management easier and produce much smaller, faster-loading JS bundles.

The [example app](https://captaincodeman.github.io/rdx-example/) is just 2.98 Kb minified and gzipped JS but contains a complete, if simplistic, client-side app consisting of:

- redux-like state store (the usual pattern of actions, reducers &amp; middleware)
- redux devtools support for debugging
- persistence middleware (using localStorage)
- async thunk middleware (examples for loading from REST API)
- client side routing (part of the state store)
- some web-components bound to the store (simple UI

The example is meant to demonstrate the pieces working - redux devtools extensions, routing with parameter extraction, loading remote data asynchronously with some logic to only fetch if not already loaded and of course a counter to show you can click on something and see it change. State is persisted to local storage so you have to clear that to see the data loading parts again (and the different logic when going from single todo item as the first view to the list view instead of vice-versa).

It gets 100's across the board on Lighthouse and [regular 3G webpagetest.org load time](https://www.webpagetest.org/result/191219_FP_6539bbcb495ead0943f0cb123a82d829/#run3) is less that 2.3 seconds, with low CPU use, so it should be mobile-friendly.

I know this isn't a _totally_ new idea - there have been other "tiny redux" implementations and I'm sure it's not the smallest, plus there are already packages such as redux-rematch (which was a big inspiration) and others such as redux-bundler. This just combines both pieces. In [Homer Simpson style](https://comb.io/V9E2bk):

&gt; "These are the people who saw an overcrowded marketplace and said, ‘Me too!'"

I've been using it on a proper app and the smaller package sizes plus the savings from having to write less store-related code have resulted in going from a bundle size of 156.5 Kb minified / 47 Kb gzipped to 102 Kb minified / 30 Kb gzipped, approximately 65% of the original. The biggest win though is probably that I'm finding the coding far more pleasant, with less time spent on repetitive redux boilerplate and more on the actual app functionality so I feel I'm getting more done (well, when I'm not writing things like this).

If you want to look at the code, it's split into three source repos:

* The [example app](https://github.com/CaptainCodeman/rdx-example)
* The [basic state store](https://github.com/CaptainCodeman/rdx) and middlewares
* The [helpers](https://github.com/CaptainCodeman/rdx-model) for defining state models

The model helpers repo contains the best explanation for what it does and how it works and the aim is for it all to be Typescript friendly, but I still have a couple of things to workout there (around making the models defined in plugins, such as routing, available)

BTW: I'm calling it "rdx" because it's "like Redux, but smaller". Thankfully, I don't have to rely on my marketing skills to put food on the table.

Any feedback, suggestions, criticisms etc... would be welcome :)
## [10][Redux shouldn't hold all of your state](https://www.reddit.com/r/reduxjs/comments/ec4ahq/redux_shouldnt_hold_all_of_your_state/)
- url: https://michaelwashburnjr.com/2019/12/09/stop-storing-data-redux/
---

