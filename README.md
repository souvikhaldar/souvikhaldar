# Souvik Haldar

[![Website](https://img.shields.io/badge/website-souvikhaldar.in-0A66C2?style=flat-square)](https://souvikhaldar.in)
[![Built with Hugo](https://img.shields.io/badge/built%20with-Hugo-FF4088?style=flat-square)](https://gohugo.io/)
[![Deployed on Netlify](https://img.shields.io/badge/deployed%20on-Netlify-00C7B7?style=flat-square)](https://www.netlify.com/)

Personal website and writing archive for notes on security, programming, networking, blockchain, tools, events, and long-form ideas.

## About

I am Souvik Haldar, founder of [Hornet](https://www.hornet.technology/). I work across software engineering, cybersecurity, distributed systems, Golang, Linux, blockchain, and infrastructure.

This repository powers [souvikhaldar.in](https://souvikhaldar.in), a Hugo static site built with the LoveIt theme and deployed through Netlify.

## Writing Topics

- Security research, exploitation notes, and operational tooling
- Golang, backend engineering, algorithms, and systems programming
- Networking, distributed systems, and infrastructure
- Blockchain fundamentals and cryptocurrency notes
- Productivity, learning, events, and personal essays

## Local Development

Install Hugo Extended, then run:

```bash
hugo server
```

Build the production site:

```bash
hugo --minify
```

The generated static site is written to `public/`.

## Deployment

This site is deployed on Netlify.

```toml
[build]
  command = "hugo --minify"
  publish = "public"

[build.environment]
  HUGO_VERSION = "0.152.2"
```

The custom domain is configured as:

- Primary domain: `souvikhaldar.in`
- Redirect domain: `www.souvikhaldar.in`

## Repository Structure

```text
content/     Site pages and articles
layouts/     Custom Hugo templates
static/      Images, CSS, and JavaScript
themes/      Hugo theme submodule
public/      Generated build output, ignored by git
```

## Contact

- Website: [souvikhaldar.in](https://souvikhaldar.in)
- Email: [mail@souvikhaldar.in](mailto:mail@souvikhaldar.in)
