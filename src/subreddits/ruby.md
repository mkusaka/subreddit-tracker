# ruby
## [1][One-liner introduction - Ruby one-liners cookbook](https://www.reddit.com/r/ruby/comments/j34eyl/oneliner_introduction_ruby_oneliners_cookbook/)
- url: https://learnbyexample.github.io/learn_ruby_oneliners/one-liner-introduction.html
---

## [2][Ruby 3.0 Freezes Range Objects](https://www.reddit.com/r/ruby/comments/j332ai/ruby_30_freezes_range_objects/)
- url: https://scriptday.com/blog/2020/10/01/ruby-3-0-freezes-range-objects
---

## [3][Gathering ideas: Ruby is the best programming language for beginners](https://www.reddit.com/r/ruby/comments/j2tu7k/gathering_ideas_ruby_is_the_best_programming/)
- url: https://www.reddit.com/r/ruby/comments/j2tu7k/gathering_ideas_ruby_is_the_best_programming/
---
I had an epiphany: Ruby is much better than Python for someone learning to program. :-) I realize those are fighting words in many corners of the Internet. Probably also in academia.

I'm putting together my thoughts for blog posts. Does anyone here have any points or anecdotes about learning Ruby as a first language, or teaching someone?

I had my realization [while coping with Python's cruft and arbitrary design decisions](https://www.reddit.com/r/learnpython/comments/ispgri/any_way_to_make_textsplitjoin_capitalize_possible/). It had been a while since I'd done much Python, and I was surprised at how bad the UX is. Python is held up as an amazing teaching language, yet it presents an uneven smorgasbord of global functions (which are misleadingly called "builtins"), classes, mutating "functions" and intentionally-difficult-to-use fp features.
## [4][What surprised us in PostgreSQL-schema based multitenancy](https://www.reddit.com/r/ruby/comments/j365oy/what_surprised_us_in_postgresqlschema_based/)
- url: https://blog.arkency.com/what-surprised-us-in-postgres-schema-multitenancy/
---

## [5][Procedure - new rails design pattern](https://www.reddit.com/r/ruby/comments/j36wjs/procedure_new_rails_design_pattern/)
- url: https://longliveruby.com/articles/rails-procedure-design-pattern
---

## [6][Arriving at Station: The Evolution of Our API Documentation Platform [Built in Ruby]](https://www.reddit.com/r/ruby/comments/j2wdo9/arriving_at_station_the_evolution_of_our_api/)
- url: https://www.nexmo.com/blog/2020/09/02/arriving-at-station-the-evolution-of-our-api-documentation-platform?utm_source=reddit&amp;utm_medium=organic&amp;utm_campaign=social_media
---

## [7][How to use rack-attack with fail2ban, with a database query?](https://www.reddit.com/r/ruby/comments/j2twhr/how_to_use_rackattack_with_fail2ban_with_a/)
- url: https://www.reddit.com/r/ruby/comments/j2twhr/how_to_use_rackattack_with_fail2ban_with_a/
---
Hello there! I need to ban malicious http requests made to \`/api/check\_user\`, only if the email in the parameters is not present in the database.Given that:- We wanted to block that IP in the first line cause we already saw many malicious attempts  in the logs,- We had to throttle the requests  to the URLs matching \`/oauth|check\_user|recovery/\`- And on top of that, we wanted to ban the IPs making too many http requests to \`/api/check\_user\` when the email was not in the db.

The 410 rule was added to avoid some noise cause we already have some 403s in the system.

    Rack::Attack.blocklist_ip("X.X.X.X")
    
    # Throttle to same email 6 reqs/minute
    Rack::Attack.throttle("limit logins per email", limit: 6, period: 60) do |request|
      if /oauth|check_user|recovery/.match?(request.path)
        request.params["email"].to_s.downcase.gsub(/\s+/, "")
      end
    end
    
    Rack::Attack.blocklisted_response = lambda do |request|
      [410, {}, ["Gone\n"]]
    end
    
    # This is the important bit
    Rack::Attack.blocklist("fail2ban pentesters") do |request|
      Rack::Attack::Fail2Ban.filter("pentesters-#{request.ip}", maxretry: 3, findtime: 10.minutes, bantime: 6.hours) do
        /check_user/.match?(request.path) &amp;&amp;
          !PanelUser.exists?(email: CGI.unescape(request.params["email"].to_s.downcase.gsub(/\s+/, "")))
      end
    end

but it looks like that the last block is blocking also requests to \`/admin/login\` (the backoffice URL). Can you help me to find the error? :)  


I know that is not very performant to keep the query there, but that's not an issue at the moment.

...And since we're here: is there a way to clean the rules and redefine them without redeploying every time we wanted to try another config?

Thanks!

F
## [8][Ruby 3.0 adds `Hash#except` and `ENV.except`](https://www.reddit.com/r/ruby/comments/j2ka43/ruby_30_adds_hashexcept_and_envexcept/)
- url: https://blog.saeloun.com/2020/09/30/ruby-adds-support-for-hash-except
---

## [9][Terraspace: The Terraform Framework](https://www.reddit.com/r/ruby/comments/j2m21y/terraspace_the_terraform_framework/)
- url: https://www.reddit.com/r/ruby/comments/j2m21y/terraspace_the_terraform_framework/
---
Hi all, created Terraspace, a Terraform Framework. It makes it more fun and easy to work with Terraform. Useful for folks building infrastructure-as-code.

Terraspace is written in Ruby. Though it's more like "Ruby Sprinkles" on top of Terraform, you still have the full power of Ruby at your fingertips. Found that Ruby is one of the most expressive languages to build tools and frameworks like Terraspace.

Sharing with the Ruby community. 

Docs: https://terraspace.cloud/
Blog Post: https://blog.boltops.com/2020/08/22/introducing-terraspace-the-terraform-framework
## [10][Matz in 2019: ".. unless extremely bad things happen [next year]"](https://www.reddit.com/r/ruby/comments/j2fcm8/matz_in_2019_unless_extremely_bad_things_happen/)
- url: https://i.redd.it/0qrcfnfgp7q51.png
---

