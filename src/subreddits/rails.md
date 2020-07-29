# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [3][Add Authenticity Token to JavaScript AJAX request](https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/)
- url: https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/
---
I need to send it manually and cannot submit it through a form. Is there a specific request param that Rails looks for in order to verify the Authenticity token? 

Perhaps, I’m going at it wrong. The other idea I had was to make it an API-only route/controller and, maybe then I wouldn’t have to use the authenticity token and just use an API auth token? Sorry, I am new to Rails :D
## [4][Why redirect_to proc?](https://www.reddit.com/r/rails/comments/hzk7uy/why_redirect_to_proc/)
- url: https://www.reddit.com/r/rails/comments/hzk7uy/why_redirect_to_proc/
---
What is the use case of redirecting to a proc?

As in 

    redirect_to proc { edit_post_url(@post) }

vs 

    redirect_to edit_post_url(@post)

?

What are some uses of this?

Edit: from https://api.rubyonrails.org/classes/ActionController/Redirecting.html
## [5][Turning a list item into an edit form](https://www.reddit.com/r/rails/comments/hzpgfb/turning_a_list_item_into_an_edit_form/)
- url: https://www.reddit.com/r/rails/comments/hzpgfb/turning_a_list_item_into_an_edit_form/
---
I have a list of employees and an edit button by each one. Instead of the edit button redirecting to a form, I want it to turn the employee's info into text boxes that are auto-filled with the current value of that attribute. The edit button would turn into a submit button. The user would make the changes needed, and after submitting, the employee's new attributes would be listed. 

Here's my understanding of how this would work. 

1. The user clicks on the edit button, which triggers a function that removes the current contents of the employee's div and replaces it with the form built as a document fragment by another function. 
2. The new submit button is a link to the update action of the employee controller with a data-remote="true" attribute. 
3. The update action sends back the new employee's instance, which somehow triggers a function that will delete the form and replace it with a new document fragment containing the updated employee information. 

My biggest concern is pulling off the Ajax call properly. I'm trying to use ujs and vanilla JavaScript, but the resources I've found for doing Ajax calls in rails use coffee script and/or jQuery. I've found a few newer articles on medium, but they don't go in enough depth for me to be able to adapt they're examples for my propose.

Do you guys know of any resources that may help me fill in the gaps?

Thanks!
## [6][Has many through association not working (ActiveRecord::HasManyThroughAssociationNotFoundError)](https://www.reddit.com/r/rails/comments/hzp1k0/has_many_through_association_not_working/)
- url: https://www.reddit.com/r/rails/comments/hzp1k0/has_many_through_association_not_working/
---
Hi I'm getting the following error when I test out either of the following associations:

    Venue.last.users
    ActiveRecord::HasManyThroughAssociationNotFoundError: Could not find the association :venue_users in model Venue

&amp;#x200B;

    User.last.venues
    ActiveRecord::HasManyThroughAssociationNotFoundError: Could not find the association :venue_users in model User

  
Here's what my models look like:

    class Venue &lt; ApplicationRecord
      has_many :users, through: :venue_users
    end

    class User &lt; ApplicationRecord
      has_many :venues, through: :venue_users
    end

  
This is what the join table looks like in my schema:

    create_table "venue_users", force: :cascade do |t|
        t.bigint "user_id", null: false
        t.bigint "venue_id", null: false
        t.integer "user_type", default: 0
        t.index ["user_id"], name: "index_venue_users_on_user_id"
        t.index ["venue_id"], name: "index_venue_users_on_venue_id"
      end

  
Can someone help me where I've gone wrong?  


Thanks.
## [7][Rails way of trimming down controller action code size?](https://www.reddit.com/r/rails/comments/hzmath/rails_way_of_trimming_down_controller_action_code/)
- url: https://www.reddit.com/r/rails/comments/hzmath/rails_way_of_trimming_down_controller_action_code/
---
I'd like to move some business logic out of the controller actions.
## [8][Is skip_before_action :verify_authenticity_token a bad idea?](https://www.reddit.com/r/rails/comments/hzm7ic/is_skip_before_action_verify_authenticity_token_a/)
- url: https://www.reddit.com/r/rails/comments/hzm7ic/is_skip_before_action_verify_authenticity_token_a/
---
Could it, say, allow any unauthenticated user to perform a request on that resource even if I use before\_action :authenticate\_user!  


What are the side effects of using skip\_before\_action :verify\_authenticity\_token?
## [9][Can I add an IF condition in a scss file?](https://www.reddit.com/r/rails/comments/hzmyve/can_i_add_an_if_condition_in_a_scss_file/)
- url: https://www.reddit.com/r/rails/comments/hzmyve/can_i_add_an_if_condition_in_a_scss_file/
---
I have to create a different style of a website for the Arabic users (text-align: right and change a lot of float).

