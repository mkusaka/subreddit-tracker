# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Difference between webpacker:install:package-name vs yarn add package-name?](https://www.reddit.com/r/rails/comments/fuq6kn/difference_between_webpackerinstallpackagename_vs/)
- url: https://www.reddit.com/r/rails/comments/fuq6kn/difference_between_webpackerinstallpackagename_vs/
---
Hi. I have been using 'yarn add' to install js packages just fine since Rails 6. Today I read a tutorial for Stimulus, and the tutorial uses webpacker:install:stimulus to install the package. I am wondering is there any difference between the two, and are they interchangeable? Thanks in advance.
## [4][GOTO 2019 • You Really Don't Need All that JavaScript, I Promise • Stuart Langridge](https://www.reddit.com/r/rails/comments/fuqrxy/goto_2019_you_really_dont_need_all_that/)
- url: https://www.reddit.com/r/rails/comments/fuqrxy/goto_2019_you_really_dont_need_all_that/
---
A nice talk about Javascript [https://www.youtube.com/watch?v=rxlJRydqmk8](https://www.youtube.com/watch?v=rxlJRydqmk8)
## [5][How do you use respon_to JavaScript haml](https://www.reddit.com/r/rails/comments/futehg/how_do_you_use_respon_to_javascript_haml/)
- url: https://www.reddit.com/r/rails/comments/futehg/how_do_you_use_respon_to_javascript_haml/
---
Anyone here using haml ? 
How do you guys using haml for JavaScript??

I always failed to respond to JS with haml ini controller
Itu always unknownFormat

I did : 
render "admin/error" in controller
Or
render template: "error.js.haml"
Or
render template: "admin/error.js.haml" 
Or
Render template: "admin/error.js.haml" l, layout: false

Many things I do but always failed 
And unknownFormat for respond_do

"Admin/error.js.haml" doest read any value even I do console.log also 

Can anyone give correct example ?
## [6][How would you "prepare" for a mobile branch?](https://www.reddit.com/r/rails/comments/fuoudm/how_would_you_prepare_for_a_mobile_branch/)
- url: https://www.reddit.com/r/rails/comments/fuoudm/how_would_you_prepare_for_a_mobile_branch/
---
I am not yet there, but I see it coming. I have a Rails 5 app in which basically users send promos that are published in many screens. So far, everything has been through browser, but I have seen situations where a mobile app would be a best fit (pull notification, instant dashboard without getting to pages..and so on. The thing is my app is *old school* \- I use default cookies , I render everything with html.erb, and so on- I really wouldn't know where to start or what things should I consider.

Can you please share any insight about that?
## [7][GOTO 2019 • You Really Don't Need All that JavaScript, I Promise • Stuart Langridge](https://www.reddit.com/r/rails/comments/fuqng6/goto_2019_you_really_dont_need_all_that/)
- url: https://www.reddit.com/r/rails/comments/fuqng6/goto_2019_you_really_dont_need_all_that/
---
A nice talk about Javascript [https://www.youtube.com/watch?v=rxlJRydqmk8](https://www.youtube.com/watch?v=rxlJRydqmk8)
## [8][Vue.js on some rails app pages. Stimulus.js for the rest. How would you approach this?](https://www.reddit.com/r/rails/comments/fud9xo/vuejs_on_some_rails_app_pages_stimulusjs_for_the/)
- url: https://www.reddit.com/r/rails/comments/fud9xo/vuejs_on_some_rails_app_pages_stimulusjs_for_the/
---
I'd like to be able to use vue-cli if possible, and yet stay close to the default webpack config. 

How would one approach this? Any examples?

Edit: I guess I am wondering what the pros and cons are of keeping the vue SPAs in a separate codebase run by vue-cli? Feels cleaner, especially given that only a couple of pages will be SPAs and the rest is vanilla rails with stimulus for added interactivity.
## [9][Integrating Flaticon fonts with Webpack/Rails 6](https://www.reddit.com/r/rails/comments/fuilla/integrating_flaticon_fonts_with_webpackrails_6/)
- url: https://www.reddit.com/r/rails/comments/fuilla/integrating_flaticon_fonts_with_webpackrails_6/
---
Hello all,

I am trying to use Flaticon fonts using Webpacker. I am getting the following error from all of my fonts when Webpacker compiles:
```
Module not found: Error: Can't resolve '../fonts/Flaticon.ttf' in '/Users/user/Documents/GitHub/my_app/app/javascript/stylesheets'
 @ ./app/javascript/stylesheets/application.scss (./node_modules/css-loader/dist/cjs.js??ref--7-1!./node_modules/postcss-loader/src??ref--7-2!./node_modules/sass-loader/dist/cjs.js??ref--7-3!./app/javascript/stylesheets/application.scss) 7:36-68
 @ ./app/javascript/stylesheets/application.scss
 @ ./app/javascript/packs/application.js

```

**Under app/assets/fonts:**

```
* Flaticon.eot
* Flaticon.svg
* Flaticon.tff
* Flaticon.woff
* Flaticon.woff2
```

**Under javascript/stylesheets/_flaticon.scss:**
```
    /*
    Flaticon icon font: Flaticon
    Creation date: 03/04/2020 21:12
    */

    @font-face {
  font-family: "Flaticon";
  src: url("../fonts/Flaticon.eot");
  src: url("../fonts/Flaticon.eot?#iefix") format("embedded-opentype"),
  url("../fonts/Flaticon.woff2") format("woff2"),
  url("../fonts/Flaticon.woff") format("woff"),
  url("../fonts/Flaticon.ttf") format("truetype"),
  url("../fonts/Flaticon.svg#Flaticon") format("svg");
  font-weight: normal;
  font-style: normal;
}

@media screen and (-webkit-min-device-pixel-ratio:0) {
  @font-face {
    font-family: "Flaticon";
    src: url("../fonts/Flaticon.svg#Flaticon") format("svg");
  }
}

    .fi:before{
        display: inline-block;
        font-family: "Flaticon";
        font-style: normal;
        font-weight: normal;
        font-variant: normal;
        line-height: 1;
        text-decoration: inherit;
        text-rendering: optimizeLegibility;
        text-transform: none;
        -moz-osx-font-smoothing: grayscale;
        -webkit-font-smoothing: antialiased;
        font-smoothing: antialiased;
    }

    .flaticon-policeman:before { content: "\f100"; }
    
    $font-Flaticon-policeman: "\f100";
```

**Under config/webpacker.yml:**
```
# Note: You must restart bin/webpack-dev-server for changes to take effect

default: &amp;default
  source_path: app/javascript
  source_entry_path: packs
  public_root_path: public
  public_output_path: packs
  cache_path: tmp/cache/webpacker
  check_yarn_integrity: false
  webpack_compile_output: true

  # Additional paths webpack should lookup modules
  # ['app/assets', 'engine/foo/app/assets']
  resolved_paths: ['app/assets']

  # Reload manifest.json on all requests so we reload latest compiled packs
  cache_manifest: false

  # Extract and emit a css file
  extract_css: false

  static_assets_extensions:
    - .jpg
    - .jpeg
    - .png
    - .gif
    - .tiff
    - .ico
    - .svg
    - .eot
    - .otf
    - .ttf
    - .woff
    - .woff2

  extensions:
    - .mjs
    - .js
    - .sass
    - .scss
    - .css
    - .module.sass
    - .module.scss
    - .module.css
    - .png
    - .svg
    - .gif
    - .jpeg
    - .jpg

development:
  &lt;&lt;: *default
  compile: true

  # Verifies that correct packages and versions are installed by inspecting package.json, yarn.lock, and node_modules
  check_yarn_integrity: true

  # Reference: https://webpack.js.org/configuration/dev-server/
  dev_server:
    https: false
    host: localhost
    port: 3035
    public: localhost:3035
    hmr: false
    # Inline should be set to true if using HMR
    inline: true
    overlay: true
    compress: true
    disable_host_check: true
    use_local_ip: false
    quiet: false
    pretty: false
    headers:
      'Access-Control-Allow-Origin': '*'
    watch_options:
      ignored: '**/node_modules/**'


test:
  &lt;&lt;: *default
  compile: true

  # Compile test packs to a separate directory
  public_output_path: packs-test

production:
  &lt;&lt;: *default

  # Production depends on precompilation of packs prior to booting for performance.
  compile: false

  # Extract and emit a css file
  extract_css: true

  # Cache manifest.json for performance
  cache_manifest: true

```

I have tried to move the fonts folder to vendor and under javascript, but to no avail. Any help would be much appreciated!
## [10][Rspec tests failing when using Rails.cache, but pass if I do a binding.pry](https://www.reddit.com/r/rails/comments/fugmzr/rspec_tests_failing_when_using_railscache_but/)
- url: https://www.reddit.com/r/rails/comments/fugmzr/rspec_tests_failing_when_using_railscache_but/
---
I'm having a weird issue where I'm testing a controller that has a procedure that uses caching. The test is failing, but if I do a binding.pry inside the method that does the caching, the test passes.

example of the method containing the caching and the binding.pry:

    def method_example:   
      data = Rails.cache.fetch(cache_key) do                                               
    ProcedureService.new(params).generate
      end
    
      binding.pry    
      data 
    end

Example of the test:

    it "reverts record amount" do     
      expect(record.amount).to eq((original_amount + other_amount).to_d) 
    end

The caching is done via redis\_store.

When done in the development environment, it works fine. What I don't understand is why it is failing but passing when adding a stopper? It seems it could be something about the time it takes to fetch the cache
## [11][Tutorial for integrating FlatIcon with Rails6/Webpack?](https://www.reddit.com/r/rails/comments/fuh2z3/tutorial_for_integrating_flaticon_with/)
- url: https://www.reddit.com/r/rails/comments/fuh2z3/tutorial_for_integrating_flaticon_with/
---
I was wondering if anyone has done this before. I want to download the font-files from FlatIcon, since their icons are much better than font-awesome. I tried to do this by putting the font files in ‘app/assets/fonts’ and precompiling, but to no avail.
## [12][Looking for multi-tenant website builder &amp; RSVP codebase ideas](https://www.reddit.com/r/rails/comments/fudupr/looking_for_multitenant_website_builder_rsvp/)
- url: https://www.reddit.com/r/rails/comments/fudupr/looking_for_multitenant_website_builder_rsvp/
---
hi, I'm a business owner (not dev) new to this r/. I currently have a business that runs on wp multisite, and I'm researching if there is a similar codebase for a multi-tenant website builder that could be used as a development base for moving my operations into ruby. 

I believe ruby would be a superior solution, but I don't know if it would be economical to develop. My needs are a not exotic, a multi-tenant website builder and rsvp system, so I thought there might be something out there to build on but hunting on GitHub hasn't worked so far.

any thoughts, links to repos, ideas most welcome!
