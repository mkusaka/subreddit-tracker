# javascript
## [1][WTF Wednesday (October 21, 2020)](https://www.reddit.com/r/javascript/comments/jffpsl/wtf_wednesday_october_21_2020/)
- url: https://www.reddit.com/r/javascript/comments/jffpsl/wtf_wednesday_october_21_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Building a React app to solve every Sudoku puzzle.](https://www.reddit.com/r/javascript/comments/jfr1ow/building_a_react_app_to_solve_every_sudoku_puzzle/)
- url: https://dev.to/jsolano/building-a-react-app-to-solve-every-sudoku-puzzle-3c95
---

## [3][Using const/let instead of var can make JavaScript code run 10× slower in Webkit?](https://www.reddit.com/r/javascript/comments/jfdiq3/using_constlet_instead_of_var_can_make_javascript/)
- url: https://github.com/evanw/esbuild/issues/478
---

## [4][How we built the Chrome DevTools WebAuthn tab | Google Developers](https://www.reddit.com/r/javascript/comments/jfz21h/how_we_built_the_chrome_devtools_webauthn_tab/)
- url: https://developers.google.com/web/updates/2020/10/webauthn-tab
---

## [5][React hooks and Ember.JS](https://www.reddit.com/r/javascript/comments/jfze3y/react_hooks_and_emberjs/)
- url: https://www.effective-ember.com/blog/react-hooks-and-ember/
---

## [6][Managing side effects with monads in JavaScript with practical examples](https://www.reddit.com/r/javascript/comments/jflrse/managing_side_effects_with_monads_in_javascript/)
- url: https://www.7urtle.com/javascript-applicative-functor-monads
---

## [7][Wouldn't it be nice to have "Symbol.promise" as standard Symbol, similar to "Symbol.asyncIterator"?](https://www.reddit.com/r/javascript/comments/jfp0tk/wouldnt_it_be_nice_to_have_symbolpromise_as/)
- url: https://dev.to/noseratio/wouldn-t-it-be-nice-to-have-symbol-promise-as-standard-symbol-similar-to-symbol-asynciterator-410o
---

## [8][Polling in React using the useInterval Custom Hook](https://www.reddit.com/r/javascript/comments/jfi42y/polling_in_react_using_the_useinterval_custom_hook/)
- url: https://blog.asayer.io/polling-in-react-using-the-useinterval-custom-hook
---

## [9][Node.js v15.0.0 is here!](https://www.reddit.com/r/javascript/comments/jes44j/nodejs_v1500_is_here/)
- url: https://medium.com/@nodejs/node-js-v15-0-0-is-here-deb00750f278
---

## [10][[AskJS] What's the pros and cons of using a non-building framework today?](https://www.reddit.com/r/javascript/comments/jfjfu4/askjs_whats_the_pros_and_cons_of_using_a/)
- url: https://www.reddit.com/r/javascript/comments/jfjfu4/askjs_whats_the_pros_and_cons_of_using_a/
---
Context: We have a huge legacy web app, with first commit 2003. It's based on the Yii PHP framework. Now we're thinking of moving beyond jQuery on the client side. We tried so far with multi-root Vue, but it ended up "suboptimal" and now we're actually reverting the changes back to vanilla PHP framework with a very thin layer of JavaScript. This "thinness" is relative, though, and we must prepare another way forward. Both Flow and TypeScript work on vanilla JS, so this can be included, as well as linting (NB: we have the luxury to assume support of ES8 in our administration interface). VueJS seem to "work" without any build system, which could be a gentle step out of jQuery/vanilla, but are there too many drawbacks? Feel free to leave any feedback. :)
## [11][[AskJS] Yarn workspaces vs importing from (aliased) directory](https://www.reddit.com/r/javascript/comments/jfa11j/askjs_yarn_workspaces_vs_importing_from_aliased/)
- url: https://www.reddit.com/r/javascript/comments/jfa11j/askjs_yarn_workspaces_vs_importing_from_aliased/
---
Hello :)

I'm trying to wrap my head around what the best approach for organizing code in a monorepo is.

Let's  say I have a \`landing-page\`, \`app\`, and \`blog\` folders which all  contain code for the respective sites. I also have a \`common\` folder  where I put utilities which I use across all the three sites.

My  current approach is that I use expressions like \`import {formatDate}  from 'common/date'\` to import code from the \`common\` folder into the  \`landing-page\`, \`app\`, or \`blog\` directories. (I use Webpack aliases to  be able to use \`import ... from 'common/date\` rather than \`import ...  from '../../../common'\`

I'd like  to know, however, what benefits would using a tool like Yarn Workspaces  bring? In what way it is better to turn my \`landing-page\`, \`app\`,  \`blog\`, and \`common\` folders into "installable" npm packages which I  would import from each other?

To  me, it seems like a lot more boilerplate than my current approach and I  don't see the benefit of using it. However, it's possible that something  just didn't click with me as for why I should use Yarn Workspaces. So  I'd love to hear from you.
