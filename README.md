# AniTsu (Anime no Tsuiseki | アニメの追跡)

> AniTsu is a CLI Anime Scheduler crafted in Go, designed to enhance your anime-watching experience. With AniTsu, effortlessly manage and organize your anime journey with a few simple commands. Track your favorite shows, schedule upcoming episodes, and stay on top of your anime viewing schedule, all from the command line.

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## Project structure:

```go
AniTsu
│
├── api
│   ├── api.go
│   └── api_test.go
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Installation

```sh
go get -u github.com/kenjitheman/anitsu
```

## Usage

```sh
anitsu -d day_of_the_week
```

## Output exaple

```
[+] Title: Sousou no Frieren
    + English Title: Frieren: Beyond Journey's End
    + Japanese Title: 葬送のフリーレン
    + Type: TV
    + Episodes: 28
    + Status: Currently Airing
    + Airing: true
    + Aired From: 2023-04-29
    + Aired To: 2024-02-01
    + Score: 9.1
    + Popularity: 550
    + Background: Sousou no Frieren was released on Blu-ray and DVD in seven volumes from January 24, 2024, to July 17, 2024.
    + Members: 407932
    + Scored By: 118438
    + Duration: 24 min per ep
    + Favorites: 11546
    + URL: https://myanimelist.net/anime/52991/Sousou_no_Frieren
    + Rank: 1
    + Year: 2023
    + MAL ID: 52991
    + Rating: PG-13 - Teens 13 or older
    + Season: fall
    + Source: Manga
----------------------------
and in the same style others that arrive this day...
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](./LICENSE)
