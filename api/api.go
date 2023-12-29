package api

import (
    "fmt"
    jikan "github.com/darenliang/jikan-go"
)

type AnimeInfo struct {
    Title           string   `json:"title"`
    TitleEnglish    string   `json:"title_english"`
    TitleJapanese   string   `json:"title_japanese"`
    Type            string   `json:"type"`
    Episodes        int      `json:"episodes"`
    Status          string   `json:"status"`
    Airing          bool     `json:"airing"`
    AiredFrom       string   `json:"aired_from"`
    AiredTo         string   `json:"aired_to"`
    Score           float64  `json:"score"`
    Broadcast       string   `json:"broadcast"`
    Licensors       []string `json:"licensors"`
    Producers       []string `json:"producers"`
    Popularity      int      `json:"popularity"`
    Background      string   `json:"background"`
    Demographics    []string `json:"demographics"`
    TitleSynonyms   []string `json:"title_synonyms"`
    ExtplicitGenres []string `json:"explicit_genres"`
    Themes          []string `json:"themes"`
    Members         int      `json:"members"`
    Studios         []string `json:"studios"`
    ScoredBy        int      `json:"scored_by"`
    Duration        string   `json:"duration"`
    Favorites       int      `json:"favorites"`
    Url             string   `json:"url"`
    Rank            int      `json:"rank"`
    Year            int      `json:"year"`
    MalId           int      `json:"mal_id"`
    Rating          string   `json:"rating"`
    Season          string   `json:"season"`
    Source          string   `json:"source"`
}

func FormatAnimeInfo(animeInfos []AnimeInfo) string {
    var formattedData string

    for _, info := range animeInfos {
        formattedData += fmt.Sprintf("[+] Title: %s\n", info.Title)
        formattedData += fmt.Sprintf("    + English Title: %s\n", info.TitleEnglish)
        formattedData += fmt.Sprintf("    + Japanese Title: %s\n", info.TitleJapanese)
        formattedData += fmt.Sprintf("    + Type: %s\n", info.Type)
        formattedData += fmt.Sprintf("    + Episodes: %d\n", info.Episodes)
        formattedData += fmt.Sprintf("    + Status: %s\n", info.Status)
        formattedData += fmt.Sprintf("    + Airing: %v\n", info.Airing)
        formattedData += fmt.Sprintf("    + Aired From: %s\n", info.AiredFrom)
        formattedData += fmt.Sprintf("    + Aired To: %s\n", info.AiredTo)
        formattedData += fmt.Sprintf("    + Score: %.2f\n", info.Score)
        formattedData += fmt.Sprintf("    + Popularity: %d\n", info.Popularity)
        formattedData += fmt.Sprintf("    + Background: %s\n", info.Background)
        formattedData += fmt.Sprintf("    + Members: %d\n", info.Members)
        formattedData += fmt.Sprintf("    + Scored By: %d\n", info.ScoredBy)
        formattedData += fmt.Sprintf("    + Duration: %s\n", info.Duration)
        formattedData += fmt.Sprintf("    + Favorites: %d\n", info.Favorites)
        formattedData += fmt.Sprintf("    + URL: %s\n", info.Url)
        formattedData += fmt.Sprintf("    + Rank: %d\n", info.Rank)
        formattedData += fmt.Sprintf("    + Year: %d\n", info.Year)
        formattedData += fmt.Sprintf("    + MAL ID: %d\n", info.MalId)
        formattedData += fmt.Sprintf("    + Rating: %s\n", info.Rating)
        formattedData += fmt.Sprintf("    + Season: %s\n", info.Season)
        formattedData += fmt.Sprintf("    + Source: %s\n", info.Source)
        formattedData += "----------------------------\n"
    }

    return formattedData
}

func GetData(filter jikan.ScheduleFilter) (string, error) {
    schedules, err := jikan.GetSchedules(filter)
    if err != nil {
        return "", fmt.Errorf("error getting schedules: %v", err)
    }

    var animeInfos []AnimeInfo

    for _, anime := range schedules.Data {
        info := AnimeInfo{
            Title:         anime.Title,
            TitleEnglish:  anime.TitleEnglish,
            TitleJapanese: anime.TitleJapanese,
            Type:          anime.Type,
            Episodes:      anime.Episodes,
            Status:        anime.Status,
            Airing:        anime.Airing,
            AiredFrom:     anime.Aired.From.Format("2000-02-02"),
            AiredTo:       anime.Aired.To.Format("2000-02-02"),
            Score:         anime.Score,
            Popularity:    anime.Popularity,
            Background:    anime.Background,
            TitleSynonyms: anime.TitleSynonyms,
            Members:       anime.Members,
            ScoredBy:      anime.ScoredBy,
            Duration:      anime.Duration,
            Favorites:     anime.Favorites,
            Url:           anime.Url,
            Rank:          anime.Rank,
            Year:          anime.Year,
            MalId:         anime.MalId,
            Rating:        anime.Rating,
            Season:        anime.Season,
            Source:        anime.Source,
        }

        animeInfos = append(animeInfos, info)
    }

    formattedData := FormatAnimeInfo(animeInfos)

    return formattedData, nil
}
