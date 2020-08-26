# youtube-cli

![youtube-cli banner](banner.png)

## List of Content

- [About](#about)
- [Usage](#usage)
- [License](#license)

## About

Youtube cli interface. Currently is only capable of generating a simple subscription feed page.

## Setup

Clone the repo with submodules using the following command:

```
git clone https://github.com/SunGrow/youtube-cli.git
```

## Usage

All of the required flags should be set before calling the command.
To see the list of all possible commands use

```sh
./youtube-cli -h

```

Subscription list file shoud be generated before the subscription feed generation.

```sh
./youtube-cli sublist

```

You can get your subscription_manager xml file from [here](https://www.youtube.com/subscription_manager?action_takeout=1)

## License

See [LICENSE](LICENSE).

