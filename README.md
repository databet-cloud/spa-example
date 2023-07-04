# Databet Go SDK

This sdk provides clients and tools to:

1. Connect to feed service and import its data
2. Apply patches from feed to existing sport events
3. Communicate with mts service and place/decline bets, etc...
4. Fetch API resources, such as sports, markets, tournaments, teams, players, ...

## Authorization

To use provided clients, you should configure your http client to pass tls certificate:

```go
// Parse tls certificate
certificate, err := tls.LoadX509KeyPair("certFile", "keyFile")
if err != nil {
    panic(err)
}

// Create http client and pass tls certificate to tls config
tr := &http.Transport{
    TLSClientConfig: &tls.Config{
        Certificates: []tls.Certificate{certificate},
    },
}

httpClient := &http.Client{
    Timeout:   0, // must be explicitly set to 0, since feed is a stream of events
    Transport: tr
}
```

## Feed
[https://feed.databet.cloud](https://mts.databet.cloud)

Feed client allows you to synchronize your system with an actual feed version

[More Info](https://docs.data.bet/feed/)

```go
client := feed.NewClientHTTP(httpClient, "https://feed.databet.cloud")

cursor, lastVersion, err := client.GetAll(context.Background())
if err != nil {
    panic(err)
}

fmt.Println(lastVersion) // 1z8hsCeNdtl000004gfQ0H

for cursor.HasMore() {
    msg, err := cursor.Next(context.Background())
    if err != nil {
        panic(err)
    }
    
    var sportEvent sportevent.SportEvent
    
    _ = json.Unmarshal(msg, &sportEvent)
    
    // Handle sport event
}
```

You should use GetAll only once, during the first synchronization.
Once you synchronized with an actual feed version, you can pass it to GetFromVersion method.

```go
cursor, err := client.GetFromVersion(context.Background(), lastVersion)
if err != nil {
    panic(err)
}

for cursor.HasMore() {
    msg, err := cursor.Next(context.Background())
    if err != nil {
        panic(err)
    }
    
    var log feed.LogEntry
    
    err = json.Unmarshal(msg, &log)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(log.Version) // 2z9hsCeNdtl000005gfQ0J
    
    // Handle log entry
}
```

To get the last feed version, simply use:

```go
version, err := client.GetFeedVersion(context.Background())
```

### Feed logs
Feed logs has 3 different types:
* match_new: has field "sport_event" with the full sport event data in it
* match_update: has field "changes", that can be handled by sportevent.Patcher
* bet_rollback: indicates that you should rollback all bets from "dt_start" to "dt_end" for the specified "market_ids" and "match_id"

### Sport Event Patcher
```
patcher := sportevent.NewPatcherSimdJSON()

err := patcher.ApplyPatches(sportEvent, log.Changes)
if err != nil {
    panic(err)
}
```

## MTS
[https://mts.databet.cloud](https://mts.databet.cloud)

The DataBet platform allows you to track bookmaker's bets and modify odds values, according to amounts of stakes. Also, the platform can tell you if a potential bet would violate any restrictions.

Note: The system is multi-currency, but to integrate correctly you have to provide the platform administrator a full list of currencies in which bets can be accepted.

[More Info](https://docs.data.bet/bets-accounting/)

## API
[https://api.databet.cloud](https://api.databet.cloud)

DataBet API client allows you to fetch markets, sports, tournaments, organizations, players and teams by their ids/filters and translates names to the given locale.

## Statistics
[https://statistics.databet.cloud](https://statistics.databet.cloud)

DataBet API client allows you to fetch sport event versioned statistics (each sport has a different set of parameters in statistics) by their id and version. 

## Common errors
* **403 Forbidden** means that you should add your ip to whitelist, using [STM](https://stm.databet.cloud) bookmaker settings
* **401 Unauthorized** indicates problems with your bookmaker tls certificate, contact support to fix it