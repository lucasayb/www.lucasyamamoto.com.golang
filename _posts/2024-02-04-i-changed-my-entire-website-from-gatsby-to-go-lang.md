---
title: I changed my entire website from Gatsby to Go lang
date: 2024-02-04 12:09:18
description: Why did I learned a bit of Go and decided to create my own SSG to
  use on my blog
thumbnail: /assets/uploads/dall·e-2024-02-04-12.42.59-design-an-eye-catching-800x400-thumbnail-for-an-article-discussing-the-go-programming-language.-the-image-should-feature-a-stylized-cartoonish-gopher.webp
category: Development
color: "#e74c3c"
---
Right now, I'm recovering from surgery at my dear friend Victor Almeida's home. While doing so, and being away from my work, I was really bored in the past couple of days. Because of this, I decided to do something: I realized that my website needed an upgrade. The main reason I have this website/blog is because I like to experience new things and conduct tests, and there's nothing better than my own controlled environment for that.

The next logical technology to use on my website would be NextJS, as I see a lot of bloggers switching from Gatsby to Next due to the performance gains, among other reasons, and the ease of implementations, as NextJS has really advanced modules.

While I like to stay informed about what technology the market is using, I really wanted to go through a different path. If you work with a bit of DevOps, you know for a fact that there are a lot of tools built with Go language. The reason why is because Go has an easy syntax (maybe not as easy as Python), an amazing quantity of standard libraries that you don't need to install, and because it's compiled, which makes it amazingly fast, much faster than interpreted languages. Some might say that Go is the modern implementation of C, but I don't have enough authority to agree with that.

Anyway, if it's still not clear, I chose Go as my new language to build my own static site generator. And yes, there are many SSGs existing in the market, with Hugo being one of those that is, by the way, built with Go.

Although I find the implementation of some SSGs, such as Hugo itself, Gatsby, and Jekyll (one that I used before and was really considering going back to), fascinating, I wanted to build my own, not to reinvent the wheel, but because I wanted a challenge so that I could learn how to program with Go. This is something that I always do to learn new technologies: I choose the technology, read a bit of the basic syntax in its documentation, and decide to build something with it, which is, in my opinion, one of the best ways to learn something new in the development world.

## Wins

What do I gain by switching to Go? Build time! Let's be fair: when you build a website with Gatsby, it does much more than just transforming Markdown files into HTML. It builds different versions of the images, creates an entire React distribution directory so that you can utilize all the potential of React, and the list goes on. However, for my blog, my personal testing environment, those features are not that crucial. They are important in contexts where you require high performance, such as for an e-commerce site or a SaaS platform, but for my blog, they are less so. What is important, however, is the build time. It used to take 2 minutes to build the entire site with Gatsby; now it takes 23 seconds, and I'm not even using the pre-compiled version of the app.

### Before

![Deployed in 2m and 9 seconds using Gatsby](/assets/uploads/screenshot-2024-02-04-at-11.05.00.jpg "Build time with GatsbyJS")

### After

![Deployed in 23s with the SSG built in Go](/assets/uploads/screenshot-2024-02-04-at-11.05.47.jpg "Build time with the SSG build in Go")

It may not seem much, but when you want to publish articles in a frequent way and fast, it is important to have such reduced time. 

## Losses

I must say that I likely lost more features than I gained with this change. Not because I built it in Go, but because I built a static site generator (SSG) from scratch. Had I used a framework such as Hugo, probably half of these issues would have already been resolved.

### Page Speed scores

#### GatsbyJS

![It shows 65% as performance, 88% for accebility, 100 for best practices and 99 for SEO](/assets/uploads/screenshot-2024-02-04-at-11.15.47.jpg "Performance before with Gatsby")

#### SSG built with Go

![It shows 57% for performance, 71% for accessibility, 96% for best practices and 91 for SEO](/assets/uploads/screenshot-2024-02-04-at-11.15.11.jpg "Performance with the SSG built using Go")

