# javascript
## [1][WTF Wednesday (November 04, 2020)](https://www.reddit.com/r/javascript/comments/jnza9x/wtf_wednesday_november_04_2020/)
- url: https://www.reddit.com/r/javascript/comments/jnza9x/wtf_wednesday_november_04_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][I built a site to instant-search 32 Million Songs in milliseconds (using InstantSearch.js, ParcelJS and Typesense)](https://www.reddit.com/r/javascript/comments/jou01q/i_built_a_site_to_instantsearch_32_million_songs/)
- url: https://songs-search.typesense.org/
---

## [3][Visual Studio Code October 2020](https://www.reddit.com/r/javascript/comments/jp25ly/visual_studio_code_october_2020/)
- url: https://code.visualstudio.com/updates/v1_51
---

## [4][JavaScript new features (ES2021).](https://www.reddit.com/r/javascript/comments/joeyp0/javascript_new_features_es2021/)
- url: https://sambat-tech.netlify.app/what-new-in-es12/
---

## [5][Why hello world?](https://www.reddit.com/r/javascript/comments/jp4h5h/why_hello_world/)
- url: https://mohamed-muneer.medium.com/why-hello-world-e7d1cb8800c
---

## [6][Why And How To Migrate A React Application to Next.js](https://www.reddit.com/r/javascript/comments/jp402z/why_and_how_to_migrate_a_react_application_to/)
- url: http://reactapplicationtonextjs.space/
---

## [7][[AskJS] Standard is a bad idea](https://www.reddit.com/r/javascript/comments/joqw5b/askjs_standard_is_a_bad_idea/)
- url: https://www.reddit.com/r/javascript/comments/joqw5b/askjs_standard_is_a_bad_idea/
---
On a surface level Standard JS sounds like a good idea, it enforces a consistent code style in your project to improve maintainability. However, I don't like it, first off the name `standard` is just misleading it is not a standard it's a custom package for a custom runtime called Node.js.

EcmaScript doesn't define a 'standard' code style, because it shouldn't exist, there can be a conventional style but not a standard code style. Standard also includes one of the most questionable style decisions which actually increases the chance of making a mistake. Take the following example:

    console.log('Hello, world!')
    
    (() =&gt; {})()

Is that valid code? It should, but it's not, JavaScript uses automatic semicolon insertion, there are specific rules for where it is triggered, and it's much more complicated trying to understand if ASI is triggered or not to just using semicolons everywhere in your code. Besides having to put a semicolon after the first statement but not elsewhere is inconsistent, or putting it before the second statement looks even worse `;(() =&gt; {})()`.

If anyone's wondering, no standard can't catch it as an error, it will accept a semicolon in between but you'd actually have to execute the code before noticing it.
## [8][I built a modern web app with React, TypeScript &amp; Redux best practices.](https://www.reddit.com/r/javascript/comments/jodnkv/i_built_a_modern_web_app_with_react_typescript/)
- url: https://github.com/itailv/awesome-todo-app
---

## [9][[AskJS] Create Amazing Web-Based Interactive Tables and Spreadsheets](https://www.reddit.com/r/javascript/comments/jozxof/askjs_create_amazing_webbased_interactive_tables/)
- url: https://www.reddit.com/r/javascript/comments/jozxof/askjs_create_amazing_webbased_interactive_tables/
---
Is jExcel, a good CE to create an amazing web-based interactive tables and online spreadsheet.
## [10][[AskJS] I like flowtype do you?](https://www.reddit.com/r/javascript/comments/jozo5e/askjs_i_like_flowtype_do_you/)
- url: https://www.reddit.com/r/javascript/comments/jozo5e/askjs_i_like_flowtype_do_you/
---
Before I start I know TS is super popular and there's always comments telling people to simply move to TS from JS. But please refrain, I'm looking for genuine opinions from people who use or used it in the past.

I genuinely like flowtype, it fits my needs and I've never encountered a tooling issue that has made me throw my laptop out the window. If anything it's been because I didn't understand the tool well enough or I had to improve my coding standards to to make it more sound.

The community feels small but wondering how many of you are there on r/javascript? Please tell me your experience if you are working it now. Love it, hate it and why?
## [11][What's the deal with SvelteKit?](https://www.reddit.com/r/javascript/comments/jofsdq/whats_the_deal_with_sveltekit/)
- url: https://svelte.dev/blog/whats-the-deal-with-sveltekit
---

