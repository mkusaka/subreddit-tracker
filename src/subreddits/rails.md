# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [3][Advice on Application Architecture](https://www.reddit.com/r/rails/comments/jh42dn/advice_on_application_architecture/)
- url: https://www.reddit.com/r/rails/comments/jh42dn/advice_on_application_architecture/
---
Hey everyone

I'm building a social-network-like platform with a Rails API and React all on AWS EBS. This isn't a hobby project, so I need to be careful about my design choices and would therefore appreciate some advice.

What's worrying me is how I will introduce realtime functionality. Like most social networks, my app will have a feed, likes, and maybe comments--all of which need to be updated in (near) realtime when something happens in the database. There is also a 1-to-1 chat component.

Because I'm kind of on a ticking clock, I'm thinking about taking the easy way out with the chat and using Firebase for that (chat is not really relevant to my biz logic). From what I'm reading, realtime chat seems to be finicky with Rails and ActionCable. I hypothesize that ActionCable will do the job for all the other realtime stuff. Would you agree?

I've seen alternatives like Pusher and Pubnub and they look great, but they're also quite pricey. Anyone run into a similar situation?
## [4][Moving a method: :delete in another page](https://www.reddit.com/r/rails/comments/jh7teq/moving_a_method_delete_in_another_page/)
- url: https://www.reddit.com/r/rails/comments/jh7teq/moving_a_method_delete_in_another_page/
---
Hi guys,

in the page in view app/view/books/edit.html.erb I have this to delete the page

    &lt;%= link_to "Destroy", book_path(id: @book.slug), method: :delete, data: { confirm: "are you sure?" } %&gt;

I want to move it in a custom popup. So I edited it in this way

    &lt;%= link_to "Destroy", warning_path(:type =&gt; 'delete_book'), remote: true, data: { load_popup: true } %&gt;

and in app/views/warning/index.html.erb I added this

    &lt;% if params[:type] == 'delete_book' %&gt;
    	&lt;%= link_to "Destroy", book_path(id: @book.slug), method: :delete, data: { confirm: "are you sure?" } %&gt;
    &lt;% end %&gt;

But I have this error: `Undefined method 'slug' for nil:NilClass`. How to fix?
## [5][[Help] Rails Server gives error "ERR_CONNECTION_REFUSED" in browser](https://www.reddit.com/r/rails/comments/jgx048/help_rails_server_gives_error_err_connection/)
- url: https://www.reddit.com/r/rails/comments/jgx048/help_rails_server_gives_error_err_connection/
---
EDIT: Solved changing port 8000 seemed to do the trick thanks for your help

I'm new to rails development and only have a couple of simple projects. All of my projects are stored in a wsl2 Ubuntu 18.04LTS vm. I have always been able to run \`rails s\` and view my project on localhost:3000. Even last night I was able to view 2 separate projects this way. Now, today when I start rails server everything looks like it works as expected within the terminal (starting on [127.0.0.1:3000](https://127.0.0.1:3000)) but when I try to load in my browser I get \`ERR\_CONNECTION\_REFUSED\`. Other than the message in the browser the server in the terminal looks just like normal and shows no error messages. Any help would be greatly appreciated.
## [6][Security : Is sending back an invalid password is a good or bad practice ?](https://www.reddit.com/r/rails/comments/jgop0h/security_is_sending_back_an_invalid_password_is_a/)
- url: https://www.reddit.com/r/rails/comments/jgop0h/security_is_sending_back_an_invalid_password_is_a/
---
If user wants to register to my app and enter a too short password, is it bad practice to fill again this password in the rendered template ? (Or for a JSON request, show this "bad" password in the JSON response).
## [7][Rails 6 vs JQuery](https://www.reddit.com/r/rails/comments/jgrfwq/rails_6_vs_jquery/)
- url: https://www.reddit.com/r/rails/comments/jgrfwq/rails_6_vs_jquery/
---
I am having THE issue, but I am cannot find the fix. My problem is intermittent and only in production. JQuery works like a champ in my development enviroment, but when I push out my app to Production, I am constantly teetering back and forth between no Jquery and it working mostly. I have dug through countless tutorials, reddit posts, and stackoverflow questions looking for my answer, I am hoping that someone can see where I went astray and help me never have to worry about this again. 

&amp;#x200B;

environment.js

    const { environment } = require('@rails/webpacker')
    const webpack = require('webpack')
    environment.plugins.append('Provide', new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      Popper: ['popper.js', 'default'],
     })
    )
    module.exports = environment

application.js

    import "bootstrap"
    require("@rails/ujs").start()
    require("turbolinks").start()
    require("channels")
    require("jquery")
    
    
    
    import flatpickr from "flatpickr"
    require("flatpickr/dist/flatpickr.css")
    
    document.addEventListener("turbolinks:load", () =&gt; {
    	flatpickr("[data-behavior='flatpickr']", {
    		
    	})
    })
    
    import 'controllers'
     
    var componentRequireContext = require.context("components", true);
    var ReactRailsUJS = require("react_ujs");
    ReactRailsUJS.useContext(componentRequireContext);
    
    $(function () {
      $('[data-toggle="tooltip"]').tooltip()
    })

Application.html.erb

    &lt;head&gt;
        &lt;%= csrf_meta_tags %&gt;
        &lt;%= csp_meta_tag %&gt;
        &lt;%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;
        &lt;%= stylesheet_pack_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;
        &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %&gt;
        &lt;script async src="location"&gt;&lt;/script&gt;
        &lt;%= javascript_include_tag 'analytics', async: true %&gt;
        &lt;%= favicon_link_tag asset_path('nsfav16.ico') %&gt;
      &lt;/head&gt;