Can I add an if condition in a scss file?

Or should I create two different scss files?

Or should I create a new scss file with !important?
## [10][i want to show employee working on the project.](https://www.reddit.com/r/rails/comments/hzz8u1/i_want_to_show_employee_working_on_the_project/)
- url: https://www.reddit.com/r/rails/comments/hzz8u1/i_want_to_show_employee_working_on_the_project/
---
i want to show employee working on the project.But i cannot find the way who to impliment that.

    &lt;li class="dropdown"&gt;
 &lt;%= link_to "path", :class =&gt; "dropdown-toggle", :data =&gt; {:toggle =&gt; "dropdown"} do %&gt; &lt;i class="fa fa-file-text"&gt;&lt;/i&gt;   Contents &lt;span class="caret"&gt;&lt;/span&gt; &lt;% end %&gt;
          &lt;ul class="dropdown-menu"&gt;
 &lt;% Project.all.each do |c| %&gt;
              &lt;li class="dropdown-item list"&gt;&lt;%= link_to c.title, summery_admin_timelogs_path(:id =&gt; c.id) %&gt;&lt;/li&gt;
 &lt;% end %&gt;
          &lt;/ul&gt;
 &lt;/li&gt; 
    def summery
project = Project.find(params[:id])
 employees = u/project.employees
end
    
    
    resources :timelogs do
 collection do
 get :summery
 end
 end
## [11][convert rails api to rails app](https://www.reddit.com/r/rails/comments/hzemhx/convert_rails_api_to_rails_app/)
- url: https://www.reddit.com/r/rails/comments/hzemhx/convert_rails_api_to_rails_app/
---
I have a Rails Api that has a Vue front end (in it's own directory called client). Now we're wanting to add admin screen but the client has indicated they'd prefer those just be rails pages like RailsAdmin. Is there a good guide for converting a Rails Api to a full Rails App while not blowing away the things you've already done? Most guides seem to indicate running rails-new again and carefully selecting things not to overwrite but that doesn't really tell you if you're missing anything once it's done.
## [12][Is Strong Params blocking headers ?](https://www.reddit.com/r/rails/comments/hzf4q2/is_strong_params_blocking_headers/)
- url: https://www.reddit.com/r/rails/comments/hzf4q2/is_strong_params_blocking_headers/
---
Hi everyone,

I m creating a rails api for a react native project, I use devise\_token\_auth gem for my user authentication.

 I try to pass params to my server but i always got this

    Unpermitted parameter: :format
    ...
    Completed 400 Bad Request in 1610ms (Views: 0.4ms | ActiveRecord: 102.4ms | Allocations: 238027)

i think the error comes from my post controller

    class Api::V0::PostsController &lt; ApplicationController
        include DeviseTokenAuth::Concerns::SetUserByToken
        before_action :authenticate_user!
        before_action :find_post, only: [:show, :update, :destroy]
    
        #GET/post
        def index
            @posts = Post.all
            render :json =&gt; @posts ,  status: 200
        end
    
        #POST/post
        def create
            @post = Post.new(post_params)
            @post.user
            if @post.save
                render json: @post , status: 201
            else
                render error: {error: 'Unable to create post.'}, status: 400
            end
        end
        ...
        private
        # Use callbacks to share common setup or constraints between actions.
        def find_post
            @post = Post.find(params[:id])
        end
      
        # Never trust parameters from the scary internet, only allow the white list through.
        def post_params
            params.permit(:title, :content, :created_by, :entry, :category_id, :rdv, :tag1, :tag2, :tag3, :user_id)
        end

i need to receive token Data from header to be identify as a user but if i do that i got the error. Here is my fetch from my serveur:

    export function onCreate(data){
        return dispatch =&gt; {
            dispatch(fetchPostsPending());
            return fetch('http://localhost/api/v0/posts',{
              method:'POST',
              headers:{
                "access-token": data.accessToken,
                "token-type":   data.tokenType,
                "client":       data.client,
                "expiry":       data.expiry,
                "uid":          data.uid
                },
                body: JSON.stringify({
                    "title":data.title,
                    "content":data.content,
                    "created_by":data.created_by,
                    "entry":data.entry,
                    "category_id":data.category_id,
                    "rdv":data.rdv,
                    "tag1":data.tag1,
                    "tag2":data.tag2,
                    "tag3":data.tag3,
                    "user_id":data.user_id
                  })
            })

my fetch totally worked with my index def because i don't send a body .

thats why i think the problem come from my def post\_params.

do i need to do something like that:

    def post_params
            params.require(:header).permit(:access-token, :token-type, :client, :expiry ,:uid)
            params.permit(:title, :content, :created_by, :entry, :category_id, :rdv, :tag1, :tag2, :tag3, :user_id)
        end
    

if i do need to do that how am I suppose to wright thing like a minus sign(-) in my params require ?
