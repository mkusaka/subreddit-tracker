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
## [3][5 ways to fix the latest-comment n+1 problem](https://www.reddit.com/r/rails/comments/jgam27/5_ways_to_fix_the_latestcomment_n1_problem/)
- url: https://www.reddit.com/r/rails/comments/jgam27/5_ways_to_fix_the_latestcomment_n1_problem/
---
Hi, 

I wrote a little guide with 5 ways to fix a kind "n+1 queries" problem, that I called the "latest-comment".

I called it that way because one example/instance of the problem, is when you want a list of posts with the last comment, but it can also be extrapolated to "the last review in a list of products", or "the cheapest price", etc..

[https://bhserna.com/5-ways-to-fix-the-latest-comment-n-1-problem.html](https://bhserna.com/5-ways-to-fix-the-latest-comment-n-1-problem.html)

I hope it can be useful for someone :)
## [4][Introducing Stimulus components with a first class support for Rails](https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/)
- url: https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/
---
Stimulus deserves to have a big and qualitative ecosystem with plug'n'play controllers like in other modern JS frameworks. 

More info here 👉 [https://guillaumebriday.fr/introducing-stimulus-components](https://guillaumebriday.fr/introducing-stimulus-components)

All the available controllers are here 👉 [https://github.com/stimulus-components](https://github.com/stimulus-components)

Feel free to open PRs and issues 🥳
## [5][What serializers do you use in rails API?](https://www.reddit.com/r/rails/comments/jg2he7/what_serializers_do_you_use_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/jg2he7/what_serializers_do_you_use_in_rails_api/
---
I was looking at [https://github.com/procore/blueprinter](https://github.com/procore/blueprinter) over ActiveModel Serializer but I'm not sure of any best practices when implementing this in a large project. I like the idea of its view and having multiple views in a single file to use for the required API.   


I have two questions, 

1. Do you have any ideas from your usage of blueprinter for JSON serialization?
2. Also, do you know of any public repo on GitHub that uses this so I can skim its code to understand how they use it?
## [6][No method error](https://www.reddit.com/r/rails/comments/jggb6t/no_method_error/)
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
## [7][How do you validate your model ? Simple Active Record validation, or dry-rb, or equivalent ?](https://www.reddit.com/r/rails/comments/jg5gx0/how_do_you_validate_your_model_simple_active/)
- url: https://www.reddit.com/r/rails/comments/jg5gx0/how_do_you_validate_your_model_simple_active/
---
For example if you use dry-rb, do you still use Active Record validations ? I try to find the way to validates models elegantly enough, without over-engineer everything. Thanks for your thoughts !
## [8][Rspec stub current user on a request spec for a rails engine with no access to Devise](https://www.reddit.com/r/rails/comments/jg3j7y/rspec_stub_current_user_on_a_request_spec_for_a/)
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
## [9][Anyone using Svelte,Elm or anything more obscure?How's your experience?](https://www.reddit.com/r/rails/comments/jfwvu2/anyone_using_svelteelm_or_anything_more/)
- url: https://www.reddit.com/r/rails/comments/jfwvu2/anyone_using_svelteelm_or_anything_more/
---

## [10][Active Admin](https://www.reddit.com/r/rails/comments/jg4coy/active_admin/)
- url: https://www.reddit.com/r/rails/comments/jg4coy/active_admin/
---
Is Active Admin a free-to-use tool? I barely came across it and since I'm working on an application for a company it looks like it would help a lot
## [11][Ruby on Rails: templates and generators in 2020](https://www.reddit.com/r/rails/comments/jfmas5/ruby_on_rails_templates_and_generators_in_2020/)
- url: https://www.reddit.com/r/rails/comments/jfmas5/ruby_on_rails_templates_and_generators_in_2020/
---
2020 is a being a very rich year for Rails **boilerplates and generators**. I've written a short article describing the most prominent ones: [https://blog.corsego.com/2020/10/ruby-on-rails-templates-and-generators.html](https://blog.corsego.com/2020/10/ruby-on-rails-templates-and-generators.html)

Hope you find it useful :)
## [12][Do you use Spree to sell digital products?](https://www.reddit.com/r/rails/comments/jg39nc/do_you_use_spree_to_sell_digital_products/)
- url: https://www.reddit.com/r/rails/comments/jg39nc/do_you_use_spree_to_sell_digital_products/
---
Is anyone doing this? I'm building a site where multiple vendors will be able to sell digital products. Seems like Spree is more for physical products. Wondering if I should try to customize it or use something else entirely. 

Thanks!
