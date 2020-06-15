# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][StimulusReflex, I was waiting for you](https://www.reddit.com/r/rails/comments/h92jon/stimulusreflex_i_was_waiting_for_you/)
- url: https://www.reddit.com/r/rails/comments/h92jon/stimulusreflex_i_was_waiting_for_you/
---
Hello to this new tool, without being really aware, I was actually waiting for it. I felt pain about this topic. There were no place were to store the front state in Ruby-on-Rails. Not properly. How to know which accordion is opened or not, how to know which tab is opened or not, etc. The only way I found until then was to use a centralized state (root object with Vue, or Vuex, or redux alone or with React). And store/restore it from localStorage. Without this centralized state, things get messy very quickly for complex scenarii. I wasn't completly convinced by StimulusJS itself because the state is assumed by the DOM, therefore it is very difficult to reason about it. Now with Reflex, it seems you can rely on a front state that is 100% assumed by the backend, removing the need to reason about a frontend state, and all the heavy mandatory JS tooling. Guys, this is a huge promise. Thanks a lot !!
## [3][[Help]Creating authentic seeds and attaching images](https://www.reddit.com/r/rails/comments/h9dvxn/helpcreating_authentic_seeds_and_attaching_images/)
- url: https://www.reddit.com/r/rails/comments/h9dvxn/helpcreating_authentic_seeds_and_attaching_images/
---
    Hi, I'm trying to create authentic seeds in 3 categories for a restaurant app and attaching a default image for each one. On the menu page, I want to just display a 'featured' item, which will be chosen at random. However, I'm not able to create seeds with the correct category id, nor iterate over to display an image. For the partial below, I get an undefined method error. If anyone could point me in the right direction, or is willing to chat over DM, would be greatly appreciated

&lt;div class="row"&gt;
		&lt;% @cats.each do |cat| %&gt;
    &lt;% if cat.products.where(featured: true).first.image.attached? %&gt;
    &lt;%= image_tag(cat.products.where(featured: true).first.image, class:'img-fluid')%&gt;
    &lt;% end %&gt;
    &lt;div class='col-12'&gt;
      &lt;h1 style="text-align:center;"&gt;&lt;%= cat.heading %&gt;&lt;/h1&gt;
      &lt;hr /&gt;
    &lt;/div&gt;
    &lt;% @products.each do |product| %&gt;
    &lt;% if product.category == cat %&gt;
    &lt;div class="col-4"&gt;
      &lt;h5&gt;
        &lt;%= product.name %&gt;
      &lt;/h5&gt;
      &lt;p&gt;
        &lt;%= product.description %&gt;
      &lt;/p&gt;
      &lt;span&gt;&lt;%= product.price %&gt; | &lt;%= "Available for delivery" if product.delivery %&gt;&lt;/span&gt;
      &lt;%= form_with model: @order_item do |f| %&gt;
      &lt;%= f.hidden_field :product_id, value: product.id %&gt;
      &lt;%= f.number_field :quantity, value: 1, min: 1%&gt;
      &lt;%= f.submit "Add to Cart" %&gt;
      &lt;% end %&gt;
    &lt;/div&gt;

    &lt;% end %&gt;
    &lt;% end %&gt;
    &lt;% end %&gt;
  &lt;/div&gt;
## [4][Which is better? Puma or Passenger?](https://www.reddit.com/r/rails/comments/h9em7o/which_is_better_puma_or_passenger/)
- url: https://www.reddit.com/r/rails/comments/h9em7o/which_is_better_puma_or_passenger/
---
Which server is best for handling high volume of light weight requests? I did load testing with both servers, Puma won the contest with very slight difference. The main reason I'm asking to reduce running servers for my production app. As of now 15 servers are running with passenger. Is this possible to reduce servers count to atleast 13 or 14 by switching to Puma from Passenger?
## [5][Best way to association one to "many manies"?](https://www.reddit.com/r/rails/comments/h985c6/best_way_to_association_one_to_many_manies/)
- url: https://www.reddit.com/r/rails/comments/h985c6/best_way_to_association_one_to_many_manies/
---
I'm wondering what's the best way to do the association for the following models. On one hand, we have `Reviewers` and on the other one we have their reviews on different platforms (Tripadvisor, Yelp, etc.). Each type of review would have their own unique characteristics, so we have models like `TripadvisorReview`, `YelpReview`, etc.

Each `Reviewer` only has one type of reviews. To keep all the reviewers in one model, we validates the uniqueness of their username on the `scope:` of their `:platform` ("tripadvisor", "yelp", etc.).

