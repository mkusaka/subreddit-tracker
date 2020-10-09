# javascript
## [1][WTF Wednesday (October 07, 2020)](https://www.reddit.com/r/javascript/comments/j6tkx9/wtf_wednesday_october_07_2020/)
- url: https://www.reddit.com/r/javascript/comments/j6tkx9/wtf_wednesday_october_07_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Develop, build, deploy, redeploy, and teardown frontend projects fast, really fast ⚡️](https://www.reddit.com/r/javascript/comments/j7tmtt/develop_build_deploy_redeploy_and_teardown/)
- url: https://github.com/reblim/fast
---

## [3][JavaScript escape room game: [code]capi Code Out!](https://www.reddit.com/r/javascript/comments/j7x2ef/javascript_escape_room_game_codecapi_code_out/)
- url: https://codeout.codecapi.com
---

## [4][Excel import and export open source library with javascript-Luckyexcel](https://www.reddit.com/r/javascript/comments/j7um06/excel_import_and_export_open_source_library_with/)
- url: https://github.com/mengshukeji/Luckyexcel
---

## [5][ESLint A year of paying contributors: Review](https://www.reddit.com/r/javascript/comments/j7cgwu/eslint_a_year_of_paying_contributors_review/)
- url: https://eslint.org/blog/2020/10/year-paying-contributors-review
---

## [6][Visual Studio Code September 2020](https://www.reddit.com/r/javascript/comments/j7hgvq/visual_studio_code_september_2020/)
- url: https://code.visualstudio.com/updates/v1_50
---

## [7][Understanding Scope: Javascript Concepts(intermittent) by PotetoTech](https://www.reddit.com/r/javascript/comments/j7tc1a/understanding_scope_javascript/)
- url: https://www.potetotech.com/2020/08/understanding-scope-javascript-concepts.html
---

## [8][How to get and set the value of a select menu with vanilla JS](https://www.reddit.com/r/javascript/comments/j7xsvf/how_to_get_and_set_the_value_of_a_select_menu/)
- url: https://gomakethings.com/how-to-get-and-set-the-value-of-a-select-menu-with-vanilla-js/
---

## [9][What happened to Immutable.JS? And how can we react?](https://www.reddit.com/r/javascript/comments/j7j9ky/what_happened_to_immutablejs_and_how_can_we_react/)
- url: https://dev.to/alekseiberezkin/what-happened-to-immutable-js-and-how-can-we-react-5c34
---

## [10][[AskJS] Has Cypress ever been using for both "unit" and e2e testing?](https://www.reddit.com/r/javascript/comments/j7qa8s/askjs_has_cypress_ever_been_using_for_both_unit/)
- url: https://www.reddit.com/r/javascript/comments/j7qa8s/askjs_has_cypress_ever_been_using_for_both_unit/
---
Whether this has been done before or not, I'm going to try to implement it. I'm just wondering if there is prior art out there to reference. Also, feedback on the idea in general would be appreciated.

**Background**

So, I've done a lot of unit testing with tools like Jest/Enzyme, and I've done e2e testing with Cypress and Selenium. I hate Selenium and love Cypress, btw. Given that I am a strong opponent of snapshot testing with Jest/Enzyme, I find myself writing similar logic in both Jest and Cypress. For example, when I have a presentational component, I'll write Jest logic that tests that the fields render properly for the data I pass in as props. Then, in my Cypress e2e testing, I'll write basically the same thing, but for Cypress.

In addition, I've been embracing the philosophy of integration tests over unit tests, relying on tests that do a more thorough job of testing the application's actual full behavior rather than just narrow pieces of it. Doesn't mean I don't write unit tests anymore, they just tend to be a smaller percentage of my overall test suite.

**My Goal**

I want to write a Cypress test suite that can serve as both "integration" style tests and my e2e tests, eliminating nearly all the need to use Jest/Enzyme. Here's a rough overview of how I'm planning to have it work. Keep in mind that a lot of the ideas here are pretty rough and I have not dove too deep into the planning of this project.

**Integration Tests**

1. My React app would need to be mounted somehow, preferably headless so this could run in a CI/CD pipeline.
2. Cypress would run headless (I know it can do that).
3. On initialization, code would run that would mock all API calls, so they wouldn't go anywhere.  

   1. This is the only piece of the app that gets mocked, everything else is the full end-to-end logic.
4. Cypress tests would then "click" through the app, relying on the mocked data to handle API responses, validating the UI in the process.

**End to End Tests**

1. Same exact Cypress suite as before, just with some configuration changes.
2. Cypress runs headless against a live version of the app with real (qa) data and real interactions with other services.
3. No APIs are mocked.
4. Cypress tests will properly validate the running application.

**Final Thoughts**

As you can see, the objective is to have one Cypress suite that can both do integration testing against mock data, and actual testing of the live app. I feel these kind of tests will be more robust than Jest/Enzyme unit tests, and since they are re-usable for the e2e tests they have the added benefit of decreasing the overall time spent on testing (unit + e2e tests being written simultaneously). Obviously I would expect some errors to occur when running one way vs another, but it would still be a time saver without sacrificing test coverage.

This is not to say I would expect this to replace all unit tests. There are still plenty of cases where normal unit tests are still essential, such as some fancy TypeScript dynamic type checking logic I wrote recently. However, I really do like the overall direction I'm trying to go in here.

Alright, this wall of text is done. It's late and I'm going to bed. Appreciate both being pointed to prior art that's out there and feedback on the overall idea.
## [11][[AskJS] How to clean up npm scripts in package json?](https://www.reddit.com/r/javascript/comments/j7vjuv/askjs_how_to_clean_up_npm_scripts_in_package_json/)
- url: https://www.reddit.com/r/javascript/comments/j7vjuv/askjs_how_to_clean_up_npm_scripts_in_package_json/
---
We have many npm scripts in our package json, and the file looks verbose and messy. Do you recommend any tools to make it cleaner?

&amp;#x200B;

In the past, I saw a library that let you define the scripts in a YML file and support things like parallel execution, but I can't remember the name.

&amp;#x200B;

Thanks in advance.
