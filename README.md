# www.lucasyamamoto.com.golang

[![Netlify Status](https://api.netlify.com/api/v1/badges/f253ccc8-baf3-4239-b39a-f8aa7b879308/deploy-status)](https://app.netlify.com/sites/lucasyamamoto-golang/deploys)

www.lucasyamamoto.com.golang is a blog built with my own static site generator written in Go. It parses markdown files, applies layout templates, and generates a complete website including assets and sitemaps, ready to be deployed.

## Features

- **Markdown Parsing**: Efficiently parses markdown files from `_posts` and `_pages` folders, transforming them into HTML while extracting frontmatter.
- **Template Loading**: Utilizes the `text/template` library to apply `_layouts` and include `_includes`, enabling reusable website components.
- **Site Generation**: Generates the entire website, including homepage, posts, pages, assets, and a sitemap, according to the specifications in `_config.yml`.
- **Dynamic Content**: Supports pagination and JSON data generation, allowing for dynamic content presentation.
- **Sitemap Generation**: Automatically generates a sitemap for SEO optimization, ensuring your site is fully indexed by search engines.

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.21).
- Basic understanding of Go programming.
- Familiarity with Markdown and YAML configurations.
- Node 18.x (used only to build the SASS)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/lucasayb/www.lucasyamamoto.com.golang.git
cd www.lucasyamamoto.com.golang
```

### Development
In the `.vscode` folder, I added a two files:
* `launch.json`: Every time that any file is saved into the repository, the command to build the site is executed
* `settings.json`: I added the `_site` as the base folder in which the live server will run. 

### Usage

1. **Configure Your Site**: Edit the `_config.yml` file to match your site's specifications.
2. **Prepare Content**: Add your markdown files to the `_posts` and `_pages` directories.
3. **Add you SASS files**: The SASS files needs to be into the `sass` folder.
4. **Build Your CSS**: Run the following command to build your css:
```bash
yarn build:css
```
5. **Build Your Site**: Run the following command to build your site:
```bash
yarn build
```
Keep in mind that both commands will run after the `yarn install` because they are into the `postinstall` scripts of the `package.json`

The site will be generated in the `_site` directory, ready to be deployed.

## Project Structure

- `parser/parser.go`: Parses markdown files, transforms them into HTML, and extracts frontmatter.
- `loader/loader.go`: Applies layout templates and includes using the `text/template` library.
- `generator/generator.go`: Responsible for generating the entire site, including assets, pages, and the sitemap.

## Contributing

We welcome contributions to www.lucasyamamoto.com.golang! Whether it's submitting bug reports, feature requests, or contributing code, here's how you can get involved:

1. **Fork the repository** and create your feature branch: `git checkout -b my-new-feature`.
2. **Commit your changes**: `git commit -am 'Add some feature'`.
3. **Push to the branch**: `git push origin my-new-feature`.
4. **Submit a pull request**: Ensure your code adheres to the project's standards and includes appropriate tests and documentation.

## License

This project is open source and available under the [MIT License](LICENSE).
