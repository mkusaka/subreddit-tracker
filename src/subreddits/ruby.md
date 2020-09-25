# ruby
## [1][Managing Bundler and Rubygems with Ansible](https://www.reddit.com/r/ruby/comments/izgtki/managing_bundler_and_rubygems_with_ansible/)
- url: https://thibautbarrere.com/2020/09/19/managing-bundler-and-rubygems-with-ansible
---

## [2][Digital Signatures for Ruby app, clueless](https://www.reddit.com/r/ruby/comments/izcm0g/digital_signatures_for_ruby_app_clueless/)
- url: https://www.reddit.com/r/ruby/comments/izcm0g/digital_signatures_for_ruby_app_clueless/
---
How do I get usable certificates for creating a simple document signing app using the below libraries? PSPDFKit for annotating the PDF with an ink signature, then Origami for applying a digital signature to the PDF.

I tried to buy an Adobe approved trustlist cert, but they only offer them on USB and and try to charge me 10x as much for aws cloudhsm support. and aws cloudhsm is super expensive as well.

How am I supposed to do this? Am I barking up the wrong tree with aws cloudhsm?

[https://pspdfkit.com/guides/web/current/digital-signatures/digital-signatures-on-web/](https://pspdfkit.com/guides/web/current/digital-signatures/digital-signatures-on-web/)

[https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e](https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e) 
## [3][Some good gui framework](https://www.reddit.com/r/ruby/comments/ize6lw/some_good_gui_framework/)
- url: https://www.reddit.com/r/ruby/comments/ize6lw/some_good_gui_framework/
---
Hello, am looking for a good gui framework or "gem" for ruby, while i know some will recommend c# i tried it 5 times and it doesnt click me...

Thanks.
## [4][How to use proper error handling flow in Ruby?](https://www.reddit.com/r/ruby/comments/iz1l3a/how_to_use_proper_error_handling_flow_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/iz1l3a/how_to_use_proper_error_handling_flow_in_ruby/
---
I'm new to developing software and trying to learn best practices.

I have a class to wrap AWS Dynamo API calls that will be used by several other classes.

```
class Dynamo &lt; AWS
    def initialize profile=nil, dryrun=true, region, table_name
      super
      @region = region
      @table_name = table_name
    end

    def put_item(item)
      begin
        get_aws_client(
            :dynamodb,
            { region: @region}
          ).put_item({
            table_name: @table_name,
            item: item
          })
      rescue  Aws::DynamoDB::Errors::ServiceError =&gt; e
        log :error, "#{e}"
      end
    end
  end
```

What is the proper error handling and logging workflow if another class uses this class to utilize the put\_item() method and it throws an error?
## [5][Improve the performance of super by eileencodes • Pull Request #3545](https://www.reddit.com/r/ruby/comments/iyp7rz/improve_the_performance_of_super_by_eileencodes/)
- url: https://github.com/ruby/ruby/pull/3545
---

## [6][Resize and optimise images on upload with ActiveStorage](https://www.reddit.com/r/ruby/comments/iyxsty/resize_and_optimise_images_on_upload_with/)
- url: https://vitobotta.com/2020/09/24/resize-and-optimise-images-on-upload-with-activestorage/
---

## [7][Packwerk is a Ruby gem used to enforce boundaries and modularize Rails applications.](https://www.reddit.com/r/ruby/comments/iyfrz7/packwerk_is_a_ruby_gem_used_to_enforce_boundaries/)
- url: https://github.com/shopify/packwerk
---

## [8][My First open-source Project: Building AmazonPay Gem](https://www.reddit.com/r/ruby/comments/iy7m2f/my_first_opensource_project_building_amazonpay_gem/)
- url: https://www.reddit.com/r/ruby/comments/iy7m2f/my_first_opensource_project_building_amazonpay_gem/
---
I built my first open-sourc project and I wrote about it and my motivations here: [https://blog.everistus.xyz/my-first-open-source-project-building-amazonpay-gem-ckff8s5ov043o65s13ngx0pil](https://blog.everistus.xyz/my-first-open-source-project-building-amazonpay-gem-ckff8s5ov043o65s13ngx0pil)
## [9][Structish Objects](https://www.reddit.com/r/ruby/comments/iykhi5/structish_objects/)
- url: https://www.reddit.com/r/ruby/comments/iykhi5/structish_objects/
---
I've been working on a gem that's very much experimental, and would love some feedback. The ReadMe is not complete, but I'm hoping there's enough to get the point across. What I'd really like to find out is:

\- Is there anything out there that already satisfies these requirements but better?

\- Would you as a dev find this data structure useful?

\- Is my implementation the ideal way to go about things?

&amp;#x200B;

[https://github.com/DylanBlakemore/structish](https://github.com/DylanBlakemore/structish)
## [10][Rails will allow a module with extend ActiveSupport::Concern to be prepended](https://www.reddit.com/r/ruby/comments/iy91jy/rails_will_allow_a_module_with_extend/)
- url: https://blog.saeloun.com/2020/09/23/prepend-class-methods-for-concerns.html
---

