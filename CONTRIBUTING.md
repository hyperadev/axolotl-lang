# Welcome!

Welcome to the Axolotl locale files, first off, thank you for taking the time to consider contributing!
All contributions are extremely helpful and greatly appreciated!

This document contains a set of guides for contributing to this project.

<details>
<summary>Table of Contents</summary>

<!-- TOC -->
* [Welcome!](#welcome)
* [Code of Conduct](#code-of-conduct)
* [Questions](#questions)
* [Contributing](#contributing)
  * [Submitting issues](#submitting-issues)
  * [Translating a message](#translating-a-message)
  * [Translating to a new language](#translating-to-a-new-language)
<!-- TOC -->
</details>

# Code of Conduct

Please help keep this project open and inclusive for all.  
Read and follow the [Code of Conduct](https://github.com/HyperaDev/.github/blob/main/CODE_OF_CONDUCT.md) before
contributing to this repository.

If you have encountered someone who is not following the Code of Conduct, please report them
to [oss@hypera.dev](mailto:oss@hypera.dev).

# Questions

> **Please do not use GitHub issues to ask questions.** You will get a faster response if you ask on Discord!

If you wish to ask a question, please contact us using Discord by joining
the [Hypera Development Discord server](https://discord.hypera.dev/), and you will get a response as soon as
someone is next available.

# Contributing

There are many ways to contribute to `hypera.dev/lib`, and they all help!  
Here are the most common types of contributions:

* [Submitting issues](#submitting-issues)
  * [Security vulnerabilities](#security-vulnerabilities)
* [Suggesting features](#suggesting-features)
* [Code contributions](#code-contributions)
* [Supporting the authors](#supporting-the-authors)

## Submitting issues

If you have found a typo, grammatical issue, missing translation, or another issue, you can help us
by [creating an issue](https://github.com/HyperaDev/axolotl-lang/issues/new)!

If you have the time and required knowledge, you can
also [create a pull request](https://github.com/HyperaDev/axolotl-lang/compare) to resolve the issue!

## Translating a message

If you have found a message that has not been translated to an already supported locale that you can translate, you can
help us greatly by [create a pull request](https://github.com/HyperaDev/axolotl-lang/compare) to translate the message!

## Translating to a new language

Do you know a language that we don't currently support? You can help us greatly by contributing a new translation file
for the language! If you only have the time to translate some of the messages, that is okay! Having some translations is
a lot better than none.

Please note that we are only able to support locales also supported by Discord.
You can see all locales supported by Discord at https://discord.com/developers/docs/reference#locales.

To translate to a new language:
 - Create a file in [`locales`](locales) for the language, using the BCP 47 language tag.
 - Translate messages from [`locales/en-GB.toml`](locales/en-GB.toml) (or the language file that you can understand).
 - Add the locale to `supportedLocales` in [`lang.go`](lang.go) (You can see all options
in [discordgo](https://github.com/bwmarrin/discordgo/blob/master/locales.go))
 - [Create a pull request](https://github.com/HyperaDev/axolotl-lang/compare) to contribute them to this repository!


Thank you to everyone who has contributed! :D