Here's the current model for `Reviewer`:
```
class Reviewer &lt; ApplicationRecord  
  has_many :tripadvisor_reviews, dependent: :destroy  
  has_many :yelp_reviews, dependent: :destroy  
  validates :username, presence: true, uniqueness: { scope: :platform }  
end
```
Here's the model for `TripadvisorReview`:
```
class TripadvisorReview &lt; ApplicationRecord
  belongs_to :reviewer
end
```
(Same for `YelpReview`, etc.)

I considered creating an association like `ReviewerPlatform` but individually this is a `one_to_many` association: each reviewer has many reviews only on one platform. I'm wondering if there would be a more efficient way to do this association, or is this one fitting "the Rails way"?
## [6][Gem for Dynamic, Per-User Banner Messages?](https://www.reddit.com/r/rails/comments/h96o36/gem_for_dynamic_peruser_banner_messages/)
- url: https://www.reddit.com/r/rails/comments/h96o36/gem_for_dynamic_peruser_banner_messages/
---
Just display a banner with a message for a given user or sets of users. Messages may differ per user. I thought there was a gem for this but could not find it. 

One must be able to set messages that will be displayed on next page load. 

Suggestions?
## [7][HTTP2 early hints in Rails / Puma](https://www.reddit.com/r/rails/comments/h8z9kz/http2_early_hints_in_rails_puma/)
- url: https://www.reddit.com/r/rails/comments/h8z9kz/http2_early_hints_in_rails_puma/
---
Hi, I've enabled early hints in puma, but when i enable http2 in nginx i get an error on mobile clients net::ERR\_HTTP2\_PROTOCOL\_ERROR

From what I've found looks like nginx still doesn't support early hints, has anyone gotten it working? Could you share your stack / configuration? 

&amp;#x200B;

Thanks
## [8][Connecting to Heroku Follower Database](https://www.reddit.com/r/rails/comments/h99i6n/connecting_to_heroku_follower_database/)
- url: https://www.reddit.com/r/rails/comments/h99i6n/connecting_to_heroku_follower_database/
---
I'm trying to setup a follower database on Heroku. What's the easiest way to configure my Rails application to connect to the follower database for reads? Do I do it on the model level? Or there might be an easier way to do it. I can't seem to find any resources online.
## [9][Edit and delete (newbie)](https://www.reddit.com/r/rails/comments/h94tr1/edit_and_delete_newbie/)
- url: https://www.reddit.com/r/rails/comments/h94tr1/edit_and_delete_newbie/
---
I am having trouble getting my app to edit and delete an actual object. I am a newbie and this is my first rails project for my bootcamp. I am able to get the links connected, but when you click the link it does not delete anything. I also must mention that I had to nest some routes, so it really confused me in terms of the logic. If anyone can point me to some good resources that would be really helpful!
## [10][Out of curiosity, how many of you are finding that your professional Rails-based environment is moving more toward JavaScript?](https://www.reddit.com/r/rails/comments/h8mdaz/out_of_curiosity_how_many_of_you_are_finding_that/)
- url: https://www.reddit.com/r/rails/comments/h8mdaz/out_of_curiosity_how_many_of_you_are_finding_that/
---
The reason I ask this is because I am - and have been for quite some time - a professional JavaScript developer.

However, I love working with Rails and am positioning myself to move into my first Rails gig within the next 6 months.

And I had a very interesting conversation with a pro Rails developer the other day.  This guy is a senior engineer who has been doing Rails for over 12 years.  Recently, though, he was laid off from a remote job due to his company taking a big economic hit because of Covid-19.

He told me that he immediately started sending out resumes and went through a battery of tests and interviews and what struck his as interesting is that they were FAR more interested in his JavaScript (React mainly) experience than they were his 12 years of Rails development.  In fact, he said that because of my JS experience, I would probably have an easier time finding a Rails job than he did - which sounds ludicrous. 

Are any of you folks starting to witness the shift as well?
## [11][Quick Demo of Reactive UI with Stimulus Reflex](https://www.reddit.com/r/rails/comments/h8j14f/quick_demo_of_reactive_ui_with_stimulus_reflex/)
- url: https://www.reddit.com/r/rails/comments/h8j14f/quick_demo_of_reactive_ui_with_stimulus_reflex/
---
[https://youtu.be/MlnNkcz-oIc](https://youtu.be/MlnNkcz-oIc)