It’s nice to say that I’m losing here because I didn’t build almost anything that Gatsby has implemented by default, like the [blur-up effect, which first loads the image in much lower quality, but blurred, and then shows it fully when the image is already loaded.](https://gracious-jennings-0e5ca9.netlify.app) This aims to reduce the FCP and the CLS. Also, SEO and best practices need to be improved and will be improved.

### Search

That's right, I didn't implement the search feature. I was so eager to deploy the site ASAP that I just didn't implement it. But it's possible, and maybe even better, because of how fast Go handles processing in general. I think I'll implement it soon.

### Apple Shortcuts

Recently, I built the Apple Shortcuts page, a place where I could post the automations that I built using Apple's Shortcuts that just help me on a daily basis. The reason why I didn’t implement it is the same reason why I didn’t implement any category listing page. It’s just that I haven’t figured out yet the best way to deal with multiple listing pages.

Let me show the `generator.GenerateHome` method that I built in the `generator.go`.

```go
func GenerateHome(config parser.Config, posts []parser.Post, pages parser.Pages, output string) string {
	createFolder(output)
	from := 0
	to := pages.PerPage


	posts = sortPosts(config, posts)
	postsCount := len(posts)
	var rendered bytes.Buffer
	var fileName string
	for page := 1; page <= pages.PagesCount; page++ {
		if to > postsCount {
			to = postsCount
		}
		postsPage := posts[from:to]
		previousPage := page - 1
		nextPage := page + 1
		paginationData := loader.PaginationData{
			PreviousPage:     previousPage,
			NextPage:         nextPage,
			Page:             page,
			Pages:            pages.PagesCount,
			ShowPreviousPage: previousPage >= 1,
			ShowNextPage:     nextPage <= pages.PagesCount,
		}
		homepageData := loader.HomepageData{
			Posts:      postsPage,
			Config:     config,
			Pagination: paginationData,
		}
		rendered = loader.Homepager(homepageData)
		if page == 1 {
			fileName = "index"
		} else {
			fileName = "index-" + strconv.FormatInt(int64(page), 10)
		}
		createFile(output, fileName, rendered.Bytes())
		to += pages.PerPage
		from += pages.PerPage
	}
	fmt.Println("Home generated successfully")
	return rendered.String()
}
```

If you look at it carefully, I just grab the information from posts and create a few different index pages so that I can have pagination.

And in the layout `home.html`, I just render this information:

```html
{{ define "Home" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{ template "Head" .Config }}
    <title>{{ .Config.Title }}</title>
  </head>
  <body class="dark">
    {{ template "Header" }}
    <main class="content">
      <div class="container">
        {{ range .Posts }}
          {{ template "PostItem" . }}
        {{ end }}
        {{ template "Pagination" . }}
      </div>
    </main>
    {{ template "Body" .Config }}
  </body>
</html>
{{ end }}
```

Every single post on this site is currently in the `_posts` folder of the project. I took the same idea that I liked from Jekyll. I could have just created another method called `GenerateAppleShortcuts` that would do the same thing as `GenerateHome` but for another page, but I didn’t think it would be as modular as I wanted. I’ll explain later why I just didn’t do it anyway.

## Deployment with Netlify

I know that in the wild, what is used the most is Vercel; at Keyrus, we use it a lot to deploy things from the startups and tests that we do, but I just love the way Netlify handles things. And it’s free! In the build phase on Netlify, I can just run my command \`bin/build\` that runs my `go run .`, because yes, Netlify also uses Go in its cloud. I’m not certain if Vercel can do the same thing; I believe it can, but the seamless integration with Decap CMS (formerly Netlify CMS) is just as easy as it could be. I just pointed out that my entire site was in the `_site` folder and done! Having worked and deployed so many systems on Digital Ocean, AWS EC2, or S3, I was just delighted that this phase was so easy that I just couldn’t believe it.

## The Decap CMS

When I implemented this site with Gatsby, I integrated Netlify CMS. Recently, it transitioned to Decap CMS, but it remains the same great CMS as before. As I retained all the folder structure and formatting that I had previously established when building this SSG, everything functioned as it did before. I merely had to add the admin folder with the `config.yml`, and that was it.

## The Project

The project can be found here. I've added a `README.md` and explained a few things (with the help of ChatGPT, obviously). The older version is available here. I've left everything open because it's beneficial to share some achievements and gather opinions from people.

I do plan to build a library from this project, not because I think I will create the new Hugo, but to experience the entire process of building a library and possibly finding a direct and simple way to implement sites like Jekyll does. I have a great admiration for Jekyll and have been contemplating building a theme converter for that SSG. Maybe? I'm not certain yet.

## I’m Not a Go Reference

I plan to continue improving in the Go language world. However, I'm still primarily a web developer trying to learn a complex language like Go to explore other possibilities. I have many ideas, and Go is included in most of them. But everything I've done is just simple Go programming. I'm not sure if any Go developer would take a look and be dismayed (I hope not; please send me a message before doing so). However, I would love to hear opinions from Go developers and learn what I can improve in my code.