I am happy to post any other info someone may need to help me figure this out.
## [8][5 ways to fix the latest-comment n+1 problem](https://www.reddit.com/r/rails/comments/jgam27/5_ways_to_fix_the_latestcomment_n1_problem/)
- url: https://www.reddit.com/r/rails/comments/jgam27/5_ways_to_fix_the_latestcomment_n1_problem/
---
Hi, 

I wrote a little guide with 5 ways to fix a kind "n+1 queries" problem, that I called the "latest-comment".

I called it that way because one example/instance of the problem, is when you want a list of posts with the last comment, but it can also be extrapolated to "the last review in a list of products", or "the cheapest price", etc..

[https://bhserna.com/5-ways-to-fix-the-latest-comment-n-1-problem.html](https://bhserna.com/5-ways-to-fix-the-latest-comment-n-1-problem.html)

I hope it can be useful for someone :)
## [9][What serializers do you use in rails API?](https://www.reddit.com/r/rails/comments/jg2he7/what_serializers_do_you_use_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/jg2he7/what_serializers_do_you_use_in_rails_api/
---
I was looking at [https://github.com/procore/blueprinter](https://github.com/procore/blueprinter) over ActiveModel Serializer but I'm not sure of any best practices when implementing this in a large project. I like the idea of its view and having multiple views in a single file to use for the required API.   


I have two questions, 

1. Do you have any ideas from your usage of blueprinter for JSON serialization?
2. Also, do you know of any public repo on GitHub that uses this so I can skim its code to understand how they use it?
## [10][Introducing Stimulus components with a first class support for Rails](https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/)
- url: https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/
---
Stimulus deserves to have a big and qualitative ecosystem with plug'n'play controllers like in other modern JS frameworks. 

More info here 👉 [https://guillaumebriday.fr/introducing-stimulus-components](https://guillaumebriday.fr/introducing-stimulus-components)

All the available controllers are here 👉 [https://github.com/stimulus-components](https://github.com/stimulus-components)

Feel free to open PRs and issues 🥳
## [11][No method error](https://www.reddit.com/r/rails/comments/jggb6t/no_method_error/)
- url: https://www.reddit.com/r/rails/comments/jggb6t/no_method_error/
---
I am following a tutorial that consumes a google translate API and allows the app to translate from one language to another. Everything is going fine until I try to run the code in the controller. Then I get "NoMethodError in TranslationsController#index"

        undefined method `[]' for nil:NilClass
        Extracted source (around line #37):


        languages = api_request('language/translate/v2/languages')

        keys = languages['data']['languages'].map { |l| l['language'].upcase }

        I18nData
        .languages


This is my controller code:

    class TranslationsController &lt; ApplicationController
    before_action :set_defaults

    def index
    end

    def translate
    end

    private

    def set_defaults
      @languages = fetch_languages
    end

    def api_request(path, method: :get, body: nil)
      params = {
        headers: {
          'x-rapidapi-key': 'YOUR_API_KEY',
          'content-type': 'application/x-www-form-urlencoded'
        }
      }

     params[:body] = body if body

      response = Excon.send(method,
        "https://google-translate1.p.rapidapi.com/#{path}",
        params
      )

    JSON.parse(response.body)
    end

    def fetch_languages
      languages = api_request('language/translate/v2/languages')

     keys = languages['data']['languages'].map { |l| l['language'].upcase }

      I18nData
        .languages
        .slice(*keys)
        .each_with_object([]) do |(iso, name), memo|
          memo &lt;&lt; [name, iso]
        end
      end
    end





I don't know what I'm doing wrong.
## [12][Rspec stub current user on a request spec for a rails engine with no access to Devise](https://www.reddit.com/r/rails/comments/jg3j7y/rspec_stub_current_user_on_a_request_spec_for_a/)
- url: https://www.reddit.com/r/rails/comments/jg3j7y/rspec_stub_current_user_on_a_request_spec_for_a/
---
I am working on a rails engine. I have no access to Devise or its helper methods since those are in the host application. We also don't have a user model so we have to mock the model

&amp;#x200B;

I have a controller defined as

&amp;#x200B;

\`\`\`

    require_dependency "sample/application_controller"
    
    module Sample
     class ExampleController &lt; ApplicationController
      def list_estates
        if current_user.has_access_to?(:communities)
           u/estates = Estate.all
        end
      end
     end
    end

\`\`\`

&amp;#x200B;

I would like to test that the method actually performs xyz if current\_user has access to communities.

&amp;#x200B;

&amp;#x200B;

\`\`\`

    require 'rails_helper'
    module PremierSaAgreementSigning
    class User
    
    	attr\_reader :email, :password
    
    	def initialize(email, password)
    
    		u/email = email
    
    		u/password = password
    
    	end
    
    end

&amp;#x200B;

    describe ExampleController, type: :request do
      context 'user is signed in' do
        let(:user) { [User.new](https://User.new)('test@example.com', 'password') }
        let(:estate) { create(:estate) }
    
         before do
            ActionView::Base.any_instance.stub(:current_user) { user }
         end
    
    
         it 'returns a list of estates' do
            get "/sample/example/list_estates"
            expect(assigns(:estates)).to include(estate)
         end
       end
    end

\`\`\`

&amp;#x200B;

Since I have no access to devise or a user model.

How do I get \`current\_user\` instance on the \`ExampleController\` to be equal to the user I have created here. Then how do I get the now new current\_user to mock the \`has\_access\_to?\` method which is also not defined in this engine.

&amp;#x200B;

My method above returns

    `ActionView::Base does not implement #current_user`

&amp;#x200B;

&amp;#x200B;

Please note that I do not need to create or expose any session/login logic here as that is handled on the host application and not required on the engine.

    I just need to mock the `current_user` object and run another fake method..`has_access_to?` on this object

How can I do this. I am a beginner and this has been quite challenging
